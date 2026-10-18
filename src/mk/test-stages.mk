# オリジナルの LuneScript のテストを行ない、結果を work/test1.resut に出力する
test-org:
	$(call exeLune,$(LNSC),.,test/start.lns,glue,work/glue $(EXE_OPT) 2>&1)
	$(MAKE) build-glue
	$(call exeLune,$(LNSC),.,test/start.lns,ast,$(EXE_OPT) 2>&1) > work/ast1.result
	$(call exeLune,$(LNSC),.,test/start.lns,lua,$(EXE_OPT) ) > work/lua1.result
# フルパス指定で動作することを確認
	$(call exeLune,$(LNSC),.,$(CUR_DIR)/test/start.lns,exe,$(EXE_OPT) ) > work/test1.result
	$(call exeLune,$(LNSC),.,test/start.lns,save,$(EXE_OPT) --depends work/1/start.dep)
	$(call exeLune,$(LNSC_LUA),.,test/imp.lns,exe,$(EXE_OPT) 2>&1) >> work/test1.result

	$(call exeLune,$(LNSC),.,test/Sub2.lns,SAVE,work/1 $(EXE_OPT))

	grep -v -e 'warning:' -e 'skip:' work/test1.result > work/test.result

	mkdir -p work/0/log/comp
	$(MAKE) test-cmpl TEST_DIR=. LOG_DIR=work/0/log/comp LOG_FILE=work/test1.result

	$(MAKE) -s TEST_DIR=. LOG_DIR=work/0/log test-error | \
		grep -e 'error:' -e 'warning:' >> work/test1.result
	$(MAKE) -s TEST_DIR=. LOG_DIR=work/0/log test-error \
		COMMON_OP=--disableGenInherit | \
		grep -e 'error:' -e 'warning:' > work/org-error.result
	LANG= $(MAKE) TEST_DIR=. test-etc |		\
		grep -v 'Entering directory' |		\
		grep -v 'Leaving directory' >> work/test1.result
	$(call exeFront,.) >> work/test1.result
ifeq ($(WITH_GO),y)
	$(call test-shebang,.,a b c,work/test1.result)
	$(call test-shebang,.,,work/test1.result)
endif
ifndef SKIP_CHECKING_TESTING
# 現状、CI 環境では --testing が正常が動かないので、ここは skip する
	echo "test -- testing.lns" >> work/test1.result
	$(call exeLune,$(LNSC),.,test/testing.lns,exe, --testing $(EXE_OPT) 2>&1) >> work/test1.result
	echo "test -- TestMini.lns" >> work/test1.result
	$(call exeLune,$(LNSC),.,test/TestMini.lns,exe,--testing $(EXE_OPT) 2>&1) >> work/test1.result
	echo "end -- testing" >> work/test1.result
endif
	$(call check_singleBin,lnsc,work/test1.result,testpack1)

# work/1-1 に start.lns を -u 指定で meta ファイルを生成する
	cp -a work/1 work/0-1
	cp -a lune work/0-1/lune
ifeq ($(WITH_GO),y)
	cp ./lnsc work/0-1
endif
	$(call exeLune,$(LNSC),work/0-1,test/start.lns,-u SAVE,$(EXE_OPT)) > /dev/null



test-org-proj:
ifndef SKIP_PROJ_TEST
	$(MAKE) -j1 -C test/proj/ test MKFILE=Makefile LUA=$(LUA) SRC_DIR= 2>&1 | \
		grep -v /test/proj > work/0/log/proj
else
	@echo skip
endif





# 新しくビルドした LuneScript のテストを行ない、結果を work/test2.resut に出力する
test-new:
#	$(call exeLune,work/3,test/start.lns,glue, ../newglue 2>&1)
#	$(MAKE) build-glue GLUE_DIR=work/newglue

	$(call exeLune,$(LNSC),work/3,test/start.lns,ast,$(EXE_OPT) 2>&1 ) > work/ast2.result
	$(call exeLune,$(LNSC),work/3,test/start.lns,lua,$(EXE_OPT) ) > work/lua2.result
