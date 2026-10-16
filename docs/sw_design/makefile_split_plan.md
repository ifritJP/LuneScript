# `src/Makefile` 分割計画

約 2,200 行ある `src/Makefile` を、**動作を変えずに** 複数のファイルに分割するための計画と検証方法をまとめる。

> 調査対象: ブランチ `dev_with_llm`。行番号は分割前 (`5a665062`) の `src/Makefile` のもの。
> フェーズ 1 は実施済み。結果は 8 章を参照。
> GNU Make 4.4.1 で確認した。

## 1. 方針

- **目的**: ファイルを分けて見通しを良くする。**動作は 1 行も変えない。**
- **やらないこと**: ターゲットの整理・改名、処理の修正、不具合の修正。
  分割中に見つけた問題は記録だけして、分割とは別のコミットで扱う。
- **2 段階で進める**:
  - **フェーズ 1 (機械的分割)**: 元のファイルを **連続した行範囲** で切り出し、**元と同じ順序で** `include` する。
    include を展開すると元の Makefile と完全に同じテキストになるので、動作が同じであることを機械的に保証できる。
  - **フェーズ 2 (任意の整理)**: 関連するターゲットを別のファイルに移す。順序が変わるので、
    フェーズ 1 より慎重な検証が必要になる (6 章)。まずはフェーズ 1 だけを行う。

## 2. 分割を難しくしている要因 (調査結果)

### 2.1 Makefile の記述順序に依存している箇所がある

make は Makefile を上から順に読み、`ifeq` / `ifdef` / `:=` / `$(shell ...)` をその時点で評価する。
レシピ内に書かれた `ifeq` も、読み込み時に評価される。そのため、記述の順序を変えると動作が変わる箇所がある。

| 箇所 | 内容 |
|---|---|
| `Makefile.local` の include (118 行) | 51–58 行の `ENABLE_LUAROCKS` / `NO_LUAROCKS` の判定はこの include より前にあるため、`Makefile.local` の設定は反映されない。一方、これより後の判定には反映される。include の位置を変えると動作が変わる |
| `diff-result` 内の `ifeq ($(RUNTIME_ORG_VER),$(RUNTIME_NEW_VER))` (1224 行) | 変数の定義 (1353–1356 行) より **前** にあるため、評価時には両方とも未定義 (空) で、常に真になる。`make -n diff-result` で、`.meta` の比較 (`work/0-1/test/` と `work/4/test/` の diff) が常に実行されることを確認した。定義を前に移すと動作が変わる |
| 既定のゴール | 最初に現れるターゲット (`depend_on` 未指定時は `help`) が既定のゴールになる。順序が変わると、引数なしの `make` の動作が変わる |
| `check_singleBin`, `dump-lns-lua-ver`, `buildGo` など | `WITH_GO` / `GOTAG` によって、`define` 自体が切り替わる |

→ フェーズ 1 では **順序をまったく変えない** ことで、これらの影響を受けないようにする。

### 2.2 Makefile 自身を作業ディレクトリにコピーしている

次の 3 か所で `Makefile` を `work/` 以下にコピーし、そこで `$(MAKE) -C work/N ...` を実行している。
分割後は `mk/` も一緒にコピーしないと、`work/2` などでの make が失敗する。

| 行 | 箇所 | コピー先 |
|---|---|---|
| 467 | `cp_src_set` | `work/1`, `work/2`, `work/3`, `work/5`, `work/6`, `work/13`, `work/formatter/work/*` など (呼び出し元が多い) |
| 1916 | `go-build` | `work/go` |
| 2140 | `test-formatter-prepare` | `work/formatter` |

`Makefile.ccc` のコピー (1474 行など) は対象外。また、`Makefile.local` は元からコピーされていない。
この動作は変えない。

### 2.3 別のディレクトリから include されている

