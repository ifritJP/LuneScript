# LuneScript でプログラムを開発する AI エージェント向けガイド

LuneScript を使って何かを開発する AI のエージェント (Claude Code など) 向けに、
コマンドの使い方・言語の要点・間違えやすい点をまとめる。

LuneScript 自体 (コンパイラ) を開発する場合は、リポジトリのルートの `CLAUDE.md` を参照。

## この文書の使い方

LuneScript を使うプロジェクトで、エージェントがこの文書を読むようにする。

- プロジェクトのルートの `CLAUDE.md` にこの文書の内容をコピーする。
- または、`CLAUDE.md` に `@path/to/lunescript_for_agents.md` と書いて読み込ませる (Claude Code の import 機能)。

## LuneScript とは

- 静的型付け・型推論・nil 安全の言語。Lua と C の構文が基調で、型の書き方は TypeScript・Swift に近い。
- コンパイラ `lnsc` で、Lua (5.1〜5.4) か Go に変換して動かす (wasm への変換も実験的に対応)。
- 言語の意味は Lua に従う (整数の除算、文字列がバイト列であることなど。下の「Lua の意味に依存する点」)。
- 詳しい仕様は、リファレンス (日本語) https://ifritjp.github.io/documents/lunescript/ を参照。

## プロジェクトの準備

- プロジェクトのルートに `lune.js` を置く。内容は最低限 `{}` で良い。
  このディレクトリがモジュールの起点になる (`foo/bar/mod1.lns` のモジュール名は `foo.bar.mod1`)。
- `lnsc` は、`lune.js` のあるディレクトリをカレントディレクトリにして実行する。
- 毎回指定するオプションは `lune.js` に書ける: `{ "cmd_option": [ "--valid-luaval" ] }`。
- Go にも変換する場合は、`go.mod` も同じディレクトリに置く (下の「Go への変換」)。

```
proj/
 +-- lune.js
 +-- main.lns
 +-- mylib/
      +-- Calc.lns      // import mylib.Calc; で使う (Calc.add() のように最後の名前でアクセス)
```

## lnsc の使い方

| 目的 | コマンド |
|---|---|
| 実行する (トップレベルの処理だけ。`__main` は呼ばない) | `lnsc main.lns exe` |
| `__main` まで実行する (引数を渡す) | `lnsc -shebang main.lns arg1 arg2` |
| テスト (`__test` ブロック) を実行する | `lnsc mylib/Calc.lns exe --testing` |
| Lua のコードを標準出力に出す | `lnsc main.lns lua` |
| Lua に変換して保存する (`main.lua`) | `lnsc main.lns save` |
| Lua とメタ情報 (`.meta`。import を速くする) を保存する | `lnsc main.lns SAVE` |
| 全ファイルを一括で変換する (`__main` を呼ぶコードも出力) | `find . -name '*.lns' \| lnsc @- save --main main` |
| 警告をエラーにする | オプションに `-Werror` を付ける |

- `save` は指定したファイルだけを変換する。import しているモジュールも変換するには、全ファイルを指定する (`@-` で標準入力から渡すのが速い)。
- 変換した Lua のコードは LuneScript なしで動く (`lua5.3 main.lua x y`)。ランタイムは各ファイルに展開される
  (`-r` を付けると `require` する形になるが、そのモジュールが必要になる)。
- コマンドのエラー (`file.lns:行:列: error: ...`) の位置を見て直す。`warning` も、ほとんどは直すべきもの。

`__main` (コマンドの引数を受け取る入口):

```lns
// @lnsFront: skip
pub fn __main( argList:&List<str> ): int {   // argList[ 1 ] はスクリプトのパス
   foreach arg in argList {
      print( arg );
   }
   return 0;                                  // 終了コード
}
```

## 開発の進め方 (エージェント向け)

1. 小さく書いて、`lnsc xxx.lns exe` で頻繁にコンパイル・実行する。コンパイラのエラーは具体的なので、そのまま直せる。
2. 警告も直す。特に `##` (引数の省略) と `**` (多値の受け渡し) の警告は、意図しない動作の原因になる。
3. テストは `__test` ブロックに書き、`--testing` で実行する。
   **NG があっても `lnsc` の終了コードは 0** なので、出力の `test total: ... (OK:n, NG:m)` の NG の数で判定する。

```lns
// @lnsFront: skip
pub fn add( a:int, b:int ): int {
   return a + b;
}
__test case1( ctrl ) {
   __t._eq( ctrl, add( 1, 2 ), 3 );     // __t は組み込みのテスト用モジュール (_eq・_isNil・_isNotNil など)
}
```

