LUA=lua5.3
#LUA=lua5.2



LUA51=lua5.1
LUA51_INC=/usr/include/lua5.1

LUA_INC_DIR=
ifneq ($(LUA_INC_DIR),)
LUA_CFLAGS=-I$(LUA_INC_DIR)
endif

ifneq ($(findstring $(MAKECMDGOALS),"install" "uninstall"),)
# LUA_DIR_SCRIPT='for path in string.gmatch( package.path, "[^;]+" ) do if not path:find( "^%./" ) then print( (path:gsub( "/%?%.lua", "" )) ); break end end'
# LUA_MOD_DIR=$(shell echo $(LUA_DIR_SCRIPT) | $(LUA) 2> /dev/null)
include lune.mk
endif

INSTALL=/usr/bin/install -c

NDIR=$(shell pwd)

.PHONY: help exec build build-glue accept-package check_lua_mod_dir install uninstall

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
	(cd $1; LUA_PATH=./?.lua $(LUA) lune/base/base.lua lune/base/$3.lns SAVE $2 $(BYTECOMP))
endef

define exeLune
	(cd $1; LUA_PATH=./?.lua $(LUA) lune/base/base.lua $2 $3 $4)
endef

define exeComp
	echo lune | cat $3 - | $(call exeLune,$1,$2,comp -i,$4) | \
			json_pp --json_opt=canonical,pretty | sort
endef

define buildLns
	$(call transwork,$1,$2,Ver)
	$(call transwork,$1,$2,LuaMod)
	$(call transwork,$1,$2,Depend)
	$(call transwork,$1,$2,Util)
	$(call transwork,$1,$2,LuaVer)
	$(call transwork,$1,$2,frontInterface)
	$(call transwork,$1,$2,Writer)
	$(call transwork,$1,$2,Parser)
	$(call transwork,$1,$2,Option)
	$(call transwork,$1,$2,Ast)
	$(call transwork,$1,$2,TransUnit)
	$(call transwork,$1,$2,convLua)
	$(call transwork,$1,$2,dumpNode)
	$(call transwork,$1,$2,OutputDepend)
	$(call transwork,$1,$2,front)
	$(call transwork,$1,$2,glueFilter)
endef


define build-testcode
	$(call exeLune,work/3,test/Sub2.lns,$1 save,) > /dev/null
	$(call exeLune,work/3,test/Sub3.lns,$1 save,) > /dev/null
	$(call exeLune,work/3,test/Sub3_2.lns,$1 save,) > /dev/null
	$(call exeLune,work/3,test/Sub4.lns,$1 save,) > /dev/null
	$(call exeLune,work/3,test/Sub5.lns,$1 save,) > /dev/null
	$(call exeLune,work/3,test/Sub6.lns,$1 save,) > /dev/null
	$(call exeLune,work/3,test/simple.lns,$1 save,) > /dev/null
	$(call exeLune,work/3,test/Class1.lns,$1 save,) > /dev/null
	$(call exeLune,work/3,test/Class2.lns,$1 save,) > /dev/null
	$(call exeLune,work/3,test/Class3.lns,$1 save,) > /dev/null
	$(call exeLune,work/3,test/Class4.lns,$1 save,) > /dev/null
	$(call exeLune,work/3,test/start.lns,$1 save,) > /dev/null
	$(call exeLune,work/3,test/mapping.lns,$1 save,) > /dev/null
	$(call exeLune,work/3,test/funcSym.lns,$1 save,) > /dev/null
endef

