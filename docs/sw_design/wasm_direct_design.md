# LuneScript から wasm への直接変換 (WasmGC 方式) の設計検討

LuneScript のコードを、Go を経由せずに wasm へ直接変換するための設計方針・課題・進め方をまとめる。
Go 経由の方式 (案 A) との比較は [wasm_support.md](wasm_support.md) を参照。

> 調査対象: ブランチ `dev_with_llm` (`f86d38c0`)。実験環境: Node.js v22.22.1 (V8 12.4)。
> wasm の各提案 (GC、例外処理、スレッドなど) の状況は執筆時点の認識であり、着手時に確認すること。

## 1. 方針 (決定事項)

| 項目 | 決定 |
|---|---|
| 変換方式 | `.lns` を wasm へ **直接** 変換する (新しいバックエンド `convWasm`、`-langWasm`) |
| メモリ管理 | **WasmGC** (GC は実行環境に任せる)。wasm のマルチスレッドはまだ仕様が固まっておらず、今使うのは非効率なため。マルチスレッドは仕様が固まってから検討する |
| 目標の順序 | (1) 簡単なテストコードが動く → (2) `make` で実行しているテストに合格する → (3) セルフホストできる |
| ランタイム | 初期は LuneScript 以外の言語で書いてよい。最終的には、ブートストラップとごく一部の例外を除いて LuneScript で書く (セルフホスト構成) |
| Lua 連携 | 序盤は不要。セルフホストには必須 (マクロの評価に Lua の VM を使うため。3.3 参照)。**C の Lua 実装を wasm に変換して取り込む** |

## 2. 目的と期待する効果

| 目的 | Go 経由 (案 A) の現状 | WasmGC 方式で期待すること |
|---|---|---|
| サイズ・ロード時間 | hello でも 7.5 MB (gzip 後 1.7 MB) | GC とアロケータを含めなくてよいので、ランタイムを小さくできる。WasmGC のサンプル (構造体と配列を使う 2 関数) は 134 バイト |
| ヒープの使用量 | 箱詰め・`[]LnsAny` の生成・全オブジェクトのヒープ割り当てが多い (6 章) | 型をそのまま wasm の struct / array にでき、`int` などを箱詰めせずに扱える。単相化・エスケープ解析で、さらに減らす |

## 3. 調査結果

### 3.1 実行環境での WasmGC の動作確認

構造体 (`struct.new` / `struct.get`) とバイト配列 (`array.new` / `array.len`) を使う wasm をバイナリで直接生成し、
Node.js v22 (V8 12.4) で実行して、正しい結果になることを確認した。

- この環境の wabt (`wat2wasm` 1.0.36) は GC の命令に対応していない。
  lnsc がバイナリを直接出力する方針なので問題ないが、デバッグ用に GC 対応のツール (Binaryen の `wasm-opt` / `wasm-dis` など) があると便利 (9 章)。

### 3.2 `make` で実行しているテスト

| テスト | 使われ方 | wasm 版での扱い |
|---|---|---|
| `test/tutorial/*.lns` (20 件) | `make test-tutorial` など | 最初の目標 (M1) |
| `test/testGo.lns` (+ `test/subGoPkg`) | `make test-go` (Go 変換のテスト)。Lua 版の実行結果 (`shebang`) と比べる | M2 の中心。Go と同じく Lua 以外への変換を前提にしたテストで、`/.nolua51./` のような印で対象外の部分を除ける |
| `test/start.lns` | `test-comp` の主なテスト (2,166 行) | M2 の後半。C のグルー (`module GlueTest require ... glue`)、`__luago`、`import.l` / `import.d` (遅延ロード)、`lune.base.Writer` の import などを含むので、Lua 連携 (M3) が必要な部分は除いて始める |
| `test/error/*` | コンパイルエラーのテスト | コンパイラ (フロントエンド) のテストなので、バックエンドには関係しない |

### 3.3 セルフホストに Lua 連携が必須である理由

マクロは、`Macro.lns` がマクロの本体を Lua のコードに変換し、`_load` (Lua の `load`) で実行して評価している
(`loadCode`、`evalFromCodeToLuaCode` / `evalToLuaCode`)。そのため、lnsc 自身を wasm で動かすには Lua の VM が必要になる。

## 4. 全体構成

```
                ┌─────────────── コンパイル時 (lnsc) ───────────────┐
.lns ─▶ 構文・意味解析 ─▶ AST (型付き) ─▶ convWasm ─▶ IR ─▶ 最適化 ─▶ wasm 出力 ─▶ app.wasm
                                                                        │
                ┌─────────────── 実行時 ─────────────────────────────┐   ▼
                │ app.wasm (WasmGC)  ◀─▶  ランタイム  ◀─▶  ホスト (Node.js / ブラウザ / WASI) │
                │                         └─(M3 以降) lua.wasm (C の Lua、線形メモリ)       │
                └────────────────────────────────────────────────────┘
```

