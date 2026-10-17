# wasm (WasmGC) 版のベンチマーク

go 版のネイティブ実行と、wasm (WasmGC) に変換して Node.js で実行した場合の速度・メモリ・GC の比較を記録する。
変換の改善 (`wasm_direct_design.md` の 6 章の最適化など) でどの程度伸びたかを確認するため、
計測のたびに日付・コミット・環境とともに結果を追記する (古い結果は消さない)。

## 1. 測り方

```sh
cd src
make bench-wasm                          # wasm-opt なし
make bench-wasm WASM_OPT=<wasm-opt のパス>  # wasm-opt -O3 した wasm も測る
```

`misc/wasm/bench.sh` が、`test/wasm_bench/b_*.lns` を次の 2 通りに変換して実行し、結果を Markdown の表で出力する。

| 種類 | 変換と実行 |
|---|---|
| go | `lnsc -langGo` で go に変換し、`go build` したネイティブの実行ファイル |
| wasm | `lnsc -langWasm` で wasm に変換し、`node misc/wasm/lnswasm.mjs` で実行 |
| wasm-opt | 上の wasm に Binaryen の `wasm-opt -O3` をかけたもの (最適化で縮まる余地の目安) |

- 実行時間は実時間 (wall clock) で、3 回の最小値。wasm の時間には Node.js の起動 (約 0.05 秒、`b_empty`) を含む。
- RSS は最大常駐メモリ。wasm の値には Node.js 自体の約 60MB を含む。
- 「出力」は、go 版と wasm 版の出力が一致したかどうか (wasm-opt を測った場合は、それも含む)。
- GC の表は、go の `GODEBUG=gctrace=1` と Node.js の `--trace-gc` の結果から集計する。
  go の停止時間は stop-the-world の時間の合計 (並行に動くマーク処理は含まない)。

ベンチマーク (`src/test/wasm_bench/`):

| 名前 | 内容 |
|---|---|
| b_int | 整数の四則演算・剰余・ビット演算・シフト (3 億回のループ) |
| b_real | 実数の演算 (3 億回のループ) |
| b_call | 関数の再帰呼び出し (`fib( 40 )`) |
| b_ctrl | `while` / `if` / `switch` / alge の `match` (500 万回の collatz) |
| b_class | 継承・オーバーライド・インタフェース経由の呼び出し・アクセッサ |
| b_closure | クロージャ・高階関数 |
| b_str | `format`・`sub`・連結・比較・`find` |
| b_coll | List・Map (int / str のキー)・Set |
| b_gc_short | 短命なインスタンスを 4,000 万個生成する |
| b_gc_list | 小さな List・Map を多数生成し、一部を保持する |
| b_gc_tree | 二分木の生成と破棄 (binary-trees。長命な木と短命な木が混在する) |
| b_empty | 何もしない (起動時間) |

## 2. 結果

### 2026-10-17 (M2 の前半、最適化なし)

| 項目 | 内容 |
|---|---|
| コミット | `1d2d8724` (support_direct_wasm) |
| 変換の状態 | M1 + Lua のランタイムの取り込み (`wasm_direct_design.md` の 14.4)。6 章の最適化は未実施 (generics・コレクションの要素は全て箱に入れる) |
| CPU | AMD Ryzen 9 9900X (12 コア / 24 スレッド) |
| メモリ | 60GB |
| OS | Linux 7.0.0-38-generic |
| go | go1.27.1 linux/amd64 |
| Node.js | v22.22.1 |
| Binaryen | version_133 |

#### 実行時間