`tools/history/Makefile.prof` が `src/Makefile` を include している
(`include $(WORKDIR)/src/Makefile` または `include Makefile`)。
分割後の `include mk/xxx.mk` を **カレントディレクトリからの相対パス** で書くと、カレントディレクトリが違う場合に失敗する。
そのため **`src/Makefile` の場所を基準にしたパス** で include する。

### 2.4 カレントディレクトリからの相対パスで include しているもの

`include Makefile.local` (119 行)、`include lune.mk` (64 行)、`include $(DEPENDS)` (95 行) は、
カレントディレクトリからの相対パスである。分割後もこのまま (書き換えない) にする。

### 2.5 Makefile をテキストとして読んでいるコードがある (実施時に判明)

`misc/go/binding/bind.lns` は `<dir>/Makefile` を **テキストとして読み**、`LNS_SRCS +=` で始まる行から
Go 版 lnsc に埋め込むモジュールの一覧を作る (35–48 行)。`LNS_SRCS` を `mk/` に移すと、
`binding.go` にモジュールが埋め込まれなくなる。その結果、単体バイナリの lnsc が `lune.base.Testing` などを読み込めなくなり、
`check_singleBin` (`test-new` の最後) で失敗する。

make の動作の検証 (A〜C) では検出できず、実際のビルド (D) で見つかった。
→ `LNS_SRCS` の定義ブロックは `Makefile` 本体に残す (順序は変わらない)。

### 2.6 特殊な機能は使っていない

`.ONESHELL`、`.SECONDEXPANSION`、`export`、`override`、`vpath`、`+` で始まるレシピ行などは使われていない。
分割による影響を受ける特殊な機能はない。

## 3. フェーズ 1 の分割案

`src/mk/` に連続した行範囲で切り出す。境界は空行やコメントの区切りに合わせて実装時に確定する。

| ファイル | 行範囲 (目安) | 内容 |
|---|---|---|
| `mk/config.mk` | 1–129 | Lua / lnsc の設定、`LANG_KIND`、`COM_OPT`、`depend_on` 時の `build_from_d`、`Makefile.local` の include、`.PHONY` |
| `mk/help.mk` | 130–169 | `help`, `exec` |
| `mk/selfhost.mk` | 170–491 | `exeLune` などの共通マクロ、`LNS_SRCS`、`selfhost_build*`、`build-testcode`、`luarocks-check`、`build`、`cp_src_set` |
| `mk/test-comp.mk` | 492–614 | `test-global`, `test-comp`, `test-comp-with-go`, `--testing` 関連 |
| `mk/test-cases.mk` | 615–814 | エラーテスト (`TEST_CASE`, `test-error`)、補完テスト (`COMP_TARGETS`, `test-cmpl`) |
| `mk/test-stages.mk` | 815–1238 | `test-org`, `test-new*`, `test-parallel`, `test-comp2*`, `check-rockspec`, `diff-result` |
| `mk/test-lua51.mk` | 1239–1305 | Lua 5.1 のクロスコンパイルテスト |
| `mk/release.mk` | 1306–1430 | `accept-package`、Docker テスト、ランタイム / メタ形式の確認、`build-glue`、`install` |
| `mk/conv-c.mk` | 1431–1605 | C 変換関連 (実験的) |
| `mk/go.mk` | 1606–1959 | `setup-pkg-lua`、Go のビルド設定、`test-go`、`go-build`、`test-lnsc-go` |
| `mk/prof.mk` | 1960–2099 | `test-govslua`、`goprof`、`luaprof` |
| `mk/tools.mk` | 2100–末尾 | `test-reference`、`update-lnstags`、フォーマッタのテスト |

`luarocks-check` が `selfhost.mk` に入るなど、内容的に不自然な箇所も残るが、フェーズ 1 では順序の維持を優先する。

分割後の `src/Makefile` は次の形になる。