4. 型で表せることは型で表す (`stem` とキャストはできるだけ使わない。generics・alge・nilable を使う)。

## 言語の要点

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
| `obj$.member`・`list$[ 1 ]`・`func$()`・`obj$.method$()` | nil 条件演算子 (結果は nilable。メソッドの呼び出しは `$()` まで付ける) |
| `x!` | エラーの伝搬 (戻り値が nilable / `__Ret` の関数の中で、nil / `.Err` なら return) |

エラー処理: 例外 (大域脱出) は無い。nilable の戻り値 (`str!, str!` でエラーのメッセージも返す) か、
`__Ret<T,E>` (`.Ok( val )` / `.Err( err )`) で返し、`!` で伝搬する。`error( "msg" )` はプログラムを止める。

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
- 公開するものには `pub` を付ける (関数・クラス・変数はモジュールのトップレベルで宣言する)。
- `&T` は immutable な型。`let` (mut なし) で型を推論した変数は immutable になる。
- generics は `fn func<T>( val:T ): List<T>`、`class Box<T:Super(IF)>`。インタフェースは generics にできない。
- クラスと JSON の相互変換: `class Data extend (Mapping) { }` で `_toMap()` と `Data._fromMap( map )` ができる。
- マクロは `macro _name( arg:int ) { { /* コンパイル時の処理 */ } /* 展開する文 */ }`。名前は `_` で始める。

## 組み込みの関数

組み込みの関数は、意図的に最小限にしている (Lua の標準ライブラリの一部だけ)。
組み込みの関数を増やすと、Go などの変換先ごとにランタイムで同じ動作を実装する必要があるためである。
**それ以外の Lua の関数が必要な場合は、`module` 宣言で宣言して使う** (下の例) のがポリシー。
組み込みでない関数を直接書くとコンパイルエラーになる。

| 種類 | 使えるもの |
|---|---|
| 全体 | `print`・`type`・`tostring`・`tonumber( txt ## )` (戻り値は `real!`)・`error`・`require`・`loadfile`・`_load`・`collectgarbage` |
| 文字列 (`txt.xxx()` / `string.xxx( txt )`) | `find`・`byte`・`format`・`rep`・`gmatch`・`gsub`・`sub`・`lower`・`upper`・`reverse`、`txt.replace( src, dst )`、長さは `#txt`。List の連結は `string._join( list, sep ## )` (Lua の `table.concat`。`_lune_control default_strict_generics;` のモジュールでは `string.__join`) |
| List | `insert`・`remove`・`sort`・`unpack`、長さは `#list` |
| Set | `add`・`del`・`has`・`and`・`or`・`sub`・`clone`・`len` |
| Map | メソッドは無い。`map[ key ] = nil` で削除、`foreach` / `forsort` で列挙。`#` は使えない |
| `io` | `io.open( path, mode )` (ファイルの `read`・`write`・`seek`・`close`)、`io.stdin`・`io.stdout`・`io.stderr`、`io.popen` |
| `os` | `clock`・`time`・`date`・`difftime`・`exit`・`remove`・`rename` |
| `math` | `random`・`randomseed` だけ (それ以外は下の `module` 宣言で使う。実数から整数への切り捨てだけなら `val@@int`) |

正確な引数と戻り値は、LuneScript のソースの `src/lune/base/Builtin.lns` に定義がある。

組み込みでない Lua の関数は、`module` 宣言で型を付けて使う (Lua 版・Go 版のどちらの lnsc でも動く):

```lns
// @lnsFront: ok
module Math require 'math' {
   pub static let pi:real;
   pub static fn floor( x:real ):int;
   pub static fn sin( x:real ):real;
}
print( Math.floor( -3.5 ), Math.sin( Math.pi / 2 ) );   // -4  1.0
```

- `require` した Lua のモジュールと、宣言した型の整合性は検査されない (合っていない場合の動作は未定義)。宣言は Lua のマニュアルに合わせて正確に書く。
- 呼び出しは Lua の VM を経由するので、Go に変換した場合は組み込みの関数より遅い。頻繁に呼ぶ処理では注意する。

## Lua との連携

