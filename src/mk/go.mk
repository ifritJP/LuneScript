GO_RUNTIME = lune/base/runtime_go
GO_MOD = github.com/ifritJP/LuneScript/src

LUA_URL=https://www.lua.org/ftp/lua-5.3.6.tar.gz

setup-pkg-lua-sub:
	mkdir -p work
	curl $(LUA_URL) -o work/lua.tar.gz
	(cd work; tar xvfz lua.tar.gz;mv lua-* lua)
	make -j -C work/lua/src linux
	cat ../etc/lua-5.3.pc |							\
		sed 's@^includedir=.*@includedir=$(shell pwd)/work/lua/src/@' |	\
		sed 's@^Libs:.*@Libs: $(shell pwd)/work/lua/src/liblua.a -lm -ldl@' > lua-5.3.pc


setup-pkg-lua:
	@if ! pkg-config lua-5.3; then			\
		$(MAKE) setup-pkg-lua-sub;		\
	else						\
		echo ready lua-5.3;			\
	fi

# go mod tidy すると、 github.com/ifritJP/LuneScript/src/lune が取ってこれてしまうので、
# 強制的に github.com/ifritJP/LuneScript/src/lune に置き換え。
GOMODPAT="s@LuneScript/src/lune.*@LuneScript/src master@g"

FORMATTER="github.com/ifritJP/LuneScript/tools/formatter"

define buildGoSetup
	(cd $1; go mod init $2)
	(cd $1; echo "require github.com/ifritJP/LuneScript/src master" >> go.mod)
	if [ "$(GOARCH)" = "wasm" ]; then \
		echo "require $(FORMATTER) master" >> $1/go.mod;		\
		echo "replace $(FORMATTER) => $(PWD)/../tools/formatter" >> $1/go.mod; \
	fi
	(cd $1; go mod tidy;sed -i $(GOMODPAT) go.mod; go mod tidy)
endef

GOBUILD_TAG=
ifdef GOTAG
GOBUILD_TAG=-tags $(GOTAG)
endif

INT2STR=
ifeq ($(GOTAG),gopherlua)
ifeq ($(LNSC),$(LNSC_LUA51))
INT2STR=--int2str unneed0
else
INT2STR=--int2str need0
endif
endif

ifdef GOTAG
define buildGo
	$(call buildGoSetup,$1,$3)
	@echo "replace github.com/ifritJP/LuneScript/src => ./" >> $1/go.mod
	(cd $1; go build $(GOBUILD_TAG) -o $2 && ./$2)
endef
else
define buildGo
	$(call buildGoSetup,$1,$3)
	@echo "replace github.com/ifritJP/LuneScript/src => ./" >> $1/go.mod
	(cd $1; PKG_CONFIG_PATH=$(shell pwd) go build $(GOBUILD_TAG) -o $2 && ./$2)
endef
endif

GO_WORK_DIR = work/go

# lns ソースを Golang へトランスコンパイルする
#
# $1: 実行する作業ディレクトリ
# $2: トランスコンパイルする lns ソースのパス
# $3: lnsc のオプション
# $4: 出力先ディレクトリ
# $5: その他オプション
define convGo
	$(call exeLune,$1,$2,$3.lns,SAVE -langGo $(LNSC_OP) $(CONVGO_OP),$4 $5)
endef

define setup-go
	mkdir -p $1
	rsync -a --include='*.lua' \
			--include='*.lns' --exclude='*.*' lune $1
#	cp -au go.mod $1
	cp -au lune/base/runtime_go $1/lune/base/
	cp -au lune/base/binding_dummy.go $1/lune/base/
	cp -au lune/base/DependLuaOnLns.go $1/lune/base
	cp -au build.counter $1
	-rm work/result-testGo.txt work/result-testGo2.txt
endef

define test-go-one-setup
	cp -au test $1
endef

define test-go-one-exec
	$(LNSC) $5 mkmain $2/main.go $6
	cp -au $4/$3.lns $2/$4
	$(call convGo,$1,$2,$4/$3,--main $5 $6 --app testgo)
#	if [ $4 != "." ]; then mv $2/$4/$3.go $2; fi
	$(call buildGo,$2,test.exe,testgo) > work/result-testGo.txt
	$(call exeLune,$1,$2,$4/$3.lns,shebang) > work/result-testGo2.txt
	if test "${NO_LUA_DIFF}" = ""; then					\
		diff -C 2 -I 'ONLY_GO' -I 'ONLY_LUA' -I 'DIFF_IGNORE'		\
				work/result-testGo.txt work/result-testGo2.txt;	\
	fi
