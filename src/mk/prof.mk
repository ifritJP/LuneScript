define test-govslua-comp
	(cd $1; /usr/bin/time -f 'lns:%e/%U/%S:$2' $(LNSC_GO) lune/base/$2.lns SAVE $3 --log log)
	(cd $1; LUA_PATH=./?.lua /usr/bin/time -f 'lua:%e/%U/%S:$2' $(LNSC_LUA) lune/base/$2.lns SAVE $4 --log log)
	@(cd $1; diff $4/lune/base/$2.lua $3/lune/base/$2.lua)
	@(cd $1; diff $4/lune/base/$2.meta $3/lune/base/$2.meta)
endef


# go-build でビルドした go 版 lns と lua 版 lns とで、
# lns ソースの変換を行ない。同じ結果になることを確認する。
test-govslua:
	-rm work/go.time work/lua.time
ifndef NO_BATCH
	/usr/bin/time -f 'go:%e/%U/%S:$2' -o work/go.time -a		\
		$(MAKE) test-govslua-go-batch
else
	/usr/bin/time -f 'go:%e/%U/%S:$2' -o work/go.time -a		\
		$(MAKE) test-govslua-go
endif
ifndef NO_BATCH
	/usr/bin/time -f 'lua:%e/%U/%S:$2' -o work/lua.time -a		\
		$(MAKE) test-govslua-lua-batch
else
	/usr/bin/time -f 'lua:%e/%U/%S:$2' -o work/lua.time -a		\
		$(MAKE) test-govslua-lua
endif
	cat work/go.time work/lua.time
	diff -u -r work/1 work/2 -I buildId				\
		-I ': Entering directory'				\
		-I ': Leaving directory' -x '$(notdir $(LNSC_GO))' -x ast


define test-buildlns-setup
	mkdir -p $1
	rm -rf $1/*
	$(call mk_lune_proj_dir,$1)
	$(call cp_src_set,$1,.)
endef

define test-buildlns-run
	$(call selfhost_build_dep,$1,.,$2,-langGo --package lnsc -ol 53)
	test "$(ONLY_CONV)" = "" && LANG= $(MAKE) -C $2 LANG_KIND=$3 tgt_dump_ast > $2/ast || echo skip
endef

test-govslua-lua:
	$(call test-buildlns-setup,work/1)
	$(call test-buildlns-run,$(LNSC),work/1,lua)

test-govslua-go:
	$(call test-buildlns-setup,work/2)
ifdef TEST_MAIN_GO
	cp $(LNSC_GO) work/2
else
	cp $(GO_WORK_DIR)/build/$(LNSC_GO) work/2
endif
	$(call test-buildlns-run,/usr/bin/time -f 'go:%e/%U/%S:$2' -o work/go.time -a work/2/$(LNSC_GO),work/2,go)


test-govslua-lua-batch:
	$(call test-buildlns-setup,work/1)
	$(call selfhost_buildOp_batch,$(GO_WORK_DIR),$(LNSC),../1,--depends ../1/ -langGo --package lnsc)
test-govslua-go-batch:
	$(call test-buildlns-setup,work/2)
	$(call selfhost_buildOp_batch,$(GO_WORK_DIR),$(LNSC_GO),../2,--depends ../2/ -langGo --package lnsc)


goprof-measure:
	@index=1;									\
	while test $${index} -lt 12 ; do						\
		echo $${index};								\
		GOMAXPROCS=$${index} make goprof LNSC_CMD=./lnsc 2>&1 | grep -e '^go:';	\
		GOMAXPROCS=$${index} make goprof LNSC_CMD=./lnsc 2>&1 | grep -e '^go:';	\
		GOMAXPROCS=$${index} make goprof LNSC_CMD=./lnsc 2>&1 | grep -e '^go:';	\
		index=$$((index+1));							\
	done


ifndef SKIP_PROF
PROF_OP=-prof
endif
goprof-main:
	$(call selfhost_buildOp_batch,.,/usr/bin/time -f 'go:%e/%U/%S:$2' $(LNSC_GO),work/go,-langGo $(PROF_OP))

ifdef PROF_SCHE
# 10msec 毎に go のスケジューラー情報をダンプ
GODEBUG:=schedtrace=$(PROF_SCHE)
endif
ifdef PROF_GC
ifdef GODEBUG
# GC 毎に go の GC 情報をダンプ
GODEBUG:=$(GODEBUG),gctrace=1
else
GODEBUG:=gctrace=1
endif
endif

goprof:
	$(call test-buildlns-setup,work/go)
	GODEBUG=$(GODEBUG) $(MAKE) goprof-main
ifndef SKIP_PROF
	(cd .; go tool pprof -dot @-.profi  ) | dot -Tsvg -odump.svg
	(cd .; go tool pprof -dot @-.profimem  ) | dot -Tsvg -odumpmem.svg
endif



PROF_ONE_SRC=lune/base/Nodes
goprof-one:
	$(call test-buildlns-setup,work/go)
	$(call selfhost_buildOp_batch,.,/usr/bin/time -f 'go:%e/%U/%S:$2' $(LNSC_GO),work/go,-langGo)
	touch work/go/$(PROF_ONE_SRC).*
	echo $(PROF_ONE_SRC).lns | (cd .; $(LNSC_GO) @- SAVE work/go $(BYTECOMP)	\
			--package lnsc $(COM_OPT) $(LOG_OPT) -prof)

	(cd .; go tool pprof -dot @-.profi  ) | dot -Tsvg -odump.svg
	(cd .; go tool pprof -dot @-.profimem  ) | dot -Tsvg -odumpmem.svg
view-goprof:
ifdef BROWSER
	go tool pprof -http=0.0.0.0:8100 lnsc @-.profi
else
	go tool pprof -http=0.0.0.0:8100 -no_browser lnsc @-.profi
endif
view-goprofmem:
ifdef BROWSER
	go tool pprof -http=0.0.0.0:8100 lnsc @-.profimem
else
	go tool pprof -http=0.0.0.0:8100 -no_browser lnsc @-.profimem
endif
view-gotrace:
ifdef BROWSER
	go tool trace -http=0.0.0.0:8100 @-.profitrace
else
	go tool trace -http=0.0.0.0:8100 -no_browser @-.profitrace
endif

luaprof:
	$(call test-buildlns-setup,work/go)
	$(call selfhost_buildOp_batch,.,/usr/bin/time -f 'lua:%e/%U/%S:$2' env $(LNSC),work/go,-langGo )


