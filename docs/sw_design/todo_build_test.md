# ビルド・テスト環境の改善点 (TODO)

`src/Makefile`、`.github/workflows/test.yml`、および [test_comp.md](test_comp.md) の解析から
見つかった改善点をまとめる。

> 調査対象: ブランチ `dev_with_llm` (`87ebeb68`)。行番号は `src/Makefile` のもの。
> いずれも静的な解析にもとづく指摘で、実際に試して確認したものではない。

優先度の目安 (指摘時点):

- **高**: 誤った結果になる、作業が失われる、CI が動かない可能性があるもの
- **中**: 手間や時間がかかる、つまずきやすいもの
- **低**: 保守性や見通しの問題

状態:

- **対応予定**: 対応する方針が決まったもの
- **対応済み**: Makefile などを修正して対応したもの
- **文書化済み**: 文書化で対応し、すでに記載したもの
- **対応不要**: 意図した仕様である、または問題にならないと判断したもの
- **未対応**: 問題はあるが、現状は手が回っていないもの
- **未判断**: まだ方針を決めていないもの

## 一覧

| # | 優先度 | 状態 | 項目 |
|---|---|---|---|
| 1 | 高 | 対応済み | `test-comp` が成功すると確認なしで生成物を上書きする |
| 2 | 高 | 対応済み | `go-build` が `src/lnsc` を確認なしで上書きする |
| 3 | 高 | 対応不要 | `luarocks-check` がユーザーの `~/.luarocks` を変更する |
| 4 | 高 | 未対応 | GitHub Actions のアクションと実行環境が古い |
| 5 | 高 | 対応不要 | CI で Go 版の `test-comp-with-go` を実行していない |
| 6 | 中 | 文書化済み | Go 版 lnsc (`src/lnsc`) が git 管理外で、作り方が文書化されていない |
| 7 | 中 | 対応不要 | Lua 版だけで accept すると `.go` と `lnsc` が古いまま残る |
| 8 | 中 | 対応不要 | `LANG_KIND=go` を `WITH_GO=y` なしで使える |
| 9 | 中 | 対応不要 | `-j` 実行時に失敗箇所が分かりにくい |
| 10 | 中 | 文書化済み | 前提ツールの確認がなく、途中で失敗する |
| 11 | 中 | 対応不要 | Go のビルドでネットワークアクセスが必要 |
| 12 | 中 | 対応不要 | `work/` を複数のターゲットが共有している |
| 13 | 中 | 対応不要 | CI で `--testing` のテストを省略している |
| 14 | 低 | 未判断 | モジュール一覧が 2 か所に重複している |
| 15 | 低 | 未判断 | rockspec のファイル一覧を手で管理している |
| 16 | 低 | 未判断 | 環境依存のパスがハードコードされている |
| 17 | 低 | 未判断 | `Makefile` が巨大で、テストの全体像を把握しにくい |
| 18 | 低 | 未判断 | `Makefile.local` の存在と役割が分かりにくい |
| 19 | 低 | 未判断 | コメントアウトされた CI ジョブが残っている |

## 詳細

### 1. `test-comp` が成功すると確認なしで生成物を上書きする (高)

- **状態: 対応済み**
- **判断**: 成功時に上書きするのは意図した仕様であり、既定の動作は変えない。
  テストだけを行うケースは基本的に不要だが、成果物を上書きしないオプションを追加してもよい。
- **対応**: `test-comp2` に `NO_ACCEPT=y` を追加した。指定すると `diff-result` まで行い、
  `accept-package` を省略する。既定の動作 (成功時に上書き) は変えていない。

- **現状**: `test-comp2` (1059) の最後で `accept-package ACCEPT=y` が実行され、
  `lune/base/*.lua`, `*.meta`, `*.go`, `src/lnsc`, `test/result/test1.result` が上書きされる。
  `ACCEPT=y` によって確認プロンプト (1305–1307 行) も飛ばされる。
- **問題**: テストを実行しただけのつもりで、コミット対象のファイルが変更される。
  コミットしていない状態で実行すると、元の状態に戻しにくい。
- **当初の改善案** (参考):
  - テストのターゲット (`test-comp`) と昇格のターゲット (`accept`) を分ける。
    昇格したい場合は `make test-comp accept-package` のように明示する。
  - 少なくとも `ACCEPT=y` を外せる変数 (例: `NO_ACCEPT=y`) を用意する。
  - 上書きした場合はその旨と `git diff --stat` を最後に表示する。

