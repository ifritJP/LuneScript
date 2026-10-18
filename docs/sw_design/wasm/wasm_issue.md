# wasm 版の課題

wasm (WasmGC) への直接変換 ([wasm_direct_design.md](wasm_direct_design.md)) の、未解決の課題・未修正の不具合・Lua 版との既知の差をまとめる。
対応した場合は、この文書から外し、`wasm_direct_design.md` の 14 章に実施内容を記録する。

2026-10-20 時点 (M4 とランタイムの高速化 (14.9)・WASI 対応 (14.10)・Lua 版 lnsc の埋め込み (14.11) の後)。

## 1. 一覧

| 番号 | 課題 | 影響 | 状況 |
|---|---|---|---|
| 2.1 | Lua の値の参照が解放されない | Lua の値を多く扱う長時間のプログラムで、メモリが増え続ける | WasmGC の仕様の制限。当面は受け入れる |
| 2.2 | C のグルーと C で書いた Lua のモジュールを読み込めない | `module ... glue` を使うプログラムが動かない | 未着手 (`skip-wasm`) |
| 2.3 | WASI のランタイムで `io.popen` が使えない | WASI では外部コマンドの出力を読めない | WASI の制限 |
| 2.4 | wasmtime での実行が遅い | wasmtime では Node.js の約 9 倍の時間がかかる | wasmtime の実装による。変換側では対処しない |
| 2.5 | 計算が主の処理が go 版より遅い | `b_class`・`b_closure`・`b_int` などが go 版の 2〜2.4 倍 | 6 章の最適化が未実施 |
| 2.6 | Lua のパターン関数を呼ぶたびに文字列をコピーする | 同じ長い文字列に何度もパターンを使う処理が遅い | 未着手 |
| 3.1 | `[ ... ]` に nil が含まれる場合の要素数が Lua 版と違う | `#[ ... ]` などの結果が違う | 未修正 |
| 3.2 | Go 版の lnsc の並列の解析で、クロージャの情報が揃わない | `-langWasm` では並列の解析をしない (変換が遅くなる) | 回避済み (フロントエンドの問題として残る) |
| 3.3 | `make bench-wasm LNSC=...` が失敗する | ベンチマークを make から別の lnsc で実行できない | 未修正 (`bench.sh` を直接実行すれば良い) |

4 章は、仕様として受け入れている Lua 版との差。

## 2. 未解決の課題

### 2.1 Lua の値の参照が解放されない

Lua の値 (テーブル・関数など) を LuneScript 側に取り出すと、Lua の registry に登録 (`luaL_ref`) し、
その参照番号を持つ `$LuaVal` を作る。`$LuaVal` が GC で回収された時に登録を解除する必要があるが、
今はプログラムの実行中に解除されない。

- 取り出した Lua の値 (`Luaval`) の数に比例して Lua の registry が増え、その値から辿れる Lua のデータとともに、プロセスの終了まで残る。
- wasm 側の `$LuaVal` は通常どおり GC で回収される。Lua の文字列関数 (`format` など) の呼び出しや、Lua の値を使わないプログラムでは増えない。
- lnsc のように短時間で終わるプロセスでは問題にならない。

| 実行環境 | 状況 |
|---|---|
| Node.js | `FinalizationRegistry` に登録して `lnsl_unref` を呼ぶ (14.7)。ただし、コールバックは JavaScript のイベントループに戻った時にだけ呼ばれ、プログラムは `wasi.start()` の中で最後まで同期的に動くので、実行中は 1 つも解放されない |
| WASI のランタイム | 解放する仕組みが無い (`lnsjs_stub.wat` の `luaval_new` は何もしない) |

原因は、WASI やランタイムではなく、WebAssembly (WasmGC) の仕様にある。

- WasmGC には、オブジェクトが回収されたことを知る仕組み (ファイナライザ・弱参照) が無い (MVP では意図的に外され、将来の拡張として議論されている段階)。
  どのランタイムでも、wasm の中だけでは回収を検知できない。
