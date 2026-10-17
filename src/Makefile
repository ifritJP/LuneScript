# -*- makefile -*-
# 実体は mk/*.mk に分割している。include の順序は元の Makefile の記述順と同じで、
# 順序を変えると読み込み時に評価される条件 (ifeq/ifdef/:=) の結果が変わるので注意すること。

# この Makefile のあるディレクトリ。
# 他のディレクトリから include されても mk/ を見つけられるように、他の include より前で求める。
LNS_MK_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

# Lua / lnsc の設定、depend_on 時のビルド、Makefile.local、.PHONY
include $(LNS_MK_DIR)mk/config.mk
# help, exec
include $(LNS_MK_DIR)mk/help.mk
# lnsc 実行用の共通マクロ (exeLune, cmd_exec_trans など)
include $(LNS_MK_DIR)mk/exec.mk

# セルフホストを構成するモジュール一覧
# misc/go/binding/bind.lns が、この Makefile から 'LNS_SRCS +=' の行を読み取って
# Go 版 lnsc に埋め込むモジュールを決めている。そのため、この定義は mk/ に移さないこと。
LNS_SRCS :=
LNS_SRCS += base/Async
LNS_SRCS += base/Runner
LNS_SRCS += base/Str
LNS_SRCS += base/Types
LNS_SRCS += base/Ver
LNS_SRCS += base/LuaMod
LNS_SRCS += base/LuneControl
LNS_SRCS += base/Meta
LNS_SRCS += base/Testing
LNS_SRCS += base/TestingMacro
LNS_SRCS += base/Code
LNS_SRCS += base/Log
LNS_SRCS += base/LuaVer
LNS_SRCS += base/Depend
LNS_SRCS += base/Util
LNS_SRCS += base/frontInterface
LNS_SRCS += base/Writer
LNS_SRCS += base/AsyncTokenizer
LNS_SRCS += base/Tokenizer
LNS_SRCS += base/Json
#LNS_SRCS += base/MetaLoader
LNS_SRCS += base/Ast
LNS_SRCS += base/Option
LNS_SRCS += base/GoMod
#LNS_SRCS += base/bind
LNS_SRCS += base/Nodes
LNS_SRCS += base/NodeIndexer
LNS_SRCS += base/Formatter
LNS_SRCS += base/Macro
LNS_SRCS += base/TransUnitIF
LNS_SRCS += base/BuiltinTransUnit
LNS_SRCS += base/Builtin
LNS_SRCS += base/AstInfo
LNS_SRCS += base/Import
LNS_SRCS += base/TransUnit
LNS_SRCS += base/DependLuaOnLns
LNS_SRCS += base/convLua
#LNS_SRCS += base/ConvNodes
LNS_SRCS += base/dumpNode
#LNS_SRCS += base/convCC
LNS_SRCS += base/convGo
LNS_SRCS += base/convPython
LNS_SRCS += base/WasmBin
LNS_SRCS += base/WasmRt
LNS_SRCS += base/convWasm
LNS_SRCS += base/OutputDepend
LNS_SRCS += base/Converter
LNS_SRCS += base/glueFilter
LNS_SRCS += base/front
LNS_SRCS := $(addprefix lune/,$(LNS_SRCS))

# セルフホストビルド、build-testcode、luarocks-check、cp_src_set
include $(LNS_MK_DIR)mk/selfhost.mk
# test-global, test-comp, --testing 関連
include $(LNS_MK_DIR)mk/test-comp.mk
# test-etc、エラーテスト、補完テスト
include $(LNS_MK_DIR)mk/test-cases.mk
# test-org, test-new*, test-comp2*, diff-result
include $(LNS_MK_DIR)mk/test-stages.mk
# Lua 5.1 クロスコンパイルテスト
include $(LNS_MK_DIR)mk/test-lua51.mk
# accept-package、Docker テスト、形式確認、glue、install
include $(LNS_MK_DIR)mk/release.mk
# C 変換関連
include $(LNS_MK_DIR)mk/conv-c.mk
# Go ビルド設定、test-go、go-build
include $(LNS_MK_DIR)mk/go.mk
# test-govslua、goprof、luaprof
include $(LNS_MK_DIR)mk/prof.mk
# test-reference、lnstags、フォーマッタ
include $(LNS_MK_DIR)mk/tools.mk
# wasm (WasmGC) への直接変換のテスト
include $(LNS_MK_DIR)mk/wasm.mk
