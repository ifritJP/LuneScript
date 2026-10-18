// LuneScript から変換した wasm (WasmGC) を Node.js で実行する。
//
// usage: node lnswasm.mjs xxx.wasm [引数...]
//
// 変換した wasm は、C のランタイム (モジュール名 "lnsrt") の memory と関数を import する。
// C のランタイムは WASI のコマンドで、Node.js の node:wasi で動かす。
//  - lnsrt.wasm: 入出力など (lnsrt.c から生成)
//  - lnslua.wasm: lnsrt.c と Lua (lnslua.c から生成)。変換した wasm が Lua の関数を import している場合に使う
// C のランタイムの main() が、変換した wasm の main (モジュール名 "lnsmain") を呼び出す。
// C のランタイムのディレクトリは環境変数 LNS_WASM_RT_DIR で指定できる (lnslua.wasm だけは LNS_WASM_LUA でも指定できる)。
//
// Node.js でだけ使える次の関数 (モジュール名 "lnsjs") も提供する。
// WASI のランタイム (wasmtime など) では、何もしない関数 (lnsjs_stub.wat) を使う。
//  - luaval_new: Lua の値 ($LuaVal) が GC で回収された時に、Lua の registry の登録を解除する
//  - popen / popen_data: io.popen
//
// 変換した wasm は、最初から最適化コンパイラ (TurboFan) でコンパイルする (V8 の --no-liftoff)。
// V8 の wasm には実行中の関数を最適化版に切り替える仕組み (OSR) が無く、
// 1 回しか呼ばれない関数 (main など) の中の重いループが、
// ベースラインのコンパイラ (Liftoff) のコードのまま動いてしまうため。
// C のランタイムは、何度も呼ばれる関数が多く段階的な最適化で十分なので、
// --no-liftoff を設定する前に (Liftoff で) コンパイルする。
// ただし、512KB 以上の大きな wasm (lnsc 自身など) は、全ての関数のコンパイルに時間がかかるので、
// V8 の標準の動作 (Liftoff から段階的に最適化) にする。
// 環境変数 LNS_WASM_LIFTOFF=1 / 0 で、どちらかに固定できる。

import { readFileSync, writeSync } from 'node:fs';
import { fileURLToPath } from 'node:url';
import { setFlagsFromString } from 'node:v8';
import { execSync } from 'node:child_process';

/** C のランタイムの wasm のパス */
function getRtWasmPath(useLua) {
  if (useLua && process.env.LNS_WASM_LUA) {
    return process.env.LNS_WASM_LUA;
  }
  const name = useLua ? 'lnslua.wasm' : 'lnsrt.wasm';
  if (process.env.LNS_WASM_RT_DIR) {
    return `${process.env.LNS_WASM_RT_DIR}/${name}`;
  }
  return fileURLToPath(new URL(`../../work/wasm-lua/${name}`, import.meta.url));
}

/** LEB128 (符号なし) を読み、値と次の位置を返す */
function readUleb(bin, pos) {
  let val = 0;
  let shift = 0;
  for (;;) {
    const byte = bin[pos++];
    val += (byte & 0x7f) * 2 ** shift;
    if ((byte & 0x80) === 0) {
      return [val, pos];
    }
    shift += 7;
  }
}

/**
変換した wasm が Lua の関数 (C のランタイムの lnsl_xxx) を import しているかどうか。

C のランタイムのどちら (lnsrt.wasm / lnslua.wasm) を使うかを決めるために使う。
--no-liftoff を設定する前にコンパイルしないように、import セクションを直接読む。
*/
function usesLua(bin) {
  let pos = 8;
  while (pos < bin.length) {
    const id = bin[pos++];
    let size;
    [size, pos] = readUleb(bin, pos);
    if (id === 2) {
      let num;
      let work = pos;
      [num, work] = readUleb(bin, work);
      for (let index = 0; index < num; index++) {
        let len;
        [len, work] = readUleb(bin, work);
        const modName = bin.toString('utf8', work, work + len);
        work += len;
        [len, work] = readUleb(bin, work);
        const name = bin.toString('utf8', work, work + len);
        work += len;
        if (modName === 'lnsrt' && name.startsWith('lnsl_')) {
          return true;
        }
        const kind = bin[work++];
        switch (kind) {
          case 0x00: // func: 型インデックス
            [, work] = readUleb(bin, work);
            break;
          case 0x02: { // memory: limits
            const flag = bin[work++];
            [, work] = readUleb(bin, work);
            if (flag & 1) {
              [, work] = readUleb(bin, work);
            }
            break;
          }
          default:
            // 変換した wasm が import するのは func と memory だけ
            throw new Error(`unsupported import kind: ${kind}`);
        }
      }
      return false;
    }
    pos += size;
  }
  return false;
}

