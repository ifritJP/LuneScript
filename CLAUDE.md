# CLAUDE.md

LuneScript は、静的型付け・nil 安全の言語を Lua (5.1〜5.4)・Go・Python・wasm (WasmGC) に変換するトランスコンパイラ。
コンパイラ `lnsc` 自身も LuneScript で書かれていて、セルフホストしている (`src/lune/base/*.lns`、約 74,000 行)。

この文書は、AI のエージェントがこのリポジトリで作業するための手順と注意点をまとめる。
言語の詳しい仕様は、リファレンス (下の「参考資料」) を参照。

## 作業の方針

- 回答・ドキュメント・コードのコメントは日本語で書く (コミットのメッセージは英語)。
- コミット・プッシュ・ブランチのマージは、指示があった時だけ行う。
- 既存の不具合の修正は、作業中のブランチではなく `dev_with_llm` で行うよう指示されることがある。ブランチの使い分けは指示に従う。
- 言語の意味は Lua 版を正とする。Go 版・wasm 版などの他の変換先は、Lua 版と同じ結果になるようにする。

## リポジトリの構成

| パス | 内容 |
|---|---|
| `src/` | ビルドとテストは全てここで行う (`cd src`) |
| `src/lune/base/*.lns` | コンパイラのソース。`.lua`・`.meta`・`.go` は生成物だが git で管理している (セルフホストのため) |
| `src/lune/base/runtime_go/` | Go 版のランタイム (Go で書いたもの) |
| `src/lune/base/base.lua` | Lua 版の lnsc の起動用 (`lua5.3 lune/base/base.lua src.lns 〜`) |
| `src/lnsc` | Go 版の lnsc (git では管理していない。`make go-build` で作る) |
| `src/test/` | テスト。`start.lns` (主な正常系)、`testGo.lns` (Go 版との比較)、`tutorial/`、`wasm/`、`wasm_bench/` |
| `src/mk/*.mk` | Makefile の分割 (`selfhost.mk`・`test-comp.mk`・`test-stages.mk`・`go.mk`・`wasm.mk` など) |
| `src/misc/` | 補助のツール (`misc/wasm/`: wasm のランタイムと実行環境) |
| `docs/sw_design/` | 設計の文書 (`wasm/` は wasm 対応、`improvement_plan.md` は改善項目) |

## ビルドとテスト

全て `src` ディレクトリで実行する。

| 目的 | コマンド | 時間の目安 |
|---|---|---|
| Go 版の lnsc を作る (最初に 1 回) | `make go-build` | |
| `.lns` を変えた後、Lua 版の lnsc (`.lua`・`.meta`) を作り直す | `make build` | 約 40 秒 |
| セルフホストの確認 (最終確認) | `make -j test-comp-with-go NO_ACCEPT=y` | 約 40 秒 |
| Go への変換のテスト | `make test-go` | 約 10 秒 |
| wasm への変換のテスト (Node.js) | `make test-wasm` | 約 10 秒 |
| wasm 版 lnsc のセルフホストの確認 | `make test-wasm-selfhost` | 約 50 秒 |
| 同上を WASI のランタイム (wasmtime) で | `make test-wasm-wasi`・`make test-wasm-selfhost-wasi` | 約 10 秒・約 90 秒 |

- `test-comp-with-go` は、現在の lnsc で 2 世代セルフホストして出力が一致すること (不動点) と、
  `test/` の正常系・異常系のテストの結果が変わらないことを確認する。合格すると、`work/3` の成果物
  (`lnsc`・`.lua`・`.meta`・`.go`) で現在のものを置き換える (`accept-package`)。
  - `NO_ACCEPT=y` を付けると、テストだけ行い置き換えない。
    この場合、生成した `.go` は `work/3/lune/base/*.go` にあるので、必要なものを `lune/base/` に写す。
  - エラーの調査は `-j` を外すと追いやすい。
  - テストの結果の差分が正しい場合 (エラーのメッセージを変えた場合など) は、作者が `make accept-package` で受け入れる。AI は勝手に受け入れない。
  - Lua 5.1 が入っている環境では Lua 5.1 のテストも行う。除くには `LUA51=` を指定する。
- ビルドのたびに `src/build.counter` (と `Option.go` のビルド番号) が増える。AI の作業では、コミットに含めないよう元に戻す。
- 変換器 (`conv〜.lns`) を変えた場合は、`make build` → 対象の変換先のテスト → 最後に `test-comp-with-go`、の順で確認する。
- Lua 版の lnsc は Go 版より 10 倍以上遅い (全モジュールの変換で約 12 秒と約 0.4 秒)。
  試行錯誤には Go 版 (`./lnsc` や `work/3/lnsc`) を使うと速い。