| ベンチマーク | go [s] | wasm [s] | wasm/go | wasm-opt [s] | wasm-opt/go | go RSS [MB] | wasm RSS [MB] | 出力 |
|---|---|---|---|---|---|---|---|---|
| b_call | 0.30 | 0.29 | 1.0 | 0.29 | 1.0 | 5 | 62 | 一致 |
| b_class | 1.19 | 4.34 | 3.6 | 3.95 | 3.3 | 5 | 64 | 一致 |
| b_closure | 0.56 | 1.11 | 2.0 | 1.08 | 1.9 | 5 | 62 | 一致 |
| b_coll | 0.28 | 0.58 | 2.1 | 0.52 | 1.9 | 82 | 157 | 一致 |
| b_ctrl | 0.73 | 3.97 | 5.4 | 2.76 | 3.8 | 5 | 63 | 一致 |
| b_empty | 0.00 | 0.05 | - | 0.05 | - | 5 | 61 | 一致 |
| b_gc_list | 0.60 | 0.51 | 0.9 | 0.46 | 0.8 | 21 | 107 | 一致 |
| b_gc_short | 0.69 | 0.63 | 0.9 | 0.46 | 0.7 | 13 | 65 | 一致 |
| b_gc_tree | 1.25 | 0.56 | 0.4 | 0.55 | 0.4 | 103 | 316 | 一致 |
| b_int | 0.41 | 2.42 | 5.9 | 1.31 | 3.2 | 5 | 62 | 一致 |
| b_real | 0.86 | 0.94 | 1.1 | 0.95 | 1.1 | 5 | 63 | 一致 |
| b_str | 0.64 | 1.62 | 2.5 | 1.60 | 2.5 | 16 | 76 | 一致 |

#### GC

| ベンチマーク | go 実時間/CPU [s] | go 1 コア (GOMAXPROCS=1) [s] | wasm 実時間/CPU [s] | go の GC 回数 / 停止時間の合計 | Node の GC 回数 (young / full) / 停止時間の合計 |
|---|---|---|---|---|---|
| b_coll | 0.30 / 0.45 | 0.35 | 0.64 / 0.75 | 14 / 0.7ms | 37 / 1 / 85.4ms |
| b_gc_list | 0.54 / 1.39 | 0.60 | 0.50 / 0.52 | 183 / 6.0ms | 347 / 1 / 24.1ms |
| b_gc_short | 0.70 / 0.94 | 0.92 | 0.64 / 0.64 | 757 / 16.2ms | 1867 / 0 / 20.5ms |
| b_gc_tree | 1.37 / 3.80 | 2.28 | 0.55 / 1.27 | 130 / 4.9ms | 236 / 3 / 216.6ms |

#### wasm のサイズ [byte]

| ベンチマーク | wasm | wasm-opt |
|---|---|---|
| b_call | 1705 | 1225 |
| b_class | 2978 | 2307 |
| b_closure | 2335 | 1751 |
| b_coll | 4934 | 4146 |
| b_ctrl | 2324 | 1862 |
| b_empty | 1641 | 1176 |
| b_gc_list | 3671 | 2970 |
| b_gc_short | 2202 | 1699 |
| b_gc_tree | 2272 | 1738 |
| b_int | 1977 | 1479 |
| b_real | 2737 | 2118 |
| b_str | 3324 | 2465 |

この他に、Lua の関数を使う wasm は `lnslua.wasm` (約 190KB) を読み込む。

#### 考察

- 計算が主の処理 (b_int・b_ctrl・b_class) は go の 3.6〜5.9 倍かかる。実数の演算 (b_real) と関数の再帰呼び出し (b_call) は go とほぼ同じ。
- 遅い原因:
  - `%`・`/`・シフトを、Lua と同じ動作 (床除算、64 以上のシフト量など) にする関数の呼び出しで実装している。
    定数で割る場合などもそのまま呼び出している。`wasm-opt -O3` だけで b_int は 5.9 倍から 3.2 倍になるので、変換時の最適化で縮められる。
  - `for` の繰り返しのたびに、増分の符号を判定している。
  - インタフェース経由の呼び出しが、`$Obj` への cast → vtable → itable の配列 → cast → `call_ref` と多段になっている。
    List の要素は anyref なので、取り出すたびに cast する。
  - alge の `match` を `ref.test` の連鎖で分岐している (tag による分岐ではない)。
  - `format`・`find` などは Lua (`lnslua.wasm`) を呼び出しており、呼び出しのたびに文字列をコピーする。
