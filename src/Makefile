LUA=lua5.3
#LUA=lua5.2


LUACHECK=~/.luarocks/bin/luacheck

PKG_NAME=lua-5.3
LUA_LIB=$(shell pkg-config --libs $(PKG_NAME))
LUA_INC_DIR=$(shell pkg-config --cflags-only-I $(PKG_NAME) | sed 's/^-I//')
ifneq ($(LUA_INC_DIR),)
LUA_CFLAGS=-I$(LUA_INC_DIR)
endif

PKG_NAME_51=lua-5.1
LUA51=lua5.1
LUA51_INC=$(shell pkg-config --cflags-only-I $(PKG_NAME_51) | sed 's/^-I//')


LUAROCKS=luarocks

ifdef LOG
LOG_OPT=--log info
endif

LNSC_LUA=$(LUA) lune/base/base.lua
LNSC_GO=./lnsc
LNSC_LUAROCKS=~/.luarocks/bin/lnsc

# make のパラメータで言語を切り替える場合、
# LNSC_CMD や LNSC を直接変更するのではなく LANG_KIND を変更する。
LANG_KIND=lua

ifeq ($(LANG_KIND),lua)
LNSC_CMD=$(LNSC_LUA)
endif
ifeq ($(LANG_KIND),go)
LNSC_CMD=$(LNSC_GO)
endif
ifeq ($(LANG_KIND),luarocks)
LNSC_CMD=$(LNSC_LUAROCKS)
endif

ifeq ($(LNSC_CMD),)
$(error "LANG_KIND is illegal -- " $(LNSC_CMD) )
endif

LNSC=LUA_PATH=./?.lua $(LNSC_CMD)

ifdef ENABLE_LUAROCKS
ifneq ($(LUA51),)
ENABLE_LUA51:=$(shell which $(LUA51))
endif
ifndef NO_LUAROCKS
LUAROCKS_VER:=$(shell $(LUAROCKS) | grep 'Lua version' | awk '//{print $$3}')
endif
endif

ifneq ($(findstring $(MAKECMDGOALS),"install" "uninstall"),)
# LUA_DIR_SCRIPT='for path in string.gmatch( package.path, "[^;]+" ) do if not path:find( "^%./" ) then print( (path:gsub( "/%?%.lua", "" )) ); break end end'
# LUA_MOD_DIR=$(shell echo $(LUA_DIR_SCRIPT) | $(LUA) 2> /dev/null)
include lune.mk
endif

INSTALL=/usr/bin/install -c

NDIR=$(shell pwd)

.PHONY: help exec build build_from_d luarocks-check luarocks-check-sub
.PHONY: release-check build test-save build-countup test-comp
.PHONY: test-registered-rockspec test-etc test-error-all test-error test-org
.PHONY: test-new test-parallel test-comp2 check-rockspec diff-result
.PHONY: test-lua51 test-lua51-sub accept-package check-runtime-format
.PHONY: check-meta-format build-glue check_lua_mod_dir install uninstall
.PHONY: test-ccc clean-ccc build-miniC test-miniC test-subC build-cc test-formatter
.PHONY: test-installed-cmd test-on-docker


help:
	@echo usage:
	@echo "  make install"
	@echo "  make test-comp"
	@echo "  make test-global"
	@echo "  make release-check"
	@echo "  make accept-package"
	@echo "  make build-glue [LUA_INC_DIR=path]"
	@echo "  make build-cc"
	@echo "  make test-miniC"
	@echo "  make test-subC"
	@echo "  make test-testC"
	@echo "  make test-formatter"
	@echo "  make test-mkbuiltin"
	@echo "  make build-testcase"
	@echo "  make test-tutorial"
	@echo "  make test-installed-cmd"
	@echo "  make test-on-docker"
	@echo "  make test-reference"
	@echo "  make test-go"
	@echo "  make test-miniGo [NO_LUA_DIFF=y]"
	@echo "  make scratch-go"
	@echo "  make test-tutorial-go"
	@echo "  make go-build"
	@echo "  make setup-pkg-lua"


exec:
	$(LUA) lune/base/base.lua test/test.lns token
	$(LUA) lune/base/base.lua test/test.lns ast
	$(LUA) lune/base/base.lua test/test.lns LUA

EXE_OPT := --valid-luaval
COM_OPT := --use-ipairs -Werror --valid-luaval
ifeq ($(WITH_GO),y)
COM_OPT += -langGo
endif


ifdef depend_on
DEPENDS := $(shell find lune -iname '*.d')
DEPENDS_META := $(patsubst %.d,%.meta,$(DEPENDS))

include $(DEPENDS)

$(DEPENDS_META):

%.meta: %.lns
	(cd $(FROM_DIR); $(LNSC) $< SAVE $(TO_DIR) \
			$(BYTECOMP) $(COM_OPT) $(LOG_OPT) $(LNS_OPT))

build_from_d: $(DEPENDS_META)
ifeq ($(WITH_GO),y)
	$(call go-build-lns,.)
endif

endif



define cutMeta
	(cd work/$2/lune/base; cat $1.lua | awk '/^----- meta/{META++;}; //{if (META != 1) {print $$0}} ' > $1.cut.lua )
endef


define exeLune
	(cd $2; $1 $3 $4 $5)
endef

define exeFront
	(cd $1; LUA_PATH=./?.lua $(LUA) -e 'print( require( "lune.base.front" ).convertLnsCode2LuaCode( "print( 1 + 10 );", "lua" ) )')
endef



define cmd_exec_trans
	$(call exeLune,$1,$2,lune/$5.lns SAVE $3 $4 $(BYTECOMP) $(COM_OPT) $(LOG_OPT))
endef

define cmd_exec_trans_dep
	$(call exeLune,$1,$2,lune/$5.lns SAVE $3 --depends $3/lune/$5.d $4 $(BYTECOMP) $(COM_OPT) $(LOG_OPT))
endef

define cmd_exec_ast
	@echo -------- ast $5 ---------- > /dev/stderr
	$(call exeLune,$1,$2,lune/$5.lns ast $3 $(LOG_OPT) $(EXE_OPT) $4)
endef


LNS_SRCS :=
LNS_SRCS += base/Async
LNS_SRCS += base/Str
LNS_SRCS += base/Types
LNS_SRCS += base/Ver
LNS_SRCS += base/LuaMod
LNS_SRCS += base/LuneControl
LNS_SRCS += base/Meta
LNS_SRCS += base/Testing
LNS_SRCS += base/Code
LNS_SRCS += base/Log
LNS_SRCS += base/LuaVer
LNS_SRCS += base/Depend
LNS_SRCS += base/Util
LNS_SRCS += base/frontInterface
LNS_SRCS += base/Writer
LNS_SRCS += base/AsyncParser
LNS_SRCS += base/Parser
LNS_SRCS += base/Json
#LNS_SRCS += base/MetaLoader
LNS_SRCS += base/Ast
LNS_SRCS += base/Option
#LNS_SRCS += base/bind
LNS_SRCS += base/Nodes
LNS_SRCS += base/Formatter
LNS_SRCS += base/Macro
LNS_SRCS += base/TransUnit
LNS_SRCS += base/DependLuaOnLns
LNS_SRCS += base/convLua
#LNS_SRCS += base/ConvNodes
LNS_SRCS += base/dumpNode
LNS_SRCS += base/convCC
LNS_SRCS += base/convGo
LNS_SRCS += base/OutputDepend
LNS_SRCS += base/glueFilter
LNS_SRCS += base/front
LNS_SRCS += Util
LNS_SRCS := $(addprefix "lune/",$(LNS_SRCS))

define selfhost_buildOp_batch
	echo $(LNS_SRCS) | sed 's/$$/.lns\n/g' | sed 's/ /.lns\n/g' | \
		(cd $1; $2 @- SAVE $3 $4 $(BYTECOMP) --package lnsc $(COM_OPT) $(LOG_OPT))
endef


