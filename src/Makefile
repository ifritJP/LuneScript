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

ifdef LOG
LOG_OPT=--log info
endif

ifdef LNSC_ENABLE
ifneq ($(LUA51),)
ENABLE_LUA51:=$(shell which $(LUA51))
endif
LUAROCKS_VER:=$(shell luarocks | grep 'Lua version' | awk '//{print $$3}')
endif

ifneq ($(findstring $(MAKECMDGOALS),"install" "uninstall"),)
# LUA_DIR_SCRIPT='for path in string.gmatch( package.path, "[^;]+" ) do if not path:find( "^%./" ) then print( (path:gsub( "/%?%.lua", "" )) ); break end end'
# LUA_MOD_DIR=$(shell echo $(LUA_DIR_SCRIPT) | $(LUA) 2> /dev/null)
include lune.mk
endif

INSTALL=/usr/bin/install -c

NDIR=$(shell pwd)

.PHONY: help exec build build-glue accept-package
.PHONY: check_lua_mod_dir install uninstall test-ccc test-comp build-cc

help:
	@echo usage:
	@echo "  make install"
	@echo "  make test-comp"
	@echo "  make release-check"
	@echo "  make build-glue [LUA_INC_DIR=path]"

exec:
	$(LUA) lune/base/base.lua test/test.lns token
	$(LUA) lune/base/base.lua test/test.lns ast
	$(LUA) lune/base/base.lua test/test.lns LUA


define cutMeta
	(cd work/$2/lune/base; cat $1.lua | awk '/^----- meta/{META++;}; //{if (META != 1) {print $$0}} ' > $1.cut.lua )
endef

define transwork
	(cd $1; LUA_PATH=./?.lua $(LUA) lune/base/base.lua lune/$3.lns SAVE $2 $(BYTECOMP) -Werror $(LOG_OPT) )
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

define exeComp
	echo lune | cat $3 - | $(call exeLune,$1,$2,comp -i,$4) | \
			json_pp --json_opt=canonical,pretty | sort
endef

define buildLns
	$(call transwork,$1,$2,base/Ver)
	$(call transwork,$1,$2,base/LuaMod)
	$(call transwork,$1,$2,base/Code)
	$(call transwork,$1,$2,base/Log)
	$(call transwork,$1,$2,base/LuaVer)
	$(call transwork,$1,$2,base/Depend)
	$(call transwork,$1,$2,base/Util)
	$(call transwork,$1,$2,base/frontInterface)
	$(call transwork,$1,$2,base/Writer)
	$(call transwork,$1,$2,base/Parser)
	$(call transwork,$1,$2,base/Ast)
	$(call transwork,$1,$2,base/Option)
	$(call transwork,$1,$2,base/Nodes)
	$(call transwork,$1,$2,base/dumpNode)
	$(call transwork,$1,$2,base/TransUnit)
	$(call transwork,$1,$2,base/convLua)
	$(call transwork,$1,$2,base/convCC)
	$(call transwork,$1,$2,base/OutputDepend)
	$(call transwork,$1,$2,base/front)
	$(call transwork,$1,$2,base/glueFilter)
	$(call transwork,$1,$2,Util)
endef


define build-testcode
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
	-(luarocks remove $1 --local)
	cat ../$1-main-1.rockspec | sed 's@"src.*/lune/@"lune/@g' \
			| sed 's@src/lnsc.lua@lnsc.lua@g' > work/8/$1-main-1.rockspec
	cp -a lnsc.lua work/8
	cp -a ../docs work/8
	(cd work/8; luarocks make $1-main-1.rockspec --local)

	$(LNSC)
	(cd work/8; $(LNSC) test/start.lns exe) > work/test8.result

	(luarocks remove $1 --local)
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
	$(call test-inst-from-luarocks,lunescript)
	diff work/test.result work/test8.result
endif

release-check:
	$(MAKE) test-comp LUA_INC_DIR=../../lctags/external/lua/lua-5.3.4/src/
	$(MAKE) test-comp LUA=lua5.2 LUA_INC_DIR=../../lctags/external/lua/lua-5.2.4/src/
