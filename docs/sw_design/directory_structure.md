# LuneScript ディレクトリ構成

LuneScript は Lua / Go を変換先とするトランスコンパイラであり、
コンパイラ自身も LuneScript で記述されたセルフホスト言語である。
本書はリポジトリのディレクトリ構成と各ディレクトリの役割をまとめる。

> 調査対象: ブランチ `dev_with_llm` (`87ebeb68`) の git 管理ファイル。
> 未追跡のローカル作業ファイル (`src/aaa/`, `src/hoge/`, `src/lnsc` 等のビルド成果物) は対象外とする。

## 1. トップレベル

```
LuneScript/
├── .github/workflows/test.yml   CI (GitHub Actions)
├── docs/                        設計・解説ドキュメント (org 形式)
│   └── sw_design/               ソフトウェア設計ドキュメント (本書)
├── etc/lua-5.3.pc               pkg-config 定義 (C ランタイムビルド用)
├── src/                         コンパイラ本体・ランタイム・テスト
├── tools/                       LuneScript で書かれた周辺ツール
├── llm/                         LLM 学習用ワークスペース (本ブランチでは未追跡)
├── lns-*.el                     Emacs 用メジャーモード・補完・flycheck 等
├── lunescript-main-1.rockspec   LuaRocks パッケージ定義 (Lua 5.2+)
├── lunescript51-main-1.rockspec LuaRocks パッケージ定義 (Lua 5.1)
├── readme.org                   README
└── LICENSE
```

### docs/

| ファイル | 内容 |
|---|---|
| `design.org` | 設計メモ |
| `introduce*.org` | 言語紹介 |
| `mapping.org` | Mapping (構造化⇔非構造化データ変換) |
| `mutable.org` | mutable/immutable の扱い |
| `completion.org` | コード補完 |
| `reason.org` | 言語開発の動機 |
| `todo.org` | TODO |

### Emacs サポート (`lns-*.el`)

`lns-mode.el` (メジャーモード) を中心に、補完 (`lns-completion.el`,
`lns-company-mode.el`, `lns-auto-complete.el`, `lns-helm.el`)、
構文チェック (`lns-flycheck.el`, `lns-flymake.el`)、
lnsc 呼び出し (`lns-command.el`, `lns-dispatch.el`, `lns-inq.el`)、設定 (`lns-conf.el`) から成る。

## 2. src/ — コンパイラ本体

```
src/
├── Makefile            ビルド・セルフホスト・テストの全ターゲット
├── Makefile.ccc        C 変換系の実験用 Makefile
├── lnsc.lua            Lua 版 lnsc のエントリ (lune.base.base を require)
├── setup.lua           インストール用セットアップ
├── go.mod / go.sum     Go モジュール定義 (module github.com/ifritJP/LuneScript/src)
├── lune.js             lnsc のプロジェクト設定 (共通コンパイルオプション等)
├── build.counter       ビルド回数カウンタ
├── lune/base/          コンパイラ本体 (後述)
├── primal/             初期ブートストラップ用の Lua 実装
├── misc/go/            Go 版 lnsc の main パッケージ等
├── check/go/           Go ビルド確認用の最小プロジェクト
└── test/               テストケース
```

### 2.1 src/lune/base/ — コンパイラ本体

各モジュールは **`.lns` (ソース) / `.lua` (Lua 変換結果) / `.go` (Go 変換結果)** の
3 点セットでコミットされている。`.lns` を正とし、`.lua`・`.go` は
LuneScript 自身で変換した生成物である (セルフホスト)。
`TransUnitCtrl.lns` などサフィックス付きのファイルは親モジュールに
`subfile use` で取り込まれる分割ソース (`subfile owner`) のため、単独の `.lua`/`.go` を持たない。

処理の流れに沿って分類すると次のとおり。

| 区分 | モジュール | 役割 |
|---|---|---|
| エントリ / ドライバ | `base.lua`, `front`, `frontInterface`, `Option`, `Runner`, `LuneControl` | コマンドライン解析、コンパイル全体の制御 |
| 字句解析 | `Tokenizer`, `AsyncTokenizer`, `Code` | トークン化 (非同期版あり) |
| 構文解析・意味解析 | `TransUnit` (+ `TransUnitCtrl/Decl/Exp/StatementList/Test/IF`), `BuiltinTransUnit` | AST 構築、型検査 |
| AST / 型 | `Nodes`, `ConvNodes`, `Ast`, `AstInfo`, `Types`, `Builtin`, `NodeIndexer` | ノード定義、型情報、組込み型 |
| マクロ | `Macro`, `TestingMacro`, `Testing` | マクロ展開、テスト用マクロ |
| モジュール / メタ情報 | `Import`, `Meta`, `MetaLoader`, `Depend`, `DependLuaOnLns`, `OutputDepend`, `LuaMod`, `GoMod`, `LuaVer` | import 解決、`.meta` 入出力、依存関係、Lua バージョン差異 |
| コード生成 | `convLua`, `convGo`, `convPython`, `Converter`, `Formatter`, `dumpNode`, `glueFilter` | Lua / Go (/ Python) 出力、整形、AST ダンプ、C グルー生成 |
| 共通ユーティリティ | `Util`, `Str`, `Json`, `Log`, `Writer`, `Async`, `Ver` | 汎用処理 |
| Lua ランタイム | `runtime.lua`, `runtime8.lua` | 変換後 Lua コードが利用するランタイム |
| Go バインディング | `binding.go`, `binding_dummy.go` | Go 版での Lua 連携 |