# フルパス指定で動作することを確認
	$(call exeLune,$(LNSC),work/3,$(CUR_DIR)/work/3/test/start.lns,exe,$(EXE_OPT) ) > work/test2.result

# meta ファイルを生成する
	-rm work/3/test/imp.meta
	$(call exeLune,$(LNSC),work/3,test/imp.lns,SAVE,$(EXE_OPT) -u) > /dev/null
# meta ファイルから正常に import できることを確認する
	-rm work/3/test/imp.meta
	$(call exeLune,$(LNSC),work/3,test/imp.lns,exe,$(EXE_OPT) -u) >> work/test2.result


	mkdir -p work/3/log/comp
	$(MAKE) test-cmpl TEST_DIR=work/3 LOG_DIR=work/3/log/comp LOG_FILE=work/test2.result

	$(MAKE) -s TEST_DIR=work/3 LOG_DIR=work/3/log  test-error | \
		grep -e 'error:' -e 'warning:' >> work/test2.result
	$(MAKE) -s TEST_DIR=work/3 LOG_DIR=work/3/log  test-error \
		COMMON_OP=--disableGenInherit | \
		grep -e 'error:' -e 'warning:' > work/new-error.result
	LANG= $(MAKE) TEST_DIR=work/3 test-etc |	\
		grep -v 'Entering directory' |		\
		grep -v 'Leaving directory' >> work/test2.result
	$(call exeFront,work/3) >> work/test2.result
ifeq ($(WITH_GO),y)
	$(call test-shebang,work/3,a b c,work/test2.result)
	$(call test-shebang,work/3,,work/test2.result)
endif

ifndef SKIP_CHECKING_TESTING
# 現状、CI 環境では --testing が正常が動かないので、ここは skip する
	echo "test -- testing.lns" >> work/test2.result
	$(call exeLune,$(LNSC),work/3,test/testing.lns,exe,--testing $(EXE_OPT) 2>&1) >> work/test2.result
	echo "test -- TestMini.lns" >> work/test2.result
	$(call exeLune,$(LNSC),work/3,test/TestMini.lns,exe,--testing $(EXE_OPT) 2>&1) >> work/test2.result
	echo "end -- testing" >> work/test2.result
endif
	$(call check_singleBin,work/3/lnsc,work/test2.result,testpack2)



test-new-proj:
ifndef SKIP_PROJ_TEST
	mkdir -p work/3/log
	$(MAKE) -j1 -C work/3/test/proj/ test MKFILE=Makefile LUA=$(LUA) SRC_DIR= 2>&1 | \
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
	$(call exeLune,$(LNSC),work/3_2,-mklunemod test_runtime.lua) > /dev/null
	$(call build-testcode,$(LNSC),work/3_2,--runtime test_runtime )
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
	-rm work/4/test/error/dis_mutable.lns
	cp -a work/3_2/lune work/4
	$(call exeLune,$(LNSC),work/3_2,test/start.lns,-u SAVE ../4 $(EXE_OPT),)
	(cd work/4; $(LUA) test/start.lua) > work/test5.result
# .meta がある状態で変換
	rm -f work/3_2/test/start.meta work/3_2/test/start.lua
	rm -f work/3_2/test/simple.meta work/3_2/test/simple.lua
	$(call exeLune,$(LNSC),work/3_2,test/start.lns,lua > /dev/null,)