- **IR**: 序盤から、型付きで、wasm の構造化制御 (block / loop / if) に近い木構造の IR を置く。
  最初は AST をほぼそのまま下ろすだけにし、最適化 (単相化、エスケープ解析、箱詰めの除去、インライン展開) は後から IR の上に追加する。
  後から IR を差し込むのは大きな手戻りになるため、最初から形だけでも用意する。
- **wasm 出力**: LuneScript でバイナリを直接出力する (LEB128、各セクション、GC の命令)。デバッグ用に、テキスト形式 (WAT 風) のダンプも出せるようにする。
- **プログラム全体の変換**: import している全モジュールを 1 つの wasm にまとめる。未使用コードの除去と単相化がプログラム全体で効くため。

## 5. 型の対応 (WasmGC)

| LuneScript | wasm での表現 | 備考 |
|---|---|---|
| `int` | `i64` | |
| `real` | `f64` | |
| `bool` | `i32` | |
| `str` | `(ref $Str)`: 不変のバイト配列 `(array i8)` を持つ struct | Lua と同じくバイト列。ハッシュ値をキャッシュする欄を持たせる |
| `nil` | `ref.null` | |
| nilable のスカラー (`int!` など) | 初期は箱詰め (`(ref null $BoxI64)`)。最適化で、局所変数では「値 + フラグ」の組にする | |
| class | struct。先頭にメソッド表 (vtable) への参照、続いて親クラスのフィールド、自クラスのフィールド | WasmGC の struct の部分型 (`sub`) で継承を表す。ダウンキャスト・型判定は `ref.cast` / `ref.test` / `br_on_cast` |
| interface | 初期は、オブジェクトのメソッド表から interface ごとの表を引く | 後で最適化する |
| `List<T>` / `Array<T>` | 長さと要素配列を持つ struct | 初期は要素を `anyref` で持つ (スカラーは箱詰め)。単相化で `(array i64)` などに特殊化する |
| `Map<K,V>` / `Set<T>` | ランタイムで実装するハッシュ表 | WasmGC には参照の同一性に基づくハッシュ値が無いので、オブジェクトにハッシュ値の欄を持たせる |
| 関数・クロージャ | 関数参照と捕捉した変数を持つ struct。呼び出しは `call_ref` | |
| ジェネリクス | 初期は型消去 (`anyref` + 箱詰め)。最適化で単相化する | ヒープ削減の主な対象 |
| alge | バリアントごとに部分型の struct。`match` は `br_on_cast` で分岐する | |
| 多値の戻り値 | wasm の複数戻り値 | |
| `error()` | 序盤はホストにメッセージを出して `unreachable` で停止する | LuneScript のコードには例外を捕まえる構文が無いので、序盤は wasm の例外処理を使わなくてよい。Lua 連携 (`pcall` など) で必要になったら検討する |

## 6. ヒープの使用を減らす最適化 (IR の上で行う)

Go 変換の出力 (`hello.lns` の例) でヒープを使っている箇所と、WasmGC 方式での対策。

| Go 変換の出力 | 対策 |
|---|---|
| `List` / `Map` の要素が `LnsAny` で、`int` も箱詰めされる | 単相化で要素型ごとの配列にする |
| 要素を取り出すたびに型アサーション | 単相化で不要になる |
| すべてのオブジェクトをヒープに置き、自分を指す `FP` フィールドを持つ | エスケープ解析で、外に出ないオブジェクトはフィールドを局所変数に分解する。継承されないクラスのメソッド呼び出しは直接呼び出しにする |
| `print` などの呼び出しのたびに `[]LnsAny` を作る (`Lns_2DDD`) | 引数の数が分かる呼び出しは展開する。多値は wasm の複数戻り値にする |
| nilable を `LnsAny` で表す | 局所変数では「値 + フラグ」の組にする |
| すべての関数に `_env` を渡す | 1 スレッドなのでグローバル変数にする |

## 7. ランタイム

### 7.1 必要な機能

| 分類 | 内容 |
|---|---|
| 文字列 | 連結、比較、`sub`、`byte`、`rep`、`lower` / `upper`、`format` (Lua の書式)、`find` / `gsub` / `gmatch` (**Lua のパターン**)、数値との変換 |
| コレクション | List / Array (`insert`、`remove`、`sort`、`unpack`)、Map、Set (`add`、`del`、`has`、`and`、`or`、`sub`、`clone`、`len`) |
| 入出力 | `print`、`io.stdout` / `stderr` / `stdin`、ファイル (`open`、`read`、`write`、`seek`、`close`)、`os` (`clock`、`time`、`date`、`exit`、`remove`、`rename`) |
| 非同期 | `__Runner`、`__Processor`、`__run` / `__join`。1 スレッドで順に実行する (Lua 版と同じ動き) |
| その他 | Mapping (`_toMap` / `_fromMap`) のための型情報、`math` |

