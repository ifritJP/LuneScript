# LuneScript の wasm 対応の検討

LuneScript で書いたプログラムを WebAssembly (wasm) で動かす方法を検討する。

> 調査対象: ブランチ `dev_with_llm` (`f86d38c0`)。
> 実験環境: Go 1.27.1、Node.js v22.22.1、Linux x86-64。

> **方針の更新**: サイズ、ヒープの使用量、将来のマルチスレッド対応を目的に、案 A ではなく **案 D (wasm の直接生成)** を目指すことになった。
> 案 D の設計検討は [wasm_direct_design.md](wasm_direct_design.md) を参照。本書の案 A の実験結果は、比較の基準 (サイズ、動作) として使う。

## 1. 結論

- **既存の Go への変換を経由する方法 (案 A) で、すでにほぼ動く。** LuneScript → Go (`-langGo`) → wasm (Go の標準ツールチェイン) の流れである。
  コンパイラ本体にはまったく手を入れずに、Go 変換の総合テスト (`test/testGo.lns`) が wasm 上で通ることを確認した (4 章)。
- 必要な作業は、新しいコード生成ではなく **ビルド手順の整備・テスト・制約の明文化** が中心になる。
- 主な制約は次の 3 点である。
  - Lua 実装に純 Go の gopher-lua (Lua 5.1 互換) を使う必要がある。
  - 外部プロセスの起動 (`io.popen` など) は使えない。
  - wasm のサイズが約 7.5 MB (gzip 後 約 1.7 MB) と大きい。
- 推奨: 案 A を採用し、5 章の手順で段階的に対応する。

## 2. 「wasm 対応」の範囲

wasm 対応には 2 つの意味がある。本書は (2) を扱う。

| # | 内容 | 現状 |
|---|---|---|
| 1 | **コンパイラ (lnsc) 自体** を wasm で動かす (ブラウザ上で変換する) | 対応済み。`make build-wasm` (`GOOS=js GOARCH=wasm`、`GOTAG=gopherlua`) で `src/lnsc.wasm` を作る。入口は `misc/go/main/lnsc_main_wasm.go` で、フォーマッタなどの関数を JavaScript に公開している |
| 2 | **LuneScript で書いたプログラム** を wasm にコンパイルして動かす | 本書の対象。専用の手順・テストはない |

wasm の実行環境も 2 種類ある。

| 種類 | Go の指定 | 実行環境 | 用途 |
|---|---|---|---|
| js/wasm | `GOOS=js GOARCH=wasm` | ブラウザ、Node.js (Go 付属の `wasm_exec.js` が必要) | Web ページへの組み込み |
| WASI | `GOOS=wasip1 GOARCH=wasm` (Go 1.21 以降) | wasmtime、wasmer、Node.js の `node:wasi` など | サーバー、CLI、プラグイン |

## 3. 実現方法の比較

| 案 | 流れ | 必要な作業 | 長所 | 短所 |
|---|---|---|---|---|
| **A. Go 経由** (推奨) | `.lns` → `-langGo` → Go → `go build` (wasm) | ビルド手順・テストの整備 | 既存の Go 変換とランタイム (`runtime_go`) をそのまま使える。テスト済みの経路。js と WASI の両方に対応 | Go ランタイムと gopher-lua を含むのでサイズが大きい。Lua 5.1 互換になる |
| B. Lua VM を wasm で動かす | `.lns` → Lua → wasm 上の Lua VM (wasmoon など) | Lua VM の組み込み | コンパイラの変更が不要。Lua 5.4 も選べる | 実行は Lua のインタプリタになるので遅い。LuneScript 側では作業がほぼない (外部ツールの組み合わせ) |
| C. C 経由 | `.lns` → C → Emscripten / wasi-sdk | C 変換 (`convCC`) の完成 | サイズを小さくでき、速い | C 変換は実験段階で、`LNS_SRCS` でもコメントアウトされている (ビルド対象外)。完成までの作業が非常に大きい |
| D. wasm を直接生成 | `.lns` → wasm | 新しいバックエンドと GC・ランタイムの作成 | 最も小さく速くできる | 作業量が最大。GC、文字列、Map などのランタイムを wasm 用に作る必要がある |

案 A は、セルフホストで鍛えられた Go 変換をそのまま使える。さらに 4 章のとおり実際に動くので、費用対効果が最も高い。

## 4. 実験結果 (案 A)

### 4.1 手順

```sh
lnsc hello.lns mkmain main.go                       # main.go を生成
lnsc hello.lns SAVE -langGo --main hello --app hello # hello.go を生成
# go.mod: require github.com/ifritJP/LuneScript/src + replace でローカルの src を参照
GOOS=wasip1 GOARCH=wasm go build -tags gopherlua -o hello.wasm .   # WASI 版
GOOS=js     GOARCH=wasm go build -tags gopherlua -o hello.wasm .   # js 版
node $(go env GOROOT)/lib/wasm/wasm_exec_node.js hello.wasm       # js 版の実行
node run_wasi.mjs hello.wasm                                       # WASI 版の実行 (node:wasi を使う 10 行程度のスクリプト)
```

