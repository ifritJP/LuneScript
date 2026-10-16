// LuneScript から変換した wasm (WasmGC) を Node.js で実行する。
//
// usage: node lnswasm.mjs xxx.wasm
//
// wasm が import する関数 (モジュール名 "lns") を提供する。
// 出力は Lua 5.3 の print と同じ形式にする。

import { readFileSync, writeSync } from 'node:fs';

const outBuf = [];
let outLen = 0;

function write(bytes) {
  outBuf.push(bytes);
  outLen += bytes.length;
  if (outLen > 64 * 1024) {
    flush();
  }
}
function writeText(txt) {
  write(Buffer.from(txt, 'latin1'));
}
function flush() {
  if (outBuf.length > 0) {
    writeSync(1, Buffer.concat(outBuf));
    outBuf.length = 0;
    outLen = 0;
  }
}

// double の正確な 10 進表現を、数字列と小数点の位置で返す。
// val = 0.d1d2d3... * 10^point (val > 0)
function exactDecimal(val) {
  const view = new DataView(new ArrayBuffer(8));
  view.setFloat64(0, val);
  const bits = view.getBigUint64(0);
  const expBits = Number((bits >> 52n) & 0x7ffn);
  let mant = bits & 0xfffffffffffffn;
  let exp;
  if (expBits === 0) {
    exp = -1074;
  } else {
    mant |= 1n << 52n;
    exp = expBits - 1075;
  }
  let digits;
  let point;
  if (exp >= 0) {
    digits = (mant << BigInt(exp)).toString();
    point = digits.length;
  } else {
    // mant / 2^k = mant * 5^k / 10^k
    digits = (mant * 5n ** BigInt(-exp)).toString();
    point = digits.length + exp;
  }
  digits = digits.replace(/0+$/, '');
  return { digits, point };
}

// 有効数字 prec 桁に丸める (最近接偶数丸め。glibc の printf と同じ)
function roundDigits(digits, point, prec) {
  if (digits.length <= prec) {
    return { digits, point };
  }
  let head = digits.slice(0, prec);
  const next = digits.charCodeAt(prec) - 48;
  const rest = digits.slice(prec + 1);
  let roundUp = false;
  if (next > 5) {
    roundUp = true;
  } else if (next === 5) {
    if (/[1-9]/.test(rest)) {
      roundUp = true;
    } else {
      roundUp = (head.charCodeAt(prec - 1) - 48) % 2 === 1;
    }
  }
  if (roundUp) {
    const num = (BigInt(head) + 1n).toString();
    if (num.length > head.length) {
      point += 1;
    }
    head = num.slice(0, prec);
  }
  return { digits: head.replace(/0+$/, '') || '0', point };
}

// C の "%.<prec>g"
function formatG(val, prec) {
  const sign = val < 0 || Object.is(val, -0) ? '-' : '';
  const abs = Math.abs(val);
  if (abs === 0) {
    return sign + '0';
  }
  const exact = exactDecimal(abs);
  const { digits, point } = roundDigits(exact.digits, exact.point, prec);
  const exp = point - 1;
  if (exp < -4 || exp >= prec) {
    let mant = digits[0];
    if (digits.length > 1) {
      mant += '.' + digits.slice(1);
    }
    const absExp = Math.abs(exp);
    return `${sign}${mant}e${exp < 0 ? '-' : '+'}${absExp < 10 ? '0' + absExp : absExp}`;
  }
  let txt;
  if (point <= 0) {
    txt = '0.' + '0'.repeat(-point) + digits;
  } else if (point >= digits.length) {
    txt = digits + '0'.repeat(point - digits.length);
  } else {
    txt = digits.slice(0, point) + '.' + digits.slice(point);
  }
  return sign + txt;
}

// Lua 5.3 の tostring( real ) と同じ形式 ("%.14g"、整数値なら ".0" を付ける)
export function formatReal(val) {
  if (Number.isNaN(val)) {
    // x86-64 の Linux では 0/0 などの NaN は符号ビットが立っており、"-nan" になる
    return '-nan';
  }
  if (val === Infinity) {
    return 'inf';
  }
  if (val === -Infinity) {
    return '-inf';
  }
  let txt = formatG(val, 14);
  if (!/[.e]/.test(txt)) {
    txt += '.0';
  }
  return txt;
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
      print_i64: (val) => writeText(val.toString()),
      print_f64: (val) => writeText(formatReal(val)),
      print_bool: (val) => writeText(val ? 'true' : 'false'),
      print_str: (ref) => write(readStr(ref)),
      print_nil: () => writeText('nil'),
      print_sep: () => writeText('\t'),
      print_end: () => writeText('\n'),
    },
  };

  const bin = readFileSync(wasmPath);
  const { instance } = await WebAssembly.instantiate(bin, imports);
  exports = instance.exports;
  try {
    exports.main();
  } finally {
    flush();
  }
}

if (process.argv[1] && import.meta.url.endsWith(process.argv[1].split('/').pop())) {
  if (process.argv.length < 3) {
    console.error('usage: node lnswasm.mjs xxx.wasm');
    process.exit(1);
  }
  await run(process.argv[2]);
}