### 7.2 何で書くか (段階的に移行する)

C や Rust は WasmGC の参照型を直接扱えないので、GC のオブジェクトを扱う部分には使えない。そこで次のように分ける。

| 段階 | GC のオブジェクトを扱う部分 (文字列、コレクションなど) | ホストとのやり取り (入出力など) |
|---|---|---|
| 初期 (M0〜M2) | **convWasm の中に、IR で書いた組み込み関数として持つ** (lnsc が wasm 出力と同時に生成する) | **JavaScript** (Node.js / ブラウザ用の薄い層) |
| 中期 (M3〜M4) | 同上 | WASI に対応する (ファイルアクセスなど、セルフホストに必要なもの) |
| 最終 (M5) | **LuneScript で書く**。メモリや wasm 命令を直接使うための最小限の組み込み関数 (intrinsic) を用意する | ホストとの境界の薄い層と、C の Lua だけを例外として残す |

Lua 互換の処理 (`string.format`、パターンなど) は、互換性を優先して C の Lua をそのまま使う (7.3)。

### 7.3 Lua のランタイムを使う処理と自前で実装する処理

Lua のランタイムの動作を前提にしている処理は、wasm にした C の Lua (`lnslua.wasm`) をそのまま使う。
ただし、Lua を使えない処理や、Lua を使う利点が無い処理は自前で実装する。

| 機能 | 実装 | 理由 |
|---|---|---|
| `string.format` (`"%s" (x)` 構文も含む) | **Lua** | 書式の細部 (`%g` / `%q` / 幅・精度・エラー) まで一致する。コンパイラ自身で 1,400 箇所近く使う構文なので、性能が問題になったら `%s` / `%d` だけの書式を自前で処理する (go 版と同じ二段構え) |
| 実数の文字列化 (`%.14g`) | **Lua** | 正確な丸めを自前で実装するのは難しい |
| `find` (パターン)・`gsub`・`gmatch`・`match` | **Lua** | lstrlib そのものを使えば移植による差が無い。`find` の `plain` 指定は自前の単純な検索で良い |
| `tonumber` | **Lua** | 16 進・指数・空白などの解釈規則が多く、実数は `strtod` の正確な丸めが必要 |
| `_load`・`require` (Lua のモジュール)・`Luaval`・`string.dump` | **Lua** (必須) | Lua そのもの |
| `error()` | 自前 | wasm の例外処理で実装する。Lua の `pcall` 経由では GC ヒープ側の状態を巻き戻せない |
| 等価・大小・ハッシュ (int と real の同一視など)、`//`・`%`・シフト、`@@@int` | 自前 | 言語の意味そのもので、全ての演算に関わる。Lua を呼ぶと遅すぎる |
| `print`・オブジェクトの `tostring` (`table: 0x...`) | 自前 | 実数の部分だけ Lua を呼ぶ |
| `sort` | 自前 | Lua 側から GC の List を扱えない。Lua の `table.sort` は不安定ソートで、同値の要素の順序は Lua でも決まっていない |
| 文字列の比較 | 自前 | C ロケールの `strcoll` (バイト比較) と同じ |
| `byte`・`rep`・`lower`・`upper`・`reverse`・`sub` | 自前 | 数行で書ける (Lua も C ロケールでは ASCII のみ)。呼び出しごとの文字列のコピーの方が高くつく |
| `math.floor` など | 自前 | wasm の命令で済む |
| `math.random` | 自前 | Lua でも libc 依存で、結果の一致は期待されていない |
| `io.*`・`os.*` | ホスト (Node.js、後に WASI) | Lua ではなく OS の機能。wasm の Lua の io ライブラリも結局 WASI が必要になる。`io.popen` はホストでのみ対応する |

Lua を使う場合の注意点:

- 呼び出しごとに、GC ヒープの文字列と Lua の線形メモリとの間でコピーが必要になる。
  `AsyncTokenizer.lns` のように同じ文字列に開始位置を変えて何度も `find` するコードでは、
  直前にコピーした文字列を再利用する仕組みが必要になる (未対応)。
- Lua を使うプログラムは `lnslua.wasm` (約 190KB) を読み込む。どの関数を使うかはコンパイル時にわかるので、
  Lua の関数を使う場合だけ import する。ただし、`print` で `stem` を出力するだけでも実数の文字列化のために Lua を import する。