### 4.2 結果

| 確認内容 | 結果 |
|---|---|
| 小さなプログラム (クラス、再帰、List、Map、`forsort`、nilable) | ネイティブ、js、WASI の出力が一致 |
| Go 変換の総合テスト `test/testGo.lns` (3,115 行) を gopher-lua 用の設定 (`/.nolua51./` の除外、`--int2str need0`) で変換 | ネイティブ (gopher-lua)、js、WASI の出力 (1,126 行) が完全に一致。ただし `io.popen` を使う 1 か所を除外した |
| `io.popen( "ls" )` (2644 行付近) | wasm では失敗 (`panic: error ls`)。外部プロセスを起動できないため |
| Lua 実装のビルドタグを指定しない場合 | ビルドエラー (`undefined: lua_rawstr` など)。`runtime_go` は Lua 実装が必須で、wasm では cgo の Lua が使えない |
| wasm のサイズ (hello) | 7.6 MB (`-ldflags='-s -w'` で 7.5 MB、gzip 後 1.7 MB) |

## 5. 対応方針 (案 A)

### ステップ 1: ビルド手順の整備 (小)

- `mk/go.mk` に、LuneScript のプログラムを wasm にビルドするターゲットを追加する
  (例: `make build-app-wasm SRC=xxx.lns GOOS=wasip1`)。中身は 4.1 の手順をまとめたもの。
- `GOTAG=gopherlua` を wasm では自動で付ける (既存の `build-wasm` と同じ)。
- 生成する `go.mod` の雛形を用意する (`buildGoSetup` の `go mod tidy` 方式を使うかどうかも決める)。

### ステップ 2: テストの追加 (小〜中)

- `test-go` の wasm 版として `test-go-wasm` を追加する。`test/testGo.lns` を wasm でビルドし、Node.js で実行して、
  ネイティブの gopher-lua 版の結果と diff する (4.2 の手順をそのまま自動化する)。
- wasm で使えない処理 (`io.popen` など) は、既存の `/.nolua51./` と同じ仕組みで除外できるようにする
  (例: `/.nowasm./` という印を付けて、wasm 用のテストでは `sed` で除外する)。
- 必要なら `test-comp` (`test-comp2-main` など) から呼ぶ。Node.js が無い環境では省略する。

### ステップ 3: 制約の明文化 (小)

ドキュメントに次の内容を書く。

- 使えない機能: 外部プロセスの起動 (`io.popen` で確認。`os.execute` なども同様と考えられるが未確認)、cgo の Lua。
- Lua 連携 (`__luago`、Lua の値の扱い) は gopher-lua (Lua 5.1 互換) の動作になる。
  整数を文字列にする時の表記 (`10` / `10.0`) などが Lua 5.3 版と異なるので、`--int2str` の指定が必要になる場合がある。
- ファイルアクセス: WASI では実行時に許可したディレクトリ (preopen) だけを使える。
  js/wasm のブラウザ環境ではファイルシステムが無い。
- スレッド: Go の wasm は 1 スレッドで動く。goroutine (LuneScript の非同期処理) は動くが、並列にはならない (`testGo.lns` には `__async` などを使うテストが約 50 か所あり、wasm でも結果が一致した)。

### ステップ 4: 使い勝手の改善 (中、必要に応じて)

- **JavaScript との連携**: 今は LuneScript から JavaScript の関数を呼ぶ手段や、LuneScript の関数を JavaScript に公開する手段がない。
  `lnsc_main_wasm.go` と同じように `syscall/js` で公開する Go のグルーコードを自動生成する仕組み、
  または Go 1.24 以降の `//go:wasmexport` (WASI) を使う仕組みを検討する。
- **サイズの削減**: Lua 連携を使わないプログラム向けに、Lua 実装を含めない (Lua を呼ぶと実行時エラーにする) ビルドタグを
  `runtime_go` に追加すると、gopher-lua の分だけ小さくできる可能性がある (未計測)。TinyGo も候補だが、未検証である
  (`runtime_go` が使う reflect や goroutine の対応状況を確認する必要がある)。
- **lnsc のオプション化**: `lnsc xxx.lns build-wasm` のように、変換から wasm の出力までを lnsc の 1 コマンドで行う。

## 6. 未確認の事項

- wasmtime / wasmer での実行 (今回は Node.js だけで確認した)。
- ブラウザでの実行 (js/wasm を Node.js で実行して確認した)。
- `__async` などの非同期処理を多用するプログラムの性能。
- 実行速度の比較 (ネイティブ、wasm、Lua)。
- TinyGo でのビルドの可否。