| やりたいこと | 方法 |
|---|---|
| Lua のモジュールを型付きで使う | `module Name require 'luamod' { pub static fn func1( val:int ):int; pub fn method( val:int ):int; }` (`.` で呼ぶ関数は `static fn`、`:` で呼ぶメソッドは `fn`) |
| Lua のコードを実行する | `__luago { let loaded = unwrap _load( code ## ); ... }` |
| Lua の値 | `Luaval<T>` 型 (`--valid-luaval` を指定した場合。Go に変換するコードで Lua と連携する場合に使う)。Lua のテーブルを LuneScript の値にするには `expandLuavalMap( val )` |
| LuneScript のモジュールを Lua から使う | 変換した `.lua` を `require` する (モジュールのテーブルが返る。別の値を返すには `provide sym;`) |

## Go への変換

```sh
lnsc main.lns SAVE -langGo --main main --app myapp        # main.go
lnsc mylib/Calc.lns SAVE -langGo --app myapp              # import するモジュールも変換する
lnsc main mkmain main_entry.go --gomod <LuneScript の src のパス> --app myapp   # 起動用の main 関数と go.mod
go mod tidy
go build -o app .
```

- 変換したコードは LuneScript の Go 版のランタイム (`github.com/ifritJP/LuneScript/src/lune/base/runtime_go`) を使う。
  `--gomod` に使っている lnsc と同じ版の LuneScript のソースを指定すると、同じ版のランタイムでビルドできる。
- 通常のビルドには cgo と C の Lua 5.3 が必要。`go build -tags gopherlua` なら不要 (Lua の部分が Lua 5.1 相当になる)。
- **`__test` ブロックを含むモジュールを Go に変換すると、テスト用のモジュールを import するコードになり、ビルドできない。**
  Go にするモジュールにはテストを書かないか、テストを別のモジュールに分ける。
- 詳しくは LuneScript のリポジトリの `docs/go_build.md` を参照。

## よくある間違い

| 症状・エラー | 原因 | 直し方 |
|---|---|---|
| `shadowing symbol of ...` | 外側のスコープと同じ名前の変数 (別のスコープでも不可) | 名前を変える |
| `argument number is mismatch` (書式の引数で) | `gsub` などの多値の関数をそのまま渡した | `( )` で囲んで先頭の値だけにする |
| `use '##', instate of 'int!'` (警告) | nilable の引数を省略した | `func( 1 ## )` のように `##` を付ける |
| `use '**' at arg(2)` (警告) | 多値を返す関数の結果を、可変長・省略可能な引数に渡した | `func()**` のように `**` を付ける |
| `not found -- ')'. actually '.'` | 文字列リテラルに直接メソッドを付けた (`"".byte()`) | 変数に入れてから呼ぶ |
| `not found field typeInfo: str.len` | `str` に `len()` は無い | `#txt` を使う |
| `not found type -- table` / `pcall` / `math.floor` | 組み込みでない Lua の API | `module` 宣言で宣言して使う (上の「組み込みの関数」)。`table.concat` の代わりは `string._join` |
| `'return' could not use here` | モジュールのトップレベルで `return` した | 関数の中に入れる |
| `... can't be external symbol` / `This type must be public` | `pub` のクラス・関数が、`pub` でない型を使っている | 使っている型も `pub` にする |
| `can't call the type -- xxx!, TypeInfoKind.Nilable` | nilable の値のメソッドを呼んだ | `obj$.method$()` か、`if!` などで nil でなくしてから呼ぶ |
| `This statement is not reached` | `while true` や `__` (戻らない関数) の後に文がある | 到達しない文を消す |
| `exe` で `__main` が動かない | `exe` はトップレベルの処理だけを実行する | `-shebang` で実行するか、`save --main` で変換する |
| 変換した `main.lua` が `module 'xxx' not found` | import しているモジュールを変換していない | 全ファイルを変換する (`find . -name '*.lns' \| lnsc @- save`) |
| テストが NG なのに終了コードが 0 | `--testing` は終了コードで結果を返さない | 出力の `NG:` の数を見る |

## Lua の意味に依存する点

- `%`・整数の `/` は床除算 (`-7 / 2` は `-4`、`-7 % 3` は `2`)。
- 整数の演算の結果が範囲外になると、Lua では実数になる (型は `int` のまま)。`@@int` で整数に戻す。
- 実数の文字列化は `%.14g` で、整数値の実数は `1.0` になる。
- 文字列はバイト列。`txt[ i ]` はバイトの値、`#txt` はバイト数。`lower`・`upper` は ASCII だけ。
- List の範囲外の添字のアクセスは未定義 (nil にならない場合がある)。Map のアクセスは nilable。
- Map のキーは int と real の同じ値を区別しない。int・real・str 以外のキーはオブジェクトの同一性で比較する。
- `List.sort` は安定ソートではない。同じ値の要素の順番が意味を持つ場合は、比較の条件で順番を決める。
