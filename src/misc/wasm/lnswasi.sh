#!/bin/bash
# LuneScript から変換した wasm (WasmGC) を、WASI のランタイム (wasmtime) で実行する。
#
# usage:
#   misc/wasm/lnswasi.sh merge xxx.wasm out.wasm   変換した wasm を、WASI の wasm にする
#   misc/wasm/lnswasi.sh run xxx.wasm [引数...]    WASI の wasm (xxx.wasi.wasm) にして wasmtime で実行する
#
# 変換した wasm は、C のランタイム (lnsrt.wasm、Lua を使う場合は lnslua.wasm) の関数と memory、
# Node.js でだけ使える関数 (lnsjs) を import する。
# Binaryen の wasm-merge で、これらを 1 つの wasm にする (lnsjs は何もしない関数 lnsjs_stub.wasm を使う)。
# import は WASI の関数 (wasi_snapshot_preview1) だけになる。
# また、wasmtime は古い形式の例外処理 (Lua の setjmp/longjmp が使う try) に対応していないので、
# 新しい形式 (exnref) に変換する。
#
# run では、ルート (/) を preopen し、カレントディレクトリをホストと同じにする (環境変数 LNS_WASM_CWD)。
# 環境変数 LUA_PATH・LNS_WASM_STACK は、そのまま渡す。
#
# 環境変数:
#   LNS_WASM_RT_DIR  C のランタイムと lnsjs_stub.wasm のディレクトリ (デフォルト: work/wasm-lua)
#   BINARYEN_BIN     Binaryen の bin ディレクトリ
#   WASMTIME         wasmtime のパス

set -e

SRC_DIR=$(cd "$(dirname "$0")/../.."; pwd)
RT_DIR=${LNS_WASM_RT_DIR:-$SRC_DIR/work/wasm-lua}
BINARYEN_BIN=${BINARYEN_BIN:-$HOME/opt/binaryen-version_133/bin}
WASMTIME=${WASMTIME:-$(command -v wasmtime || echo "$HOME/opt/wasmtime-v49.0.4-x86_64-linux/wasmtime")}
FEATURES="--enable-gc --enable-reference-types --enable-multivalue --enable-bulk-memory
  --enable-nontrapping-float-to-int --enable-sign-ext --enable-mutable-globals
  --enable-exception-handling"

# 変換した wasm $1 が使う C のランタイムのパス
rtPath() {
    # Lua の関数 (lnsl_xxx) を import している場合は、Lua を含む C のランタイムを使う
    if grep -q lnsl_ "$1"; then
        echo "$RT_DIR/lnslua.wasm"
    else
        echo "$RT_DIR/lnsrt.wasm"
    fi
}

# 変換した wasm $1 を、WASI の wasm $2 にする
merge() {
    local tmp="$2.tmp"
    "$BINARYEN_BIN/wasm-merge" $FEATURES "$(rtPath "$1")" lnsrt "$1" lnsmain	\
        "$RT_DIR/lnsjs_stub.wasm" lnsjs -o "$tmp"
    "$BINARYEN_BIN/wasm-opt" $FEATURES --translate-to-exnref "$tmp" -o "$2"
    rm -f "$tmp"
}

case "$1" in
    merge)
        merge "$2" "$3"
        ;;
    run)
        wasm=$2
        shift 2
        out="${wasm%.wasm}.wasi.wasm"
        rt=$(rtPath "$wasm")
        if [ ! "$out" -nt "$wasm" ] || [ ! "$out" -nt "$rt" ]; then
            merge "$wasm" "$out"
        fi
        envOpts=(--env "LNS_WASM_CWD=$PWD")
        for name in LUA_PATH LNS_WASM_STACK; do
            if [ -n "${!name+x}" ]; then
                envOpts+=(--env "$name=${!name}")
            fi
        done
        # 引数の先頭 (wasm のパス) は、Node.js (lnswasm.mjs) と同じく変換した wasm のパスにする
        exec "$WASMTIME" run -W gc=y,exceptions=y,function-references=y --dir / "${envOpts[@]}" \
             --argv0 "$wasm" "$out" "$@"
        ;;
    *)
        echo "usage: $0 merge xxx.wasm out.wasm | run xxx.wasm [args...]" >&2
        exit 1
        ;;
esac