```make
# src/Makefile の場所 (他のディレクトリから include されても mk/ を見つけられるようにする)
LNS_MK_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

include $(LNS_MK_DIR)mk/config.mk
include $(LNS_MK_DIR)mk/help.mk
include $(LNS_MK_DIR)mk/selfhost.mk
...
include $(LNS_MK_DIR)mk/tools.mk
```

- `LNS_MK_DIR` は、他の include (`Makefile.local` など) より **前** の、ファイルの先頭で求める
  (include すると `MAKEFILE_LIST` の末尾が変わるため)。
- `-include` ではなく `include` を使う。ファイルが無い場合にエラーで止まるようにして、コピー漏れにすぐ気付けるようにする。
- `Makefile` をコピーしている 3 か所 (2.2) に、`mk/` のコピーを追加する。

```make
# cp_src_set (467 行)
	cp $2/Makefile $1/
	cp -a $2/mk $1/
```

## 4. 検証方法

検証は次の 5 段階で行う。A〜C は数分で終わる静的な検証、D は実際のビルドとテスト、E は周辺の利用箇所の確認である。

| # | 検証 | 確認すること | 所要時間 |
|---|---|---|---|
| A | include を展開して元と比較 | 分割したテキストが元と完全に一致する | 数秒 |
| B | make の内部データベースの比較 | 変数・ルール・レシピが完全に一致する | 数分 |
| C | 全ターゲットの dry run の比較 | 実行されるコマンドが一致する | 数分 |
| D | 実際のビルドとテスト | 結果ファイルと成否が一致する、`work/` にコピーされた Makefile が動く | test-comp 数回分 |
| E | 周辺の利用箇所 | CI、Docker、`tools/history` が動く | 数分〜 |

### A. include を展開して元と比較 (フェーズ 1 の中心)

`include $(LNS_MK_DIR)mk/...` の行を対応するファイルの中身に置き換え、元の Makefile と比べる。
差分は、追加した `LNS_MK_DIR` の行と `cp -a ... mk` の行だけになるはずである。

```sh
cd src
awk '/^include \$\(LNS_MK_DIR\)mk\//{ f=$2; sub(/^\$\(LNS_MK_DIR\)/, "", f);
       while ((getline l < f) > 0) print l; close(f); next } { print }' Makefile > /tmp/expanded.mk
git show HEAD:src/Makefile | diff - /tmp/expanded.mk
```

これが一致すれば、make が読むテキストは (追加行を除いて) 元と同じであり、2.1 の順序依存も保たれる。

### B. make の内部データベースの比較

`make -pq` で、make が読み込んだ後の変数・ルール・レシピの一覧を出力し、分割前後で比べる。
定義元のファイル名と行番号のコメントや `MAKEFILE_LIST` は必ず変わるので、取り除いてから比べる。

```sh
dump() {  # $@: make に渡す引数
  LC_ALL=C make -pq "$@" 2>&1 |
    grep -v -e "^# makefile (from '" -e "(from '.*', line [0-9]*)" \
            -e '^MAKEFILE_LIST' -e '^LNS_MK_DIR' -e '^# [0-9]* .*files' || true
}
```

次の組み合わせで、分割前後の出力を比べる (読み込み時に評価される条件を一通り通すため)。

| 引数 | 通す条件 |
|---|---|
| `help` | 既定 (`LANG_KIND=lua`) |
| `help LANG_KIND=go WITH_GO=y` | Go 版 (`COM_OPT` への `-langGo` 追加、`check_singleBin` など) |
| `help LANG_KIND=luarocks ENABLE_LUAROCKS=y` | luarocks 版、`LUAROCKS_VER` |
| `help GOARCH=wasm GOTAG=gopherlua` | wasm、gopher-lua |
| `help NO_ACCEPT=y TEST_SERIAL=y SKIP_PROJ_TEST=y LUA51=` | `test-comp2` 内の条件 |
| `install`, `check-runtime-format`, `diff-result`, `check-meta-format` | `MAKECMDGOALS` による条件 (60, 1353, 1369 行) |
| `-C work/2 build_from_d depend_on=y FROM_DIR=../1 TO_DIR=../2` | `depend_on` と `.d` の include (`work/2` は D の後にある状態で) |

