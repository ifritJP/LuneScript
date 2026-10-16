test-global:
# 意図しないシンボルが global になっていないか調べる
	find lune/base/ -iname '*.lua' | \
		xargs $(LUACHECK) | grep global | \
		grep -v _lune | \
		grep -v -e 'field.*unpack.*of.*global.*table' && exit 1 || echo ok
	find work/3/test/ -iname '*.lua' | \
		grep -v ImpSub.lua | \
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

# .# ファイルがあると途中でテストが失敗するので、最初に確認する。
# なお、 .# ファイルは emacs の自動バックアップファイル
	@if test "$$(find lune -iname '.#*.lns')" != ""; then	\
		echo "error found invalid file file.";		\
		exit 1;						\
	fi


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
ifneq ($(batch),n)
	$(call selfhost_buildOp_batch,.,/usr/bin/time -f 'first:%e/%U/%S' $(LNSC_GO),work/1,--depends work/1/ )
else
	$(call selfhost_build_dep,$(LNSC),.,work/1,,--package lnsc)
endif
	$(call binding-lns-go,work/1)
	$(call go-build-lns,work/1)
	$(call mk_runtime,$(LNSC),work/1,.,)
endif


	cp -au test work/1
	$(call build-testcode,$(LNSC),.,,work/1)
	$(call cp_src_set,work/2,work/1)
	$(call cp_src_set,work/3,work/1)
	$(call cp_src_set,work/5,work/1)
ifndef SKIP_RUNTIME_VER_CHECK
	$(MAKE) check-runtime-format
endif
	$(MAKE) test-comp2


	@echo finish to build 
	@LANG= date




test-comp-with-go:
	$(MAKE) test-comp LANG_KIND=go WITH_GO=y



LNS_TEST_SRCS := $(addprefix $(LOG_DIR)/testing/,$(addsuffix .testing,$(LNS_SRCS)))

$(LOG_DIR)/testing/%.testing: %.lns
	$(LNSC) $< --testing ast > $@ 2>&1 && rm $@ || exit 1;

_test-testing-para: $(LNS_TEST_SRCS)

# 単体テストコードを確認する
_test-testing-exe:
	mkdir -p $(LOG_DIR)/testing/lune/base
	$(MAKE) _test-testing-para
test-testing:
	$(MAKE) _test-testing-exe LOG_DIR=work/0/log
ifeq ($(WITH_GO),y)
TESTING_JOBS=-j
else
TESTING_JOBS=-j5
endif
_test-testing_new:
	$(MAKE) $(TESTING_JOBS) -C work/3 _test-testing-exe LOG_DIR=log


test-registered-rockspec:
	curl https://luarocks.org/manifests/ifritjp/lunescript-main-1.rockspec -o work/rockspec
	@diff ../lunescript-main-1.rockspec work/rockspec || echo "====== need upload rockspec ======" && echo "ok"