## 8. Lua 連携の設計 (M3)

| 項目 | 方針 |
|---|---|
| Lua の取り込み | C の Lua 5.3 を clang (wasi-sdk) で wasm にコンパイルする。線形メモリを使う別モジュール `lnslua.wasm` になる (14.4 で実装した) |
| 結合 | ホストが 2 つのモジュールをインスタンス化して、互いの関数を import / export でつなぐ。配布を 1 ファイルにしたい場合は Binaryen の `wasm-merge` でまとめる |
| LuneScript → Lua の値 | Lua 側のレジストリの参照番号 (`luaL_ref`) を持つ struct で表す (`Luaval`) |
| Lua → LuneScript の値 | LuneScript 側のハンドル表 (GC オブジェクトの参照を入れる配列) の添字を、Lua の userdata に入れる |
| 文字列の受け渡し | Lua の線形メモリと WasmGC のバイト配列の間でコピーする。変換した wasm が Lua の memory を import し、`i32.load8_u` / `i32.store8` で直接コピーする (JavaScript を経由しない) |
| Lua のエラー処理 | Lua の C 実装は `setjmp` / `longjmp` を使う。clang の `-mllvm -wasm-enable-sjlj` と `-lsetjmp` で wasm の例外処理にコンパイルし、Node.js v22 で動くことを確認した |
| GC をまたぐ循環参照 | Lua の GC と WasmGC をまたいで循環した参照は回収されない。lnsc のように短時間で終わるプロセスでは問題になりにくい。長時間動くアプリでは注意が必要 |

## 9. 必要なツール

| ツール | 用途 | 状況 |
|---|---|---|
| Node.js (V8 11.9 以降) | WasmGC の実行・テスト | あり (v22.22.1)。動作確認済み |
| Binaryen (`wasm-opt`、`wasm-dis`、`wasm-merge`) | GC 対応の検証・逆アセンブル・最適化・モジュールの結合 | なし。デバッグのために早めに入れることを推奨 |
| wasi-sdk (clang + wasi-libc) | C の Lua を wasm にコンパイル | あり (wasi-sdk 34、`~/opt/wasi-sdk-34.0-x86_64-linux`。`make` の `WASI_SDK` で変更できる) |
| wasmtime (GC 対応版) | WASI での実行 (M4 以降) | なし。GC 対応の状況を確認する |

## 10. 進め方

| 段階 | 内容 | 完了の条件 |
|---|---|---|
| **M0. 基盤** | wasm バイナリの出力 (LEB128、セクション、GC の命令)、IR の骨組み、デバッグ用ダンプ、Node.js で動かすための JavaScript の層と実行スクリプト。`-langWasm` オプション | `int` / `real` / `bool`、関数、`if` / `while` / `for`、`print` を使う `fib` などが動き、Lua 版と同じ出力になる |
| **M1. 簡単なテスト** | クラス (継承・メソッド)、文字列、List / Map / Set、nilable、`foreach` / `forsort`、ジェネリクス (型消去)、クロージャ、interface | `test/tutorial/*.lns` (20 件) が Lua 版と同じ出力になる |
| **M2. make のテスト** | 前半で C の Lua を wasm にして取り込み、Lua 互換の処理を切り替える (7.3、M3 から前倒し)。alge / `match`、Mapping、多値、標準ライブラリ、複数モジュールの import、`__Runner` (1 スレッド) | `test/testGo.lns` (+ `subGoPkg`) と、`test/start.lns` のうち Lua 連携を使わない部分が、Lua 版と同じ出力になる。`make test-wasm` を追加する |
| **M3. Lua 連携** | `Luaval`、`__luago`、`_load` (C の Lua の wasm 化と結合は M2 の前半で行う) | `test/start.lns` の Lua 連携の部分も合格する |
| **M4. セルフホスト** | lnsc (`lune/base` の 42 モジュール) を wasm に変換し、wasm 版の lnsc で自分自身をコンパイルする。WASI によるファイルアクセス、ランタイムの Lua ファイルの埋め込み | wasm 版の lnsc で `test-comp` と同じ確認 (2 世代の出力が一致する不動点の確認と、テスト結果の一致) に合格する |
| **M5. ランタイムのセルフホスト化** | ランタイム (IR の組み込み関数、JavaScript) を LuneScript で書き直す | ブートストラップ、ホストとの境界、C の Lua 以外が LuneScript で書かれている |

最適化 (6 章) は M2 の後から並行して進め、各段階でサイズ・ヒープの割り当て量・速度を計測する (11 章)。

## 11. 検証方法

