ifeq ($(WITH_GO),y)
# lnsc がシングルファイルでも、バインドしているコードがロードされることを確認する
define check_singleBin
	mkdir -p work/$3
	cp $1 work/$3
# 並列に動く test-error が test/error に一時ファイルを作って消すので、
# コピー中に消えたファイル (rsync の終了コード 24) は無視する
	rsync -a test work/$3 || [ $$? = 24 ]
	(cd work/$3; ./lnsc test/testing.lns exe --testing 2>&1) >> $2
endef
else
define check_singleBin
endef
endif



define cutMeta
	(cd work/$2/lune/base; cat $1.lua | awk '/^----- meta/{META++;}; //{if (META != 1) {print $$0}} ' > $1.cut.lua )
endef


define exeLune
	(cd $2; $1 $(COMMON_OP) $3 $4 $5)
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