- `make` のコマンドラインで `LNSC=〜` を指定するのは良いが、環境変数として渡すと
  `mk/config.mk` の `LNSC` (`LUA_PATH=./?.lua lua5.3 ...`) がそのまま子のプロセスに渡り、壊れることがある。
- `LUA_PATH` は、Lua のモジュールだけでなく、lnsc が import する `.lns` を探すのにも使われる。
  別のディレクトリで lnsc を動かす場合は `LUA_PATH="./?.lua;../?.lua"` のように、変換するディレクトリを先に探すようにする。

単体のファイルの実行と変換:

```sh
cd src
LUA_PATH=./?.lua lua5.3 lune/base/base.lua path/to/file.lns exe        # Lua 版で実行
LUA_PATH=./?.lua lua5.3 lune/base/base.lua path/to/file.lns save       # file.lua を出力
./lnsc path/to/file.lns exe                                            # Go 版の lnsc で実行
./lnsc path/to/file.lns save -langGo                                   # Go に変換
./lnsc path/to/file.lns save -langWasm && node misc/wasm/lnswasm.mjs path/to/file.wasm  # wasm
```

テストのファイルの書き方:

- 期待する出力は、Lua 版の lnsc で実行した結果 (`exe`)。他の変換先のテストは、Lua 版の出力と diff する。
- 変換先で未対応の行は、行番号を保ったまま空行にする印で除く (`// skip-wasm`・`// skip-cross` など。`mk/wasm.mk` を参照)。

## LuneScript の書き方 (要点)

Lua と C の構文が基調で、型の書き方は TypeScript・Swift に近い。

```lns
// コメントは // と /* */ (Lua の // (切り捨て除算) は無い。int 同士の / が切り捨て (床除算) になる)
let val = 1;                 // immutable (再代入できない)。型は推論
let mut count:int = 0;       // mutable
let name:str! = nil;         // nilable (! が付いた型だけ nil を持てる)
let list:List<int> = [ 1, 2 ];           // 添字は 1 から
let map:Map<str,int> = { "a": 1 };       // map.a / map[ "a" ] の値は nilable (int!)
let set:Set<int> = (@ 1, 2 );
let tuple = (= 1, "a" );
let flag = false;
if val ~= 1 and not flag { }             // 不等号は ~= (!= は無い)。論理演算は and / or / not
foreach item, index in list { }          // 値が先、添字・キーが後
forsort item, key in map { }             // キーの順に列挙 (foreach の Map の順番は不定)
for index = 1, 10 { }                    // continue は無い。break は一番内側のループを抜ける
switch val { case 1, 2 { } default { } } // break は不要 (break はループを抜ける)
while true { break; }                    // 無限ループは即値の true で書く (フロー解析のため)

fn func( a:int, b:str! ): int, str {     // 多値を返せる
   return a, unwrap b default "";
}
let num, txt = func( 1 ## );             // nilable の引数の省略は ## で明示する
print( "%d %s" ( 1, "x" ) );             // 文字列リテラルの直後の ( ) は string.format
let first = (txt.gsub( "a", "b" ));      // 多値の先頭だけ使う時は ( ) で囲む
print( func( 1, "y" )** );               // 多値を可変長・省略可能な引数に渡す時は ** を付ける (無いと警告)
```

nilable の扱い:

| 書き方 | 意味 |
|---|---|
| `unwrap x` / `unwrap x default v` | nil なら実行時エラー / v にする |
| `if! x { _exp }` / `if! let a = x { }` | nil でなければ実行 (`_exp`・`a` は非 nilable) |
| `when! a, b { }` | 変数が全て nil でなければ実行 (変数だけ。メンバや式は不可) |
| `let! a = x { /* nil の時: 代入か脱出 */ };` | nil でない初期値の変数を宣言 |
| `unwrap! a = x { };` | 既存の変数に代入 |
| `obj$.member`・`list$[ 1 ]`・`func$()`・`obj$.method$()` | nil 条件演算子 (結果は nilable。メソッドの呼び出しは `$()` まで付ける) |
| `x!` | エラーの伝搬 (戻り値が nilable / `__Ret` の関数の中で、nil / `.Err` なら return) |

クラス・型:

