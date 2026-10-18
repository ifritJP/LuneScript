// LuneScript から変換した wasm (WasmGC) を Node.js で実行する。
//
// usage: node lnswasm.mjs xxx.wasm
//
// wasm が import する関数 (モジュール名 "lns") を提供する。
// wasm が Lua の関数 (モジュール名 "lua") を import している場合は、
// wasm にした Lua (lnslua.wasm、lnslua.c から生成) を読み込んで渡す。
// lnslua.wasm のパスは環境変数 LNS_WASM_LUA で指定できる。
// 出力は Lua 5.3 の print と同じ形式にする。
//
//
// 変換した wasm は、最初から最適化コンパイラ (TurboFan) でコンパイルする (V8 の --no-liftoff)。
// V8 の wasm には実行中の関数を最適化版に切り替える仕組み (OSR) が無く、
// 1 回しか呼ばれない関数 (main など) の中の重いループが、
// ベースラインのコンパイラ (Liftoff) のコードのまま動いてしまうため。
// lnslua.wasm は、何度も呼ばれる関数が多く段階的な最適化で十分なので、
// --no-liftoff を設定する前に読み込んで初期化する (初期化が遅くならないように)。
// 環境変数 LNS_WASM_LIFTOFF=1 で、V8 の標準の動作 (Liftoff から段階的に最適化) にする。

import {
  readFileSync, writeSync, openSync, readSync, closeSync, fstatSync,
} from 'node:fs';
import { fileURLToPath } from 'node:url';
import { setFlagsFromString } from 'node:v8';

const outBuf = [];
let outLen = 0;

function write(bytes) {
  outBuf.push(bytes);
  outLen += bytes.length;
  if (outLen > 64 * 1024) {
    flush();
  }
}
function flush() {
  if (outBuf.length > 0) {
    writeSync(1, Buffer.concat(outBuf));
    outBuf.length = 0;
    outLen = 0;
  }
}

/** lnslua.wasm のパス */
function getLuaWasmPath() {
  if (process.env.LNS_WASM_LUA) {
    return process.env.LNS_WASM_LUA;
  }
  return fileURLToPath(new URL('../../work/wasm-lua/lnslua.wasm', import.meta.url));
}

/** lnslua.wasm を読み込み、初期化した instance の exports を返す */
async function loadLua() {
  // node:wasi の ExperimentalWarning は、実行結果の出力に混ざるので出さない
  const emitWarning = process.emitWarning;
  process.emitWarning = (warning, ...args) => {
    if (String(warning).startsWith('WASI is an experimental feature')) {
      return;
    }
    emitWarning.call(process, warning, ...args);
  };
  // Lua の io・os ライブラリは開いていないので、WASI の機能はほぼ使わない
  const { WASI } = await import('node:wasi');
  const wasi = new WASI({ version: 'preview1', returnOnExit: true });
  const bin = readFileSync(getLuaWasmPath());
  const { instance } = await WebAssembly.instantiate(bin, wasi.getImportObject());
  wasi.initialize(instance);
  instance.exports.lnsl_init();
  return instance.exports;
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
wasm の import セクションから、import しているモジュール名の Set を返す。

コンパイルする前に、Lua (lnslua.wasm) が必要かどうかを知るために使う。
*/
function getImportModules(bin) {
  const modules = new Set();
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
        modules.add(bin.toString('utf8', work, work + len));
        work += len;
        [len, work] = readUleb(bin, work);
        work += len;
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
      return modules;
    }
    pos += size;
  }
  return modules;
}