/**
@param wasmPath 実行する wasm のパス
@param args コマンドの引数 (__main に渡す。先頭は wasm のパス)
@return 終了コード
*/
export async function run(wasmPath, args = [wasmPath]) {
  const bin = readFileSync(wasmPath);
  // node:wasi の ExperimentalWarning は、実行結果の出力に混ざるので出さない
  const emitWarning = process.emitWarning;
  process.emitWarning = (warning, ...rest) => {
    if (String(warning).startsWith('WASI is an experimental feature')) {
      return;
    }
    emitWarning.call(process, warning, ...rest);
  };
  const { WASI } = await import('node:wasi');
  /*
    ルートを preopen し、C のランタイムのカレントディレクトリを
    ホストと同じにする (環境変数 LNS_WASM_CWD)。
  */
  const wasi = new WASI({
    version: 'preview1', returnOnExit: true, args,
    env: { ...process.env, LNS_WASM_CWD: process.cwd() }, preopens: { '/': '/' },
  });

  let mainExports = null;
  let rtExports = null;
  let popenData = null;
  const rtImports = {
    ...wasi.getImportObject(),
    lnsmain: { main: () => mainExports.main() },
    lnsjs: {
      // io.popen: コマンドを実行し、出力の長さを返す (出力は popen_data で書き込む)
      popen: (cmdAddr, cmdLen) => {
        const cmd = Buffer.from(rtExports.memory.buffer, cmdAddr, cmdLen).toString('utf8');
        try {
          popenData = execSync(cmd, { stdio: ['ignore', 'pipe', 'inherit'] });
        } catch (err) {
          // 終了コードが 0 以外でも、Lua と同じく出力は読める
          popenData = err.stdout ?? Buffer.alloc(0);
        }
        return popenData.length;
      },
      popen_data: (dst) => {
        new Uint8Array(rtExports.memory.buffer, dst, popenData.length).set(popenData);
        popenData = null;
      },
    },
  };
  /*
    Lua の値 ($LuaVal) が GC で回収された時に、Lua の registry の登録 (luaL_ref) を解除する。
    コールバックは wasm の実行中には呼ばれない (イベントループに戻った時に呼ばれる) ので、
    Lua の処理の途中に割り込むことはない。
  */
  const luaValRegistry = new FinalizationRegistry((ref) => {
    rtExports.lnsl_unref(ref);
  });

  /*
    C のランタイムは、--no-liftoff を設定する前に、全ての関数を Liftoff でコンパイルする。
    V8 は関数を最初に呼ばれた時にコンパイルする (遅延コンパイル) ので、そのままでは
    main() の中で行う Lua の初期化などが --no-liftoff の設定後に TurboFan でコンパイルされ、
    起動が約 40ms 遅くなる。
  */
  const noLiftoff = !useLiftoff(bin);
  if (noLiftoff) {
    setFlagsFromString('--no-wasm-lazy-compilation');
  }
  const { instance: rtInstance } = await WebAssembly.instantiate(
    readFileSync(getRtWasmPath(usesLua(bin))), rtImports);
  rtExports = rtInstance.exports;
  if (noLiftoff) {
    setFlagsFromString('--wasm-lazy-compilation');
    // 変換した wasm をコンパイルする前に設定する必要がある
    setFlagsFromString('--no-liftoff');
  }
  const { instance } = await WebAssembly.instantiate(bin, {
    lnsrt: rtExports,
    lnsjs: {
      luaval_new: (luaVal, ref) => {
        luaValRegistry.register(luaVal, ref);
      },
    },
  });
  mainExports = instance.exports;
  try {
    return wasi.start(rtInstance);
  } catch (err) {
    if (err instanceof WebAssembly.RuntimeError && process.env.LNS_WASM_STACK === '1') {
      // デバッグ用: LNS_WASM_STACK=1 の場合、error() で trap した時の wasm の関数のスタック
      writeSync(2, `${err.stack}\n`);
      return 1;
    }
    throw err;
  }
}

/**
変換した wasm を、V8 の標準の動作 (Liftoff から段階的に最適化) でコンパイルするかどうか。

大きな wasm (lnsc 自身など) は、最初から全ての関数を TurboFan でコンパイルすると
時間がかかる (lnsc で約 6 秒) ので、標準の動作にする。
環境変数 LNS_WASM_LIFTOFF=1 / 0 で、どちらかに固定できる。
*/
function useLiftoff(bin) {
  if (process.env.LNS_WASM_LIFTOFF === '1') {
    return true;
  }
  if (process.env.LNS_WASM_LIFTOFF === '0') {
    return false;
  }
  return bin.length >= 512 * 1024;
}

if (process.argv[1] && import.meta.url.endsWith(process.argv[1].split('/').pop())) {
  if (process.argv.length < 3) {
    console.error('usage: node lnswasm.mjs xxx.wasm');
    process.exit(1);
  }
  process.exitCode = await run(process.argv[2], process.argv.slice(2));
}
