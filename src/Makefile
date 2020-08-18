LUA=lua5.3
#LUA=lua5.2



LUA51=lua5.1
LUA51_INC=/usr/include/lua5.1

LUA_LIB=-llua5.3
LUA_INC_DIR=/usr/include/lua5.3
ifneq ($(LUA_INC_DIR),)
LUA_CFLAGS=-I$(LUA_INC_DIR)
endif

LNSC=~/.luarocks/bin/lnsc

LUAROCKS=luarocks

ifdef LOG
LOG_OPT=--log info
endif

ifdef LNSC_ENABLE
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
	@echo "  make test-go"
	@echo "  make test-miniGo"

exec:
	$(LUA) lune/base/base.lua test/test.lns token
	$(LUA) lune/base/base.lua test/test.lns ast
	$(LUA) lune/base/base.lua test/test.lns LUA

ifdef depend_on
DEPENDS := $(shell find lune -iname '*.d')
DEPENDS_META := $(patsubst %.d,%.meta,$(DEPENDS))

include $(DEPENDS)

$(DEPENDS_META):

COM_OPT := --use-ipairs -Werror

%.meta: %.lns
	(cd $(FROM_DIR); \
		LUA_PATH=./?.lua $(LUA) lune/base/base.lua $< SAVE $(TO_DIR) \
			$(BYTECOMP) $(COM_OPT) $(LOG_OPT) $(LNS_OPT))

build_from_d: $(DEPENDS_META)

endif



define cutMeta
	(cd work/$2/lune/base; cat $1.lua | awk '/^----- meta/{META++;}; //{if (META != 1) {print $$0}} ' > $1.cut.lua )
endef


ifndef LNSC_ENABLE
define exeLune
	(cd $1; LUA_PATH=./?.lua $(LUA) lune/base/base.lua $2 $3 $4)
endef
else
define exeLune
	(cd $1; $(LNSC) $2 $3 $4)
endef
endif

define exeFront
	(cd $1; LUA_PATH=./?.lua $(LUA) -e 'print( require( "lune.base.front" ).convertLnsCode2LuaCode( "print( 1 + 10 );", "lua" ) )')
endef



define transwork
	$(call exeLune,$1,lune/$4.lns SAVE $2 $3 $(BYTECOMP) $(COM_OPT) $(LOG_OPT))
endef

define transwork_dep
	$(call exeLune,$1,lune/$4.lns SAVE $2 --depends $2/lune/$4.d $3 $(BYTECOMP) $(COM_OPT) $(LOG_OPT))
endef


define buildLnsOp
	$(call $4,$1,$2,$3,base/Ver)
	$(call $4,$1,$2,$3,base/LuaMod)
	$(call $4,$1,$2,$3,base/LuneControl)
	$(call $4,$1,$2,$3,base/Meta)
	$(call $4,$1,$2,$3,base/Testing)
	$(call $4,$1,$2,$3,base/Code)
	$(call $4,$1,$2,$3,base/Log)
	$(call $4,$1,$2,$3,base/LuaVer)
	$(call $4,$1,$2,$3,base/Depend)
	$(call $4,$1,$2,$3,base/Util)
	$(call $4,$1,$2,$3,base/frontInterface)
	$(call $4,$1,$2,$3,base/Writer)
	$(call $4,$1,$2,$3,base/Parser)
	$(call $4,$1,$2,$3,base/Ast)
	$(call $4,$1,$2,$3,base/Option)
	$(call $4,$1,$2,$3,base/Nodes)
	$(call $4,$1,$2,$3,base/Formatter)
	$(call $4,$1,$2,$3,base/dumpNode)
	$(call $4,$1,$2,$3,base/Macro)
	$(call $4,$1,$2,$3,base/TransUnit)
	$(call $4,$1,$2,$3,base/convLua)
	$(call $4,$1,$2,$3,base/convCC)
	$(call $4,$1,$2,$3,base/convGo)
	$(call $4,$1,$2,$3,base/OutputDepend)
	$(call $4,$1,$2,$3,base/TestCtrl)
	$(call $4,$1,$2,$3,base/front)
	$(call $4,$1,$2,$3,base/glueFilter)
	$(call $4,$1,$2,$3,Util)
	$(call exeLune,$1/$2,-mklunemod,lune/base/runtime.lua) > /dev/null
	cat $1/$2/$(VER_PATH) | $(RUNTIME_VER_PICK) | \
		awk '//{printf( "$1/$2/lune/base/runtime%s.lua",$$0);}' | \
		xargs cp $1/$2/lune/base/runtime.lua 