- インスタンスを多く生成する処理 (b_gc_*) は、wasm の方が go と同等か速い。
  V8 の世代別 GC は短命なオブジェクトを安く回収できる。go の GC は世代別ではない。
  go は並行 GC に複数のコアを使っており、b_gc_tree では実時間 1.37 秒に対して CPU 時間が 3.80 秒になる (1 コアでは 2.28 秒)。
- GC による停止時間の合計は go の方が短い (b_gc_tree で go 4.9ms、Node 217ms。Node の full GC の 1 回の停止は最大 8ms 程度)。
- メモリ (RSS) は wasm の方が 3〜5 倍多い (Node.js 自体の約 60MB を含む)。理由は次の「メモリの内訳」を参照。

#### メモリの内訳

RSS が大きいのは、wasm の値の表現が大きいからではなく、V8 がヒープを大きく広げてから回収する方針で動くためである。
full GC の後に残る (生存している) データの量は、go とほぼ同じか少ない。

| ベンチマーク | wasm の RSS | Node.js 自体 | V8 のヒープの最大 | full GC 後に残るデータ | go の生存データ / ヒープの目標 |
|---|---|---|---|---|---|
| b_gc_tree | 316MB | 約 62MB | 252MB | 43〜49MB | 48MB / 97MB |
| b_coll | 157MB | 約 62MB | 82MB | 18MB | 38MB / 76MB |
| b_gc_list | 107MB | 約 62MB | 49MB | 14MB | 8MB / 16MB |

- Node.js 自体: `node -e 0` で 57MB。`lnslua.wasm` を読み込むと 62MB (`b_empty`)。
- V8 のヒープの値は `--trace-gc` の確保済みサイズ、go の値は `GODEBUG=gctrace=1` の GC 後のサイズとヒープの目標 (goal) の最大値。
- go は標準の GC 設定 (`GOGC=100`) で、ヒープを生存データの 2 倍までに抑える。その代わり GC の回数が多い (b_gc_tree で 130 回、CPU 時間 3.8 秒)。
- V8 は、GC にかかる時間の割合が小さい間は、full GC までヒープを大きく広げる。
  b_gc_tree では、プログラムが実行に使えた時間の割合 (mu) が 0.97 と高く、生存データ 49MB に対してヒープを 252MB まで広げてから回収していた。
- b_gc_tree では、一時的な木 (最大 52 万ノード) を作り終わるまでに若い世代の GC を何度か生き延びて古い世代へ移り、
  不要になった後も次の full GC まで古い世代に残る。
- b_coll と b_gc_list では、若い世代の領域 (V8 の初期値で最大 16MB の semi-space が 2 つ以上) と、
  List や Map を広げるときの大きな配列が効いている。

V8 のフラグでヒープの広げ方を変えた場合 (wasm-opt なしの wasm、3 回の最小値):

| フラグ | b_gc_tree (時間 / RSS) | b_coll | b_gc_list |
|---|---|---|---|
| なし | 0.53s / 297MB | 0.60s / 154MB | 0.50s / 106MB |
| `--heap-growing-percent=100` (ヒープの拡大を 2 倍まで) | 0.54s / 246MB | 0.60s / 154MB | 0.51s / 107MB |
| 上 + `--max-semi-space-size=4` (若い世代を小さく) | 0.72s / 212MB | 0.65s / 127MB | 0.50s / 80MB |
| `--optimize-for-size` | 2.15s / 138MB | 0.89s / 149MB | 0.50s / 75MB |
| (参考) go | 1.27s / 106MB | 0.28s / 82MB | 0.43s / 21MB |

メモリを減らすと時間が増える (b_gc_tree は `--optimize-for-size` で go より遅くなる)。
lnsc のように短時間で終わるプロセスでは、V8 の初期値のままでも問題になりにくい。

変換側で減らせる部分 (`wasm_direct_design.md` の 6 章の最適化に含まれる):

- 全てのオブジェクトが、先頭にハッシュ用の `i32` を持っている (Map のキーと `tostring` 用)。必要な型だけに持たせれば、1 個あたり 4 バイト減る。
- `List<int>` などの要素を 1 個ずつ箱に入れている。型ごとに専用の配列を使えば、要素ごとのオブジェクトが無くなる。