# batch モードでビルドして、
# ビルド前後で同じになるか確認する
test-new3:
	mkdir -p work/3_3
	cp -a work/1/* work/3_3
	find work/3_3/lune -name '*.meta' | grep -v runtime_go | xargs rm
	$(MAKE) -C work/3_3 test-new3-sub
	@date +'=========== complete test-new3 %c============'

# lune 以下を batch ビルドし、 $2 に出力する
define test-new3-batch
	mkdir -p $2
	cp -a $1/lune $2
	cp -a $1/misc $2
	-(find $2/lune -name '*.meta' | grep -v runtime_go | xargs rm)
	$(call selfhost_buildOp_batch,.,$(LNSC),$2,--depends $2/ $3)
endef



test-new3-sub:
# カレントから work/3_3_2 に batch ビルド
	$(call test-new3-batch,.,work/3_3_2,$(TEST_OPT))
ifeq ($(WITH_GO),y)
	$(call go-build-lns,work/3_3_2)
	cp work/3_3_2/lnsc .
endif

# work/3_3_2 から work/3_3_3 に batch ビルド
	$(call test-new3-batch,work/3_3_2,work/3_3_3,$(TEST_OPT))
# ifeq ($(WITH_GO),y)
# 	$(call go-build-lns,work/3_3_3)
# 	cp work/3_3_3/lnsc .
# endif

# 出力結果の確認
	diff -u -r work/3_3_2/lune work/3_3_3/lune -x '*.lns'




# test-parallel / test-serial で実行するテスト
TEST_STAGE_TARGETS = test-org test-org-proj \
			test-comp-go-new-2 test-comp-go-org-2 \
			test-new test-new-proj test-new2 \
			test-comp-go-new test-comp-go-org \
			_test-testing_new \
			test-formatter

# テストの出力とその結果 (終了コード)
TEST_LOG_DIR = work/log

# logged-<テスト>: テストを実行し、出力を $(TEST_LOG_DIR)/<テスト>.log に保存する。
#
# -j で並列に実行しても、どのテストの出力か分かるように、端末には行頭に [<テスト>] を付けて出力する。
# 失敗しても止めずに終了コードを $(TEST_LOG_DIR)/<テスト>.status に記録し、
# 全てのテストの後で test-stage-result が失敗したテストをまとめて表示する。
logged-%:
	@mkdir -p $(TEST_LOG_DIR)
	@{ $(MAKE) $* 2>&1; echo $$? > $(TEST_LOG_DIR)/$*.status; } |	\
		tee $(TEST_LOG_DIR)/$*.log | sed -u 's/^/[$*] /'
	@if [ "$$(cat $(TEST_LOG_DIR)/$*.status)" = 0 ]; then		\
		echo "ok -- $*";						\
	else								\
		echo "NG -- $* (exit $$(cat $(TEST_LOG_DIR)/$*.status), log: $(TEST_LOG_DIR)/$*.log)"; \
	fi

# test-parallel / test-serial の結果をまとめて表示し、失敗したテストがあれば失敗する
test-stage-result:
	@ng="";								\
	for target in $(TEST_STAGE_TARGETS); do				\
		if [ "$$(cat $(TEST_LOG_DIR)/$$target.status 2>/dev/null)" != 0 ]; then \
			ng="$$ng $$target";					\
		fi;							\
	done;								\
	if [ "$$ng" != "" ]; then					\
		echo "========== NG:$$ng";					\
		for target in $$ng; do					\
			echo "---------- $$target ($(TEST_LOG_DIR)/$$target.log)"; \
			tail -n 20 $(TEST_LOG_DIR)/$$target.log;		\
		done;							\
		exit 1;							\
	fi;								\
	echo "========== ok: $(TEST_STAGE_TARGETS)"

test-parallel: $(addprefix logged-,$(TEST_STAGE_TARGETS))
	@$(MAKE) -s test-stage-result
test-serial:
	@for target in $(TEST_STAGE_TARGETS); do			\
		$(MAKE) logged-$$target;				\
	done
	@$(MAKE) -s test-stage-result


# 元の lnsc で test-go を実行
test-comp-go-org:
ifeq ($(WITH_GO),y)
	(cd .; $(MAKE) test-go LNSC=$(CUR_DIR)/lnsc)
endif

test-comp-go-org-2:
	$(MAKE) -C work/err_org test-comp-go-org COMMON_OP=--disableGenInherit



# 新しい lnsc で test-go を実行
test-comp-go-new:
ifeq ($(WITH_GO),y)
	(cd work/3; $(MAKE) test-go LNSC=$(CUR_DIR)/work/3/lnsc NO_LUA_DIFF=y)
endif

test-comp-go-new-2:
	$(MAKE) -C work/err_new test-comp-go-org COMMON_OP=--disableGenInherit



test-comp2:
ifeq ($(LANG_KIND),lua)
	$(MAKE) test-comp2-main
else
	$(MAKE) test-comp2-para
endif
# diff で、修正前と修正後の差分がないことを確認
	$(MAKE) diff-result
# NO_ACCEPT=y の場合はテストのみ行ない、成果物を上書きしない
ifdef NO_ACCEPT
	@echo skip accept-package
else ifdef LUA51
	$(MAKE) accept-package ACCEPT=y ENABLE_LUA51=$(shell which $(LUA51))
else
	$(MAKE) accept-package ACCEPT=y
endif

# lua 版と go 版を並列処理させるためのターゲット
test-comp2-para: test-comp2-main test-comp2-sub test-new3
	@date +'=========== complete test-comp2-para %c============'

test-comp2-sub:
	$(call cp_src_set,work/13,work/1)
	$(MAKE) -C work/13 test-comp2-sub-2 LANG_KIND=lua LANG_KIND=lua
	@date +'=========== complete test-comp2-sub %c============'

# 一度ビルドしたソース work/1 から、 lua 版の lns でビルドして work/13 に出力する。
# 後工程の diff-result で、 go 版の lns のビルド結果と比較して lua と go で
# 同じ出力結果になることを確認する
test-comp2-sub-2:
#	$(MAKE) -C . build_from_d depend_on=y \
#		FROM_DIR=../1 TO_DIR=../13 LNS_OPT="--package lnsc" \
#		SKIP_GO_BUILD=y
	$(call selfhost_buildOp_batch,../1,$(LNSC),../13,-langGo $(TEST_OPT))

test-comp2-main:
#	$(MAKE) test-org


# 新しいトランコンパイラで再度自分自身をコンパイルした結果が同じになることを確認する。
## work/1 から work/2
	/usr/bin/time -f 'work1:%e/%U/%S' $(MAKE) -C work/2 build_from_d \
		depend_on=y FROM_DIR=../1 TO_DIR=../2 LNS_OPT="--package lnsc"
	@date +'=========== complete work/2 %c============'

## work/2 から work/3。
## ビルド時間短縮のため、 work/3 はトランスコンパイルだけ行ない、
## go のビルドをしないように SKIP_GO_BUILD=y を指定する。
## work/3 のトランスコンパイル結果が正しければ、 
## work/2 と work/3 のビルド結果は work/2 と同じになるので、
## work/3 をビルドする必要はない。
	/usr/bin/time -f 'work2:%e/%U/%S' $(MAKE) -C work/3 build_from_d \
		depend_on=y FROM_DIR=../2 TO_DIR=../3 LNS_OPT="--package lnsc"	\
		SKIP_GO_BUILD=y
	@date +'=========== complete work/3 %c============'

	diff -u -r -I '_moduleObj.__buildId.*' -I '.*buildId = .*'	\
			-x $(notdir $(LNSC_GO)) -x binding.go		\
			work/2/lune work/3/lune

# SKIP_GO_BUILD=y で、 work/3 の go ビルドを skip したので、 work2 からコピー
ifeq ($(WITH_GO),y)
	cp work/2/lnsc work/3
endif


	cp -a test work/3
#	$(call exeLune,work/3,-mklunemod lune/base)

#	$(MAKE) test-new

	$(call exeLune,$(LNSC),work/3,test/start.lns,glue, ../newglue $(EXE_OPT) 2>&1)
	$(MAKE) build-glue GLUE_DIR=work/newglue
	mv test/glueTest.so work/3/test/glueTest.so 
	cp -a work/3 work/3_2


	mkdir -p work/0/log
	mkdir -p work/3/log

	cp -a work/3 work/err_org
ifeq ($(WITH_GO),y)
	cp ./lnsc work/err_org
endif
	cp -a work/3 work/err_new


ifndef TEST_SERIAL
	$(MAKE) test-parallel
else
# -j で全てのテストを並列で実行すると、
# エラーした際のエラー個所の特定が困難になる。
# とはいえ、 -j 指定しないと今度はテストに時間がかかり過ぎる。
# その場合に、   TEST_SERIAL=y を指定し、カテゴリの異なるテストは並列に行なわない
# ように制御する。
	$(MAKE) test-serial
endif

ifndef SKIP_PROJ_TEST
	cat work/0/log/proj >> work/test1.result
	cat work/3/log/proj >> work/test2.result
endif


ifdef LUA51
ifneq ($(shell which $(LUA51)),)
	$(MAKE) test-lua51 WITH_GO=
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


	@date +'=========== complete test-comp2-main %c============'


define check-rockspec-op
	@grep 'lune\.' ../$1 | sed 's@.*src/@@g' | sed 's@\.lua.*@.lua@g' | $(SORT) > work/$1.list
	@(cd $2; find lune -iname '*.lua' ) | grep -v runtime_go | sed 's@^@$3@g' | $(SORT) > work/$1.2.list
	diff work/$1.list work/$1.2.list
	@rm work/$1.list work/$1.2.list
endef

check-rockspec:
	$(call check-rockspec-op,lunescript-main-1.rockspec,work/3,)
ifdef LUA51
	$(call check-rockspec-op,lunescript51-main-1.rockspec,work/6,lune/base/runtime_go/legacy/lua51/)
endif


diff-result:
	(cd work/1; find test/ -iname '*.lua' -exec diff -I " _lune" -U 0 {} ../4/{} \; ) | tee work/diff.work1
	@if test -s work/diff.work1; then echo "======diff work/diff.work1====== "; exit 1; fi
	diff -u -r test/proj/  work/3/test/proj/ \
		-I '_moduleObj.__buildId.*' -I '.*buildId = .*'
ifneq ($(LANG_KIND),lua)
	diff -u -r -I '_moduleObj.__buildId' -I '.*buildId = .*'	\
			-x $(notdir $(LNSC_GO))				\
			-x 'log' -x 'test' -x binding.go		\
			work/13/lune work/3_3/work/3_3_3/lune
endif
ifndef SKIP_PROJ_TEST
ifeq ($(WITH_GO),y)
	diff test/result/test1.result work/test2.result
endif
endif
	diff work/test1.result work/test2.result
	diff work/org-error.result work/new-error.result
	diff work/ast1.result work/ast2.result
	diff work/lua1.result work/lua2.result -I " _lune"
	diff work/test.result work/test3.result
	diff work/test3.result work/test4.result
	diff work/test4.result work/test4-2.result
	diff work/test4.result work/test5.result
	diff work/1/start.dep work/3_2/start.dep
	diff -u -r work/glue work/newglue
ifeq ($(RUNTIME_ORG_VER),$(RUNTIME_NEW_VER))
# meta が等しいか確認する
	diff -x '*.so' -x 'proj' -x '*.meta2' -I 'buildId =' \
		-I '_moduleObj.__formatVersion' -r work/0-1/test/ work/4/test/
endif
ifeq ($(WITH_GO),y)
	diff -C 2 -I 'ONLY_LUA' -I 'DIFF_IGNORE' \
		work/result-testGo.txt work/3/work/result-testGo.txt
	diff -C 2 -I 'ONLY_LUA' -I 'DIFF_IGNORE' \
		work/err_org/work/result-testGo.txt work/err_new/work/result-testGo.txt
endif

	$(MAKE) diff-formatter


