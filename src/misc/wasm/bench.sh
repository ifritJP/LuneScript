#!/bin/bash
# go 版のネイティブ実行と wasm 実行の速度を比べるベンチマーク。
#
# usage: misc/wasm/bench.sh   (src ディレクトリで実行する。make bench-wasm でも良い)
#
# test/wasm_bench/b_*.lns を、go (-langGo + go build) と wasm (-langWasm) に変換して実行し、
# 実行時間 (3 回の最小値)、最大 RSS、GC の回数・停止時間を Markdown の表で出力する。
# 結果は docs/sw_design/wasm/wasm_benchmark.md に記録する。
#
# 環境変数:
#   LNSC       変換に使う lnsc (デフォルト: ./lnsc)
#   WORK       作業ディレクトリ (デフォルト: work/wasm-bench)
#   WASM_OPT   Binaryen の wasm-opt。指定すると wasm-opt -O3 した wasm も測る
#   BENCH      測るベンチマーク名 (空白区切り。デフォルト: 全て)

set -e
SRC=$(pwd)
LNSC=${LNSC:-$SRC/lnsc}
WORK=${WORK:-$SRC/work/wasm-bench}
WASM_OPT_FEATURES="--enable-gc --enable-reference-types --enable-multivalue
  --enable-bulk-memory --enable-nontrapping-float-to-int --enable-sign-ext
  --enable-mutable-globals"

if [ -z "$BENCH" ]; then
   BENCH=$(cd test/wasm_bench; ls b_*.lns | sed 's/\.lns$//')
fi

mkdir -p $WORK/src

# ======== ビルド ========
for name in $BENCH; do
   cp test/wasm_bench/$name.lns $WORK/src/
   (cd $WORK/src; $LNSC $name.lns save -langWasm > /dev/null 2> $name.conv.err)
   if [ -n "$WASM_OPT" ]; then
      $WASM_OPT $WASM_OPT_FEATURES -O3 $WORK/src/$name.wasm -o $WORK/src/$name.opt.wasm
   fi

   dir=$WORK/go/$name
   rm -rf $dir
   mkdir -p $dir
   rsync -a --include='*.lua' --include='*.lns' --exclude='*.*' lune $dir
   cp -a lune/base/runtime_go $dir/lune/base/
   cp -a lune/base/binding_dummy.go lune/base/DependLuaOnLns.go $dir/lune/base/
   # 生成したコードは github.com/ifritJP/LuneScript/src の runtime_go を import するので、
   # 同じモジュール名の go.mod を使う
   cp build.counter go.mod go.sum $dir
   cp test/wasm_bench/$name.lns $dir/
   (cd $dir;
    $LNSC $name mkmain main.go -u --package main;
    $LNSC $name.lns SAVE -langGo --main $name -u --package main --app testgo \
          > /dev/null 2> conv.err;
    PKG_CONFIG_PATH=$SRC go build -o bench.exe .)
done

# ======== 計測 ========

# 3 回実行し、最小の実行時間 [s] と最大の RSS [KB] を出力する。
# $1: 出力ファイル, 残り: コマンド
measure() {
   local out=$1
   shift
   local best=""
   local rss=""
   for _ in 1 2 3; do
      /usr/bin/time -f "%e %M" -o $WORK/time.txt "$@" > $out 2>&1
      read sec kb < $WORK/time.txt
      if [ -z "$best" ] || awk "BEGIN{ exit !($sec < $best) }"; then
         best=$sec
      fi
      if [ -z "$rss" ] || [ $kb -gt $rss ]; then
         rss=$kb
      fi
   done
   echo "$best $rss"
}
ratio() {
   awk "BEGIN{ if ($2 > 0) printf \"%.1f\", $1 / $2; else print \"-\" }"
}
mb() {
   awk "BEGIN{ printf \"%.0f\", $1 / 1024 }"
}

runWasm="node misc/wasm/lnswasm.mjs"

echo "### 実行時間"
echo
if [ -n "$WASM_OPT" ]; then
   echo "| ベンチマーク | go [s] | wasm [s] | wasm/go | wasm-opt [s] | wasm-opt/go | go RSS [MB] | wasm RSS [MB] | 出力 |"
   echo "|---|---|---|---|---|---|---|---|---|"