### 2. `go-build` が `src/lnsc` を確認なしで上書きする (高)

- **状態: 対応済み**
- **判断**: 1 と同じ。上書きするのは意図した仕様であり、既定の動作は変えない。
- **対応**: `go-build` に `NO_ACCEPT=y` を追加した (1 と同じ変数名)。指定すると
  `work/go/build/lnsc` を `src/lnsc` (wasm の場合は `lnsc.wasm`) にコピーしない。

- **現状**: `go-build` は最後に `work/go/build/lnsc` を `src/lnsc` にコピーする。
  `src/lnsc` は git 管理外なので、上書きすると元に戻せない。
- **問題**: 動作実績のある lnsc を、壊れた可能性のある新しい lnsc で上書きしてしまう。
  `NO_TEST=y` を付けた場合はテストせずに上書きされる。
- **当初の改善案** (参考):
  - 上書きの前に `src/lnsc.bak` に退避する。
  - `NO_TEST=y` の場合は `src/lnsc` にコピーせず、`work/go/build/lnsc` に残すだけにする
    (コピーする場合は `INSTALL_LNSC=y` などで明示する)。

### 3. `luarocks-check` がユーザーの `~/.luarocks` を変更する (高)

- **状態: 対応不要**
- **判断**: 意図した仕様である。

- **現状**: `luarocks-check` (374) は `luarocks remove/make --local` で、ユーザーの
  `~/.luarocks` にある `lunescript` を削除・インストールし、最後に削除する。
- **問題**: ユーザーが luarocks で LuneScript を入れていた場合、テスト後にはそれが消えている。
  ローカルでは `Makefile.local` の `NO_LUAROCKS=y` で回避しているが、既定では実行される。
- **改善案**:
  - `luarocks --tree work/8/rocks` のように、`work/` 以下の専用ツリーにインストールする。
  - 既定を無効にし、`release-check` などのリリース前確認でだけ有効にする。

### 4. GitHub Actions のアクションと実行環境が古い (高)

- **状態: 未対応**
- **判断**: CI は 3 年間保守できていない。

- **現状**: `.github/workflows/test.yml` で次の古いバージョンを使っている。
  - `actions/checkout@v2`, `actions/setup-python@v2`, `actions/setup-go@v2`
  - `actions/upload-artifact@v2`
  - Python 3.8
- **問題**: v2 系のアクションは非推奨の Node.js ランタイムで動く。
  `upload-artifact` の旧バージョンはすでに使えなくなっている。Python 3.8 もサポートが終了している。
  そのため CI が失敗するか、警告が出ている可能性が高い (最近の実行結果は未確認)。
- **改善案**: 各アクションを最新のメジャーバージョン (`checkout@v4`, `setup-go@v5`,
  `upload-artifact@v4` など) に上げ、Python も現行バージョンにする。
  Python は `hererocks` のためだけに使っているので、`leafo/gh-actions-lua` などで Lua 環境を直接用意する方法もある。

### 5. CI で Go 版の `test-comp-with-go` を実行していない (高)

- **状態: 対応不要**
- **判断**: 特に問題ない。

- **現状**: CI は Lua 版の `test-comp` と `go-build` を別々に実行している。
- **問題**: 普段の開発で使う `test-comp-with-go` でだけ行われる確認が CI では行われない。
  - Lua 版と Go 版の出力一致 (`work/13` ⇔ `work/3_3/...`)
  - batch ビルドでも出力が安定するかの確認 (`test-new3`)
  - `test-go` の新旧比較
  - `test/result/test1.result` との比較
- **改善案**: CI に `test-comp-with-go` のジョブを追加する。Go 版の org (`src/lnsc`) が無いので、
  先に `go-build NO_TEST=y` で作ってから実行する。

### 6. Go 版 lnsc (`src/lnsc`) が git 管理外で、作り方が文書化されていない (中)

- **状態: 文書化済み**
- **判断**: 文書化で対応する。
- **対応**: [test_comp.md](test_comp.md) に記載した。
  - 3.1: Go 版の前提として `src/lnsc` が必要で、clone 直後には存在しないこと
  - 3.3: 既存の lns コードから `make go-build` で Go 版 lnsc をビルドする手順