# meta 情報の更新時に正常に解析できるかテスト。
define test-meta-update
	-(cd $1; rm test/ext/*.lua test/ext/*.meta )
	$(call exeLune,$1,test/ext/sub3.lns,SAVE,)
	$(call exeLune,$1,test/ext/sub2.lns,SAVE,)
	(cd $1; rm test/ext/sub3.lua test/ext/sub3.meta )
	$(call exeLune,$1,test/ext/sub1.lns,SAVE,)
endef

release-check:
	$(MAKE) test-comp LUA_INC_DIR=../../lctags/external/lua/lua-5.3.4/src/
	$(MAKE) test-comp LUA=lua5.2 LUA_INC_DIR=../../lctags/external/lua/lua-5.2.4/src/
#	$(MAKE) test-comp LUA=lua5.1 LUA_INC_DIR=/usr/include/lua5.1/

build:
	$(call buildLns,$(LNSDIR),.)

test-save:
	$(call test-meta-update,.)


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
# - 上記の比較が全て同じなら、 candidate を orginal とする。
test-comp:
	rm -rf work
	mkdir -p work/glue
	mkdir -p work/1/lune/base
	mkdir -p work/1/test
	mkdir -p work/2/lune/base
	mkdir -p work/3/lune/base
	mkdir -p work/5/lune/base
	mkdir -p work/newglue
	@cp lune/base/base*.lua lune/base/*.lns work/1/lune/base
	@cp lune/base/base*.lua lune/base/*.lns work/2/lune/base
	@cp lune/base/base*.lua lune/base/*.lns work/3/lune/base
	@cp lune/base/base*.lua lune/base/*.lns work/5/lune/base
	$(call buildLns,.,work/1)
	$(MAKE) test-comp2


test-error:
	-$(call exeLune,$(TEST_DIR),test/error/nilable1.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/access1.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/access2.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/access3.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/access4.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/access5.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/list.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/map.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/scope1.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/scope2.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/func.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/class.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/provide1.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/provide2.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/provide3.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/shadowing.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/enum.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/eof.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/mutable.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/var.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/arg.lns,exe,2>&1)
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
	$(MAKE) -C test/proj/ test MKFILE=Makefile SRC_DIR= | \
		grep -v test/proj >> work/test1.result

#		sed 's/.*error:/error:/' >> work/test1.result


# 新しいトランコンパイラで再度自分自身をコンパイルした結果が同じになることを確認する。
	$(call buildLns,work/1,../2)
	$(call buildLns,work/2,../3)
	diff -r work/2 work/3

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
	$(MAKE) -C work/3/test/proj/ test MKFILE=Makefile SRC_DIR= | \
			grep -v test/proj >> work/test2.result
#sed 's/.*error:/error:/' | tee -a work/test2.result





## テストスクリプトを lua 化して、 lunescript なしで動かすテスト。
	$(call build-testcode)
	(cd work/3; $(LUA) test/start.lua) > work/test3.result

# テストスクリプトを lua 化して、 _lune をロードしてテスト。
	$(call build-testcode,-r)
	(cd work/3; $(LUA) test/start.lua) > work/test4.result

	$(call exeLune,work/3,test/start.lns,save,--depends start.dep)

# オプション -u 付きで save して実行テスト
	mkdir -p work/4
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


	$(call test-meta-update,work/3)

# メタフォーマットバージョンを上げなくて良いか確認
	$(MAKE) check-meta-format

# diff で、修正前と修正後の差分がないことを確認
	$(MAKE) diff-result

ifdef LUA51
	$(MAKE) accept-package ACCEPT=y ENABLE_LUA51=$(shell which $(LUA51))
else
	$(MAKE) accept-package ACCEPT=y
endif


diff-result:
	diff work/test1.result work/test2.result
	diff work/test.result work/test3.result
	diff work/test3.result work/test4.result
	diff work/test4.result work/test5.result
	diff work/1/start.dep work/3/start.dep
	diff -r work/glue work/newglue


test-lua51:
# LuneScript を Lua 5.1 にクロスコンパイルして実行
	$(call buildLns,work/3,../5 -ol 51)
	$(call exeLune,work/5,-mklunemod,lune/base)
	$(MAKE) build-glue GLUE_DIR=work/newglue LUA_INC_DIR=$(LUA51_INC) LUA_GLUE=51
	cp -a test work/5
	mv work/5/test/glueTest51.so work/5/test/glueTest.so 
	grep -v skip-lua51 test/start.lns > work/5/test/start.lns
	(cd work/5; LUA_PATH=./?.lua $(LUA51) lune/base/base.lua test/start.lns -u exe ) > work/test6.result

# 比較用にクロスコンパイル前のテスト結果を取得
	$(MAKE) test-lua51-sub LUA=$(LUA51) 

	diff -r work/5/lune work/6/lune

# サンプルを lua5.1 にクロスコンパイルして実行テスト
	mkdir -p work/7
	cp -a work/3/lune work/7
	$(MAKE) build-glue GLUE_DIR=work/newglue LUA_INC_DIR=$(LUA51_INC) LUA_GLUE=51
	cp -a test work/7
	mv work/7/test/glueTest51.so work/7/test/glueTest.so 
	grep -v skip-lua51 test/start.lns > work/7/test/start.lns
	$(call exeLune,work/7,test/start.lns,-u -ol 51 SAVE,)
	(cd work/7; $(LUA51) test/start.lua) | sed 's/\.0//g' > work/test8.result

	diff work/test6.result  work/test7.result 
	diff work/test7.result  work/test8.result


test-lua51-sub:
	mkdir -p work/6/lune/base
	@cp lune/base/base.lua lune/base/_lune.lua lune/base/*.lns work/6/lune/base
	$(call buildLns,work/5,../6)
	cp -a test work/6
	mv work/6/test/glueTest51.so work/6/test/glueTest.so 
	grep -v skip-lua51 test/start.lns > work/6/test/start.lns
	$(call exeLune,work/6,-mklunemod,lune/base)
	$(call exeLune,work/6,test/start.lns,exe,) | sed 's/\.0//g' > work/test7.result



accept-package:
ifndef ACCEPT
	@read -p "accept? $1 (y/n): " ANS; \
	if [ $${ANS}x != yx ]; then exit 1; fi;
endif
	cp work/3/lune/base/*.lua lune/base/
	cp work/3/lune/base/*.meta lune/base/
ifneq ($(ENABLE_LUA51),)
	cp work/6/lune/base/*.lns legacy/lua51/lune/base/
	cp work/6/lune/base/*.lua legacy/lua51/lune/base/
	cp work/6/lune/base/*.meta legacy/lua51/lune/base/
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
	@echo installed -- $(LUA_MOD_DIR)/lune

uninstall: check_lua_mod_dir
	rm /usr/bin/lnsc
	rm -rf $(LUA_MOD_DIR)/lune