確認済み: `make -pq help` を 2 回実行して出力が一致する (出力は決定的である)。

### C. 全ターゲットの dry run の比較

全ターゲットについて `make -n` の出力を分割前後で比べる。ターゲット一覧は B の出力から取る。

```sh
targets=$(LC_ALL=C make -pq help 2>/dev/null |
  awk '/^# Not a target/{skip=1; next} /^[a-zA-Z0-9_.\/-]+:/{ if(!skip){sub(/:.*/,""); print} skip=0 }' |
  grep -v -e '^\.' -e '%' | sort -u)
for t in $targets; do
  timeout 120 make -n -j1 "$t" > dryrun/$t.txt 2>&1
done
```

- `-n` を付けても、`$(MAKE)` を含む行はサブ make として実行される (サブ make にも `-n` が伝わるので、コマンドは表示だけで実行されない)。
- `work/` の状態で出力が変わるので、分割前後で **同じ `work/` の状態** で実行する。
  同じ作業ツリーで、分割前のコミットと分割後のコミットを切り替えて実行するとよい。
- `work/2` などにある Makefile は古いコピーのままなので、C ではコピー処理そのものは検証できない。これは D で確認する。

### D. 実際のビルドとテスト

分割前後で同じコマンドを実行し、成否と結果ファイルを比べる。
`NO_ACCEPT=y` を付けるので、成果物は上書きされない。

```sh
make -C src -j$(nproc) test-comp-with-go NO_ACCEPT=y     # 普段の開発の流れ
make -C src -j$(nproc) test-comp NO_ACCEPT=y             # Lua 版 (CI と同じ流れ)
make -C src go-build NO_TEST=y NO_ACCEPT=y               # Go 版 lnsc のビルド
make -C src go-build NO_ACCEPT=y                         # テスト付き
```

確認すること:

- 分割前後で成否 (終了コード) が同じである。
- `work/*.result` と `work/**/result-testGo*.txt` が一致する
  (時間計測の行 `first:` / `work1:` / `work2:` などは除外して比べる)。
- `work/1`, `work/2`, `work/3`, `work/13`, `work/go`, `work/formatter` などに `mk/` がコピーされていて、
  そこでのサブ make が動いている (2.2)。
- 実行後に `git status` で、コミット対象のファイルが変わっていない (`NO_ACCEPT=y` の確認)。

失敗した場合は `-j1` で再実行して、失敗した箇所を確認する。

### E. 周辺の利用箇所

| 対象 | 確認方法 |
|---|---|
| CI (`.github/workflows/test.yml`) | CI と同じコマンドをローカルで実行する: `make -C src -j6 test-comp LUA=lua5.3 LUA51= SKIP_CHECKING_TESTING=y SKIP_PROJ_TEST=y NO_ACCEPT=y`、`make -C src go-build LUA=lua5.3 LUA51=`、`make -C src LUACHECK=luacheck test-global` |
| Docker (`test/docker/compose/*/build/Dockerfile`) | リポジトリを clone して `make go-build` するので、`mk/` がコミットされていれば動く。`docker build src/test/docker/compose/linux/build` で確認する |
| `tools/history/Makefile.prof` | `src` で `make -f Makefile.prof -n goprof-test-main` を実行し、`LNS_MK_DIR` で `mk/` を見つけられることを確認する |
| 別のディレクトリからの実行 | `make -C src help` と、`cd src && make help` の両方で動くことを確認する |

## 5. 作業手順