- WASI は OS の機能 (ファイル・時刻・引数など) のインタフェースで、GC の動作は範囲外。
- Node.js の `FinalizationRegistry` は JavaScript の機能で、JavaScript が wasm のオブジェクトを登録できるので使える。

対応の候補 (いずれも保留):

| 候補 | 内容 | 問題 |
|---|---|---|
| `FinalizationRegistry.prototype.cleanupSome()` | 回収済みのオブジェクトのコールバックを、その場で同期的に呼ぶ。試作では、登録 1 万回ごとに呼ぶと、10 万個の `$LuaVal` のうち 90,834 個が `main` の実行中に解放された | Node.js v22 では `--harmony-weak-refs-with-cleanup-some` が必要。実験的な機能で、標準化の提案が止まっているので、将来消える可能性がある。Node.js だけ |
| JSPI (JavaScript Promise Integration) | wasm から async の関数を呼ぶと、wasm の実行を一時停止してイベントループに戻り、Promise の解決後に再開する。定期的にイベントループに戻せば、コールバックが呼ばれる | Node.js v22 では `--experimental-wasm-jspi` が必要 (標準化が進んでいて、新しい V8 では既定で有効)。`main` の呼び出し方の変更と、停止・再開の負担が必要 (未試作)。Node.js だけ |
| wasmtime の `externref` のホストのデータ | wasmtime を Rust の API で組み込む場合、ホストのデータを持たせた `externref` は、GC で回収された時にホストのデータの後始末 (`Drop`) が実行される。`$LuaVal` にこれを持たせれば `lnsl_unref` を呼べる見込み | 専用の実行プログラムが必要 (`wasmtime run` では使えない)。wasmtime 専用 |
| Lua の値の生存を自前で管理する | ランタイムに依存しない唯一の方法 | `$LuaVal` を参照している箇所を全て把握する必要があり、難しい (14.7 で検討) |

### 2.2 C のグルーと C で書いた Lua のモジュール

C のグルー (`module ... glue`) と、C で書いた Lua のモジュールは、wasm にした Lua では読み込めない。
`test/start.lns` の `GlueTest` (宣言と呼び出し) には `skip-wasm` を付けて、wasm のテストから除いている。

### 2.3 WASI のランタイムで `io.popen` が使えない

WASI にはプロセスを起動する機能が無いので、WASI のランタイムでは `io.popen` が失敗する (nil を返す)。
Node.js では、`lnsjs` の `popen` (`execSync`) で使える。
`test/testGo.lns` の該当する section に `// nowasi` を付け、`test-wasm-wasi` では除いている。
lnsc の中で `io.popen` を使っていた `Depend.getFileLastModifiedTime` は、ホストを使わない処理に置き換え済み (14.9)。

### 2.4 wasmtime での実行が遅い

wasm 版の lnsc で lnsc の全モジュールを Lua に変換する処理は、Node.js では 1.84 秒、wasmtime 49.0.4 では 17.3 秒かかる (14.10)。

| wasmtime の GC (`-C collector=`) | 時間 |
|---|---|
| copying (標準) | 17.3 秒 |
| drc | 27.6 秒 |
| null (回収しない) | 12.1 秒 (メモリ 1.7GB) |

GC をしなくても Node.js の約 6 倍かかるので、wasmtime (Cranelift) の WasmGC のコード生成と GC の実装が、V8 ほど最適化されていないためと考えられる。
変換側では対処しない。新しい wasmtime で改善されるかを確認する。

### 2.5 計算が主の処理が go 版より遅い

`wasm_direct_design.md` の 6 章の最適化 (generics・コレクションの要素の特殊化など) は未実施。
ベンチマーク (`wasm_benchmark.md` の 2026-10-20 (2)) では、`b_class`・`b_closure` が go 版の 2.4 倍、`b_int` が 2.0 倍、`b_ctrl` が 1.5 倍かかる。
主な要因 (`wasm_benchmark.md` の「b_class の残りの要因」):