endef

define buildLns
	$(call buildLnsOp,$1,$2,$3,transwork)
endef

define buildLns_dep
	$(call buildLnsOp,$1,$2,$3,transwork_dep)
endef



define build-testcode
	$(call exeLune,$1,test/macro.lns,$2 save,) > /dev/null
	$(call exeLune,$1,test/Sub2.lns,$2 save,) > /dev/null
	$(call exeLune,$1,test/Sub3.lns,$2 save,) > /dev/null
	$(call exeLune,$1,test/Sub3_2.lns,$2 save,) > /dev/null
	$(call exeLune,$1,test/Sub4.lns,$2 save,) > /dev/null
	$(call exeLune,$1,test/Sub5.lns,$2 save,) > /dev/null
	$(call exeLune,$1,test/Sub6.lns,$2 save,) > /dev/null
	$(call exeLune,$1,test/simple.lns,$2 save,) > /dev/null
	$(call exeLune,$1,test/Class1.lns,$2 save,) > /dev/null
	$(call exeLune,$1,test/Class2.lns,$2 save,) > /dev/null
	$(call exeLune,$1,test/Class3.lns,$2 save,) > /dev/null
	$(call exeLune,$1,test/Class4.lns,$2 save,) > /dev/null
	$(call exeLune,$1,test/start.lns,$2 save,) > /dev/null
	$(call exeLune,$1,test/mapping.lns,$2 save,) > /dev/null
	$(call exeLune,$1,test/funcSym.lns,$2 save,) > /dev/null
	$(call exeLune,$1,test/test2/test2/test.lns,$2 save,) > /dev/null
endef