export async function run(wasmPath) {
  let exports = null;

  function readStr(ref) {
    const len = exports.lns_str_len(ref);
    const bytes = Buffer.alloc(len);
    for (let i = 0; i < len; i++) {
      bytes[i] = exports.lns_str_get(ref, i);
    }
    return bytes;
  }

  /** バイト列から wasm の文字列を作る */
  function newStr(bytes) {
    const ref = exports.lns_str_new(bytes.length);
    for (let i = 0; i < bytes.length; i++) {
      exports.lns_str_set(ref, i, bytes[i]);
    }
    return ref;
  }

  // io.open で開いたファイル (番号 → { fd, pos, append })
  const files = new Map();
  let nextFileNo = 0;
  let lastFileError = '';
  /** Lua の io.open のモード → Node.js の flags */
  const modeMap = {
    r: 'r', w: 'w', a: 'a', 'r+': 'r+', 'w+': 'w+', 'a+': 'a+',
  };
  /** Lua (C の strerror) と同じエラーのメッセージ */
  const errMessMap = {
    ENOENT: 'No such file or directory',
    EACCES: 'Permission denied',
    EISDIR: 'Is a directory',
    EEXIST: 'File exists',
  };
  /** file の pos から最大 size バイトを読む */
  function readBytes(file, size) {
    const buf = Buffer.alloc(size);
    const num = readSync(file.fd, buf, 0, size, file.pos);
    file.pos += num;
    return buf.subarray(0, num);
  }
  /** ファイルを読む (file_read の本体) */
  function fileRead(file, kind, num) {
    switch (kind) {
      case 0: { // "a": 残りの全て (終端でも "")
        const size = fstatSync(file.fd).size;
        return newStr(readBytes(file, Math.max(0, size - file.pos)));
      }
      case 1: // "l": 1 行 (改行を除く)
      case 2: { // "L": 改行付きの 1 行
        const chunks = [];
        for (;;) {
          const bytes = readBytes(file, 1);
          if (bytes.length === 0) {
            if (chunks.length === 0) {
              return null;
            }
            break;
          }
          if (bytes[0] === 0x0a) {
            if (kind === 2) {
              chunks.push(bytes);
            }
            break;
          }
          chunks.push(bytes);
        }
        return newStr(Buffer.concat(chunks));
      }
      default: { // バイト数
        const bytes = readBytes(file, Number(num));
        if (bytes.length === 0 && Number(num) > 0) {
          return null;
        }
        return newStr(bytes);
      }
    }
  }

  const fileFuncs = {
    file_open: (pathRef, modeRef) => {
      const path = readStr(pathRef).toString('utf8');
      const mode = readStr(modeRef).toString('latin1').replace('b', '');
      const flags = modeMap[mode];
      if (!flags) {
        lastFileError = `invalid mode '${mode}'`;
        return -1;
      }
      try {
        const fd = openSync(path, flags);
        const fileNo = nextFileNo++;
        files.set(fileNo, { fd, pos: 0, append: flags.startsWith('a') });
        return fileNo;
      } catch (err) {
        lastFileError = `${path}: ${errMessMap[err.code] ?? err.message}`;
        return -1;
      }
    },
    file_error: () => newStr(Buffer.from(lastFileError, 'utf8')),
    file_read: (fileNo, kind, num) => {
      try {
        return fileRead(files.get(fileNo), kind, num);
      } catch (err) {
        // 書き込み用に開いたファイルの読み込みなど。Lua 版と同じく nil にする
        return null;
      }
    },
    file_write: (fileNo, ref) => {
      const file = files.get(fileNo);
      const bytes = readStr(ref);
      if (file.append) {
        writeSync(file.fd, bytes, 0, bytes.length, null);
      } else {
        writeSync(file.fd, bytes, 0, bytes.length, file.pos);
        file.pos += bytes.length;
      }
    },
    file_seek: (fileNo, whence, offset) => {
      const file = files.get(fileNo);
      let base = 0;
      if (whence === 1) {
        base = file.pos;
      } else if (whence === 2) {
        base = fstatSync(file.fd).size;
      }
      const pos = base + Number(offset);
      if (pos < 0) {
        return -1n;
      }
      file.pos = pos;
      return BigInt(pos);
    },
    file_close: (fileNo) => {
      const file = files.get(fileNo);
      if (file) {
        closeSync(file.fd);
        files.delete(fileNo);
      }
    },
  };

  const imports = {
    lns: {
      ...fileFuncs,
      write: (ref) => write(readStr(ref)),
      write_err: (ref) => {
        flush();
        writeSync(2, readStr(ref));
      },
      error: (ref) => {
        throw new LnsError(readStr(ref).toString('latin1'));
      },
    },
  };

  const bin = readFileSync(wasmPath);
  if (getImportModules(bin).has('lua')) {
    imports.lua = await loadLua();
  }
  if (process.env.LNS_WASM_LIFTOFF !== '1') {
    // 変換した wasm をコンパイルする前に設定する必要がある
    setFlagsFromString('--no-liftoff');
  }
  const module = await WebAssembly.compile(bin);
  const instance = await WebAssembly.instantiate(module, imports);
  exports = instance.exports;
  try {
    exports.main();
  } catch (err) {
    flush();
    if (err instanceof LnsError) {
      writeSync(2, `lua: ${err.message}\n`);
      process.exitCode = 1;
      return;
    }
    throw err;
  } finally {
    flush();
  }
}

/** LuneScript の error() による例外 */
class LnsError extends Error {}

if (process.argv[1] && import.meta.url.endsWith(process.argv[1].split('/').pop())) {
  if (process.argv.length < 3) {
    console.error('usage: node lnswasm.mjs xxx.wasm');
    process.exit(1);
  }
  await run(process.argv[2]);
}