- **出力の比較**: 各テストを wasm で実行し、Lua 版の結果と diff する (`test-go` と同じ仕組み)。
  wasm で使えない部分は、`/.nolua51./` と同じように `/.nowasm./` のような印を付けて除外する。
- **wasm の妥当性**: 出力した wasm を `WebAssembly.validate` (と、あれば `wasm-opt --validate`) で検証する。
- **指標の計測**: Go 経由版を基準に、wasm のサイズ (gzip 後も)、ロード時間 (`WebAssembly.compile` / `instantiate`)、
  ヒープの使用量 (Node.js の `--trace-gc` やメモリ使用量)、実行速度を測る。
  go 版のネイティブ実行との速度・GC の比較は `make bench-wasm` で測り、結果を `wasm_benchmark.md` に記録する。
- **セルフホスト (M4)**: `test-comp` と同じく、wasm 版 lnsc で 2 世代の出力が一致すること、および Lua 版 / Go 版 lnsc と同じ出力になることを確認する。

## 12. リスク

| リスク | 内容 | 対策 |
|---|---|---|
| 作業量 | IR、wasm 出力、ランタイム、Lua 連携をすべて作る必要がある | 段階ごとに動くものを作り、テストで確認しながら進める |
| Lua との互換性 | 文字列パターン、`string.format`、数値の文字列化が Lua とずれる | テストの diff で検出する。必要なら `lstrlib.c` を使う |
| Lua の C 実装の wasm 化 | `setjmp` / `longjmp` に wasm の例外処理が必要 | M3 の最初に試作して確認する |
| 実行環境の差 | WasmGC や例外処理の対応状況が実行環境ごとに違う | 主な実行環境を Node.js に絞って始める。WASI ランタイムは M4 で確認する |
| デバッグの難しさ | GC 対応のツールが少ない | Binaryen を入れる。IR / wasm のダンプ機能を早めに作る |

## 13. 決めてほしいこと

1. 序盤の主な実行環境を Node.js にしてよいか (ブラウザ・WASI は後から)。
2. Binaryen (デバッグ・検証用) を開発環境に入れてよいか。M3 では wasi-sdk も必要になる。
3. 初期のランタイムの分担 (7.2): GC のオブジェクトを扱う部分は convWasm の組み込み関数、ホストとのやり取りは JavaScript、でよいか。
4. M0 から始めてよいか。

## 14. 実施状況

### 14.1 M0 (基盤) — 完了

| 項目 | 内容 |
|---|---|
| `src/lune/base/WasmBin.lns` | wasm バイナリの出力 (LEB128、f64 のビット列、セクション、GC の命令) と、命令の木構造の IR (`Ins`)、モジュール (`Module`)、関数 (`Func`) |
| `src/lune/base/convWasm.lns` | AST → IR の変換 (`-langWasm`)。int / real / bool / str、モジュールレベルの関数と変数 (wasm のグローバル変数)、`let`、代入、`if` / `while` / `repeat` / `for` / `return` / `break`、算術・比較・論理・ビット演算、文字列の連結・比較・長さ、`print`。対象外のノードはエラーにする |
| 組み込みのランタイム関数 | 整数の床除算・剰余 (Lua と同じ)、real の剰余、文字列の連結・比較・一致。IR で組み立てて、変換結果の wasm に含める |
| `src/misc/wasm/lnswasm.mjs` | Node.js で実行するための JavaScript の層。`print` を Lua 5.3 と同じ形式で出力する (real は `%.14g` を最近接偶数丸めで再現) |
| コンパイラへの組み込み | `Types.Lang.Wasm`、`-langWasm` オプション、`Converter.WasmConverter`、`front.lns`。`LNS_SRCS` と rockspec にモジュールを追加 |
| テスト | `src/test/wasm/m0_*.lns` (5 件)、`make test-wasm` (`mk/wasm.mk`)。Lua で実行した結果と wasm の実行結果を比べる |

結果:

- `make test-wasm` で 5 件とも Lua 版と同じ出力になった (Lua 版 lnsc と Go 版 lnsc の両方で変換して確認)。
- 変換後の wasm のサイズは 630〜1,226 バイト (Go 経由版の hello は 7.5 MB)。
- `make -j1 test-comp-with-go` (セルフホストの確認) に合格した。

### 14.2 M1 (簡単なテスト) — 完了

`test/tutorial/*.lns` (Lua 版でもコンパイルエラーになる `class1.lns` を除く 19 件) が Lua 版と同じ出力になった。
`make test-wasm` に追加した (M0 のテストと合わせて 24 件)。

構成:

