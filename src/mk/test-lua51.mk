define setup-lua51-test
	cp -a test $1
	$(MAKE) build-glue GLUE_DIR=work/newglue LUA_INC_DIR=$(LUA51_INC) LUA_GLUE=51
	mv test/glueTest51.so $1/test/glueTest.so 
	grep -v -e 'skip-lua51' $2 test/start.lns > $1/test/start.lns
	grep -v -e 'skip-lua51' $2 test/Sub11.lns > $1/test/Sub11.lns
endef

define setup-lua51-test-cross
	$(call setup-lua51-test,$1,-e 'skip-cross')
endef

test-lua51-mkruntime:
	$(call exeLune,$(LNSC_LUA51),work/5,-mklunemod lune/base/runtime.lua)

test-lua51:
# LuneScript を Lua 5.1 にクロスコンパイルして実行
	$(MAKE) -C work/5 build_from_d depend_on=y \
			FROM_DIR=../3 TO_DIR=../5 LNS_OPT="-ol 51 --package lnsc"
	$(MAKE) test-lua51-mkruntime LUA=$(LUA51)
	$(call setup-lua51-test,work/5)
	(cd work/5; LUA_PATH=./?.lua $(LUA51) lune/base/base.lua test/start.lns -u exe $(EXE_OPT)) > work/test6.result

# 比較用にクロスコンパイル前のテスト結果を取得
	$(MAKE) test-lua51-sub LANG_KIND=lua LUA=$(LUA51) 

	diff -u -r work/5/lune work/6/lune -x binding.go

# gopher-lua で test-go 実行
	(cd work/6; $(MAKE) test-go LUA=$(LUA51) LNSC="$(LNSC_LUA51)" GOTAG=gopherlua SKIP_SETUP=y SETUP_LUA=y)

# サンプルを lua5.1 にクロスコンパイルして実行テスト
	$(call mk_lune_proj_dir,work/7)
	cp -a work/3/lune work/7
ifeq ($(WITH_GO),y)
	cp $(LNSC_GO) work/7
endif
	$(call setup-lua51-test-cross,work/7)
	$(call exeLune,$(LNSC_LUA),work/7,test/start.lns,-u -ol 51 SAVE,$(EXE_OPT))
	(cd work/7; $(LUA51) test/start.lua) | sed 's/\.0//g' > work/test7-3.result

	diff work/test6.result  work/test7.result 
	diff work/test7_2.result  work/test7-3.result


# lua5.1 用にクロスコンパイルしたものから、 Lua5.1 で自分自身をコンパイル
test-lua51-sub:
	$(call mk_lune_proj_dir,work/6)
	$(call cp_src_set,work/6,work/1)
	$(MAKE) -C work/6 build_from_d depend_on=y \
		FROM_DIR=../5 TO_DIR=../6 LNS_OPT="--package lnsc"
	$(MAKE) -C work/6 binding-lns-go-tgt
#	$(call selfhost_build,work/5,../6)
	$(call setup-lua51-test,work/6)
	$(call exeLune,$(LNSC),work/6,-mklunemod lune/base/runtime.lua)
	$(call exeLune,$(LNSC),work/6,test/start.lns,exe,$(EXE_OPT)) | sed 's/\.0//g' > work/test7.result

	cp -a work/6 work/6_2
	rm -rf work/6_2/test
	$(call setup-lua51-test-cross,work/6_2)
	$(call exeLune,$(LNSC),work/6_2,test/start.lns,exe,$(EXE_OPT)) | sed 's/\.0//g' > work/test7_2.result