- **現状**: `test-comp-with-go` の org は `src/lnsc` だが、git 管理外 (`.gitignore` にもない) である。
  clone 直後には存在しない。
- **問題**: 新しい環境で `test-comp-with-go` を実行すると、`./lnsc` が見つからずに失敗する。
  `git status` には常に `?? src/lnsc` と表示される。
- **改善案**:
  - `src/lnsc` を `.gitignore` に追加する。
  - `test-comp-with-go` の冒頭で `src/lnsc` の有無を確認し、無ければ `go-build NO_TEST=y` を
    実行する (または手順を表示して止まる)。
  - readme に、clone 後に Go 版を準備する手順を書く。

### 7. Lua 版だけで accept すると `.go` と `lnsc` が古いまま残る (中)

- **状態: 対応不要**
- **判断**: 意図した仕様である。

- **現状**: `accept-package` は `work/3/lnsc` がある場合にだけ `*.go` と `lnsc` を更新する (1314–1317 行付近)。
  Lua 版の `test-comp` では `.lua` だけが更新される。
- **問題**: コミット済みの `lune/base/*.go` と `*.lua` で、元にした `.lns` の版が食い違う。
  その後の Go 版のビルドやテストの結果が、Lua 版と異なる可能性がある。
- **改善案**:
  - Lua 版の accept 時に警告を出す。
  - `lune/base/*.go` 内の buildId と `*.lua` の buildId を比べ、食い違いを検出するターゲットを追加する。

### 8. `LANG_KIND=go` を `WITH_GO=y` なしで使える (中)

- **状態: 対応不要**
- **判断**: 特に問題ない。

- **現状**: `LANG_KIND=go` だけを指定すると、Go 版の流れ (`test-comp2-para`) を通るが、
  `COM_OPT` に `-langGo` が入らない。
- **問題**: `work/13` (`-langGo` 付き) と `work/3_3` (`-langGo` なし) の比較で前提が揃わない。
- **改善案**: `Makefile` の冒頭で、`LANG_KIND=go` かつ `WITH_GO` が空の場合に `$(error ...)` で止める。
  または `LANG_KIND=go` のときは `WITH_GO=y` を自動で設定する。

### 9. `-j` 実行時に失敗箇所が分かりにくい (中)

- **状態: 対応不要**
- **判断**: エラー時は `-j1` で再実行して回避している ([test_comp.md](test_comp.md) 3.4)。

- **現状**: 速度のために `-j` を付けて実行すると、ログが混ざって失敗箇所が分かりにくい。
  エラー時は `-j1` で再実行して確認している (test_comp.md 3.4)。
- **問題**: `-j1` での再実行には時間がかかる。
- **改善案**:
  - `make --output-sync=target` (GNU make 4 以降) を既定にし、ターゲットごとにログをまとめて出力する。
  - 各テストターゲットのログを `work/log/<target>.log` にも保存し、失敗したターゲット名を最後に表示する。
  - `diff-result` の各 diff の前に、何を比較しているのかを `echo` する。

### 10. 前提ツールの確認がなく、途中で失敗する (中)

- **状態: 文書化済み**
- **判断**: 文書化で対応する。
- **対応**: 必要なパッケージ・ツールを [test_comp.md](test_comp.md) の 2 (ビルド環境設定) と 3.1 (前提条件) に記載した。

- **現状**: 必要なツール (`lua5.3`, `pkg-config lua-5.3`, `gcc`, `go`, `rsync`, `json_pp`,
  `python3`, `/usr/bin/time`) が無くても、`test-comp` は実行を開始する。
- **問題**: 長時間実行した後で、途中のステップで失敗する。
- **改善案**: `check-env` ターゲットを追加して必要なツールとバージョンを確認し、`test-comp` の冒頭で呼ぶ。

### 11. Go のビルドでネットワークアクセスが必要 (中)

- **状態: 対応不要**
- **判断**: 特に問題ない。

- **現状**: `buildGoSetup` (1631) が毎回 `go mod init` と `go mod tidy` を実行し、`sed` で
  `go.mod` を書き換えている (`GOMODPAT`)。
- **問題**: オフライン環境やモジュールキャッシュが無い環境ではビルドできない。
  また、`go mod tidy` の結果が実行時期によって変わる可能性がある。
- **改善案**: 作業ディレクトリ用の `go.mod` / `go.sum` を固定のテンプレートとしてリポジトリに置き、
  コピーして使う。または `go work` を使ってローカルモジュールを参照する。