| ファイル | 内容 |
|---|---|
| `src/lune/base/WasmBin.lns` | 型を 1 つの rec group にまとめて出力する (struct / array / func、部分型)。GC の命令、`call_ref`、`ref.func` 用の宣言セグメント、デバッグ用の `name` セクション (関数名) |
| `src/lune/base/WasmRt.lns` (新規) | ランタイムの型 (文字列、箱、List、Map/Set、オブジェクト、vtable、クロージャ、セル) と、IR で組み立てる関数 (文字列の連結・比較・部分文字列、数値の文字列化、Lua と同じ等価比較・大小比較・ハッシュ、ハッシュ表、List の操作、shell sort、print)。使われた関数だけを生成する |
| `src/lune/base/convWasm.lns` | 全面的に書き直した。値の表現の変換 (箱詰め・cast・関数型が異なるクロージャのアダプタ)、クラス (継承・vtable・コンストラクタ・アクセッサ)、インタフェース (itable)、クロージャ (セルによる変数の共有)、generics (型消去)、alge / match、enum、Mapping、多値・`...`、`let!` / `if!` / `when!` / `unwrap!`、List / Set / Map / str のメソッド |
| `src/misc/wasm/lnswasm.mjs` | ホストの関数 (出力、実数の文字列化、`string.format`、`error`、`find` / `gsub`)。実数の文字列化・`string.format`・`find` / `gsub` は 14.4 で Lua に切り替えた |
| `src/misc/wasm/luapattern.mjs` (新規) | Lua のパターンマッチ (lstrlib.c の移植)。14.4 で削除した |

値の表現:

| LuneScript | wasm |
|---|---|
| int / char / int の enum | `i64` |
| real | `f64` |
| bool | `i32` |
| str | `(ref null $Str)` (バイト配列) |
| stem、nilable のスカラー、generics の型パラメータ | `anyref` (スカラーは箱に入れる) |
| List / Array / `...` | `$List` (要素は anyref) |
| Map / Set | `$Map` (オープンアドレス法のハッシュ表。Set は値にキーを入れる) |
| クラス | `$Obj` の部分型 (先頭に vtable) |
| インタフェース | `$Obj` (itable は vtable が持つ配列からインタフェースの番号で引く) |
| 関数 | `$Clo` の部分型 (関数参照と、捕捉した変数のセル) |
| alge | 値ごとの部分型 (tag で分岐する。パラメータの無い値はシングルトン) |

M2 以降の課題 (未対応):

- import (複数モジュール)、Lua 連携 (`Luaval` など、M3)
- 例外処理 (`error()` はホストで停止する)
- `gmatch`、`string.byte` (複数)、`lower` / `upper` などの文字列の関数の多く
- generics の型パラメータを特殊化する最適化を入れる場合、インタフェースのメソッドと実装するクラスのメソッドとで、wasm 上の関数型 (箱に入れた `anyref` か、`i64` などか) が異なり得る。そのときは itable にアダプタ (箱詰め・取り出し) を登録する必要がある。現状は型パラメータを常に `anyref` で表現しており、LuneScript の型検査でもメソッドの型の一致が求められるため、このずれは起きない (`finalizeClasses()` で検出したらエラーにしている)
- 最適化 (6 章)。現状は generics・コレクションの要素を全て箱に入れている

### 14.3 M0 の作業中に見つかった既存の問題

すべて `dev_with_llm` で修正した。

| 問題 | 原因 | 修正 |
|---|---|---|
| go 版ランタイムの `string.format( "%c" )` が、128 以上の値を UTF-8 の 2 バイトで出力していた (Lua は 1 バイト) | go の `fmt` の `%c` は rune として出力する | `runtime_go/string.go` で 1 バイトを出力するようにした |
| go 版 lnsc で、2^53 を超える整数リテラルの値 (`LiteralIntNode.$num`) が不正確になる | `tonumber()` は real を返すため、go 版では float64 を経由する | `Util.parseIntLiteral()` を追加し、トークンの文字列から正確に求めるようにした (`TransUnitExp.lns`) |
| `while` の条件に `and` を使い、ループ内に `break` が無いと、ループの後の文が「到達しない」と誤判定される | 条件式の型を `== Ast.builtinTypeBool` (参照の一致) で判定しており、`a and b` の型が別オブジェクトのときに「bool でも nilable でもない式」= 無限ループと扱っていた | `equals()` で判定するようにした (`TransUnit.lns`) |
| 後で定義する関数を参照するコードはコンパイルできるが、Lua に変換したコードが実行時にエラーになる (`attempt to call a nil value`) | Lua の出力が `local function f` を定義順に並べるだけで、それより前では `f` が未定義だった | 宣言より前で参照されているモジュールレベルのローカル関数を、先頭で `local f` と事前宣言するようにした (`convLua.lns`)。コンパイラ自身のモジュールでも 12 モジュール・60 関数が該当した (呼び出し順の都合で、これまでは動いていた) |
| フォーマッタ (`tools/formatter`) が `x = -x;` (変数への単項マイナス) を解析できない | 文法 (`EbnfDef.lns`) の `<op1>` に `"-"` が無かった | `<op1>` に `"-"` を追加した。あわせて、文字リテラル `?-` が終端記号 `"-"` に一致してしまう問題も直した (`Ebnf.lns`: 文字リテラルは終端記号と照合しない) |
| 右シフト演算子 (`\|>>`) が Lua では論理シフト、go では算術シフトになる。シフト量が 64 以上や負の場合も異なる | convGo が go の `>>` / `<<` をそのまま出力していた | go はランタイムの `Lns_shiftL` / `Lns_shiftR` で Lua と同じ動作にした。wasm も同じ動作のランタイム関数にした |
| go への変換で、2 項の `~` (排他的論理和) が go の文法にない `a ~ b` のまま出力される | convGo が `~` を `^` に置き換えていなかった | `^` に置き換えるようにした (`convGo.lns`) |
| `make build` を `LNSDIR` を指定せずに実行すると `cd ;` (ホームディレクトリへの移動) になる | `LNSDIR` は `make install` 時に生成される `lune.mk` でしか設定されない | `mk/selfhost.mk` で `LNSDIR ?= .` を既定値にした |