endef


define test-go-one
	$(call test-go-one-setup,$2)
	$(call test-go-one-exec,$1,$2,$3,$4,$5,$6)
endef

define test-go-some-setup
	cp -au $4/$3.lns $2/$4
	touch $2/$4/$3.lns
	$(call convGo,$1,$2,$4/$3,$5 --uptodate forceAll)
#	mv $2/$4/$3.go $2
endef


define test-go-some
	$(call test-go-one-exec,$1,$2,$3,$4,$5,$6 --uptodate forceAll)
endef


test-miniGo:
	$(call setup-go,$(GO_WORK_DIR))
	$(call selfhost_buildOp_batch,.,$(LNSC_GO),$(GO_WORK_DIR),)
	-rm $(GO_WORK_DIR)/*.*
	cp build.counter $(GO_WORK_DIR)
ifdef OTHER_GO
	cp $(OTHER_GO) $(GO_WORK_DIR)
endif
	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),miniGo,.,miniGo,$(INT2STR) -u --package main)
	cat work/result-testGo.txt

ifndef NO_TEST
	$(MAKE) test-go SKIP_SETUP=y
endif

miniGo-go:
	$(call buildGo,$(GO_WORK_DIR),miniGo.exe,minigo)

test-go:
ifndef SKIP_SETUP
	$(call setup-go,$(GO_WORK_DIR))
	$(call selfhost_buildOp_batch,.,$(LNSC_GO),$(GO_WORK_DIR),)
endif
ifdef SETUP_LUA
	$(call setup-go,$(GO_WORK_DIR))
	$(call selfhost_buildOp_batch,.,$(LNSC),$(GO_WORK_DIR),)
endif
	-rm $(GO_WORK_DIR)/*.*
#	rm $(GO_WORK_DIR)/go.mod
	$(call test-go-one-setup,$(GO_WORK_DIR))
ifeq ($(GOTAG),gopherlua)
# gopherlua は lua 5.1 互換なので lua51 依存のテストは除外
	sed -i.bak 's@/.nolua51./@//@g' $(GO_WORK_DIR)/test/testGo.lns
endif
	-rm $(GO_WORK_DIR)/test/*.c
	$(call test-go-some-setup,$(LNSC),$(GO_WORK_DIR),SubPkg,test/subGoPkg,--app testgo)
	$(call test-go-some,$(LNSC),$(GO_WORK_DIR),testGo,test,test.testGo,--app testgo $(INT2STR))


test-tutorial-go:
	$(call cmd_exec_trans,$(LNSC),.,.,,base/convGo)
	$(call setup-go,$(GO_WORK_DIR))

	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),match,test/tutorial,match)
	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),loop,test/tutorial,loop)
	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),interface,test/tutorial,interface)
	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),generics,test/tutorial,generics)
	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),enum,test/tutorial,enum)
	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),classOverride,test/tutorial,classOverride)
	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),classMapping,test/tutorial,classMapping)
	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),classInherit,test/tutorial,classInherit)
	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),class2accessor,test/tutorial,class2accessor)
	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),class1,test/tutorial,class1)
	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),cast,test/tutorial,cast)
	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),branch,test/tutorial,branch)
	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),nilable,test/tutorial,nilable)
	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),defaultArg,test/tutorial,defaultArg)
	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),multipleRetVal,test/tutorial,multipleRetVal)
	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),arg,test/tutorial,arg)
	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),func,test/tutorial,func)
	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),variable,test/tutorial,variable)
	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),set,test/tutorial,set)
	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),value,test/tutorial,value)

	$(call test-go-one,$(LNSC),$(GO_WORK_DIR),ccc,.)


define test-lns-test
	@echo "local _lune = require( 'lune.base.runtime' )" > $1/testRun.lua
	@echo "local __t = _lune.loadModule( 'lune.base.$2' )" >> $1/testRun.lua
	@echo "local __t = _lune.loadModule( 'lune.base.Testing' )" >> $1/testRun.lua
	@echo "__t.run( 'lune.base.$2');"			>> $1/testRun.lua
	@echo "__t.outputAllResult( io.stdout );"		>> $1/testRun.lua
	(cd $1; LUA_PATH=./?.lua $(LUA) testRun.lua)
endef


define test-convGo-test
#	cat misc/go/main/runtime_main.go |		\
#		sed 's/Lns_init/Lns_$3_init/' |			\
#		sed 's@//TEST:@@' |				\
#		sed 's/run( "" )/run( "lune.base.$3" )/' |				\
#		sed 's@//IMPORT@import . "$(GO_MOD)/src/lune/base"@' > $2/$5/runtime_main.go
	$(LNSC) --testing $3 mkmain $2/$5/runtime_main.go


	$(call convGo,$1,$2,$4/$3,--testing --package lnsc,$5)
	$(call buildGo,$2/$5,test.exe,$(GO_MOD)) >> work/result-testGo.txt
	$(call test-lns-test,$2/$5,$3) >> work/result-testGo2.txt
	diff -C 2 -I 'ONLY_LUA' work/result-testGo.txt work/result-testGo2.txt
endef

test-testGo:
	rm -rf $(GO_WORK_DIR)/*.* $(GO_WORK_DIR)/lune $(GO_WORK_DIR)/test
	$(call setup-go,$(GO_WORK_DIR))
	$(call setup-go,$(GO_WORK_DIR)/test)

	$(call cmd_exec_trans,$(LNSC),.,.,,base/convGo)

	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/Testing,--testing --package lnsc,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/Types,--testing --package lnsc,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/Async,--testing --package lnsc,./test)
	$(call test-convGo-test,$(LNSC),$(GO_WORK_DIR),Ver,lune/base,./test)
	$(call test-convGo-test,$(LNSC),$(GO_WORK_DIR),Code,lune/base,./test)
	$(call test-convGo-test,$(LNSC),$(GO_WORK_DIR),LuaMod,lune/base,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/LuaVer,--testing --package lnsc,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/LuneControl,--testing --package lnsc,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/Meta,--testing --package lnsc,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/Log,--testing --package lnsc,./test)
	$(call test-convGo-test,$(LNSC),$(GO_WORK_DIR),Util,lune/base,./test)
	$(call test-convGo-test,$(LNSC),$(GO_WORK_DIR),AsyncTokenizer,lune/base,./test)
	$(call test-convGo-test,$(LNSC),$(GO_WORK_DIR),Tokenizer,lune/base,./test)
	$(call test-convGo-test,$(LNSC),$(GO_WORK_DIR),Json,lune/base,./test)
	$(call test-convGo-test,$(LNSC),$(GO_WORK_DIR),Writer,lune/base,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/frontInterface,--testing --package lnsc,./test)
	$(call test-convGo-test,$(LNSC),$(GO_WORK_DIR),Ast,lune/base,./test)
	$(call test-convGo-test,$(LNSC),$(GO_WORK_DIR),Option,lune/base,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/GoMod,--testing --package lnsc,./test)
	$(call test-convGo-test,$(LNSC),$(GO_WORK_DIR),Nodes,lune/base,./test)
	$(call test-convGo-test,$(LNSC),$(GO_WORK_DIR),Formatter,lune/base,./test)
	$(call test-convGo-test,$(LNSC),$(GO_WORK_DIR),Macro,lune/base,./test)
	$(call test-convGo-test,$(LNSC),$(GO_WORK_DIR),TransUnit,lune/base,./test)
	$(call test-convGo-test,$(LNSC),$(GO_WORK_DIR),convLua,lune/base,./test)
	$(call test-convGo-test,$(LNSC),$(GO_WORK_DIR),dumpNode,lune/base,./test)
#	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/convCC,--testing --package lnsc,./test)
	$(call test-convGo-test,$(LNSC),$(GO_WORK_DIR),convGo,lune/base,./test)
	$(call test-convGo-test,$(LNSC),$(GO_WORK_DIR),OutputDepend,lune/base,./test)
	$(call convGo,$(LNSC),$(GO_WORK_DIR),lune/base/glueFilter,--testing --package lnsc,./test)
	$(call test-convGo-test,$(LNSC),$(GO_WORK_DIR),front,lune/base,./test)

test-subgo:
	$(call cmd_exec_trans,$(LNSC),.,.,,base/convGo)
	$(call setup-go,$(GO_WORK_DIR))

	$(call test-convGo-test,$(LNSC),$(GO_WORK_DIR),front,lune/base,./test)

define binding-lns-go
#	if diff $1/ver_lns_lua $1/ver_lua; then					\
#		$(LNSC_LUA) misc/go/binding/bind.lns shebang -b $1;		\
#	else									\
#		$(LNSC_LUA) misc/go/binding/bind.lns shebang $1;		\
#	fi
# luac は bind するように binding.go を生成すると、
# git に luac を登録しないといけなくなる。
# luac は扱い難いので bytecompile はしない。
	$(LNSC_LUA) misc/go/binding/bind.lns shebang $1

	mv output.go $1/lune/base/binding.go
endef

binding-lns-go-tgt:
	$(call binding-lns-go,.)

ifeq ($(GOARCH),wasm)
MAIN_GO=misc/go/main/lnsc_main_wasm.go
else
MAIN_GO=misc/go/main/lnsc_main.go
endif

define go-build-lns
	cp $(MAIN_GO) $1
	$(call buildGoSetup,$1,$(GO_MOD))
	(cd $1; PKG_CONFIG_PATH=$(shell pwd) time go build -o lnsc $(GOBUILD_TAG))
endef

build-wasm:
	GOOS=js GOARCH=wasm $(MAKE) go-build NO_TEST=y GOTAG=gopherlua

build-wasm-quick:
	GOOS=js GOARCH=wasm $(MAKE) GOTAG=gopherlua build-wasm-quick-sub
build-wasm-quick-sub:
	-rm $(GO_WORK_DIR)/build/go.mod
	$(call go-build-lns,$(GO_WORK_DIR)/build)
	cp $(GO_WORK_DIR)/build/$(LNSC_GO) lnsc.wasm



# セルフホストのソースを go に変換し、 go 版 lns をビルドする
go-build:
# ./lune のソースから $(GO_WORK_DIR)/lune 以下に変換結果を出力
	rm -rf $(GO_WORK_DIR)
	$(call setup-go,$(GO_WORK_DIR))
	$(call cp_src_set,$(GO_WORK_DIR),.)
	cp Makefile $(GO_WORK_DIR)
	cp -a mk $(GO_WORK_DIR)/
	$(call selfhost_buildOp_batch,.,$(LNSC),$(GO_WORK_DIR),-langGo,)

# ビルドするため $(GO_WORK_DIR)/lune を $(GO_WORK_DIR)/build/lune にコピー
	$(call setup-go,$(GO_WORK_DIR)/build)
	$(call cp_src_set,$(GO_WORK_DIR)/build,$(GO_WORK_DIR))
	cp -a $(GO_WORK_DIR)/lune $(GO_WORK_DIR)/build

# $(GO_WORK_DIR)/build をビルド
	$(call binding-lns-go,$(GO_WORK_DIR)/build)
	$(call go-build-lns,$(GO_WORK_DIR)/build)
	$(call mk_runtime,$(LNSC),$(GO_WORK_DIR)/build,.,)

ifndef NO_TEST
	cp $(GO_WORK_DIR)/build/$(LNSC_GO) $(GO_WORK_DIR)
	$(MAKE) test-govslua LNSC_OP="$(TEST_OPT)"
	$(MAKE) LNSC=$(shell pwd)/$(GO_WORK_DIR)/build/$(LNSC_GO) test-lnsc-go \
		SKIP_SETUP=y SETUP_LUA=y


	$(MAKE) -s TEST_DIR=$(GO_WORK_DIR) LOG_DIR=work/0/log test-error | \
		grep -e 'error:' -e 'warning:' > work/test1.result
	$(MAKE) -s TEST_DIR=$(GO_WORK_DIR) LOG_DIR=work/1/log \
		LNSC=./$(LNSC_GO) test-error | \
			grep -e 'error:' -e 'warning:' > work/test2.result
	diff work/test1.result work/test2.result -I 'Entering directory' -I 'Leaving directory'
endif

# NO_ACCEPT=y の場合は $(GO_WORK_DIR)/build に残し、カレントの lnsc を上書きしない
ifdef NO_ACCEPT
	@echo skip copying $(GO_WORK_DIR)/build/$(LNSC_GO)
else ifneq ($(GOARCH),wasm)
	cp $(GO_WORK_DIR)/build/$(LNSC_GO) .
else
	cp $(GO_WORK_DIR)/build/$(LNSC_GO) lnsc.wasm
endif

test-lnsc-go:
	-$(MAKE) test-go
# LuneScript のソース lune/base にアクセスできなくても lnsc コマンドが動くか確認する
	(cd $(GO_WORK_DIR); $(LNSC) test/testGo.lns exe) > work/result-testGo3.txt
	diff -I 'fib_l' work/result-testGo2.txt work/result-testGo3.txt 