| 手順 | 内容 | 検証 |
|---|---|---|
| 0 | 現在の未コミットの変更 (`NO_ACCEPT` など) をコミットし、基準となるコミットを決める | — |
| 1 | 基準のコミットで、B〜D の出力を保存する (`baseline/`) | — |
| 2 | 隔離した作業環境を用意する: `git worktree add ../LuneScript-split`、git 管理外の `src/lnsc` と `src/Makefile.local` をコピーする | — |
| 3 | 分割スクリプトで `mk/*.mk` を生成し、`Makefile` を include だけにする。`LNS_MK_DIR` の追加と、コピー 3 か所の修正を行う | A |
| 4 | B, C を実行して基準と比べる | B, C |
| 5 | D を実行して基準と比べる | D |
| 6 | E を実行する | E |
| 7 | 1 つのコミットにまとめる (手作業による変更が入らないように、分割スクリプトもコミットメッセージか `docs/` に残す) | — |
| 8 | `NO_ACCEPT` なしで普段どおり `test-comp-with-go` を実行し、成果物の更新まで含めて動くことを確認する | — |

- 分割は **スクリプトで機械的に行う** (`sed -n 'N,Mp' Makefile > mk/xxx.mk`)。手作業でコピー & ペーストしない。
  境界の行番号を表にしてスクリプトに渡すと、抜け (どのファイルにも入らない行) や重複を検証できる
  (行範囲がすき間なく、重ならずに 1〜最終行を覆っているかを確認する)。
- 元に戻す場合は、そのコミットを `git revert` すればよい (1 コミットにまとめるため)。

## 6. フェーズ 2 (任意の整理) の注意

フェーズ 1 の後で、内容的に不自然な配置 (`luarocks-check` が `selfhost.mk` にあるなど) を直す場合のルール。

- 移してよいもの: **レシピだけを持つターゲット**、**`define`** (切り替えのない `define`)、
  **`=` で定義され、レシピ内でしか使われない変数**。これらは読み込み順に依存しない。
- 移す場合に注意が必要なもの: **`ifeq` / `ifdef` の判定**、**`:=` と `$(shell ...)`**、**include**、
  **最初のターゲット**、**`WITH_GO` などで切り替わる `define`**。これらを移すと、評価のタイミングが変わる可能性がある。
- 2.1 の `RUNTIME_ORG_VER` の箇所のように、現状の動作が順序に依存している箇所は、
  移動で直ってしまう (= 動作が変わる) 可能性がある。直す場合は別のコミットにする。
- 移動ごとに B〜D を実行する。A (テキストの一致) は使えないので、B (データベースの一致) が主な検証になる。

## 7. 分割中に見つかった問題 (分割とは別に扱う)

| 箇所 | 内容 |
|---|---|
| `diff-result` (1224 行) | `RUNTIME_ORG_VER` / `RUNTIME_NEW_VER` を定義前に参照しているため、ランタイムのバージョンによらず `.meta` の比較が常に行われる。意図した動作かどうかを確認する |
| `test/proj` のテスト (`test-org-proj`, `test-new-proj`) | `-j` で実行すると、`test/proj` のサブ make が「副次 make で -j1 を強制指定しました」という警告を出し、さらに `Mod1`〜`Mod4` の出力順序が実行ごとに変わる。そのため `diff test/result/test1.result work/test2.result` で失敗し、`make -j24 test-comp-with-go` が分割前から失敗する (`SKIP_PROJ_TEST=y` なら成功する) |

## 8. 実施結果 (フェーズ 1)

### 8.1 分割後の構成

分割は [split_makefile.py](split_makefile.py) で機械的に行った (`python3 split_makefile.py src`)。