確認したテスト:

- `make -j1 test-comp-with-go` (セルフホスト)。修正はコンパイラ自身の動作を変えるので、
  2 段階で行った (修正前のコンパイラで修正後のコンパイラを作って正式版にした後、修正を前提にしたコードとテストを追加して再度確認)。
- `test/testGo.lns` にシフトと排他的論理和のテストを追加 (go 版と Lua 版の結果が一致することを `test-go` で確認)。
- `test/wasm/m0_*.lns` にシフト、`and` を条件にした `while`、後で定義する関数の参照のテストを追加。`make test-wasm` で、Lua 版・go 版どちらの lnsc で変換しても合格。

### 14.4 M2 前半: Lua のランタイムの取り込み — 完了

7.3 の方針に従い、C の Lua 5.3.6 を wasm にして取り込み、JavaScript で実装していた Lua 互換の処理を切り替えた。
`make test-wasm` の 24 件が Lua 版・go 版の lnsc の両方で合格し、セルフホストのテスト (`make -j1 test-comp-with-go`) も合格した。

| ファイル | 内容 |
|---|---|
| `src/misc/wasm/lnslua.c` (新規) | Lua を呼び出すための関数を export する (スタック操作の API。builtin の関数を番号で push し、`lua_pcall` で呼び出す) |
| `src/mk/wasm.mk` | `make wasm-lua` で `work/wasm-lua/lnslua.wasm` を作る (lua.org から取得し、sha256 を確認してから wasi-sdk でコンパイルする)。`test-wasm` はこれに依存する |
| `src/lune/base/WasmBin.lns` | 関数を ID で指定し、エンコード時に関数インデックスに変換するようにした (関数を定義した後でも import できる)。memory の import、`i32.load8_u` / `i32.store8`。import する関数の型を rec group の外に出力する |
| `src/lune/base/WasmRt.lns` | Lua の関数を使う時だけ import する。値の受け渡し (`luaPushAny` / `luaGetAny`)、builtin の呼び出し (`luaCall`、Lua のエラーはそのメッセージでエラーにする)。実数の文字列化・`format`・`find`・`gsub` を Lua で実装した |
| `src/misc/wasm/lnswasm.mjs` | 変換した wasm が Lua を import している場合だけ `lnslua.wasm` を読み込んで渡す。JavaScript の `format` / 実数の文字列化 / パターンマッチを削除した |

作業中にわかったこと:

- WasmGC では、rec group の中の関数型は、別のモジュールの同じ構造の関数型と一致しない (型は rec group 単位で比べる)。
  そのため、他の wasm から import する関数の型は rec group の外に定義する必要がある。
- `lnslua.wasm` は `-Os` で約 190KB。WASI の関数を import するが、io・os ライブラリを開いていないのでほぼ使わない。
- 多値を返す `find` を引数の途中に書いた場合 (`print( s.find( a ), s.find( b ) )`) に、全ての値を展開していた問題を直した
  (`ExpMultiTo1` で `...` の先頭の要素を取り出す)。

Lua 版との既知の差:

- Lua の関数でのエラーのメッセージが少し異なる。Lua 版は `[string "..."]:75: bad argument #2 to 'format' ...` だが、
  wasm 版は位置が無く、関数名が `'string.format'` になる (C から呼んでいるので呼び出し元の名前が無い)。