```lns
interface IF1 { pub fn func(); }             // 使う前に宣言する (相互参照は proto で前方宣言)
abstract class Base { pub abstract fn name(): str; }
class Point extend Base (IF1) {              // 単一継承 + インタフェース (pub class なら Base・IF1 も pub が必要)
   pri let x:int {pub};                      // getter (obj.$x で呼ぶ)
   pri let mut y:int {pub,pub};              // getter と setter (set_y)
   pub fn __init( x:int, y:int ) { super(); self.x = x; self.y = y; }
   pub fn move( dx:int ) mut { self.y = self.y + dx; }   // メンバを変えるメソッドは mut
   pub override fn name(): str { return "point"; }      // 親クラスのメソッドの上書きは override
   pub fn func() { }                                    // インタフェースの実装には override を付けない
   pub static fn create(): Point { return new Point( 0, 0 ); }
}
let mut pt = Point.create();
pt.move( 1 );
alge Shape { Circle( r:real ), Rect( w:real, h:real ) }
let shape = Shape.Circle( 1.0 );
match shape { case .Circle( r ) { } case .Rect( w, h ) { } }   // 全ての値の網羅を強制するには _match (_switch も同様)
enum Color { Red, Green }                    // Color.Red / .Red、$_txt で名前 ("Color.Green")
let val:stem = 1;
let n = val@@int;                            // 検査しないキャスト。@@@ は検査して nil、@@= はメソッドの無いクラス
```

- メンバの宣言で初期値は書けない。コンストラクタを書かない場合は、全メンバを引数に取るデフォルトコンストラクタができる。
- `&T` は immutable な型。`let` (mut なし) で型を推論した変数は immutable になる。
- generics は `fn func<T>( val:T ): List<T>`、`class Box<T:Super(IF)>`。インタフェースは generics にできない。
- マクロは `macro _name( arg:int ) { { /* コンパイル時の処理 */ } /* 展開する文 */ }`。名前は `_` で始める。
  `,,var` (展開)、`,,,var` (文字列 → シンボル)、`,,,,var` (シンボル → 文字列)、`` `{ } `` (文の値)。マクロは Lua で実行される。
- `_lune_control default_async_all;` のモジュールでは、関数が `__async` になる。`__async` の関数から、外側の mutable な変数や `__noasync` の関数は使えない。
- Lua を直接使うブロック: `__luago { }`・`__luaLock { }`・`__luaDepend { }`。Lua の値は `Luaval<T>` 型。
- 組み込みの関数・メソッドの引数と戻り値は `src/lune/base/Builtin.lns` (データとして定義している) を見るのが確実。
- 組み込みの関数は意図的に最小限にしている (増やすと Go などの変換先ごとにランタイムの対応が必要になるため)。
  それ以外の Lua の関数は、`module Math require 'math' { pub static fn floor( x:real ):int; }` のように `module` 宣言で使うのがポリシー。
  組み込みの関数を安易に増やさない。

## よくある間違い

| 症状・エラー | 原因 | 直し方 |
|---|---|---|
| `shadowing symbol of ...` / `This symbol is shadowed` | 外側のスコープと同じ名前の変数 (別のスコープでも不可) | 名前を変える |
| `argument number is mismatch` (書式の引数で) | `gsub` などの多値の関数をそのまま渡した | `( )` で囲んで先頭の値だけにする |
| `can't access the mutable type's symbol(...) from async` | `default_async_all` のモジュールで、局所関数から外側の mutable な変数を使った | モジュールの関数にして引数で渡す |
| `can't access noasync function in async` | `__async` の関数から `__noasync` の関数を呼んだ | 呼び出し側も `__noasync` にするか、構成を変える |
| `It can't call prototype function from '__init'` | コンストラクタから、プロトタイプ宣言だけのメソッドを呼んだ | 初回の使用時に生成するなど、コンストラクタの外で呼ぶ |
| `not found -- ')'. actually '.'` | 文字列リテラルに直接メソッドを付けた (`"".byte()`) | 変数に入れてから呼ぶ |
| `not found field typeInfo: str.len` | `str` に `len()` メソッドは無い | `#txt` を使う |
| `not found type -- table` / `pcall` | 組み込みでない Lua の API | `module` 宣言で宣言して使う (`table.concat` の代わりは `string._join`。組み込みの一覧は `Builtin.lns`) |
| `argument(1) type must be int` (`"%d"` の書式で) | 書式と引数の型をコンパイル時に検査している | 型を合わせる (書式が変数の場合は検査しない) |
| `use '##', instate of 'int!'` (警告) | nilable の引数を省略した | `func( 1 ## )` のように `##` を付ける (`-Werror` ではエラー) |
| `use '**' at arg(2)` (警告) | 多値を返す関数の結果を、可変長・省略可能な引数に渡した | `func()**` のように `**` を付ける |
| `type mismatch ... use the option '--legacy-mutable-control'` | 引数・代入の型の不一致 (末尾にオプションの案内が付くが、多くは単なる型の誤り。`&` と `mut` の違いも含む) | 型を見直す |
| `'return' could not use here` | モジュールのトップレベルで `return` した | 関数の中に入れる |
| `... can't be external symbol` / `This type must be public` | `pub` のクラス・関数が、`pub` でない型を使っている | 使っている型も `pub` にする |
| 大きな文字列の生成が遅い | `txt = txt .. item` の繰り返しは、要素数 × 全体の長さに比例する | List に集めて `string._join` (`default_strict_generics` のモジュールでは `string.__join`) でまとめて連結する |