define selfhost_buildOp
	$(call $5,$1,$2,$3,$4,base/Async)
	$(call $5,$1,$2,$3,$4,base/Str)
	$(call $5,$1,$2,$3,$4,base/Types)
	$(call $5,$1,$2,$3,$4,base/Ver)
	$(call $5,$1,$2,$3,$4,base/LuaMod)
	$(call $5,$1,$2,$3,$4,base/LuneControl)
	$(call $5,$1,$2,$3,$4,base/Meta)
	$(call $5,$1,$2,$3,$4,base/Testing)
	$(call $5,$1,$2,$3,$4,base/Code)
	$(call $5,$1,$2,$3,$4,base/Log)
	$(call $5,$1,$2,$3,$4,base/LuaVer)
	$(call $5,$1,$2,$3,$4,base/Depend)
	$(call $5,$1,$2,$3,$4,base/Util)
	$(call $5,$1,$2,$3,$4,base/frontInterface)
	$(call $5,$1,$2,$3,$4,base/Writer)
	$(call $5,$1,$2,$3,$4,base/AsyncParser)
	$(call $5,$1,$2,$3,$4,base/Parser)
	$(call $5,$1,$2,$3,$4,base/Json)
	$(call $5,$1,$2,$3,$4,base/Ast)
	$(call $5,$1,$2,$3,$4,base/Option)
	$(call $5,$1,$2,$3,$4,base/Nodes)
	$(call $5,$1,$2,$3,$4,base/Formatter)
	$(call $5,$1,$2,$3,$4,base/Macro)
	$(call $5,$1,$2,$3,$4,base/TransUnit)
	$(call $5,$1,$2,$3,$4,base/DependLuaOnLns)
	$(call $5,$1,$2,$3,$4,base/convLua)
	$(call $5,$1,$2,$3,$4,base/dumpNode)
	$(call $5,$1,$2,$3,$4,base/convCC)
	$(call $5,$1,$2,$3,$4,base/convGo)
	$(call $5,$1,$2,$3,$4,base/OutputDepend)
	$(call $5,$1,$2,$3,$4,base/glueFilter)
	$(call $5,$1,$2,$3,$4,base/front)
	$(call $5,$1,$2,$3,$4,Util)
endef

define mk_runtime
	$(call exeLune,$1,$2,$3/$4,-mklunemod,lune/base/runtime.lua) > /dev/null
	cat $2/$3/$(VER_PATH) | $(RUNTIME_VER_PICK) | \
		awk '//{printf( "$2/$3/lune/base/runtime%s.lua",$$0);}' | \
		xargs cp $2/$3/lune/base/runtime.lua 
endef

define selfhost_build
	$(call selfhost_buildOp,$1,$2,$3,$4,cmd_exec_trans)
	$(call mk_runtime,$1,$2,$3,$4)
endef

define selfhost_build_dep
	$(call selfhost_buildOp,$1,$2,$3,$4,cmd_exec_trans_dep)
	$(call mk_runtime,$1,$2,$3,$4)
endef

define dump_ast
	$(call selfhost_buildOp,$1,$2,$3,$4,cmd_exec_ast)
endef

tgt_dump_ast:
	@echo $(LNSC) > /dev/stderr
	@$(call dump_ast,$(LNSC),.,,)



define build-testcode
	$(call exeLune,$1,$2,test/macro.lns,$3 save,$4) > /dev/null
	$(call exeLune,$1,$2,test/Sub2.lns,$3 save,$4) > /dev/null
	$(call exeLune,$1,$2,test/Sub3.lns,$3 save,$4) > /dev/null
	$(call exeLune,$1,$2,test/Sub3_2.lns,$3 save,$4) > /dev/null
	$(call exeLune,$1,$2,test/Sub4.lns,$3 save,$4) > /dev/null
	$(call exeLune,$1,$2,test/Sub5.lns,$3 save,$4) > /dev/null
	$(call exeLune,$1,$2,test/Sub6.lns,$3 save,$4) > /dev/null
	$(call exeLune,$1,$2,test/Sub8.lns,$3 save,$4) > /dev/null
	$(call exeLune,$1,$2,test/Sub10.lns,$3 save,$4) > /dev/null
	$(call exeLune,$1,$2,test/simple.lns,$3 save,$4) > /dev/null
	$(call exeLune,$1,$2,test/Class1.lns,$3 save,$4) > /dev/null
	$(call exeLune,$1,$2,test/Class2.lns,$3 save,$4) > /dev/null
	$(call exeLune,$1,$2,test/Class3.lns,$3 save,$4) > /dev/null
	$(call exeLune,$1,$2,test/Class4.lns,$3 save,$4) > /dev/null
	$(call exeLune,$1,$2,test/start.lns,$3 save,$4) > /dev/null
	$(call exeLune,$1,$2,test/mapping.lns,$3 save,$4) > /dev/null
	$(call exeLune,$1,$2,test/funcSym.lns,$3 save,$4) > /dev/null
	$(call exeLune,$1,$2,test/test2/test2/test.lns,$3 save,$4) > /dev/null
endef