#### runtime_go/

Go 変換結果が利用するランタイム (`package runtimelns`)。

- `lua.go`, `luaCore.go`, `luaFunc.go`, `luaMisc.go`, `gopherluaCore.go` — Lua VM 連携 (C Lua / gopher-lua)
- `collection*.go`, `string*.go`, `stream*.go`, `tuple.go`, `err.go` — 組込み型の実装
- `asyncCtrl*.go`, `depend*.go`, `builtin*.go` — 非同期制御、環境依存処理
- `*_env.go` / `*_noenv.go` — 実行環境 (env) の有無でビルドタグにより切り替わる実装
- `legacy/lua51/` — Lua 5.1 向けのレガシー実装

#### runtime_c/

C 変換 (実験的) 向けランタイム。`lns_base.c`, `lns_builtin.c`, `lns_gc.c`,
`lns_collection.cxx`, `lns_alloc.cxx`, `lns_luaWrapper.c`, `lns_main.c` 等。

### 2.2 src/primal/

LuneScript がセルフホストする以前の、Lua で直接書かれた初期実装
(`Parser.lua`, `TransUnit.lua`, `convLua.lua`, `dumpNode.lua`, `base.lua`)。
歴史的なブートストラップ用コードで、現在の `Makefile` からは参照されていない。

### 2.3 src/misc/go/

- `main/lnsc_main.go`, `main/common_main.go` — Go 版 lnsc の `main`
- `main/lnsc_main_wasm.go` — WebAssembly 版 lnsc の `main`
- `binding/bind.lns` — Go バインディング定義

### 2.4 src/test/

| パス | 内容 |
|---|---|
| `*.lns` (直下 約 47 件) | 言語機能ごとの正常系テスト (`test.lns` 等) |
| `*.lua`, `*.c` | Lua 連携・C グルーのテスト補助 |
| `error/` (約 100 件) | コンパイルエラーになるべきコードのテスト |
| `tutorial/` | チュートリアル例題 (`class1.lns`, `generics.lns`, `match.lns` 等) |
| `proj/` | 複数モジュール構成のプロジェクトテスト |
| `ext/`, `test2/`, `subGoPkg/` | 外部モジュール・サブパッケージ参照のテスト |
| `result/` | 期待結果 |
| `docker/` | インストール手順のテスト (`compose/linux`, `alpine`, `centos7/8` の Dockerfile) |

## 3. tools/ — 周辺ツール

いずれも LuneScript で記述され、`lune.js` を持つ独立プロジェクトになっている。

| ディレクトリ | 内容 |
|---|---|
| `formatter/` | EBNF ベースのコードフォーマッタ (`lns/Ebnf*.lns`, `Parser.lns`, `Rule.lns` 等) |
| `history/` | ビルド時間等の履歴を集計・プロット |
| `runnerAnalyzer/` | Runner (非同期処理) のログ解析 |
| `scheduleProf/` | スケジューリングのプロファイル |

## 4. llm/ — LLM 学習用ワークスペース

LuneScript コードを生成する LLM のファインチューニング用作業領域。
`dev_with_llm` ブランチでは未追跡 (`llm` ブランチで管理)。

- `artifacts/dataset/` — 学習データセット (サンプル LuneScript コード、AtCoder ABC 問題、検証スクリプト)
- `res/` — 参照資料 (リファレンス、セルフホストソース等)
- `train_lora.py`, `export_gguf.py` — Unsloth による LoRA 学習・GGUF 出力
- `.agent/rules/lunescript.md` — AI エージェント向け作業ルール

## 5. ビルドとセルフホスト

`src/Makefile` が中心となる。主なターゲット:

| ターゲット | 内容 |
|---|---|
| `make build` | 既存の lnsc で `lune/base/*.lns` を変換し `.lua` を再生成 (セルフホストビルド) |
| `make go-build` | `.lns` を `-langGo` で Go に変換し、Go 版 lnsc をビルド |
| `make test-comp` | `work/1` 以下にセルフホストビルドし、生成したコンパイラでテストを実行 |
| `make test-go`, `test-tutorial-go` | Go 変換結果のテスト |
| `make test-error` | `test/error/` のエラーテスト |
| `make install` | Lua モジュールディレクトリへのインストール |
| `make build-wasm` | WebAssembly 版のビルド |

セルフホストの流れ:

```
lune/base/*.lns ──(既存 lnsc)──▶ *.lua ──▶ Lua 版 lnsc (lnsc.lua → lune.base.base → front)
                └─(既存 lnsc -langGo)──▶ *.go ──▶ Go 版 lnsc (misc/go/main + runtime_go)
```

CI (`.github/workflows/test.yml`) では Lua 5.3 環境で `test-comp`、`go-build`、
`test-global` (luacheck によるグローバル変数チェック) と、Docker でのインストール確認を実行する。

詳細は [make test-comp 解析](test_comp.md) を参照。