| ファイル | 分割前の行 | 内容 |
|---|---|---|
| `Makefile` | — | `LNS_MK_DIR` の定義と include |
| `mk/config.mk` | 1–131 | Lua / lnsc の設定、`depend_on` 時のビルド、`Makefile.local`、`.PHONY` |
| `mk/help.mk` | 132–170 | `help`, `exec` |
| `mk/exec.mk` | 171–214 | lnsc 実行用の共通マクロ (`exeLune`, `cmd_exec_trans` など) |
| `Makefile` (本体) | 215–263 | `LNS_SRCS` (2.5 のため本体に残す) |
| `mk/selfhost.mk` | 264–491 | セルフホストビルド、`build-testcode`、`luarocks-check`、`cp_src_set` |
| `mk/test-comp.mk` | 492–610 | `test-global`, `test-comp`, `--testing` 関連 |
| `mk/test-cases.mk` | 611–813 | `test-etc`、エラーテスト、補完テスト |
| `mk/test-stages.mk` | 814–1238 | `test-org`, `test-new*`, `test-comp2*`, `diff-result` |
| `mk/test-lua51.mk` | 1239–1305 | Lua 5.1 クロスコンパイルテスト |
| `mk/release.mk` | 1306–1430 | `accept-package`、Docker テスト、形式確認、`build-glue`、`install` |
| `mk/conv-c.mk` | 1431–1605 | C 変換関連 |
| `mk/go.mk` | 1606–1959 | Go のビルド設定、`test-go`、`go-build` |
| `mk/prof.mk` | 1960–2099 | `test-govslua`、`goprof`、`luaprof` |
| `mk/tools.mk` | 2100–2179 | `test-reference`、`update-lnstags`、フォーマッタ |

元の Makefile に対する変更は、次の追加だけである (削除・変更した行はない)。

- `LNS_MK_DIR` の定義と include 行、説明コメント
- `mk/` をコピーする 3 行 (`cp_src_set`、`go-build`、`test-formatter-prepare`)

### 8.2 検証結果

すべて `git worktree` で隔離した作業ツリーで行った (Go 1.27.1, GNU Make 4.4.1, Lua 5.3.6)。

| # | 検証 | 結果 |
|---|---|---|
| A | include を展開して元と比較 | 削除行 0。追加行は 8.1 の追加分だけ |
| B | `make -pq` の比較 (11 通りの条件) | 一致 (`mk/*.mk` 自体の項目・統計行・項目の並び順を除く) |
| C | 全 98 ターゲット × 2 条件 (既定 / `LANG_KIND=go WITH_GO=y`) の `make -n` の比較 | 一致 (追加したコピー行、エラー位置の表示、時間計測、並列実行による出力の混ざりを除く) |
| — | `bind.lns` が生成する `binding.go` の比較 | 一致 (`go:embed` 44 件) |
| D | 実ビルド 5 種類の比較 (下表) | 終了コードと結果ファイルが一致 (`DIFF_IGNORE` の時間計測行を除く) |
| E | `make -C src`、`src` 内での `make`、別ディレクトリから `make -f`、`tools/history/Makefile.prof` からの include | いずれも `mk/` を正しく見つける |

D で実行したコマンド (いずれも `NO_ACCEPT=y`、実行後に `git status` で成果物が変わっていないことも確認):

| コマンド | 分割前 | 分割後 |
|---|---|---|
| `make -j24 test-comp-with-go` | 失敗 (7 章の `test/proj` の問題) | 同じ箇所で失敗 |
| `make -j24 test-comp-with-go SKIP_PROJ_TEST=y` | 成功 | 成功 |
| `make -j24 test-comp LUA=lua5.3 LUA51= SKIP_CHECKING_TESTING=y SKIP_PROJ_TEST=y` (CI と同じ) | 成功 | 成功 |
| `make go-build NO_TEST=y` | 成功 | 成功 |
| `make go-build` (テスト付き) | 成功 | 成功 |
| `make -j1 test-comp-with-go` (参考) | — | 成功。`-j24` での失敗は並列実行が原因であることを確認 |

検証の過程で、`LNS_SRCS` を `mk/` に移した最初の案は D で失敗した (2.5)。
`LNS_SRCS` を本体に残す形に直し、A〜E をすべてやり直した。