### 14.5 M2 後半: `test/testGo.lns` — 完了 (Lua 連携の部分を除く)

`test/testGo.lns` (+ `test/subGoPkg`) が、Lua 連携 (`_load` / `loadfile` / `io.popen`) を使う 15 個の section を除いて、
Lua 版と同じ出力 (1,067 行) になった。
`make test-wasm` に追加した (`test-wasm-testgo`)。除く section は `_section(); // nowasm: ...` と印を付けておき、
テストの際に行番号を保ったまま空行にする (Lua 版の期待値も同じファイルから作る)。

対応した機能:

| 機能 | wasm での実現方法 |
|---|---|
| `$.` / `$[` (nil なら nil になるアクセス) | prefix を局所変数に入れ、null なら nil、そうでなければ通常のアクセス。式文の `obj$.func()` は null でなければ呼び出す |
| `!` (CondRet) | 文の前に値を局所変数に入れ、nil (`__Ret` は `Err`) なら関数から戻る。nilable と `__Ret`、エラー情報を伴う 2 値の 3 種類 |
| タプル (`(= a, b)`、`tuple...`) | Lua 版と同じく値のリスト (`$List`) で表し、`...` として展開する |
| `__serr` / `__Er` | `__Er` を実装し、メッセージ (`txt`) を持つクラスを生成する (Lua 版のランタイムの `LnsErr` に相当) |
| `apply` | `str.gmatch()` は Lua で全ての一致を求めてから処理する。それ以外は Lua の汎用の for と同じく、関数・状態・制御変数で繰り返す |
| `continue` | ループの更新 (カウンタの加算など) を本体の先頭で行っているので、`loop` の先頭への分岐にする |
| `_toMap` / `_fromMap` / `_fromStem` | オブジェクトが struct なので、Mapping のクラスとそれを含む List / Map を再帰的に変換する。`_fromMap` は Lua 版の `_toInt` / `_toList` などと同じ検査をして、エラーの位置 (`list.1.val.1` など) を返す |
| `advertise` | メンバのメソッドを呼び出す関数を生成し、vtable に登録する |
| generics のインタフェースを具体的な型で実装 | itable に、型パラメータ (anyref) と実装の型 (i64 など) を変換するアダプタを登録する (14.2 の課題) |
| `__run` / `__join` / `__lns.Sync` | Lua 版と同じく、`__run` はその場で `run()` を呼び出す (Skip なら呼ばない)。`createFlag` / `_createPipe` は nil |
| 文字列の関数 | `byte` / `rep` / `lower` / `upper` / `reverse` / `replace` と、`string.find` などのモジュールの関数は Lua の関数を呼び出す。`string._join` は wasm で実装 |
| import (複数のモジュール) | import したモジュール (間接的なものを含む) も 1 つの wasm にまとめる。各モジュールのトップレベルの処理は初期化の関数にし、import 文の位置で呼び出す (Lua の require と同じく 1 回だけ)。front が依存の順に AST を集めて convWasm に渡すので、`--uptodate forceAll` で AST を作る必要がある (`.meta` から読み込んだモジュールは AST が無い) |
| `io.open` とファイルの `read` / `write` / `seek` / `close` | ホスト (`lnswasm.mjs`) の関数 (`file_open` など) で、Node.js の fs を使う。ファイルの値は `$Obj` の部分型でホストのファイル番号を持つ。ホストで文字列を作るため、`lns_str_new` / `lns_str_set` を export する |
| その他 | `alias`、Map の `#`、`for` の実数の増分、`__func__`、`type()`、`tonumber()`、stem の値の `[]`、`__asyncLock`、enum / alge の `get__txt()`、関数を持つメンバの呼び出し、static メンバの getter、インタフェースの getter、`__lns.runMode` などの builtin の定数 |

作業中に見つかった既存の問題:

| 問題 | 修正 |
|---|---|
| マクロの中の `__line__` の `LiteralIntNode` の値 (`$num`) が、マクロを定義した行になっていた (トークンの文字列は展開した行) | `TransUnitExp.lns` で、値も展開した行にした |

Lua 版との既知の差:

- generics のクラスの `_fromMap` で、型パラメータのメンバは型を検査しない (型を消去しているため)。
- ファイルの `read` が失敗した場合 (書き込み用に開いたファイルの読み込みなど)、Lua 版は `nil, メッセージ, 番号` を返すが、wasm 版は `nil` だけを返す。
- 標準入力からの `read` は未対応 (nil を返す)。

残りの作業:

- M2: `test/start.lns` のうち、Lua 連携を使わない部分 (3.2)
- M3: `_load` / `loadfile` / `io.popen` などの Lua 連携 (`Luaval`、`__luago`)。testGo の `// nowasm` の 15 個の section