- インタフェース経由の呼び出しが、vtable → itable の配列 (anyref) → `ref.cast` → `call_ref` と多段になっている。
- メソッドの中で、フィールドに触るたびに `self` を `ref.cast` し直している。
- `%`・`/`・シフトを、Lua と同じ動作にする関数の呼び出しで実装している (定数で割る場合などもそのまま呼んでいる)。
- List の要素は anyref で、取り出すたびに cast する。`List<int>` などの要素を 1 個ずつ箱に入れている。

### 2.6 Lua のパターン関数を呼ぶたびに文字列をコピーする

`gsub`・`gmatch`・`match`、特殊文字を含むパターンの `find` は、Lua (lstrlib) で処理する。
呼び出しごとに、対象の文字列を GC ヒープから Lua の線形メモリにコピーする。
`AsyncTokenizer.lns` のように、同じ長い文字列に開始位置を変えて何度もパターンを使う処理では、
直前にコピーした文字列を再利用する仕組みがあると良い。
14.9 の高速化で、lnsc の処理全体に占める Lua の呼び出しは 2〜3% になったので、優先度は低い。

## 3. 見つかった不具合 (未修正)

### 3.1 `[ ... ]` に nil が含まれる場合の要素数

可変長引数から List を作る `[ ... ]` で、値に nil が含まれる場合の要素数が Lua 版と違う。
Lua 版は `{...}` の `#` になるので、`f( nil )` の `#[ ... ]` が 0 になるが、wasm 版は 1 になる (14.9 の確認中に見つけた)。

### 3.2 Go 版の lnsc の並列の解析で、クロージャの情報が揃わない

Go 版の lnsc で関数の本体を並列に解析する (`threadPerUnitThread`) と、クロージャが捕捉する変数の情報 (`Scope.closureSymList`) が揃わない。
Lua・Go への変換はクロージャの捕捉を言語に任せるので影響しないが、wasm への変換はこの情報を使う。
`-langWasm` では並列に解析しないことで回避している (14.8)。フロントエンドの問題として残っている。

### 3.3 `make bench-wasm LNSC=...` が失敗する

環境変数 `LNSC` を付けて make を実行すると、make は `mk/config.mk` の `LNSC` (`LUA_PATH=./?.lua lua5.3 ...`) を
環境変数として `bench.sh` に渡すので、`bench.sh` が `LUA_PATH=./?.lua` をコマンドとして実行して失敗する (終了コード 127)。
`LNSC=<lnsc のパス> misc/wasm/bench.sh` を直接実行すれば良い (`wasm_benchmark.md`)。

## 4. Lua 版との既知の差 (仕様として受け入れているもの)

| 差 | 理由 | 記録 |
|---|---|---|
| `__free` (ファイナライザ) が呼ばれない | WasmGC にはファイナライザが無い (Lua 5.1 と同じ。テストでは `skip-wasm`) | 14.6 |
| `return 1##` / `return ##` のように省略した戻り値の数 | Lua は省略した戻り値の分だけ値の数が減るが、wasm は関数の戻り値の数が固定 | 14.6 |
| generics のクラスの `_fromMap` で、型パラメータのメンバの型を検査しない | 型パラメータを消去しているため | 14.5 |
| List を Map に cast した場合のキーの順番 | キーは添字の文字列なので、要素が 10 個以上だと整列の順番が Lua 版と異なる | 14.6 |
| ファイルの `read` が失敗した場合の戻り値 | Lua 版は `nil, メッセージ, 番号` を返すが、wasm 版は `nil` だけを返す | 14.5 |
| Lua の関数でのエラーのメッセージ | Lua 版は `[string "..."]:75: bad argument #2 to 'format' ...` だが、wasm 版は位置が無く、関数名が `'string.format'` になる (C から呼ぶので呼び出し元の名前が無い) | 14.4 |
| 同じ Lua のテーブルでも、受け取るたびに別の `$LuaVal` になる | `==` では等しくならない | 14.7 |
| LuneScript の値を Lua に渡すとコピーになる | Lua 側の変更は LuneScript の値に反映されない (Go 版と同じ)。クラスのオブジェクトと関数は Lua に渡せない (文字列にして渡す) | 14.7 |
