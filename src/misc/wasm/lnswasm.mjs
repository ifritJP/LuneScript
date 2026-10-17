// LuneScript から変換した wasm (WasmGC) を Node.js で実行する。
//
// usage: node lnswasm.mjs xxx.wasm
//
// wasm が import する関数 (モジュール名 "lns") を提供する。
// wasm が Lua の関数 (モジュール名 "lua") を import している場合は、
// wasm にした Lua (lnslua.wasm、lnslua.c から生成) を読み込んで渡す。
// lnslua.wasm のパスは環境変数 LNS_WASM_LUA で指定できる。
// 出力は Lua 5.3 の print と同じ形式にする。

import { readFileSync, writeSync } from 'node:fs';
import { fileURLToPath } from 'node:url';

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

  const imports = {
    lns: {
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
  const module = await WebAssembly.compile(bin);
  if (WebAssembly.Module.imports(module).some((item) => item.module === 'lua')) {
    imports.lua = await loadLua();
  }
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
