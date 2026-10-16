#!/usr/bin/env python3
"""src/Makefile を連続した行範囲で mk/*.mk に機械的に分割する (フェーズ 1)。

usage: split_makefile.py <src ディレクトリ>

- 各ファイルの開始位置は、アンカー (行全体の正規表現) で決める。
  アンカー行の直前に空行を挟まずに続くコメント行は、アンカー側のファイルに含める。
- 行範囲がすき間なく・重ならずに 1 行目から最終行までを覆っていることを確認する。
- 元の Makefile を、LNS_MK_DIR の定義と include だけのファイルに置き換える。
- Makefile をコピーしている 3 か所に、mk/ のコピーを追加する。
"""
import os
import re
import sys

# (ファイル名, 開始行のアンカー, 内容)
PARTS = [
    ("config.mk", None, "Lua / lnsc の設定、depend_on 時のビルド、Makefile.local、.PHONY"),
    ("help.mk", r"help:", "help, exec"),
    ("exec.mk", r"ifeq \(\$\(WITH_GO\),y\)", "lnsc 実行用の共通マクロ (exeLune, cmd_exec_trans など)"),
    # LNS_SRCS は misc/go/binding/bind.lns が Makefile をテキストとして読んで
    # 'LNS_SRCS +=' の行を探すので、分割せずに Makefile 本体に残す (None = 本体に残す)。
    (None, r"LNS_SRCS\s*:=", "セルフホストを構成するモジュール一覧"),
    ("selfhost.mk", r"define selfhost_buildOp_batch", "セルフホストビルド、build-testcode、luarocks-check、cp_src_set"),
    ("test-comp.mk", r"test-global:", "test-global, test-comp, --testing 関連"),
    ("test-cases.mk", r"test-etc:", "test-etc、エラーテスト、補完テスト"),
    ("test-stages.mk", r"test-org:", "test-org, test-new*, test-comp2*, diff-result"),
    ("test-lua51.mk", r"define setup-lua51-test", "Lua 5.1 クロスコンパイルテスト"),
    ("release.mk", r"accept-package:", "accept-package、Docker テスト、形式確認、glue、install"),
    ("conv-c.mk", r"test-ccc:", "C 変換関連"),
    ("go.mk", r"GO_RUNTIME = lune/base/runtime_go", "Go ビルド設定、test-go、go-build"),
    ("prof.mk", r"define test-govslua-comp", "test-govslua、goprof、luaprof"),
    ("tools.mk", r"test-reference:", "test-reference、lnstags、フォーマッタ"),
]

# Makefile をコピーしている箇所: (対象ファイル, 元の行, 直後に追加する行)
COPY_PATCHES = [
    ("selfhost.mk", "\tcp $2/Makefile $1/\n", "\tcp -a $2/mk $1/\n"),
    ("go.mk", "\tcp Makefile $(GO_WORK_DIR)\n", "\tcp -a mk $(GO_WORK_DIR)/\n"),
    ("tools.mk", "\tcp Makefile work/formatter\n", "\tcp -a mk work/formatter/\n"),
]

HEADER = """\
# -*- makefile -*-
# 実体は mk/*.mk に分割している。include の順序は元の Makefile の記述順と同じで、
# 順序を変えると読み込み時に評価される条件 (ifeq/ifdef/:=) の結果が変わるので注意すること。

# この Makefile のあるディレクトリ。
# 他のディレクトリから include されても mk/ を見つけられるように、他の include より前で求める。
LNS_MK_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

"""


NOTE_INLINE = """\
# misc/go/binding/bind.lns が、この Makefile から 'LNS_SRCS +=' の行を読み取って
# Go 版 lnsc に埋め込むモジュールを決めている。そのため、この定義は mk/ に移さないこと。
"""


def find_starts(lines):
    starts = [0]
    for name, anchor, _ in PARTS[1:]:
        hits = [i for i, l in enumerate(lines) if re.fullmatch(anchor, l.rstrip("\n"))]
        hits = [i for i in hits if i > starts[-1]]
        if not hits:
            sys.exit(f"anchor not found: {name}: {anchor}")
        i = hits[0]
        while i - 1 > starts[-1] and lines[i - 1].startswith("#"):
            i -= 1
        starts.append(i)
    return starts


def main():
    src = sys.argv[1]
    path = os.path.join(src, "Makefile")
    with open(path, encoding="utf-8") as f:
        lines = f.readlines()
    if lines[0].startswith("# -*- makefile -*-"):
        sys.exit("already split")

    starts = find_starts(lines)
    ends = starts[1:] + [len(lines)]
    # 範囲の検証: すき間なく、重ならずに全行を覆う
    assert starts[0] == 0 and ends[-1] == len(lines)
    for s, e in zip(starts, ends):
        assert s < e
    assert all(ends[k] == starts[k + 1] for k in range(len(starts) - 1))
    assert sum(e - s for s, e in zip(starts, ends)) == len(lines)

    chunks = {}
    for (name, _, desc), s, e in zip(PARTS, starts, ends):
        chunks[name or "<Makefile>"] = "".join(lines[s:e])
        print(f"{name or '<Makefile>':16s} {s + 1:5d}-{e:5d} ({e - s:4d} lines)  {desc}")

    for name, old, add in COPY_PATCHES:
        if chunks[name].count(old) != 1:
            sys.exit(f"copy site not unique: {name}: {old!r}")
        chunks[name] = chunks[name].replace(old, old + add)

    os.makedirs(os.path.join(src, "mk"), exist_ok=True)
    for name, text in chunks.items():
        if name == "<Makefile>":
            continue
        with open(os.path.join(src, "mk", name), "w", encoding="utf-8") as f:
            f.write(text)
    with open(path, "w", encoding="utf-8") as f:
        f.write(HEADER)
        for name, _, desc in PARTS:
            if name is None:
                f.write(f"\n# {desc}\n" + NOTE_INLINE + chunks["<Makefile>"])
            else:
                f.write(f"# {desc}\ninclude $(LNS_MK_DIR)mk/{name}\n")


if __name__ == "__main__":
    main()