# meta 情報の更新時に正常に解析できるかテスト。
define test-meta-update
	-(cd $1; rm test/ext/*.lua test/ext/*.meta )
	$(call exeLune,$1,test/ext/sub3.lns,SAVE,)
	$(call exeLune,$1,test/ext/sub2.lns,SAVE,)
	(cd $1; rm test/ext/sub3.lua test/ext/sub3.meta )
	$(call exeLune,$1,test/ext/sub1.lns,SAVE,)
endef

luarocks-check:
	$(MAKE) luarocks-check-sub LNSC_ENABLE=y 

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
	(cd work/8; $(LNSC) test/start.lns exe) > work/test8.result

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
	$(call buildLns,$(LNSDIR),.)

test-save:
	$(call test-meta-update,.)


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
endef

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
	$(call buildLns_dep,.,work/1,)
	$(call cp_src_set,work/2,work/1)
	$(call cp_src_set,work/3,work/1)
	$(call cp_src_set,work/5,work/1)
	$(MAKE) check-runtime-format
	$(MAKE) test-comp2


	@echo finish to build 
	@LANG= date

test-registered-rockspec:
	curl https://luarocks.org/manifests/ifritjp/lunescript-main-1.rockspec -o work/rockspec
	@diff ../lunescript-main-1.rockspec work/rockspec || echo "====== need upload rockspec ======" && echo "ok"


test-etc:
# 存在しないファイルを指定した時に、不正なファイルが作成されないことを確認する
	-@$(call exeLune,$(TEST_DIR),test/error/__.lns,SAVE,2>&1) > /dev/null
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
TEST_CASE_SRCS := $(addprefix test/error/,$(TEST_CASE))
TEST_CASE_LOGS := $(addprefix $(LOG_DIR)/,$(addsuffix .log,$(TEST_CASE)))



# テストケースを実行して、テストケース毎にログをログディレクトリに出力する
$(LOG_DIR)/%.log: test/error/%
	-$(call exeLune,$(TEST_DIR),$<,exe,2>&1) | grep -e 'error:' -e 'warning:' > $@

test-error-all: $(TEST_CASE_LOGS)
	cat $(TEST_CASE_LOGS)

# エラー系テストケースを実行し、結果を stdout に出力する
define test-error-all-case
	mkdir -p $1
	$(MAKE) test-error-all LOG_DIR=$1 TEST_DIR=$2
endef

test-error:
	$(call test-error-all-case,$(LOG_DIR),$(TEST_DIR))

	-$(call exeLune,$(TEST_DIR),test/error/noexist.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/abbr.lns,exe $(COM_OPT),2>&1) && echo error:NG

	cat test/error/mutable.lns | sed 's@//_lune@_lune@' > $(TEST_DIR)/test/error/dis_mutable.lns
	-$(call exeLune,$(TEST_DIR),test/error/dis_mutable.lns,exe,2>&1)
	rm $(TEST_DIR)/test/error/dis_mutable.lns


define exeComp
	echo lune | cat $3 - | $(call exeLune,$1,$2,comp -i,$4) | \
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
test-compl: $(COMP_TARGETS)
	find $(LOG_DIR) -type f | sort | xargs cat >> $(LOG_FILE)


# オリジナルの LuneScript のテストを行ない、結果を work/test1.resut に出力する
test-org:
	$(call exeLune,.,test/start.lns,glue,work/glue 2>&1)
	$(MAKE) build-glue
	$(call exeLune,.,test/start.lns,exe,2>&1) > work/test1.result
	$(call exeLune,.,test/start.lns,save,--depends work/1/start.dep)

	$(call exeLune,.,test/Sub2.lns,SAVE,work/1)


	grep -v 'warning:' work/test1.result > work/test.result

	mkdir -p work/0/log/comp
	$(MAKE) test-compl TEST_DIR=. LOG_DIR=work/0/log/comp LOG_FILE=work/test1.result

	$(MAKE) -s TEST_DIR=. LOG_DIR=work/0/log test-error | \
		grep -e 'error:' -e 'warning:' >> work/test1.result
	$(MAKE) TEST_DIR=. test-etc >> work/test1.result
	$(call exeFront,.) >> work/test1.result

test-org-proj:
ifndef SKIP_PROJ_TEST
	$(MAKE) -C test/proj/ test MKFILE=Makefile SRC_DIR= 2>&1 | \
		grep -v /test/proj > work/0/log/proj
else
	@echo skip
endif





# 新しくビルドした LuneScript のテストを行ない、結果を work/test2.resut に出力する
test-new:
#	$(call exeLune,work/3,test/start.lns,glue, ../newglue 2>&1)
#	$(MAKE) build-glue GLUE_DIR=work/newglue

	$(call exeLune,work/3,test/start.lns,ast,) > work/ast2.result
	$(call exeLune,work/3,test/start.lns,lua,) > work/lua.result
	$(call exeLune,work/3,test/start.lns,exe,2>&1) > work/test2.result


	mkdir -p work/3/log/comp
	$(MAKE) test-compl TEST_DIR=work/3 LOG_DIR=work/3/log/comp LOG_FILE=work/test2.result

	$(MAKE) -s TEST_DIR=work/3 LOG_DIR=work/3/log  test-error | \
		grep -e 'error:' -e 'warning:' >> work/test2.result
	$(MAKE) TEST_DIR=work/3 test-etc >> work/test2.result
	$(call exeFront,work/3) >> work/test2.result

test-new-proj:
ifndef SKIP_PROJ_TEST
	$(MAKE) -C work/3/test/proj/ test MKFILE=Makefile SRC_DIR= 2>&1 | \
		grep -v /test/proj > work/3/log/proj
else
	@echo skip
endif



test-new2:
## テストスクリプトを lua 化して、 lunescript なしで動かすテスト。
	$(call build-testcode,work/3_2)
	(cd work/3_2; $(LUA) test/start.lua) > work/test3.result

# テストスクリプトを lua 化して、 runtime をロードしてテスト。
	$(call build-testcode,work/3_2,-r)
	(cd work/3_2; $(LUA) test/start.lua) > work/test4.result

# 使用するランタイム名指定してテスト
	$(call exeLune,work/3_2,-mklunemod,runtime.lua) > /dev/null
	$(call build-testcode,work/3_2,--runtime runtime )
# ランタイムを使用していることを確認するため runtime.lua を移動しておく
	mkdir -p work/backup
	mv work/3_2/lune/base/runtime*.lua work/backup
	(cd work/3_2; $(LUA) test/start.lua) > work/test4-2.result
# runtime.lua を戻す
	mv work/backup/* work/3_2/lune/base/
	rmdir work/backup

	$(call exeLune,work/3_2,test/start.lns,save,--depends start.dep)

# オプション -u 付きで save して実行テスト
	$(call mk_lune_proj_dir,work/4)
	cp -a test work/4
	cp -a work/3_2/lune work/4
	$(call exeLune,work/3_2,test/start.lns,-u SAVE ../4,)
	(cd work/4; $(LUA) test/start.lua) > work/test5.result



test-parallel: test-org test-org-proj \
			test-new test-new-proj test-new2


test-comp2:
#	$(MAKE) test-org


# 新しいトランコンパイラで再度自分自身をコンパイルした結果が同じになることを確認する。
#	$(call buildLns,work/1,../2)
#	$(call buildLns,work/2,../3)
	$(MAKE) -C work/2 build_from_d depend_on=y FROM_DIR=../1 TO_DIR=../2
	$(MAKE) -C work/3 build_from_d depend_on=y FROM_DIR=../2 TO_DIR=../3
	diff -r -I '_moduleObj.__buildId.*' -I '.*buildId = .*' work/2 work/3


	cp -a test work/3
#	$(call exeLune,work/3,-mklunemod,lune/base)

#	$(MAKE) test-new

	$(call exeLune,work/3,test/start.lns,glue, ../newglue 2>&1)
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




ifeq ($(MAKECMDGOALS),test-comp2)
ifdef LUA51
ifneq ($(shell which $(LUA51)),)
	$(MAKE) test-lua51
endif
endif
endif

ifndef NO_LUAROCKS
	$(MAKE) luarocks-check
endif


	$(call test-meta-update,work/3)

# rockspec の更新が必要かどうか
	$(MAKE) check-rockspec

# メタフォーマットバージョンを上げなくて良いか確認
	$(MAKE) check-meta-format

# diff で、修正前と修正後の差分がないことを確認
	$(MAKE) diff-result

ifdef LUA51
	$(MAKE) accept-package ACCEPT=y ENABLE_LUA51=$(shell which $(LUA51))
else
	$(MAKE) accept-package ACCEPT=y
endif

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
	mv $1/test/glueTest51.so $1/test/glueTest.so 
	grep -v -e 'skip-lua51' $2 test/start.lns > $1/test/start.lns
endef

define setup-lua51-test-cross
	$(call setup-lua51-test,$1,-e 'skip-cross')
endef

test-lua51-mkruntime:
	$(call exeLune,work/5,-mklunemod,lune/base/runtime.lua)

test-lua51:
# LuneScript を Lua 5.1 にクロスコンパイルして実行
	$(MAKE) -C work/5 build_from_d depend_on=y \
			FROM_DIR=../3 TO_DIR=../5 LNS_OPT="-ol 51"
	$(MAKE) test-lua51-mkruntime LUA=$(LUA51)
	$(call setup-lua51-test,work/5)
	(cd work/5; LUA_PATH=./?.lua $(LUA51) lune/base/base.lua test/start.lns -u exe ) > work/test6.result

# 比較用にクロスコンパイル前のテスト結果を取得
	$(MAKE) test-lua51-sub LUA=$(LUA51) 

	diff -r work/5/lune work/6/lune

# サンプルを lua5.1 にクロスコンパイルして実行テスト
	$(call mk_lune_proj_dir,work/7)
	cp -a work/3/lune work/7
	$(call setup-lua51-test-cross,work/7)
	$(call exeLune,work/7,test/start.lns,-u -ol 51 SAVE,)
	(cd work/7; $(LUA51) test/start.lua) | sed 's/\.0//g' > work/test7-3.result

	diff work/test6.result  work/test7.result 
	diff work/test7_2.result  work/test7-3.result


# lua5.1 用にクロスコンパイルしたものから、 Lua5.1 で自分自身をコンパイル
test-lua51-sub:
	$(call mk_lune_proj_dir,work/6)
	$(call cp_src_set,work/6,work/1)
	$(MAKE) -C work/6 build_from_d depend_on=y FROM_DIR=../5 TO_DIR=../6
#	$(call buildLns,work/5,../6)
	$(call setup-lua51-test,work/6)
	$(call exeLune,work/6,-mklunemod,lune/base/runtime.lua)
	$(call exeLune,work/6,test/start.lns,exe,) | sed 's/\.0//g' > work/test7.result

	cp -a work/6 work/6_2
	rm -rf work/6_2/test
	$(call setup-lua51-test-cross,work/6_2)
	$(call exeLune,work/6_2,test/start.lns,exe,) | sed 's/\.0//g' > work/test7_2.result






accept-package:
ifndef ACCEPT
	$(MAKE) -i diff-result
	@read -p "accept? $1 (y/n): " ANS; \
	if [ $${ANS}x != yx ]; then exit 1; fi;
endif
	rm lune/base/*.lua
	cp work/3/lune/base/*.lua lune/base/
	cp work/3/lune/base/*.meta lune/base/
	cp work/3/lune/*.* lune/
ifneq ($(ENABLE_LUA51),)
	cp work/6/lune/base/*.lns legacy/lua51/lune/base/
	cp work/6/lune/base/*.lua legacy/lua51/lune/base/
	cp work/6/lune/base/*.meta legacy/lua51/lune/base/
	cp work/6/lune/*.* legacy/lua51/lune/
endif


test-installed-cmd:
	$(MAKE) LNSC_ENABLE=y LNSC=lnsc test-comp -j

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
	mkdir -p work/0
	cp -a lune work/0
	$(call buildLns,.,work/0,)
	diff -r lune work/0/lune -x '*.meta' -x '*.d' -x '*.*~' -x '*.h' -x '*.c' -x '*.tmp'
	cp work/0/lune/base/*.lua lune/base/
	cp work/0/lune/base/*.meta lune/base/
	cp work/0/lune/*.* lune/

	$(MAKE) test-testC




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
	$(call exeLune,$1,$4.lns,bootC -oc work/boot.c)
	LANG= $(MAKE) -C $1 -f Makefile.ccc "SRCS=$2 work/boot.c" TEST_OUT=test.exe \
		LUA_INC_DIR=$(LUA_INC_DIR) LUA_LIB=$(LUA_LIB)
	@# C へのトランスコンパイラ結果と、
	@# Lua へのトランスコンパイラ結果で出力が同じになることを確認
	@(cd $1; ./test.exe | grep -v -e '^:debug:' > work/$3.c.log)
	$(call exeLune,$1,$4.lns,$5,exe > work/$3.lua.log)
	diff $1/work/$3.c.log $1/work/$3.lua.log
endef

# lns ソースを C へトランスコンパイルする
#
# $1: 実行する作業ディレクトリ
# $2: トランスコンパイルする lns ソースのパス
define convC
	$(call exeLune,$1,$2,save -langC,$3)
endef

build-miniC:
	$(call test-convC-exe,.,miniC.lns,miniC,miniC)

test-miniC:
	$(call transwork,.,.,,base/convCC)
	$(LUA) lune/base/base.lua --mkbuiltin lune/base/runtime_c
	mkdir -p work/0
	rm -rf work/0/*
	mkdir -p work/0/work/c
	cp -a lune work/0
	cp Makefile.ccc work/0
	$(call convC,work/0,lune/base/Testing.lns,--testing)
#	-$(call test-convC-test,work/0,lune/base/Testing)
#	cp -a work/c/* work/0/work/c
#	cp -a work/0/work/* work/
	cp miniC.lns work/0
#	$(call test-convC-test,.,miniC)
	$(call test-convC-test,work/0,miniC)
	cp work/0/test.exe .


test-subC:
	$(call transwork,.,.,,base/convCC)
	$(call convC,.,miniC.lns)
	$(call convC,.,subC.lns)
	$(call test-convC-exe,.,miniC.lns subC.lns,subc,miniC)


build-cc:
	$(MAKE) test-miniC
	$(call convC,.,ccc.lns)
	$(call convC,.,subC.lns)
	$(call test-convC-exe,.,ccc.lns subC.lns,ccc,ccc)

test-formatter:
	$(call transwork,.,.,,base/Formatter)
	$(call exeLune,.,ccc.lns,format,) > work/format.lns

define test-convC-test
	$(call convC,$1,$2.lns,--testing)
	$(call test-convC-exe,$1,$2.c lune/base/Testing.c $3,$(notdir $2),$2,--testing)
endef

test-testC:
	$(call transwork,.,.,,base/convCC)
	$(LUA) lune/base/base.lua --mkbuiltin lune/base/runtime_c
	mkdir -p work/0
	rm -rf work/0/*
	mkdir -p work/0/work/c
	cp -a lune work/0
	cp Makefile.ccc work/0
	$(call convC,work/0,lune/base/Testing.lns,--testing)
	$(call test-convC-test,work/0,lune/base/Ver)
	$(call test-convC-test,work/0,lune/base/Code)
	$(call test-convC-test,work/0,lune/base/LuaMod)
	$(call convC,work/0,lune/base/LuaVer.lns,--testing)
	$(call convC,work/0,lune/base/LuneControl.lns,--testing)
	$(call convC,work/0,lune/base/Depend.lns,--testing)
	$(call test-convC-test,work/0,lune/base/Util, \
		lune/base/Depend.sub.c lune/base/Depend.c \
		lune/base/LuneControl.c lune/base/LuaMod.c )
	$(call test-convC-test,work/0,lune/base/Parser, \
		lune/base/Depend.sub.c lune/base/Depend.c \
		lune/base/LuneControl.c lune/base/LuaMod.c \
		lune/base/Util.c)
	$(call test-convC-test,work/0,lune/base/Writer, \
		lune/base/Depend.sub.c lune/base/Depend.c \
		lune/base/LuneControl.c lune/base/LuaMod.c \
		lune/base/Util.c)
	$(call test-convC-test,work/0,lune/base/Ast, \
		lune/base/Depend.sub.c lune/base/Depend.c \
		lune/base/LuneControl.c lune/base/LuaMod.c \
		lune/base/Util.c lune/base/Writer.c \
		lune/base/Parser.c lune/base/Code.c )
	$(call convC,work/0,lune/base/frontInterface.lns,--testing)

test-mkbuiltin:
	$(call transwork,.,.,,base/convCC)
	$(LUA) lune/base/base.lua --mkbuiltin lune/base/runtime_c
	$(MAKE) test-miniC

lctags:
	$(MAKE) -f Makefile.ccc lctags


test-start:
	$(call transwork,.,.,,base/convCC)
	$(LUA) lune/base/base.lua --mkbuiltin lune/base/runtime_c
	mkdir -p work/0
	rm -rf work/0/*
	mkdir -p work/0/work/c
	cp -a lune work/0
	cp -a test work/0
	cp Makefile.ccc work/0
	$(call convC,work/0,test/Sub4.lns)
	$(call convC,work/0,test/Sub2.lns)
	$(call test-convC-exe,work/0,test/Sub2.lns,test/Sub2,test/Sub2)


define test-tutorialC
	$(call convC,$1,$(basename $2).lns)
	$(call test-convC-exe,$1,$2,$(notdir $2),$(basename $2))
endef


test-tutorial:
	$(call transwork,.,.,,base/convCC)
	$(LUA) lune/base/base.lua --mkbuiltin lune/base/runtime_c
	mkdir -p work/0
	rm -rf work/0/*
	mkdir -p work/0/work/c
	cp -a lune work/0
	cp -a test work/0
	cp Makefile.ccc work/0
	$(call test-tutorialC,work/0,test/tutorial/match.lns)
	$(call test-tutorialC,work/0,test/tutorial/loop.lns)
	$(call test-tutorialC,work/0,test/tutorial/interface.lns)
	$(call test-tutorialC,work/0,test/tutorial/generics.lns)
	$(call test-tutorialC,work/0,test/tutorial/enum.lns)
	$(call test-tutorialC,work/0,test/tutorial/classOverride.lns)
	$(call test-tutorialC,work/0,test/tutorial/classMapping.lns)
	$(call test-tutorialC,work/0,test/tutorial/classInherit.lns)
	$(call test-tutorialC,work/0,test/tutorial/class2accessor.lns)
	$(call test-tutorialC,work/0,test/tutorial/class1.lns)
	$(call test-tutorialC,work/0,test/tutorial/cast.lns)
	$(call test-tutorialC,work/0,test/tutorial/branch.lns)
	$(call test-tutorialC,work/0,test/tutorial/nilable.lns)
	$(call test-tutorialC,work/0,test/tutorial/defaultArg.lns)
	$(call test-tutorialC,work/0,test/tutorial/multipleRetVal.lns)
	$(call test-tutorialC,work/0,test/tutorial/arg.lns)
	$(call test-tutorialC,work/0,test/tutorial/func.lns)
	$(call test-tutorialC,work/0,test/tutorial/variable.lns)
	$(call test-tutorialC,work/0,test/tutorial/set.lns)
	$(call test-tutorialC,work/0,test/tutorial/value.lns)

test-cAll:
	$(MAKE) build-cc
	$(MAKE) test-tutorialC
	$(MAKE) test-testC


test-gc:
	mkdir -p work/c
	$(MAKE) -f Makefile.ccc test-gc

GO_SRC += runtime_main.go
GO_SRC += runtime_lib.go
GO_SRC := $(addprefix lune/base/runtime_go/,$(GO_SRC))

test-go:
	go run $(GO_SRC) lune/base/runtime_go/miniGo.go


# lns ソースを Golang へトランスコンパイルする
#
# $1: 実行する作業ディレクトリ
# $2: トランスコンパイルする lns ソースのパス
define convGo
	$(call exeLune,$1,$2,save -langGo,$3)
endef

test-miniGo:
	$(call transwork,.,.,,base/convGo)
	$(call exeLune,.,lune/base/runtime_go/miniGo.lns,lua -langGo,)
	$(call convGo,.,lune/base/runtime_go/miniGo.lns)
	go build lune/base/runtime_go/miniGo.go $(GO_SRC)
	./miniGo

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
	$(call exeLune,.,test/checkReference.lns exe)
