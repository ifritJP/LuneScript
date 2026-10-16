define selfhost_buildOp_batch
	echo $(LNS_SRCS) | sed 's/$$/.lns\n/g' | sed 's/ /.lns\n/g' |	\
		(cd $1; $2 @- SAVE $3 $(BYTECOMP)			\
			--package lnsc $(COM_OPT) $(LOG_OPT) $4 )
endef


define selfhost_buildOp
	$(call $5,$1,$2,$3,$4,base/Async)
	$(call $5,$1,$2,$3,$4,base/Runner)
	$(call $5,$1,$2,$3,$4,base/Str)
	$(call $5,$1,$2,$3,$4,base/Types)
	$(call $5,$1,$2,$3,$4,base/Ver)
	$(call $5,$1,$2,$3,$4,base/LuaMod)
	$(call $5,$1,$2,$3,$4,base/LuneControl)
	$(call $5,$1,$2,$3,$4,base/Meta)
	$(call $5,$1,$2,$3,$4,base/Testing)
	$(call $5,$1,$2,$3,$4,base/TestingMacro)
	$(call $5,$1,$2,$3,$4,base/Code)
	$(call $5,$1,$2,$3,$4,base/Log)
	$(call $5,$1,$2,$3,$4,base/LuaVer)
	$(call $5,$1,$2,$3,$4,base/Depend)
	$(call $5,$1,$2,$3,$4,base/Util)
	$(call $5,$1,$2,$3,$4,base/frontInterface)
	$(call $5,$1,$2,$3,$4,base/Writer)
	$(call $5,$1,$2,$3,$4,base/AsyncTokenizer)
	$(call $5,$1,$2,$3,$4,base/Tokenizer)
	$(call $5,$1,$2,$3,$4,base/Json)
	$(call $5,$1,$2,$3,$4,base/Ast)
	$(call $5,$1,$2,$3,$4,base/Option)
	$(call $5,$1,$2,$3,$4,base/GoMod)
	$(call $5,$1,$2,$3,$4,base/Nodes)
	$(call $5,$1,$2,$3,$4,base/NodeIndexer)
	$(call $5,$1,$2,$3,$4,base/Formatter)
	$(call $5,$1,$2,$3,$4,base/Macro)
	$(call $5,$1,$2,$3,$4,base/TransUnitIF)
	$(call $5,$1,$2,$3,$4,base/BuiltinTransUnit)
	$(call $5,$1,$2,$3,$4,base/Builtin)
	$(call $5,$1,$2,$3,$4,base/AstInfo)
	$(call $5,$1,$2,$3,$4,base/Import)
	$(call $5,$1,$2,$3,$4,base/TransUnit)
	$(call $5,$1,$2,$3,$4,base/DependLuaOnLns)
	$(call $5,$1,$2,$3,$4,base/convLua)
	$(call $5,$1,$2,$3,$4,base/dumpNode)
	$(call $5,$1,$2,$3,$4,base/convGo)
	$(call $5,$1,$2,$3,$4,base/convPython)
	$(call $5,$1,$2,$3,$4,base/WasmBin)
	$(call $5,$1,$2,$3,$4,base/convWasm)
	$(call $5,$1,$2,$3,$4,base/OutputDepend)
	$(call $5,$1,$2,$3,$4,base/Converter)
	$(call $5,$1,$2,$3,$4,base/glueFilter)
	$(call $5,$1,$2,$3,$4,base/front)
endef

define mk_runtime
	$(call exeLune,$1,$2,$3/$4,-mklunemod,lune/base/runtime.lua $5) > /dev/null
	cat $2/$3/$(VER_PATH) | $(RUNTIME_VER_PICK) | \
		awk '//{printf( "$2/$3/lune/base/runtime%s.lua",$$0);}' | \
		xargs cp $2/$3/lune/base/runtime.lua 
endef

define selfhost_build
	$(call selfhost_buildOp,$1,$2,$3,$4,cmd_exec_trans)
	$(call mk_runtime,$1,$2,$3,$4)
endef

define selfhost_build_dep
	$(call selfhost_buildOp,$1,$2,$3,$4 $5,cmd_exec_trans_dep)
	$(call mk_runtime,$1,$2,$3,$4 $5)
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
	$(call exeLune,$1,$2,test/Sub11.lns,$3 save,$4) > /dev/null
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

	-$(LNSC)
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
		-iname $2 -exec cp -p {} $3 \;
endef

ifeq ($(WITH_GO),y)
define dump-lns-lua-ver
	-($1 --version | sed 's/.*Lua\([0-9\.]*\).*/\1/g') > $2
endef
else
define dump-lns-lua-ver
	echo skip dump-lns-lua-ver
endef
endif


# lns のビルドに必要なソース群をコピーする
#
# $1: コピー先ディレクトリ
# $2: コピー元ディレクトリ
define cp_src_set
	mkdir -p $1
	cp $2/Makefile $1/
	cp -a $2/mk $1/
	$(call mk_lune_proj_dir,$1)
	$(call cp_src_set_sub,$2/lune/base,'base*.lua',$1/lune/base)
	$(call cp_src_set_sub,$2/lune/base,'runtime*.lua',$1/lune/base)
	$(call cp_src_set_sub,$2/lune/base,'*.lns',$1/lune/base)
	$(call cp_src_set_sub,$2/lune/base,'*.d',$1/lune/base)
#	@cp lune/base/base*.lua lune/base/*.lns lune/base/*.d $1/lune/base
	$(call cp_src_set_sub,$2/lune,'*.lns',$1/lune)
	$(call cp_src_set_sub,$2/lune,'*.d',$1/lune)
	$(call cp_src_set_sub,$2/lune,'go.*',$1/lune)
	cp $2/lune.js $1

	cp -au $2/lune/base/runtime_go $1/lune/base/runtime_go
	cp -au $2/lune/base/binding_dummy.go $1/lune/base
	cp -au $2/lune/base/DependLuaOnLns.go $1/lune/base

	mkdir -p $1/misc/go
	cp -au $2/misc/go/binding $1/misc/go
	mkdir -p $1/misc/go/main
	cp -au $(MAIN_GO) $1/misc/go/main

	$(call dump-lns-lua-ver,$(LNSC_GO),$1/ver_lns_lua)
	$(call dump-lns-lua-ver,$(LNSC_LUA),$1/ver_lua)
endef