#	$(MAKE) test-comp LUA=lua5.1 LUA_INC_DIR=/usr/include/lua5.1/

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
	rm -rf work
	mkdir -p work/glue
	$(call mk_lune_proj_dir,work/1)
	mkdir -p work/1/test
	$(call mk_lune_proj_dir,work/2)
	$(call mk_lune_proj_dir,work/3)
	$(call mk_lune_proj_dir,work/5)
	mkdir -p work/newglue
	@cp lune/base/base*.lua lune/base/*.lns work/1/lune/base
	@cp lune/*.lns work/1/lune
	@cp lune/base/base*.lua lune/base/*.lns work/2/lune/base
	@cp lune/*.lns work/2/lune
	@cp lune/base/base*.lua lune/base/*.lns work/3/lune/base
	@cp lune/*.lns work/3/lune
	@cp lune/base/base*.lua lune/base/*.lns work/5/lune/base
	@cp lune/*.lns work/5/lune
	$(call buildLns,.,work/1)
	$(MAKE) test-comp2

	@echo finish to build 
	@LANG= date

test-etc:
# 存在しないファイルを指定した時に、不正なファイルが作成されないことを確認する
	-@$(call exeLune,$(TEST_DIR),test/error/__.lns,SAVE,2>&1) > /dev/null
	@echo -- not exist file access --
	@(cd $(TEST_DIR); find test/error -iname '__.*')



test-error:
	-$(call exeLune,$(TEST_DIR),test/error/nilable1.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/access1.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/access2.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/access3.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/access4.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/access5.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/list.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/map.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/set.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/scope1.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/scope2.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/func.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/class.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/class2.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/class3.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/class4.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/class5.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/class6.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/provide1.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/provide2.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/provide3.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/shadowing.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/enum.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/eof.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/mutable.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/var.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/arg.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/arg2.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/unwrap.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/operand.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/loop.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/cast.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/prototype.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/statement.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/flow.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/import.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/strformat.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/Alge.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/Alge2.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/Alge3.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/Alge4.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/exp.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/symbol.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/symbolEmpty.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/noexist.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/static.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/generics.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/generics2.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/generics3.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/generics4.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/box.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/form.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/ddd.lns,exe,2>&1)


	-$(call exeLune,$(TEST_DIR),test/error/abbr.lns,exe -Werror,2>&1) && echo error:NG

	cat test/error/mutable.lns | sed 's@//_lune@_lune@' > test/error/dis_mutable.lns
	-$(call exeLune,$(TEST_DIR),test/error/dis_mutable.lns,exe,2>&1)
	rm test/error/dis_mutable.lns

test-comp2:
	$(call exeLune,.,test/start.lns,glue,work/glue 2>&1)
	$(MAKE) build-glue
	$(call exeLune,.,test/start.lns,exe,2>&1) > work/test1.result
	$(call exeLune,.,test/start.lns,save,--depends work/1/start.dep)

	$(call exeLune,.,test/Sub2.lns,SAVE,work/1)


	grep -v 'warning:' work/test1.result > work/test.result
	$(call exeComp,.,test/comp.lns,test/comp.lns,test.comp 17 7) >> work/test1.result
	$(call exeComp,.,test/comp-2.lns,test/comp-2.lns,test.comp-2 17 7) >> work/test1.result
	$(call exeComp,.,test/comp-3.lns,test/comp-3.lns,test.comp-3 17 7) >> work/test1.result
	$(call exeComp,.,test/comp-4.lns,test/comp-4.lns,test.comp-4 17 12) >> work/test1.result
	$(call exeComp,.,test/comp-5.lns,test/comp-5.lns,test.comp-5 17 15) >> work/test1.result
	$(call exeComp,.,test/comp2.lns,test/compSub.lns,test.compSub 5 6) >> work/test1.result
	$(call exeComp,.,test/comp4.lns,test/comp4.lns,test.comp4 9 8) >> work/test1.result
	$(call exeComp,.,test/comp5.lns,test/comp5.lns,test.comp5 7 7) >> work/test1.result
	$(MAKE) -s TEST_DIR=. test-error | grep -e 'error:' -e 'warning:' >> work/test1.result
	$(MAKE) TEST_DIR=. test-etc >> work/test1.result
	$(MAKE) -C test/proj/ test MKFILE=Makefile SRC_DIR= 2>&1 | \
		grep -v /test/proj >> work/test1.result

#		sed 's/.*error:/error:/' >> work/test1.result


# 新しいトランコンパイラで再度自分自身をコンパイルした結果が同じになることを確認する。
	$(call buildLns,work/1,../2)
	$(call buildLns,work/2,../3)
	diff -r -I '_moduleObj.__buildId.*' -I '.*buildId = .*' work/2 work/3

	cp -a test work/3
	$(call exeLune,work/3,-mklunemod,lune/base)


	$(call exeLune,work/3,test/start.lns,glue, ../newglue 2>&1)
	$(MAKE) build-glue GLUE_DIR=work/newglue

	$(call exeLune,work/3,test/start.lns,ast,)
	$(call exeLune,work/3,test/start.lns,lua,)
	$(call exeLune,work/3,test/start.lns,exe,2>&1) | tee work/test2.result
	$(call exeComp,work/3,test/comp.lns,test/comp.lns,test.comp 17 7) | tee -a work/test2.result
	$(call exeComp,work/3,test/comp-2.lns,test/comp-2.lns,test.comp-2 17 7) | tee -a work/test2.result
	$(call exeComp,work/3,test/comp-3.lns,test/comp-3.lns,test.comp-3 17 7) | tee -a work/test2.result
	$(call exeComp,work/3,test/comp-4.lns,test/comp-4.lns,test.comp-4 17 12) | tee -a work/test2.result
	$(call exeComp,work/3,test/comp-5.lns,test/comp-5.lns,test.comp-5 17 15) | tee -a work/test2.result
	$(call exeComp,work/3,test/comp2.lns,test/compSub.lns,test.compSub 5 6) | tee -a work/test2.result
	$(call exeComp,work/3,test/comp4.lns,test/comp4.lns,test.comp4 9 8) | tee -a work/test2.result
	$(call exeComp,work/3,test/comp5.lns,test/comp5.lns,test.comp5 7 7) | tee -a work/test2.result
	$(MAKE) -s TEST_DIR=work/3 test-error | grep -e 'error:' -e 'warning:' | \
				tee -a work/test2.result
	$(MAKE) TEST_DIR=work/3 test-etc >> work/test2.result
	$(MAKE) -C work/3/test/proj/ test MKFILE=Makefile SRC_DIR= 2>&1 | \
			grep -v /test/proj >> work/test2.result
#sed 's/.*error:/error:/' | tee -a work/test2.result





## テストスクリプトを lua 化して、 lunescript なしで動かすテスト。
	$(call build-testcode,work/3)
	(cd work/3; $(LUA) test/start.lua) > work/test3.result

# テストスクリプトを lua 化して、 _lune をロードしてテスト。
	$(call build-testcode,work/3,-r)
	(cd work/3; $(LUA) test/start.lua) > work/test4.result

	$(call exeLune,work/3,test/start.lns,save,--depends start.dep)

# オプション -u 付きで save して実行テスト
	$(call mk_lune_proj_dir,work/4)
	cp -a test work/4
	cp -a work/3/lune work/4
	$(call exeLune,work/3,test/start.lns,-u SAVE ../4,)
	(cd work/4; $(LUA) test/start.lua) > work/test5.result

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
	diff work/test4.result work/test5.result
	diff work/1/start.dep work/3/start.dep
	diff -r work/glue work/newglue


define setup-lua51-test
	cp -a test $1
	$(MAKE) build-glue GLUE_DIR=work/newglue LUA_INC_DIR=$(LUA51_INC) LUA_GLUE=51
	mv $1/test/glueTest51.so $1/test/glueTest.so 
	grep -v skip-lua51 test/start.lns > $1/test/start.lns
endef

test-lua51:
# LuneScript を Lua 5.1 にクロスコンパイルして実行
	$(call buildLns,work/3,../5 -ol 51)
	$(call exeLune,work/5,-mklunemod,lune/base)
	$(call setup-lua51-test,work/5)
	(cd work/5; LUA_PATH=./?.lua $(LUA51) lune/base/base.lua test/start.lns -u exe ) > work/test6.result

# 比較用にクロスコンパイル前のテスト結果を取得
	$(MAKE) test-lua51-sub LUA=$(LUA51) 

	diff -r work/5/lune work/6/lune

# サンプルを lua5.1 にクロスコンパイルして実行テスト
	$(call mk_lune_proj_dir,work/7)
	cp -a work/3/lune work/7
	$(call setup-lua51-test,work/7)
	$(call exeLune,work/7,test/start.lns,-u -ol 51 SAVE,)
	(cd work/7; $(LUA51) test/start.lua) | sed 's/\.0//g' > work/test7-2.result

	diff work/test6.result  work/test7.result 
	diff work/test7.result  work/test7-2.result


test-lua51-sub:
	$(call mk_lune_proj_dir,work/6)
	@cp lune/base/base.lua lune/base/_lune.lua lune/base/*.lns work/6/lune/base
	@cp lune/*.lns work/6/lune
	$(call buildLns,work/5,../6)
	$(call setup-lua51-test,work/6)
	$(call exeLune,work/6,-mklunemod,lune/base)
	$(call exeLune,work/6,test/start.lns,exe,) | sed 's/\.0//g' > work/test7.result



accept-package:
ifndef ACCEPT
	$(MAKE) -i diff-result
	@read -p "accept? $1 (y/n): " ANS; \
	if [ $${ANS}x != yx ]; then exit 1; fi;
endif
	cp work/3/lune/base/*.lua lune/base/
	cp work/3/lune/base/*.meta lune/base/
	cp work/3/lune/*.* lune/
ifneq ($(ENABLE_LUA51),)
	cp work/6/lune/base/*.lns legacy/lua51/lune/base/
	cp work/6/lune/base/*.lua legacy/lua51/lune/base/
	cp work/6/lune/base/*.meta legacy/lua51/lune/base/
	cp work/6/lune/*.* legacy/lua51/lune/
endif

ifneq ($(findstring $(MAKECMDGOALS),"check-meta-format" ),)
VER_PICK=grep 'local metaVersion' | awk '//{print $$4}'
ORG_VER=$(shell cat lune/base/Ver.lua | $(VER_PICK))
NEW_VER=$(shell cat work/3/lune/base/Ver.lua | $(VER_PICK))
endif

check-meta-format:
# メタフォーマットバージョンが同じ場合は、メタ情報が一致しているかどうか比較
ifeq ($(ORG_VER),$(NEW_VER))
	@diff work/1/test/Sub2.meta work/4/test/Sub2.meta > /dev/null || \
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
	@$(INSTALL) $(LNSDIR)/lune/base/* $(LUA_MOD_DIR)/lune/base
	@$(INSTALL) $(LNSDIR)/lune/*.* $(LUA_MOD_DIR)/lune
	@echo installed -- $(LUA_MOD_DIR)/lune

uninstall: check_lua_mod_dir
	rm /usr/bin/lnsc
	rm -rf $(LUA_MOD_DIR)/lune

test-ccc:
	LANG= $(MAKE) -f Makefile.ccc LUA_INC_DIR=$(LUA_INC_DIR) LUA_LIB=$(LUA_LIB)

build-cc:
	$(call transwork,.,.,base/convCC)
	$(LUA) lune/base/base.lua ccc.lns save -langC
	cat ccc.c
	LANG= $(MAKE) -f Makefile.ccc SRCS=ccc.c TEST_OUT=test.exe \
		LUA_INC_DIR=$(LUA_INC_DIR) LUA_LIB=$(LUA_LIB)
# C へのトランスコンパイラ結果と、
# Lua へのトランスコンパイラ結果で出力が同じになることを確認
	@./test.exe | grep -v -e '^:debug:' > work/ccc.c.log
	@$(LUA) lune/base/base.lua ccc.lns exe > work/ccc.lua.log
	diff work/ccc.c.log work/ccc.lua.log
