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
	$(call test-convC-test,$(LNSC),work/0,lune/base/Tokenizer, \
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
		lune/base/Tokenizer.c lune/base/Code.c )
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