## Lua の意味に依存する点

変換先の言語の常識と違う点が多い。新しい変換先を作る場合は、Lua 版と同じ結果になるようにする。

- `%`・整数の `/` は床除算 (負の数で C 系と違う)。シフト `|<<`・`|>>` は論理シフト (シフト量が 64 以上・負の場合も Lua と同じ)。
- 実数の文字列化は `%.14g` で、整数値の実数は `1.0` になる。
- 文字列はバイト列。`txt[ i ]` はバイトの値。`lower`・`upper` は ASCII だけ。比較はバイト単位。
- Map のキーは int と real の同じ値を区別しない。int・real・str 以外のキーはオブジェクトの同一性で比較する。
- 多値・可変長引数 (`...`) は Lua と同じ。`[ ... ]` に nil を含む場合の要素数も Lua の `{...}` の `#` に従う。
- `table.sort` (`List.sort`) は安定ソートではない。出力の順番に影響する場合は、比較の条件で順番を決める。
- `Str.startsWith( txt, ptn )` は Lua 版の実装では「`ptn` を含むかどうか」(`find` で判定) になっている (Go 版は `HasPrefix`)。

## コンパイラの構成

処理の流れ: `front.lns` (全体の制御) → `Tokenizer.lns` (字句解析) → `TransUnit*.lns` (構文・意味解析。AST と型情報を作る) → `Converter.lns` と `conv〜.lns` (コードの出力)。

| モジュール | 役割 |
|---|---|
| `Option.lns` | コマンドラインのオプション (`-langGo`・`-langWasm` など) |
| `front.lns` | 全体の制御、モジュールの読み込み、`exe`・`save` などのモード |
| `TransUnit.lns`・`TransUnitExp.lns`・`TransUnitDecl.lns` など | 構文解析・型検査・フロー解析・マクロの展開。言語機能の変更で一番多く修正する |
| `Nodes.lns` | AST のノードの定義 (`Nodes.Filter` を継承して AST をたどる) |
| `Ast.lns` | 型情報 (`TypeInfo`)・スコープ (`Scope`)・シンボル (`SymbolInfo`) |
| `Builtin.lns` | 組み込みの型・関数の定義 |
| `Macro.lns`・`DependLuaOnLns.lns` | マクロの実行 (Lua の VM で動かす) |
| `Converter.lns` | 変換器の呼び出し (Lua・Go・Python・wasm) |
| `convLua.lns`・`convGo.lns`・`convPython.lns`・`convWasm.lns` | 各変換先の出力。新しい変換先は `conv〜.lns` を追加する |
| `LuaMod.lns` | Lua に出力するランタイムのコード (`_lune`) |
| `Str.lns`・`Depend.lns`・`DependLuaOnLns.lns` | `_lune_control limit_conv_code lua;` のモジュール。Lua 以外の変換先では、その言語で書いたランタイムに置き換える (Go 版は `runtime_go`) |

- 新しい変換先を作る手順と注意点は `docs/sw_design/wasm/wasm_summary.md` の 4 章にまとめている
  (Lua 版を正解にした差分テスト → 小さいテストから順に広げる → セルフホストの不動点)。
- wasm の変換器の設計と課題は `docs/sw_design/wasm/` (`wasm_direct_design.md`・`wasm_issue.md`) を参照。

## 参考資料

- LuneScript のリファレンス (日本語): https://ifritjp.github.io/documents/lunescript/
  (原稿の org ファイルは、作者の環境の `/proj/work/ifritJP.github.io/hugo/content/LuneScript/all.ja.org`)。
  言語の仕様、各機能の例 (`@lnsFront: ok` / `error` の印付き)、開発方法 (「LuneScript の開発方法」の章) がある。
- 書き方の例: `src/lune/base/*.lns` (コンパイラ自身)、`src/test/tutorial/*.lns`、`src/test/start.lns`・`src/test/testGo.lns`。
- 生成した Lua のコード (`.lua`) を見ると、構文の意味が分かりやすい。