else
   echo "| ベンチマーク | go [s] | wasm [s] | wasm/go | go RSS [MB] | wasm RSS [MB] | 出力 |"
   echo "|---|---|---|---|---|---|---|"
fi
for name in $BENCH; do
   read goSec goKb <<< $(measure $WORK/out.go.$name $WORK/go/$name/bench.exe)
   read wSec wKb <<< $(measure $WORK/out.wasm.$name $runWasm $WORK/src/$name.wasm)
   same=一致
   if ! cmp -s $WORK/out.go.$name $WORK/out.wasm.$name; then
      same=不一致
   fi
   if [ -n "$WASM_OPT" ]; then
      read oSec oKb <<< $(measure $WORK/out.opt.$name $runWasm $WORK/src/$name.opt.wasm)
      if ! cmp -s $WORK/out.go.$name $WORK/out.opt.$name; then
         same=不一致
      fi
      echo "| $name | $goSec | $wSec | $(ratio $wSec $goSec) | $oSec | $(ratio $oSec $goSec) | $(mb $goKb) | $(mb $wKb) | $same |"
   else
      echo "| $name | $goSec | $wSec | $(ratio $wSec $goSec) | $(mb $goKb) | $(mb $wKb) | $same |"
   fi
done

echo
echo "### GC"
echo
echo "| ベンチマーク | go 実時間/CPU [s] | go 1 コア (GOMAXPROCS=1) [s] | wasm 実時間/CPU [s] | go の GC 回数 / 停止時間の合計 | Node の GC 回数 (young / full) / 停止時間の合計 |"
echo "|---|---|---|---|---|---|"
for name in $BENCH; do
   case $name in
      b_gc_*|b_coll) ;;
      *) continue ;;
   esac
   /usr/bin/time -f "%e %U %S" -o $WORK/time.txt $WORK/go/$name/bench.exe > /dev/null
   read goSec goUser goSys < $WORK/time.txt
   GOMAXPROCS=1 /usr/bin/time -f "%e" -o $WORK/time.txt $WORK/go/$name/bench.exe > /dev/null
   read go1Sec < $WORK/time.txt
   /usr/bin/time -f "%e %U %S" -o $WORK/time.txt $runWasm $WORK/src/$name.wasm > /dev/null
   read wSec wUser wSys < $WORK/time.txt

   # go: "gc N @...: A+B+C ms clock" の A と C が停止時間
   GODEBUG=gctrace=1 $WORK/go/$name/bench.exe 2> $WORK/gc.go.$name > /dev/null
   goGc=$(grep -c "^gc " $WORK/gc.go.$name)
   goStw=$(grep "^gc " $WORK/gc.go.$name |
              sed -E 's/.*: ([0-9.]+)\+[0-9.]+\+([0-9.]+) ms clock.*/\1 \2/' |
              awk '{ sum += $1 + $2 } END { printf "%.1f", sum }')
   # Node: "Scavenge ... MB, X / Y ms" の X が停止時間
   node --trace-gc misc/wasm/lnswasm.mjs $WORK/src/$name.wasm > $WORK/gc.node.$name 2>&1
   young=$(grep -c "Scavenge" $WORK/gc.node.$name || true)
   full=$(grep -c "Mark-Compact" $WORK/gc.node.$name || true)
   nodePause=$(grep -E "Scavenge|Mark-Compact" $WORK/gc.node.$name |
                  sed -E 's/.*MB, ([0-9.]+) \/ [0-9.]+ ms.*/\1/' |
                  awk '{ sum += $1 } END { printf "%.1f", sum }')
   goCpu=$(awk "BEGIN{ printf \"%.2f\", $goUser + $goSys }")
   wCpu=$(awk "BEGIN{ printf \"%.2f\", $wUser + $wSys }")
   echo "| $name | $goSec / $goCpu | $go1Sec | $wSec / $wCpu | $goGc / ${goStw}ms | $young / $full / ${nodePause}ms |"
done
