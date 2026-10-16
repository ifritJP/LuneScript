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
	$(call exeLune,$(LNSC),.,test/checkReference.lns shebang)

update-lnstags:
	(../../lnstags/lnstags build lune/base/front.lns --log debug)


test-formatter:
	-rm -rf work/formatter
	mkdir -p work/formatter/apply/lune/base
	$(MAKE) test-formatter-prepare
	$(MAKE) -C work/formatter test-formatter-sub


work/formatter/apply/lune/base/%.lns: lune/base/%.lns
	@echo formatter: "$< ==> $@"
	$(LNSC) ../tools/formatter/lns/main.lns -shebang $< $@ -log > $@.log

FORMAT_SRCS := $(shell find lune/base -maxdepth 1 -iname '*.lns' )
FORMAT_SRCS := $(addprefix work/formatter/apply/,$(FORMAT_SRCS))

# work/formatter/apply に formatter で変換したコードを出力し、
# 変換後のソースをビルドする
test-formatter-prepare: $(FORMAT_SRCS)
# work/formatter/test/work/1 にビルド環境をコピー
	cp Makefile work/formatter
	cp -a mk work/formatter/
	cp -a test work/formatter
	mkdir -p work/formatter/work/1
	$(call cp_src_set,work/formatter/work/1,.)
	cp -a lune work/formatter/work/1
# work/formatter/apply/lune を work/formatter/work/1 に上書き
	cp -a work/formatter/apply/lune work/formatter/work/1

# work/formatter/work/1 から work/formatter/work/3 にビルド
ifeq ($(WITH_GO),y)
	cp $(LNSC_GO) work/formatter/work/1
endif
	mkdir -p work/formatter/work/3/lune/base
	$(call cp_src_set,work/formatter/work/3,work/formatter/work/1)
	$(call selfhost_buildOp_batch,work/formatter/work/1,$(LNSC),../3,-langGo)

ifeq ($(WITH_GO),y)
# work/formatter/work/3 で lnsc をビルド
	$(call binding-lns-go,work/formatter/work/3)
	-rm work/formatter/work/3/go.mod
	$(call go-build-lns,work/formatter/work/3)
endif

	cp -a test work/formatter/work/3

test-formatter-sub: test-new test-new-proj
ifndef SKIP_PROJ_TEST
	cat work/3/log/proj >> work/test2.result
endif


diff-formatter:
	diff work/test1.result work/formatter/work/test2.result
	diff work/org-error.result work/formatter/work/new-error.result
	diff work/ast1.result work/formatter/work/ast2.result
	diff work/lua1.result work/formatter/work/lua2.result -I " _lune"
# formatter をかけた後のファイルとの比較
	diff -r lune/base/ work/formatter/apply/lune/base/ \
		-x '*.lua' -x '*.meta' -x '*~' -x '*.go' \
		-x '*.lns.log' -x runtime_* -x '*.profi' -x '.#*' -w -B