# meta 情報の更新時に正常に解析できるかテスト。
define test-meta-update
	-(cd $2; rm -f test/ext/*.lua test/ext/*.meta )
	$(call exeLune,$1,$2,test/ext/sub3.lns,SAVE,)
	$(call exeLune,$1,$2,test/ext/sub2.lns,SAVE,)
	(cd $2; rm test/ext/sub3.lua test/ext/sub3.meta )
	$(call exeLune,$1,$2,test/ext/sub1.lns,SAVE,)
endef

luarocks-check:
	$(MAKE) luarocks-check-sub ENABLE_LUAROCKS=y LANG_KIND=luarocks

# luarocks でインストールして動作を確認する
#
# - 一旦アンインストール
# - インストール
# - 実行
# - アンインストール
define test-inst-from-luarocks
	-($(LUAROCKS) remove $1 --local)
	cat ../$1-main-1.rockspec | sed 's@"src.*/lune/@"lune/@g' \
			| sed 's@src/lnsc.lua@lnsc.lua@g' > work/8/$1-main-1.rockspec
	cp -a lnsc.lua work/8
	cp -a ../docs work/8
	(cd work/8; $(LUAROCKS) make $1-main-1.rockspec --local)

	$(LNSC)
	(cd work/8; $(LNSC) test/start.lns exe $(EXE_OPT)) > work/test8.result

	($(LUAROCKS) remove $1 --local)
endef

luarocks-check-sub:
	mkdir -p work/8
ifeq ($(LUAROCKS_VER),5.1)
ifneq ($(ENABLE_LUA51),)
	cp -a work/6/lune work/8
	$(call setup-lua51-test,work/8)
	$(call test-inst-from-luarocks,lunescript51)
	diff work/test7.result work/test8.result
else
	echo skip luarocks test
endif
else
	cp -a work/3/lune work/8
	cp -a work/3/test work/8
	$(call test-inst-from-luarocks,lunescript)

	diff work/test.result work/test8.result
endif

release-check:
	$(MAKE) test-comp LUA_INC_DIR=../../lctags/external/lua/lua-5.3.4/src/
	$(MAKE) test-comp LUA=lua5.2 LUA_INC_DIR=../../lctags/external/lua/lua-5.2.4/src/
#	$(MAKE) test-comp LUA=lua5.1 LUA_INC_DIR=/usr/include/lua5.1/
	$(MAKE) test-registered-rockspec


build:
	$(call selfhost_build,$(LNSC),$(LNSDIR),.)

test-save:
	$(call test-meta-update,$(LNSC),.)


BUILD_COUNTER=build.counter
build-countup:
	@echo $$(($$(head -n 1 $(BUILD_COUNTER)) + 1)) > $(BUILD_COUNTER)
	@echo build -- $$(cat $(BUILD_COUNTER))



define mk_lune_proj_dir
	mkdir -p $1/lune/base
	cp $(BUILD_COUNTER) $1
endef


# 第1引数のディレクト直下にある、第2引数のパターンファイルを、
# 第3引数のディレクトリにコピーする
define cp_src_set_sub
	@find $1 -maxdepth 1 -type f \
		-iname $2 -exec cp {} $3 \;
endef

# lns のビルドに必要なソース群をコピーする
#
# $1: コピー先ディレクトリ
# $2: コピー元ディレクトリ
define cp_src_set
	mkdir -p $1
	cp $2/Makefile $1/
	$(call mk_lune_proj_dir,$1)
	$(call cp_src_set_sub,$2/lune/base,'base*.lua',$1/lune/base)
	$(call cp_src_set_sub,$2/lune/base,'runtime*.lua',$1/lune/base)
	$(call cp_src_set_sub,$2/lune/base,'*.lns',$1/lune/base)
	$(call cp_src_set_sub,$2/lune/base,'*.d',$1/lune/base)
#	@cp lune/base/base*.lua lune/base/*.lns lune/base/*.d $1/lune/base
	$(call cp_src_set_sub,$2/lune,'*.lns',$1/lune)
	$(call cp_src_set_sub,$2/lune,'*.d',$1/lune)
	cp $2/lune.js $1

	cp -au $2/lune/base/runtime_go $1/lune/base/runtime_go
endef

test-global:
# 意図しないシンボルが global になっていないか調べる
	find lune/base/ -iname '*.lua' | \
		xargs $(LUACHECK) | grep global | \
		grep -v _lune | \
		grep -v -e 'field.*unpack.*of.*global.*table' && exit 1 || echo ok
	find work/3/test/ -iname '*.lua' | \
		xargs $(LUACHECK) | grep global | \
		grep -v _lune | \
		grep -v -e 'field.*unpack.*of.*global.*table' && exit 1 || echo ok

# LuneScript 自身のソースを Lua に変換する。
#
# 次の手順で実行する。
# - lune/base/ 内にある Lua に変換済みの LuneScript (以降orginalとする)を使用して、
#   lune/base/ 内の .lns ファイルを .lua に変換し、 work/1/lune/base に格納する。
# - orginal を利用してテスト用コードを実行し、
#   出力を work/test.result, work/test1.result に格納する
# - work/1/lune/base に格納した新しい LuneScript を利用して、
#   再度 .lns ファイルを .lua に変換し、work/2/lune/base に格納する。
# - work/2/lune/base に格納した新しい LuneScript を利用して、
#   再度 .lns ファイルを .lua に変換し、work/3/lune/base に格納する。
#   これを candidate とする
# - work/2 と work/3 を比較し、同じになっているか確認する。
# - candidate を利用してテスト用コードを実行し、
#   出力を work/test2.result, work/test3.result に格納する
# - work/test.result, work/test1.result と、
#   work/test2.result, work/test3.result を比較し、同じになっているか確認する。
#
#
# - 上記の比較が全て同じなら、 candidate を orginal とする。
test-comp:
	@echo start to build 
	@LANG= date

	$(MAKE) build-countup
	rm -rf work/*
	mkdir -p work/glue
	$(call mk_lune_proj_dir,work/1)
	mkdir -p work/1/test
	mkdir -p work/newglue
	$(call cp_src_set,work/1,.)
ifeq ($(LANG_KIND),lua)
# lua は batch 処理が何故か遅いので、通常処理で行なう
	$(call selfhost_build_dep,$(LNSC),.,work/1,)
ifeq ($(WITH_GO),y)
	$(call binding-lns-go,work/1)
	$(call go-build-lns,work/1)
endif
else
	$(call selfhost_buildOp_batch,.,/usr/bin/time -f 'first:%e/%U/%S' $(LNSC_GO),work/1,--depends work/1/ )
ifeq ($(WITH_GO),y)
	$(call binding-lns-go,work/1)
	$(call go-build-lns,work/1)
endif
	$(call mk_runtime,$(LNSC),work/1,.,)
endif


	cp -au test work/1
	$(call build-testcode,$(LNSC),.,,work/1)
	$(call cp_src_set,work/2,work/1)
	$(call cp_src_set,work/3,work/1)
	$(call cp_src_set,work/5,work/1)
	$(MAKE) check-runtime-format
	$(MAKE) test-comp2


	@echo finish to build 
	@LANG= date

test-comp-with-go:
	$(MAKE) test-comp LANG_KIND=go WITH_GO=y LANG_KIND=go


test-registered-rockspec:
	curl https://luarocks.org/manifests/ifritjp/lunescript-main-1.rockspec -o work/rockspec
	@diff ../lunescript-main-1.rockspec work/rockspec || echo "====== need upload rockspec ======" && echo "ok"


test-etc:
# 存在しないファイルを指定した時に、不正なファイルが作成されないことを確認する
	-@$(call exeLune,$(LNSC),$(TEST_DIR),test/error/__.lns,SAVE,2>&1) > /dev/null
	@echo -- not exist file access --
	@(cd $(TEST_DIR); find test/error -iname '__.*')


TEST_CASE += nilable1.lns
TEST_CASE += access1.lns
TEST_CASE += access2.lns
TEST_CASE += access3.lns
TEST_CASE += access4.lns
TEST_CASE += access5.lns
TEST_CASE += list.lns
TEST_CASE += map.lns
TEST_CASE += set.lns
TEST_CASE += scope1.lns
TEST_CASE += scope2.lns
TEST_CASE += scope3.lns
TEST_CASE += func.lns
TEST_CASE += class.lns
TEST_CASE += class2.lns
TEST_CASE += class3.lns
TEST_CASE += class4.lns
TEST_CASE += class5.lns
TEST_CASE += class6.lns
TEST_CASE += provide1.lns
TEST_CASE += provide2.lns
TEST_CASE += provide3.lns
TEST_CASE += shadowing.lns
TEST_CASE += enum.lns
TEST_CASE += eof.lns
TEST_CASE += mutable.lns
TEST_CASE += var.lns
TEST_CASE += arg.lns
TEST_CASE += arg2.lns
TEST_CASE += unwrap.lns
TEST_CASE += operand.lns
TEST_CASE += loop.lns
TEST_CASE += cast.lns
TEST_CASE += prototype.lns
TEST_CASE += statement.lns
TEST_CASE += flow.lns
TEST_CASE += flow_switch.lns
TEST_CASE += flow_match.lns
TEST_CASE += import.lns
TEST_CASE += import2.lns
TEST_CASE += import3.lns
TEST_CASE += import4.lns
TEST_CASE += import5.lns
TEST_CASE += strformat.lns
TEST_CASE += Alge.lns
TEST_CASE += Alge2.lns
TEST_CASE += Alge3.lns
TEST_CASE += Alge4.lns
TEST_CASE += exp.lns
TEST_CASE += symbol.lns
TEST_CASE += symbolEmpty.lns
TEST_CASE += static.lns
TEST_CASE += generics.lns
TEST_CASE += generics2.lns
TEST_CASE += generics3.lns
TEST_CASE += generics4.lns
TEST_CASE += box.lns
TEST_CASE += form.lns
TEST_CASE += ddd.lns
TEST_CASE += immediate.lns
TEST_CASE += macro.lns
TEST_CASE += macro2.lns
TEST_CASE += macro3.lns
TEST_CASE += macro4.lns
TEST_CASE += apply.lns
TEST_CASE += multiple-class.lns
TEST_CASE += multiple-if.lns
TEST_CASE += multiple.lns
TEST_CASE += switch.lns
TEST_CASE += lua.lns
TEST_CASE_SRCS := $(addprefix test/error/,$(TEST_CASE))
TEST_CASE_LOGS := $(addprefix $(LOG_DIR)/,$(addsuffix .log,$(TEST_CASE)))



# テストケースを実行して、テストケース毎にログをログディレクトリに出力する
$(LOG_DIR)/%.log: test/error/%
	-$(call exeLune,$(LNSC),$(TEST_DIR),$<,exe,$(EXE_OPT) 2>&1) | grep -e 'error:' -e 'warning:' > $@

test-error-all: $(TEST_CASE_LOGS)
	cat $(TEST_CASE_LOGS)

# エラー系テストケースを実行し、結果を stdout に出力する
define test-error-all-case
	mkdir -p $1
	$(MAKE) test-error-all LOG_DIR=$1 TEST_DIR=$2
endef

test-error:
	$(call test-error-all-case,$(LOG_DIR),$(TEST_DIR))

	-$(call exeLune,$(LNSC),$(TEST_DIR),test/error/noexist.lns,exe,2>&1)
	-$(call exeLune,$(LNSC),$(TEST_DIR),test/error/abbr.lns,exe $(COM_OPT),2>&1) && echo error:NG

	cat test/error/mutable.lns | sed 's@//_lune@_lune@' > $(TEST_DIR)/test/error/dis_mutable.lns
	-$(call exeLune,$(LNSC),$(TEST_DIR),test/error/dis_mutable.lns,exe,2>&1)
	rm $(TEST_DIR)/test/error/dis_mutable.lns


define exeComp
	echo lune | cat $3 - | $(call exeLune,$(LNSC),$1,$2,comp -i,$4) | \
			json_pp --json_opt=canonical,pretty | sort
endef


#>>>>>>>>>>>>>>>>>>>>
# 補完系テストを並列実行するための設定郡

define path2mod
	$(shell echo $1 | sed 's:\.lns$$::g' | sed 's:/:.:g')
endef

define exec_compl-wrap2
	$(call exeComp,$(TEST_DIR),$1,$2,$(call path2mod,$2) $3)
endef

# ダミーのターゲットから目的の文字列を切り出す処理
define strip_1
	$(shell echo $1 | sed 's:@1_.*::')
endef
define strip_2
	$(shell echo $1 | sed 's:.*@1_\(.*\)@2_.*:\1:')
endef
define strip_3
	$(shell echo $1 | sed 's:.*@2_\(.*\).comp:\1:' | sed 's:\.: :' )
endef

%.comp:
	$(call exec_compl-wrap2,$(call strip_1,$@),$(call strip_2,$@),$(call strip_3,$@)) > \
		$(LOG_DIR)/$(notdir $(call strip_1,$@))
#<<<<<<<<<<<<<<<<<<<<

# 補完にかけるファイルのパターン。
# PATH1@1_PATH2@2_POS.comp
#  PATH1: 解析にかける lns ファイルパス
#  PATH2: 補完対象の lns ファイルパス
#  POS: 補完対象の位置 (line.column)
COMP_TARGETS += test/comp.lns@1_test/comp.lns@2_17.7.comp
COMP_TARGETS += test/comp-2.lns@1_test/comp-2.lns@2_17.7.comp
COMP_TARGETS += test/comp-3.lns@1_test/comp-3.lns@2_17.7.comp
COMP_TARGETS += test/comp-4.lns@1_test/comp-4.lns@2_17.12.comp
COMP_TARGETS += test/comp-5.lns@1_test/comp-5.lns@2_17.15.comp
COMP_TARGETS += test/comp2.lns@1_test/compSub.lns@2_5.6.comp
COMP_TARGETS += test/comp4.lns@1_test/comp4.lns@2_9.8.comp
COMP_TARGETS += test/comp5.lns@1_test/comp5.lns@2_7.7.comp
COMP_TARGETS += test/comp6.lns@1_test/comp6.lns@2_3.2.comp
COMP_TARGETS += test/comp7.lns@1_test/comp7.lns@2_6.7.comp
COMP_TARGETS += test/comp8.lns@1_test/comp8.lns@2_4.15.comp
COMP_TARGETS += test/comp9.lns@1_test/comp9.lns@2_4.3.comp



# コード補完のテスト
test-cmpl: $(COMP_TARGETS)
	find $(LOG_DIR) -type f | sort | xargs cat >> $(LOG_FILE)


# オリジナルの LuneScript のテストを行ない、結果を work/test1.resut に出力する
test-org:
	$(call exeLune,$(LNSC),.,test/start.lns,glue,work/glue $(EXE_OPT) 2>&1)
	$(MAKE) build-glue
	$(call exeLune,$(LNSC),.,test/start.lns,exe,$(EXE_OPT) 2>&1) > work/test1.result
	$(call exeLune,$(LNSC),.,test/start.lns,save,$(EXE_OPT) --depends work/1/start.dep)

	$(call exeLune,$(LNSC),.,test/Sub2.lns,SAVE,work/1 $(EXE_OPT))


	grep -v 'warning:' work/test1.result > work/test.result

	mkdir -p work/0/log/comp
	$(MAKE) test-cmpl TEST_DIR=. LOG_DIR=work/0/log/comp LOG_FILE=work/test1.result

	$(MAKE) -s TEST_DIR=. LOG_DIR=work/0/log test-error | \
		grep -e 'error:' -e 'warning:' >> work/test1.result
	$(MAKE) TEST_DIR=. test-etc >> work/test1.result
	$(call exeFront,.) >> work/test1.result

test-org-proj:
ifndef SKIP_PROJ_TEST
	$(MAKE) -j0 -C test/proj/ test MKFILE=Makefile LUA=$(LUA) SRC_DIR= 2>&1 | \
		grep -v /test/proj > work/0/log/proj
else
	@echo skip
endif





# 新しくビルドした LuneScript のテストを行ない、結果を work/test2.resut に出力する
test-new:
#	$(call exeLune,work/3,test/start.lns,glue, ../newglue 2>&1)
#	$(MAKE) build-glue GLUE_DIR=work/newglue

	$(call exeLune,$(LNSC),work/3,test/start.lns,ast,$(EXE_OPT)) > work/ast2.result
	$(call exeLune,$(LNSC),work/3,test/start.lns,lua,$(EXE_OPT)) > work/lua.result
	$(call exeLune,$(LNSC),work/3,test/start.lns,exe,$(EXE_OPT) 2>&1) > work/test2.result


	mkdir -p work/3/log/comp
	$(MAKE) test-cmpl TEST_DIR=work/3 LOG_DIR=work/3/log/comp LOG_FILE=work/test2.result

	$(MAKE) -s TEST_DIR=work/3 LOG_DIR=work/3/log  test-error | \
		grep -e 'error:' -e 'warning:' >> work/test2.result
	$(MAKE) TEST_DIR=work/3 test-etc >> work/test2.result
	$(call exeFront,work/3) >> work/test2.result

test-new-proj:
ifndef SKIP_PROJ_TEST
	$(MAKE) -j0 -C work/3/test/proj/ test MKFILE=Makefile LUA=$(LUA) SRC_DIR= 2>&1 | \
		grep -v /test/proj > work/3/log/proj
else
	@echo skip
endif



test-new2:
## テストスクリプトを lua 化して、 lunescript なしで動かすテスト。
	$(call build-testcode,$(LNSC),work/3_2)
	(cd work/3_2; $(LUA) test/start.lua) > work/test3.result

# テストスクリプトを default-lazy を指定してテスト。
	$(call build-testcode,$(LNSC),work/3_2,--default-lazy)
	(cd work/3_2; $(LUA) test/start.lua) > work/test3_2.result
# default-lazy を有効にすると、遅延ロードによって結果が異なる
	diff work/test3.result work/test3_2.result && exit 1 || echo ok

# テストスクリプトを lua 化して、 runtime をロードしてテスト。
	$(call build-testcode,$(LNSC),work/3_2,-r)
	(cd work/3_2; $(LUA) test/start.lua) > work/test4.result

# 使用するランタイム名指定してテスト
	$(call exeLune,$(LNSC),work/3_2,-mklunemod,runtime.lua) > /dev/null
	$(call build-testcode,$(LNSC),work/3_2,--runtime runtime )
# ランタイムを使用していることを確認するため runtime.lua を移動しておく
	mkdir -p work/backup
	mv work/3_2/lune/base/runtime*.lua work/backup
	(cd work/3_2; $(LUA) test/start.lua) > work/test4-2.result
# runtime.lua を戻す
	mv work/backup/* work/3_2/lune/base/
	rmdir work/backup

	$(call exeLune,$(LNSC),work/3_2,test/start.lns,save,$(EXE_OPT) --depends start.dep)

# オプション -u 付きで save して実行テスト
	$(call mk_lune_proj_dir,work/4)
	cp -a test work/4
	cp -a work/3_2/lune work/4
	$(call exeLune,$(LNSC),work/3_2,test/start.lns,-u SAVE ../4 $(EXE_OPT),)
	(cd work/4; $(LUA) test/start.lua) > work/test5.result



test-parallel: test-org test-org-proj \
			test-new test-new-proj test-new2




test-comp2:
ifeq ($(LANG_KIND),lua)
	$(MAKE) test-comp2-main
else
	$(MAKE) test-comp2-para
endif
# diff で、修正前と修正後の差分がないことを確認
	$(MAKE) diff-result
ifdef LUA51
	$(MAKE) accept-package ACCEPT=y ENABLE_LUA51=$(shell which $(LUA51))
else
	$(MAKE) accept-package ACCEPT=y
endif

# lua 版と go 版を並列処理させるためのターゲット
test-comp2-para: test-comp2-main test-comp2-sub

test-comp2-sub:
	$(MAKE) test-comp2-sub-2 LANG_KIND=lua LANG_KIND=lua

# 一度ビルドしたソース work/ から、 lua 版の lns でビルドして work/13 に出力する。
# 後工程の diff-result で、 go 版の lns のビルド結果と比較して lua と go で
# 同じ出力結果になることを確認する
test-comp2-sub-2:
	$(call cp_src_set,work/13,work/1)
	$(MAKE) -C work/13 build_from_d depend_on=y \
		FROM_DIR=../1 TO_DIR=../13 LNS_OPT="--package lnsc"

test-comp2-main:
#	$(MAKE) test-org


# 新しいトランコンパイラで再度自分自身をコンパイルした結果が同じになることを確認する。
	/usr/bin/time -f 'work1:%e/%U/%S' $(MAKE) -C work/2 build_from_d \
		depend_on=y FROM_DIR=../1 TO_DIR=../2 LNS_OPT="--package lnsc"
	/usr/bin/time -f 'work2:%e/%U/%S' $(MAKE) -C work/3 build_from_d \
		depend_on=y FROM_DIR=../2 TO_DIR=../3 LNS_OPT="--package lnsc"
	diff -r -I '_moduleObj.__buildId.*' -I '.*buildId = .*' \
			-x $(notdir $(LNSC_GO)) work/2 work/3


	cp -a test work/3
#	$(call exeLune,work/3,-mklunemod,lune/base)

#	$(MAKE) test-new

	$(call exeLune,$(LNSC),work/3,test/start.lns,glue, ../newglue $(EXE_OPT) 2>&1)
	$(MAKE) build-glue GLUE_DIR=work/newglue
	mv test/glueTest.so work/3/test/glueTest.so 
	cp -a work/3 work/3_2


	mkdir -p work/0/log
	mkdir -p work/3/log

	$(MAKE) test-parallel
ifndef SKIP_PROJ_TEST
	cat work/0/log/proj >> work/test1.result
	cat work/3/log/proj >> work/test2.result
endif


ifdef LUA51
ifneq ($(shell which $(LUA51)),)
	$(MAKE) test-lua51
endif
endif

ifndef NO_LUAROCKS
	$(MAKE) luarocks-check
endif
	$(call test-meta-update,$(LNSC),work/3)

# rockspec の更新が必要かどうか
	$(MAKE) check-rockspec

# メタフォーマットバージョンを上げなくて良いか確認
	$(MAKE) check-meta-format


define check-rockspec-op
	@grep 'lune\.' ../$1 | sed 's@.*src/@@g' | sed 's@\.lua.*@.lua@g' | sort > work/$1.list
	@(cd $2; find lune -iname '*.lua' ) | sed 's@^@$3@g' | sort > work/$1.2.list
	diff work/$1.list work/$1.2.list
	@rm work/$1.list work/$1.2.list
endef

check-rockspec:
	$(call check-rockspec-op,lunescript-main-1.rockspec,work/3,)
ifdef LUA51
	$(call check-rockspec-op,lunescript51-main-1.rockspec,work/6,legacy/lua51/)
endif


diff-result:
	(cd work/1; find test/ -iname '*.lua' -exec diff {} ../4/{} \; )
	diff -r test/proj/  work/3/test/proj/ \
		-I '_moduleObj.__buildId.*' -I '.*buildId = .*'
ifneq ($(LANG_KIND),lua)
	diff -r -I '_moduleObj.__buildId' -I '.*buildId = .*'		\
			-x $(notdir $(LNSC_GO))				\
			-x 'log' -x 'test' work/13 work/3
endif
	diff work/test1.result work/test2.result
	diff work/test.result work/test3.result
	diff work/test3.result work/test4.result
	diff work/test4.result work/test4-2.result
	diff work/test4.result work/test5.result
	diff work/1/start.dep work/3_2/start.dep
	diff -r work/glue work/newglue


define setup-lua51-test
	cp -a test $1
	$(MAKE) build-glue GLUE_DIR=work/newglue LUA_INC_DIR=$(LUA51_INC) LUA_GLUE=51
	mv test/glueTest51.so $1/test/glueTest.so 
	grep -v -e 'skip-lua51' $2 test/start.lns > $1/test/start.lns
endef

define setup-lua51-test-cross
	$(call setup-lua51-test,$1,-e 'skip-cross')
endef

test-lua51-mkruntime:
	$(call exeLune,$(LNSC),work/5,-mklunemod,lune/base/runtime.lua)

test-lua51:
# LuneScript を Lua 5.1 にクロスコンパイルして実行
	$(MAKE) -C work/5 build_from_d depend_on=y \
			FROM_DIR=../3 TO_DIR=../5 LNS_OPT="-ol 51 --package lnsc"
	$(MAKE) test-lua51-mkruntime LUA=$(LUA51)
	$(call setup-lua51-test,work/5)
	(cd work/5; LUA_PATH=./?.lua $(LUA51) lune/base/base.lua test/start.lns -u exe $(EXE_OPT)) > work/test6.result

# 比較用にクロスコンパイル前のテスト結果を取得
	$(MAKE) test-lua51-sub LANG_KIND=lua LUA=$(LUA51) 

	diff -r work/5/lune work/6/lune

# サンプルを lua5.1 にクロスコンパイルして実行テスト
	$(call mk_lune_proj_dir,work/7)
	cp -a work/3/lune work/7
ifeq ($(WITH_GO),y)
	cp $(LNSC_GO) work/7
endif
	$(call setup-lua51-test-cross,work/7)
	$(call exeLune,$(LNSC),work/7,test/start.lns,-u -ol 51 SAVE,$(EXE_OPT))
	(cd work/7; $(LUA51) test/start.lua) | sed 's/\.0//g' > work/test7-3.result

	diff work/test6.result  work/test7.result 
	diff work/test7_2.result  work/test7-3.result


# lua5.1 用にクロスコンパイルしたものから、 Lua5.1 で自分自身をコンパイル
test-lua51-sub:
	$(call mk_lune_proj_dir,work/6)
	$(call cp_src_set,work/6,work/1)
	$(MAKE) -C work/6 build_from_d depend_on=y \
		FROM_DIR=../5 TO_DIR=../6 LNS_OPT="--package lnsc"
#	$(call selfhost_build,work/5,../6)
	$(call setup-lua51-test,work/6)
	$(call exeLune,$(LNSC),work/6,-mklunemod,lune/base/runtime.lua)
	$(call exeLune,$(LNSC),work/6,test/start.lns,exe,$(EXE_OPT)) | sed 's/\.0//g' > work/test7.result

	cp -a work/6 work/6_2
	rm -rf work/6_2/test
	$(call setup-lua51-test-cross,work/6_2)
	$(call exeLune,$(LNSC),work/6_2,test/start.lns,exe,$(EXE_OPT)) | sed 's/\.0//g' > work/test7_2.result






accept-package:
ifndef ACCEPT
	$(MAKE) -i diff-result
	@read -p "accept? $1 (y/n): " ANS; \
	if [ $${ANS}x != yx ]; then exit 1; fi;
endif
#	rm lune/base/*.lua
	cp work/3/lune/base/*.lua $(ORG_DIR)lune/base/
	cp work/3/lune/base/*.meta $(ORG_DIR)lune/base/
	cp work/3/lune/*.* $(ORG_DIR)lune/
	if [ -f work/3/$(LNSC_GO) ]; then			\
		cp work/3/$(LNSC_GO) $(ORG_DIR).;		\
		cp work/3/lune/base/*.go $(ORG_DIR)lune/base/;	\
		cp work/3/lune/base/runtime_go/binding.go	\
				$(ORG_DIR)lune/base/runtime_go;	\
	fi
ifneq ($(ENABLE_LUA51),)
	cp work/6/lune/base/*.lns $(ORG_DIR)legacy/lua51/lune/base/
	cp work/6/lune/base/*.lua $(ORG_DIR)legacy/lua51/lune/base/
	cp work/6/lune/base/*.meta $(ORG_DIR)legacy/lua51/lune/base/
	cp work/6/lune/*.* $(ORG_DIR)legacy/lua51/lune/
endif


test-installed-cmd:
	$(MAKE) ENABLE_LUAROCKS=y LANG_KIND=luarocks test-comp -j

test-on-docker:
	sudo docker build -t testlns -f test/docker/test_install test/docker
	sudo docker run --name test -it testlns \
		/bin/ash -c "make test-installed-cmd NO_LUAROCKS=y SKIP_PROJ_TEST=y"
	sudo docker container rm test
	sudo docker image rm testlns

test-on-docker_luarocks:
	sudo docker build -t testlns -f test/docker/test_install_with_luarock test/docker
	sudo docker run --name test -it testlns \
		/bin/ash -c 'PATH=/usr/local/bin:${PATH} make test-installed-cmd SKIP_PROJ_TEST=y LUAROCKS=/usr/bin/luarocks-5.3'
	sudo docker container rm test
	sudo docker image rm testlns

VER_PATH=lune/base/Ver.lua
RUNTIME_VER_PICK=grep 'local luaModVersion' | sed 's/[^0-9]//g'
ifneq ($(findstring $(MAKECMDGOALS),"check-runtime-format" ),)
RUNTIME_ORG_VER=$(shell cat $(VER_PATH) | $(RUNTIME_VER_PICK))
RUNTIME_NEW_VER=$(shell cat work/1/$(VER_PATH) | $(RUNTIME_VER_PICK))
endif

check-runtime-format:
# ランタイムバージョンが同じ場合は、ランタイムが一致しているかどうか比較
ifeq ($(RUNTIME_ORG_VER),$(RUNTIME_NEW_VER))
	@diff lune/base/runtime.lua work/1/lune/base/runtime.lua > /dev/null || \
	(echo need to change the runtime version.; exit 1;)
else
	@echo runtime version ok.
endif



ifneq ($(findstring $(MAKECMDGOALS),"check-meta-format" ),)
VER_PICK=grep 'local metaVersion' | awk '//{print $$4}'
ORG_VER=$(shell cat $(VER_PATH) | $(VER_PICK))
NEW_VER=$(shell cat work/3/$(VER_PATH) | $(VER_PICK))
endif

check-meta-format:
# メタフォーマットバージョンが同じ場合は、メタ情報が一致しているかどうか比較
ifeq ($(ORG_VER),$(NEW_VER))
	cat work/1/test/Sub2.meta | grep -v -e '^_moduleObj.__version' > work/1/test/Sub2.meta2
	cat work/4/test/Sub2.meta | grep -v -e '^_moduleObj.__version' > work/4/test/Sub2.meta2
	@diff work/1/test/Sub2.meta2 work/4/test/Sub2.meta2 > /dev/null || \
	(echo need to change the meta format version.; exit 1;)
else
	@echo meta format version ok.
endif



GLUE_DIR=work/glue
build-glue:
	gcc $(GLUE_DIR)/test_glueTest_glue.c test/glue.c -I$(GLUE_DIR) \
		-std=c99 -fPIC -shared -o test/glueTest$(LUA_GLUE).so $(LUA_CFLAGS)


check_lua_mod_dir:
ifeq ($(LUA_MOD_DIR),)
	@echo not found lua command.
	@echo please retry following command.
	@echo 
	@echo "    make install LUA='your lua command'".
	@echo 
	@exit 1
endif

install: check_lua_mod_dir
	$(MAKE) build LUA=$(LUA) LNSDIR=$(LNSDIR)
	@mkdir -p $(LUA_MOD_DIR)/lune/base
	@cat lnsc.lua | sed "s@#!.*@#! /usr/bin/env $(abspath $(LUA))@g" > /tmp/lnsc
	@chmod +x /tmp/lnsc
	@$(INSTALL) /tmp/lnsc /usr/bin
	@echo installed -- /usr/bin/lnsc
	@rm /tmp/lnsc
	@$(INSTALL) $(LNSDIR)/lune/base/*.* $(LUA_MOD_DIR)/lune/base
	@$(INSTALL) $(LNSDIR)/lune/*.* $(LUA_MOD_DIR)/lune
	@echo installed -- $(LUA_MOD_DIR)/lune

uninstall: check_lua_mod_dir
	rm /usr/bin/lnsc
	rm -rf $(LUA_MOD_DIR)/lune


# テストケースだけ修正した時に、通常のビルドチェックは重すぎるので、
# Lua に変換した結果が同じになっているかどうかを確認し、同じなら OK にする
build-testcase:
	$(call cp_src_set,work/1,.)
	$(call selfhost_build,$(LNSC),.,work/1,--testing)




test-ccc:
	LANG= $(MAKE) -f Makefile.ccc LUA_INC_DIR=$(LUA_INC_DIR) LUA_LIB=$(LUA_LIB)

clean-ccc:
	LANG= $(MAKE) -f Makefile.ccc clean


# トランスコンパイル済みの C ソースをビルド、実行、テストする
#
# $1
# $2: トランスコンパイラ済みの C ソースパスのリスト
# $3: 実行結果のログファイル名
# $4: 実行するメインモジュールファイルパス (拡張子 .lns は付けないファイルパス)
# $5: 比較用に実行する Lns のオプション
define test-convC-exe
	$(call exeLune,$1,$2,$5.lns,bootC -oc work/boot.c)
	LANG= $(MAKE) -C $2 -f Makefile.ccc "SRCS=$3 work/boot.c" TEST_OUT=test.exe \
		LUA_INC_DIR=$(LUA_INC_DIR) LUA_LIB=$(LUA_LIB)
	@# C へのトランスコンパイラ結果と、
	@# Lua へのトランスコンパイラ結果で出力が同じになることを確認
	@(cd $2; ./test.exe | grep -v -e '^:debug:' > work/$4.c.log)
	$(call exeLune,$2,$5.lns,$6,exe > work/$4.lua.log)
	diff $2/work/$4.c.log $2/work/$4.lua.log
endef

# lns ソースを C へトランスコンパイルする
#
# $1: 実行する作業ディレクトリ
# $2: トランスコンパイルする lns ソースのパス
define convC
	$(call exeLune,$1,$2,$3,save -langC,$4)
endef

build-miniC:
	$(call test-convC-exe,$(LNSC),.,miniC.lns,miniC,miniC)

test-miniC:
	$(call cmd_exec_trans,$(LNSC),.,.,,base/convCC)
	$(LUA) lune/base/base.lua --mkbuiltin lune/base/runtime_c
	mkdir -p work/0
	rm -rf work/0/*
	mkdir -p work/0/work/c
	cp -a lune work/0
	cp Makefile.ccc work/0
	$(call convC,$(LNSC),work/0,lune/base/Testing.lns,--testing)
#	-$(call test-convC-test,work/0,lune/base/Testing)
#	cp -a work/c/* work/0/work/c
#	cp -a work/0/work/* work/
	cp miniC.lns work/0
#	$(call test-convC-test,.,miniC)
	$(call test-convC-test,$(LNSC),work/0,miniC)
	cp work/0/test.exe .


test-subC:
	$(call cmd_exec_trans,$(LNSC),.,.,,base/convCC)
	$(call convC,$(LNSC),.,miniC.lns)
	$(call convC,$(LNSC),.,subC.lns)
	$(call test-convC-exe,$(LNSC),.,miniC.lns subC.lns,subc,miniC)


build-cc:
	$(MAKE) test-miniC
	$(call convC,$(LNSC),.,ccc.lns)
	$(call convC,$(LNSC),.,subC.lns)
	$(call test-convC-exe,$(LNSC),.,ccc.lns subC.lns,ccc,ccc)

test-formatter:
	$(call cmd_exec_trans,$(LNSC),.,.,,base/Formatter)
	$(call exeLune,$(LNSC),.,ccc.lns,format,) > work/format.lns

define test-convC-test
	$(call convC,$1,$2,$3.lns,--testing)
	$(call test-convC-exe,$1,$2,$3.c lune/base/Testing.c $4,$(notdir $3),$3,--testing)
endef

test-testC:
	$(call cmd_exec_trans,$(LNSC),.,.,,base/convCC)
	$(LUA) lune/base/base.lua --mkbuiltin lune/base/runtime_c
	mkdir -p work/0
	rm -rf work/0/*
	mkdir -p work/0/work/c
	cp -a lune work/0
	cp Makefile.ccc work/0
	$(call convC,$(LNSC),work/0,lune/base/Testing.lns,--testing)
	$(call test-convC-test,$(LNSC),work/0,lune/base/Ver)
	$(call test-convC-test,$(LNSC),work/0,lune/base/Code)
	$(call test-convC-test,$(LNSC),work/0,lune/base/LuaMod)
	$(call convC,$(LNSC),work/0,lune/base/LuaVer.lns,--testing)
	$(call convC,$(LNSC),work/0,lune/base/LuneControl.lns,--testing)
	$(call convC,$(LNSC),work/0,lune/base/Depend.lns,--testing)
	$(call test-convC-test,$(LNSC),work/0,lune/base/Util, \
		lune/base/Depend.sub.c lune/base/Depend.c \
		lune/base/LuneControl.c lune/base/LuaMod.c )
	$(call test-convC-test,$(LNSC),work/0,lune/base/Parser, \
		lune/base/Depend.sub.c lune/base/Depend.c \
		lune/base/LuneControl.c lune/base/LuaMod.c \
		lune/base/Util.c)
	$(call test-convC-test,$(LNSC),work/0,lune/base/Writer, \
		lune/base/Depend.sub.c lune/base/Depend.c \
		lune/base/LuneControl.c lune/base/LuaMod.c \
		lune/base/Util.c)
	$(call test-convC-test,$(LNSC),work/0,lune/base/Ast, \
		lune/base/Depend.sub.c lune/base/Depend.c \
		lune/base/LuneControl.c lune/base/LuaMod.c \
		lune/base/Util.c lune/base/Writer.c \
		lune/base/Parser.c lune/base/Code.c )
	$(call convC,$(LNSC),work/0,lune/base/frontInterface.lns,--testing)

test-mkbuiltin:
	$(call cmd_exec_trans,$(LNSC),.,.,,base/convCC)
	$(LUA) lune/base/base.lua --mkbuiltin lune/base/runtime_c
	$(MAKE) test-miniC

lctags:
	$(MAKE) -f Makefile.ccc lctags


test-start:
	$(call cmd_exec_trans,$(LNSC),.,.,,base/convCC)
	$(LUA) lune/base/base.lua --mkbuiltin lune/base/runtime_c
	mkdir -p work/0
	rm -rf work/0/*
	mkdir -p work/0/work/c
	cp -a lune work/0
	cp -a test work/0
	cp Makefile.ccc work/0
	$(call convC,$(LNSC),work/0,test/Sub4.lns)
	$(call convC,$(LNSC),work/0,test/Sub2.lns)
	$(call test-convC-exe,$(LNSC),work/0,test/Sub2.lns,test/Sub2,test/Sub2)


define test-tutorialC
	$(call convC,$1,$2,$(basename $3).lns)
	$(call test-convC-exe,$1,$2,$3,$(notdir $3),$(basename $3))
endef


test-tutorial:
	$(call cmd_exec_trans,$(LNSC),.,.,,base/convCC)
	$(LUA) lune/base/base.lua --mkbuiltin lune/base/runtime_c
	mkdir -p work/0
	rm -rf work/0/*
	mkdir -p work/0/work/c
	cp -a lune work/0
	cp -a test work/0
	cp Makefile.ccc work/0
	$(call test-tutorialC,$(LNSC),work/0,test/tutorial/match.lns)
	$(call test-tutorialC,$(LNSC),work/0,test/tutorial/loop.lns)
	$(call test-tutorialC,$(LNSC),work/0,test/tutorial/interface.lns)
	$(call test-tutorialC,$(LNSC),work/0,test/tutorial/generics.lns)
	$(call test-tutorialC,$(LNSC),work/0,test/tutorial/enum.lns)
	$(call test-tutorialC,$(LNSC),work/0,test/tutorial/classOverride.lns)
	$(call test-tutorialC,$(LNSC),work/0,test/tutorial/classMapping.lns)
	$(call test-tutorialC,$(LNSC),work/0,test/tutorial/classInherit.lns)
	$(call test-tutorialC,$(LNSC),work/0,test/tutorial/class2accessor.lns)
	$(call test-tutorialC,$(LNSC),work/0,test/tutorial/class1.lns)
	$(call test-tutorialC,$(LNSC),work/0,test/tutorial/cast.lns)
	$(call test-tutorialC,$(LNSC),work/0,test/tutorial/branch.lns)
	$(call test-tutorialC,$(LNSC),work/0,test/tutorial/nilable.lns)
	$(call test-tutorialC,$(LNSC),work/0,test/tutorial/defaultArg.lns)
	$(call test-tutorialC,$(LNSC),work/0,test/tutorial/multipleRetVal.lns)
	$(call test-tutorialC,$(LNSC),work/0,test/tutorial/arg.lns)
	$(call test-tutorialC,$(LNSC),work/0,test/tutorial/func.lns)
	$(call test-tutorialC,$(LNSC),work/0,test/tutorial/variable.lns)
	$(call test-tutorialC,$(LNSC),work/0,test/tutorial/set.lns)
	$(call test-tutorialC,$(LNSC),work/0,test/tutorial/value.lns)

test-cAll:
	$(MAKE) build-cc
	$(MAKE) test-tutorial
	$(MAKE) test-testC


test-gc:
	mkdir -p work/c
	$(MAKE) -f Makefile.ccc test-gc


GO_RUNTIME = lune/base/runtime_go
GO_MOD = lnsc

LUA_URL=https://www.lua.org/ftp/lua-5.3.6.tar.gz

setup-pkg-lua-sub:
	mkdir -p work
	curl $(LUA_URL) -o work/lua.tar.gz
	(cd work; tar xvfz lua.tar.gz;mv lua-* lua)
	make -j -C work/lua/src linux
	cat ../etc/lua-5.3.pc |							\
		sed 's@^includedir=.*@includedir=$(shell pwd)/work/lua/src/@' |	\
		sed 's@^Libs:.*@Libs: -lm -ldl $(shell pwd)/work/lua/src/liblua.a@' > lua-5.3.pc


setup-pkg-lua:
	@if ! pkg-config lua-5.3; then			\
		$(MAKE) setup-pkg-lua-sub;		\
	else						\
		echo ready lua-5.3;			\
	fi


define buildGoSetup
	(cd $1; if [ ! -f go.mod ]; then go mod init $(GO_MOD); fi)
endef

define buildGo
	$(call buildGoSetup,$1)
	(cd $1; PKG_CONFIG_PATH=$(shell pwd) go build && $(LNSC_GO))
endef

GO_WORK_DIR = work/go

# lns ソースを Golang へトランスコンパイルする
#
# $1: 実行する作業ディレクトリ
# $2: トランスコンパイルする lns ソースのパス
# $3: lnsc のオプション
# $4: 出力先ディレクトリ
define convGo
	$(call exeLune,$1,$2,$3.lns,SAVE -langGo,$4 $5)
endef

define setup-go
	mkdir -p $1
	rsync -a --include='*.lua' \
			--include='*.lns' --exclude='*.*' lune $1
	cp -au lune/base/runtime_go $1/lune/base/
	cp -au build.counter $1
	-rm work/result-testGo.txt work/result-testGo2.txt
endef

define test-go-one-setup
	cp -au test $1
endef

define test-go-one-exec
	cat lune/base/runtime_go/main/runtime_main.go |		\
		sed 's/Lns_init/Lns_$3_init/;' > $2/runtime_main.go
	cp -au $4/$3.lns $2/$4
	$(call convGo,$1,$2,$4/$3,--package main)
	mv $2/$4/$3.go $2
	$(call buildGo,$2) >> work/result-testGo.txt
	if test "${NO_LUA_DIFF}" = ""; then					\
		$(call exeLune,$1,$2,$4/$3.lns,exe) >> work/result-testGo2.txt &&	\
		diff -C 2 -I 'ONLY_LUA'						\
				work/result-testGo.txt work/result-testGo2.txt;	\
	fi
endef


define test-go-one
	$(call test-go-one-setup,$2)
	$(call test-go-one-exec,$1,$2,$3,$4)
endef

define test-go-some-setup
	cp -au $4/$3.lns $2/$4
	$(call convGo,$1,$2,$4/$3,--package main)
	mv $2/$4/$3.go $2
endef


define test-go-some
	$(call test-go-one-exec,$1,$2,$3,$4)
endef


test-miniGo:
	$(call cmd_exec_trans,$(LNSC),.,.,,base/convGo)
	$(call setup-go,$(GO_WORK_DIR))
	-rm $(GO_WORK_DIR)/*.*
	cp build.counter $(GO_WORK_DIR)
	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),miniGo,.)
	cat work/result-testGo.txt

ifndef NO_TEST
	$(MAKE) test-go
endif

miniGo-go:
	$(call buildGo,$(GO_WORK_DIR))

test-go:
	$(call cmd_exec_trans,$(LNSC),.,.,,base/convGo)
	$(call setup-go,$(GO_WORK_DIR))
	$(call test-go-one-setup,$(GO_WORK_DIR))
	$(call test-go-some-setup,$(LNSC),$(GO_WORK_DIR),SubGo,test)
	$(call test-go-some,$(LNSC),$(GO_WORK_DIR),testGo,test)


test-tutorial-go:
	$(call cmd_exec_trans,$(LNSC),.,.,,base/convGo)
	$(call setup-go,$(GO_WORK_DIR))

	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),match,test/tutorial)
	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),loop,test/tutorial)
	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),interface,test/tutorial)
	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),generics,test/tutorial)
	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),enum,test/tutorial)
	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),classOverride,test/tutorial)
	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),classMapping,test/tutorial)
	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),classInherit,test/tutorial)
	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),class2accessor,test/tutorial)
	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),class1,test/tutorial)
	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),cast,test/tutorial)
	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),branch,test/tutorial)
	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),nilable,test/tutorial)
	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),defaultArg,test/tutorial)
	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),multipleRetVal,test/tutorial)
	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),arg,test/tutorial)
	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),func,test/tutorial)
	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),variable,test/tutorial)
	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),set,test/tutorial)
	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),value,test/tutorial)

	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),ccc,.)


define test-lns-test
	@echo "local _lune = require( 'lune.base.runtime' )" > $1/testRun.lua
	@echo "local __t = _lune.loadModule( 'lune.base.$2' )" >> $1/testRun.lua
	@echo "local __t = _lune.loadModule( 'lune.base.Testing' )" >> $1/testRun.lua
	@echo "__t.run( 'lune.base.$2');"			>> $1/testRun.lua
	@echo "__t.outputAllResult( io.stdout );"		>> $1/testRun.lua
	(cd $1; LUA_PATH=./?.lua $(LUA) testRun.lua)
endef


define test-convGo-test
	cat lune/base/runtime_go/main/runtime_main.go |		\
		sed 's/Lns_init/Lns_$3_init/' |			\
		sed 's@//TEST:@@' |				\
		sed 's/run( "" )/run( "lune.base.$3" )/' |				\
		sed 's@//IMPORT@import . "$(GO_MOD)/lune/base"@' > $2/$5/runtime_main.go

	$(call convGo,$1,$2,$4/$3,--testing --package lnsc,$5)
	$(call buildGo,$2/$5) >> work/result-testGo.txt
	$(call test-lns-test,$2/$5,$3) >> work/result-testGo2.txt
	diff -C 2 -I 'ONLY_LUA' work/result-testGo.txt work/result-testGo2.txt
endef

test-testGo:
	rm -rf $(GO_WORK_DIR)/*.* $(GO_WORK_DIR)/lune $(GO_WORK_DIR)/test
	$(call setup-go,$(GO_WORK_DIR))
	$(call setup-go,$(GO_WORK_DIR)/test)

	$(call cmd_exec_trans,$(LNSC),.,.,,base/convGo)

	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/Testing,--testing --package lnsc,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/Types,--testing --package lnsc,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/Async,--testing --package lnsc,./test)
	$(call test-convGo-test,$(LNSC),$(GO_WORK_DIR),Ver,lune/base,./test)
	$(call test-convGo-test,$(LNSC),$(GO_WORK_DIR),Code,lune/base,./test)
	$(call test-convGo-test,$(LNSC),$(GO_WORK_DIR),LuaMod,lune/base,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/LuaVer,--testing --package lnsc,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/LuneControl,--testing --package lnsc,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/Meta,--testing --package lnsc,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/Log,--testing --package lnsc,./test)
	$(call test-convGo-test,$(LNSC),$(GO_WORK_DIR),Util,lune/base,./test)
	$(call test-convGo-test,$(LNSC),$(GO_WORK_DIR),AsyncParser,lune/base,./test)
	$(call test-convGo-test,$(LNSC),$(GO_WORK_DIR),Parser,lune/base,./test)
	$(call test-convGo-test,$(LNSC),$(GO_WORK_DIR),Json,lune/base,./test)
	$(call test-convGo-test,$(LNSC),$(GO_WORK_DIR),Writer,lune/base,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/frontInterface,--testing --package lnsc,./test)
	$(call test-convGo-test,$(LNSC),$(GO_WORK_DIR),Ast,lune/base,./test)
	$(call test-convGo-test,$(LNSC),$(GO_WORK_DIR),Option,lune/base,./test)
	$(call test-convGo-test,$(LNSC),$(GO_WORK_DIR),Nodes,lune/base,./test)
	$(call test-convGo-test,$(LNSC),$(GO_WORK_DIR),Formatter,lune/base,./test)
	$(call test-convGo-test,$(LNSC),$(GO_WORK_DIR),Macro,lune/base,./test)
	$(call test-convGo-test,$(LNSC),$(GO_WORK_DIR),TransUnit,lune/base,./test)
	$(call test-convGo-test,$(LNSC),$(GO_WORK_DIR),convLua,lune/base,./test)
	$(call test-convGo-test,$(LNSC),$(GO_WORK_DIR),dumpNode,lune/base,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/convCC,--testing --package lnsc,./test)
	$(call test-convGo-test,$(LNSC),$(GO_WORK_DIR),convGo,lune/base,./test)
	$(call test-convGo-test,$(LNSC),$(GO_WORK_DIR),OutputDepend,lune/base,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/glueFilter,--testing --package lnsc,./test)
	$(call test-convGo-test,$(LNSC),$(GO_WORK_DIR),front,lune/base,./test)

test-subgo:
	$(call cmd_exec_trans,$(LNSC),.,.,,base/convGo)
	$(call setup-go,$(GO_WORK_DIR))

	$(call test-convGo-test,$(LNSC),$(GO_WORK_DIR),front,lune/base,./test)


define binding-lns-go
	$(LNSC_LUA) lune/base/runtime_go/binding/bind.lns exe --testing
	mv output.go $1/lune/base/runtime_go/binding.go 
endef

define go-build-lns
	cp lune/base/runtime_go/main/lnsc_main.go $1
	$(call buildGoSetup,$1)
	(cd $1; PKG_CONFIG_PATH=$(shell pwd) time go build)
endef

# セルフホストのソースを go に変換し、 go 版 lns をビルドする
go-build:
	rm -rf $(GO_WORK_DIR)
	$(call setup-go,$(GO_WORK_DIR))
	$(call setup-go,$(GO_WORK_DIR)/test)

	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/Testing,,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/Async,,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/Types,,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/Ver,,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/Code,,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/LuaMod,,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/LuaVer,,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/LuneControl,,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/Meta,,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/Log,,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/Util,,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/AsyncParser,,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/Parser,,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/Json,,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/Writer,,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/frontInterface,,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/Ast,,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/Option,,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/Nodes,,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/Formatter,,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/Macro,,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/TransUnit,,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/convLua,,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/dumpNode,,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/convCC,,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/convGo,,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/OutputDepend,,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/glueFilter,,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/front,,./test)

	$(call go-build-lns,$(GO_WORK_DIR)/test)

ifndef NO_TEST
	$(MAKE) test-govslua
	$(MAKE) LNSC=$(shell pwd)/$(GO_WORK_DIR)/test/$(LNSC_GO) test-lnsc-go
endif
	cp $(GO_WORK_DIR)/test/$(LNSC_GO) .

test-lnsc-go:
	-$(MAKE) test-go
# LuneScript のソース lune/base にアクセスできなくても lnsc コマンドが動くか確認する
	(cd $(GO_WORK_DIR); $(LNSC) test/testGo.lns exe) > work/result-testGo3.txt
	diff -I 'fib_l' work/result-testGo2.txt work/result-testGo3.txt 


define test-govslua-comp
	(cd $1; /usr/bin/time -f 'lns:%e/%U/%S:$2' $(LNSC_GO) lune/base/$2.lns SAVE $3 --log log)
	(cd $1; LUA_PATH=./?.lua /usr/bin/time -f 'lua:%e/%U/%S:$2' $(LNSC_LUA) lune/base/$2.lns SAVE $4 --log log)
	@(cd $1; diff $4/lune/base/$2.lua $3/lune/base/$2.lua)
	@(cd $1; diff $4/lune/base/$2.meta $3/lune/base/$2.meta)
endef


# go-build でビルドした go 版 lns と lua 版 lns とで、
# lns ソースの変換を行ない。同じ結果になることを確認する。
test-govslua:
	-rm work/go.time work/lua.time
ifdef BATCH
	/usr/bin/time -f 'go:%e/%U/%S:$2' -o work/go.time --append		\
		$(MAKE) test-govslua-go-batch
else
	/usr/bin/time -f 'go:%e/%U/%S:$2' -o work/go.time --append		\
		$(MAKE) test-govslua-go
endif
	/usr/bin/time -f 'lua:%e/%U/%S:$2' -o work/lua.time --append		\
		$(MAKE) test-govslua-lua
	cat work/go.time work/lua.time
	diff -r work/1 work/2 -I buildId				\
		-I ': Entering directory'				\
		-I ': Leaving directory' -x '$(notdir $(LNSC_GO))'


define test-buildlns-setup
	mkdir -p $1
	rm -rf $1/*
	$(call mk_lune_proj_dir,$1)
	$(call cp_src_set,$1,.)
endef

define test-buildlns-run
	$(call selfhost_build_dep,$1,.,$2,-langGo --package lnsc -ol 53)
	test "$(ONLY_CONV)" = "" && LANG= $(MAKE) -C $2 LANG_KIND=$3 tgt_dump_ast > $2/ast || echo skip
endef

test-govslua-lua:
	$(call test-buildlns-setup,work/1)
	$(call test-buildlns-run,$(LNSC),work/1,lua)

test-govslua-go:
	$(call test-buildlns-setup,work/2)
ifdef TEST_MAIN_GO
	cp $(LNSC_GO) work/2
else
	cp $(GO_WORK_DIR)/test/$(LNSC_GO) work/2
endif
	$(call test-buildlns-run,/usr/bin/time -f 'go:%e/%U/%S:$2' -o work/go.time --append work/2/$(LNSC_GO),work/2,go)


test-govslua-lua-batch:
	$(call test-buildlns-setup,work/2)
	$(call selfhost_buildOp_batch,.,$(LNSC),work/2,--depends work/2/ -langGo --package lnsc)
test-govslua-go-batch:
	$(call test-buildlns-setup,work/2)
	$(call selfhost_buildOp_batch,.,$(LNSC_GO),work/2,--depends work/2/ -langGo --package lnsc)


goprof:
	(cd work/go/test/; go tool pprof -dot cpu.pprof  ) | dot -Tsvg -odump.svg



test-reference:
	if [ -d ~/work/ifritJP.github.io ]; then					\
		if [ ! -d work/ifritJP.github.io ]; then				\
			ln -s ~/work/ifritJP.github.io work/ifritJP.github.io;		\
		fi									\
	else										\
		if [ -d work/ifritJP.github.io ]; then					\
			(cd work/ifritJP.github.io; git pull);				\
		else									\
			(cd work; git clone https://github.com/ifritJP/ifritJP.github.io.git);	\
		fi									\
	fi
	$(MAKE) test-reference-sample


test-reference-sample:
	$(call exeLune,$(LNSC),.,test/checkReference.lns exe)