### 12. `work/` を複数のターゲットが共有している (中)

- **状態: 対応不要**
- **判断**: 特に問題ない。

- **現状**: `test-comp` は `rm -rf work/*` で始まる。`go-build` も `work/1`, `work/2` などを使う。
- **問題**: 別のターゲットを実行すると、直前の `test-comp` の結果 (失敗の調査に必要なファイル) が消える。
  同じ作業ツリーで 2 つのターゲットを同時に実行することもできない。
- **改善案**: ターゲットごとにディレクトリを分ける (`work/test-comp/`, `work/go-build/` など)。
  `WORK_DIR` 変数で場所を変えられるようにする。

### 13. CI で `--testing` のテストを省略している (中)

- **状態: 対応不要**
- **判断**: `make test-comp` がローカルで正常に実行できていることが前提なので、特に問題ない。

- **現状**: CI は `SKIP_CHECKING_TESTING=y` を付けている。
  `Makefile` のコメントには「CI 環境では `--testing` が正常に動かない」とある (846)。
- **問題**: 単体テストの仕組み (`--testing`) が CI で確認されない。
- **改善案**: CI で動かない原因を調べて修正し、`SKIP_CHECKING_TESTING` を外す。

### 14. モジュール一覧が 2 か所に重複している (低)

- **現状**: 同じモジュール一覧が `LNS_SRCS` (215–262) と `selfhost_buildOp` (271–314) の 2 か所に書かれている。
- **問題**: モジュールを追加・削除するときに片方を直し忘れると、Lua 版と Go 版でビルド対象がずれる。
- **改善案**: `selfhost_buildOp` を `$(foreach src,$(LNS_SRCS),$(call $5,...))` で生成し、一覧を 1 か所にする。

### 15. rockspec のファイル一覧を手で管理している (低)

- **現状**: `lunescript-main-1.rockspec` などに Lua ファイルの一覧を手で書き、
  `check-rockspec` で実ファイルとの差分を確認している。
- **改善案**: `LNS_SRCS` から rockspec のファイル一覧を生成するスクリプトを用意する。

### 16. 環境依存のパスがハードコードされている (低)

- **現状**:
  - `release-check` (417–418): `../../lctags/external/lua/...`
  - `LUACHECK` / `LNSC_LUAROCKS` (5, 27): `~/.luarocks/bin/...`
  - 2095 行付近: `~/work/ifritJP.github.io`
  - 2113 行: `../../lnstags/lnstags`
- **問題**: 開発者の環境以外では、これらのターゲットが動かない。
- **改善案**: 変数にして `Makefile.local` で上書きできるようにし、既定値は `which` などで探す。

### 17. `Makefile` が巨大で、テストの全体像を把握しにくい (低)

- **現状**: `src/Makefile` は約 2,200 行あり、ビルド、セルフホスト、各種テスト、Go、wasm、
  プロファイル、ドキュメント関連のターゲットが混在している。
- **改善案**:
  - `mk/selfhost.mk`, `mk/test.mk`, `mk/go.mk` などに分割して `include` する。
  - `help` ターゲットを、主要ターゲットの説明と推奨コマンドを表示する内容に更新する。
  - 使われていないターゲット (`Makefile.ccc` 関連の C 変換や `primal/` など) を整理する。

### 18. `Makefile.local` の存在と役割が分かりにくい (低)

- **現状**: `Makefile` は `Makefile.local` があれば `include` する (118–120)。
  `.gitignore` の `*.local` で除外されている。手元では `LUA51=` と `NO_LUAROCKS=y` が設定されている。
- **問題**: `Makefile.local` の有無で `test-comp` の内容が変わるが、そのことに気付きにくい。
- **改善案**: `Makefile.local.sample` をリポジトリに置き、よく使う設定例を示す。
  `test-comp` の開始時に、有効な主要変数 (`LANG_KIND`, `WITH_GO`, `LUA51`, `NO_LUAROCKS` など) を表示する。

### 19. コメントアウトされた CI ジョブが残っている (低)

- **現状**: `test.yml` に、CentOS 7/8 のジョブと Google Drive へのアップロードがコメントアウトされたまま残っている。
- **改善案**: 不要なら削除する。必要なら別ファイルにして `workflow_dispatch` で手動実行できるようにする。
