# wasm (WasmGC) への直接変換のテスト
#
# test/wasm/*.lns と test/tutorial/*.lns を、Lua で実行した結果と、
# wasm に変換して Node.js で実行した結果で比べる。
# 設計は docs/sw_design/wasm/wasm_direct_design.md を参照。

NODE=node
WASM_WORK_DIR=work/wasm
# 変換した wasm を実行するコマンド。デフォルトは Node.js。
# WASI のランタイム (wasmtime) で実行する場合は test-wasm-wasi などを使う (WASM_RUN=$(WASM_RUN_WASI))
WASM_RUN=$(NODE) $(CURDIR)/misc/wasm/lnswasm.mjs
WASM_RUN_WASI=$(CURDIR)/misc/wasm/lnswasi.sh run
# lnswasi.sh が使う Binaryen と wasmtime
BINARYEN_BIN ?= $(HOME)/opt/binaryen-version_133/bin
WASMTIME ?= $(HOME)/opt/wasmtime-v49.0.4-x86_64-linux/wasmtime
export BINARYEN_BIN WASMTIME

# 変換した wasm が使う C のランタイム (WASI のコマンド)。ビルドには wasi-sdk が必要。
#  - lnsrt.wasm: 入出力など (misc/wasm/lnsrt.c)。Lua を使わない wasm が使う
#  - lnslua.wasm: lnsrt.c と Lua。変換した wasm は、Lua 互換の処理 (string.format やパターンマッチなど) にも使う
# Lua のソースは lua.org から取得する。
WASI_SDK ?= $(HOME)/opt/wasi-sdk-34.0-x86_64-linux
WASM_LUA_VER=5.3.6
WASM_LUA_SHA256=fc5fd69bb8736323f026672b1b7235da613d7177e72558893a0bdcd320466d60
WASM_LUA_DIR=work/wasm-lua
WASM_LUA_SRC_DIR=$(WASM_LUA_DIR)/lua-$(WASM_LUA_VER)/src
WASM_LUA=$(WASM_LUA_DIR)/lnslua.wasm
WASM_RT=$(WASM_LUA_DIR)/lnsrt.wasm
WASM_RTS=$(WASM_LUA) $(WASM_RT)
# WASI のランタイム用の、Node.js でだけ使える関数 (lnsjs) の代わり
WASM_JS_STUB=$(WASM_LUA_DIR)/lnsjs_stub.wasm
WASM_RT_SRCS=misc/wasm/lnsrt.c misc/wasm/lnsrt.h
WASM_RT_CFLAGS=-Os -D_WASI_EMULATED_PROCESS_CLOCKS -Wl,--strip-all
WASM_RT_LIBS=-lwasi-emulated-process-clocks
# io・os・package・debug は、Lua で動かす LuneScript (マクロの実行) が使う。
# ファイルアクセスなどは WASI で行う
WASM_LUA_MODS=lapi lcode lctype ldebug ldo ldump lfunc lgc llex lmem lobject	\
	lopcodes lparser lstate lstring ltable ltm lundump lvm lzio		\
	lauxlib lbaselib ltablib lstrlib lmathlib lutf8lib			\
	loadlib liolib loslib ldblib
# l_signalT: wasm には signal.h が無い
# l_randomizePivot: table.sort の pivot の乱数化に clock() を使わない
WASM_LUA_CFLAGS=$(WASM_RT_CFLAGS) -DLUA_COMPAT_5_2 -Dl_signalT=int '-Dl_randomizePivot()=0'	\
	-DL_tmpnam=32 -mllvm -wasm-enable-sjlj

$(WASM_LUA_SRC_DIR)/lua.h:
	mkdir -p $(WASM_LUA_DIR)
	curl -sSL -o $(WASM_LUA_DIR)/lua-$(WASM_LUA_VER).tar.gz \
		https://www.lua.org/ftp/lua-$(WASM_LUA_VER).tar.gz
	echo "$(WASM_LUA_SHA256)  $(WASM_LUA_DIR)/lua-$(WASM_LUA_VER).tar.gz" | sha256sum -c
	tar xzf $(WASM_LUA_DIR)/lua-$(WASM_LUA_VER).tar.gz -C $(WASM_LUA_DIR)
	touch $@

$(WASM_LUA): misc/wasm/lnslua.c $(WASM_RT_SRCS) $(WASM_LUA_SRC_DIR)/lua.h
	$(WASI_SDK)/bin/clang --target=wasm32-wasip1 $(WASM_LUA_CFLAGS)	\
		-I$(WASM_LUA_SRC_DIR)						\
		$(addprefix $(WASM_LUA_SRC_DIR)/,$(addsuffix .c,$(WASM_LUA_MODS)))	\
		misc/wasm/lnslua.c misc/wasm/lnsrt.c -lsetjmp $(WASM_RT_LIBS) -o $@

$(WASM_RT): $(WASM_RT_SRCS)
	mkdir -p $(WASM_LUA_DIR)
	$(WASI_SDK)/bin/clang --target=wasm32-wasip1 $(WASM_RT_CFLAGS)	\
		misc/wasm/lnsrt.c $(WASM_RT_LIBS) -o $@

$(WASM_JS_STUB): misc/wasm/lnsjs_stub.wat
	mkdir -p $(WASM_LUA_DIR)
	$(BINARYEN_BIN)/wasm-as --enable-gc --enable-reference-types $< -o $@

wasm-lua: $(WASM_RTS)
# test/tutorial/class1.lns は Lua 版でもコンパイルエラーになるテストなので除外する
WASM_TEST_SRCS:=$(wildcard test/wasm/*.lns) \
	$(filter-out test/tutorial/class1.lns,$(wildcard test/tutorial/*.lns))

test-wasm: $(WASM_RTS)
	rm -rf $(WASM_WORK_DIR)
	mkdir -p $(WASM_WORK_DIR)
	@for src in $(WASM_TEST_SRCS); do					\
		name=$$(basename $$src .lns);					\
		dir=$$(dirname $$src);						\
		mkdir -p $(WASM_WORK_DIR)/$$dir;				\
		$(LNSC) $$src exe $(EXE_OPT)					\
			> $(WASM_WORK_DIR)/$$name.lua.result			\
			2> $(WASM_WORK_DIR)/$$name.lua.err || exit 1;		\
		$(LNSC) $$src SAVE $(WASM_WORK_DIR) -langWasm			\
			> /dev/null 2> $(WASM_WORK_DIR)/$$name.conv.err || exit 1; \
		$(WASM_RUN) $(WASM_WORK_DIR)/$$dir/$$name.wasm			\
			> $(WASM_WORK_DIR)/$$name.wasm.result 2>&1 || exit 1;	\
		diff -u $(WASM_WORK_DIR)/$$name.lua.result			\
			$(WASM_WORK_DIR)/$$name.wasm.result || exit 1;		\
		echo "ok -- $$name ($$(stat -c %s $(WASM_WORK_DIR)/$$dir/$$name.wasm) bytes)"; \
	done
	$(MAKE) test-wasm-testgo
	$(MAKE) test-wasm-start

# test-wasm・test-wasm-selfhost を、WASI のランタイム (wasmtime) で実行する。
# 変換した wasm は、lnswasi.sh で C のランタイムと 1 つの WASI の wasm にしてから実行する。
# Binaryen (BINARYEN_BIN) と wasmtime (WASMTIME) が必要。
test-wasm-wasi: $(WASM_RTS) $(WASM_JS_STUB)
	$(MAKE) test-wasm WASM_RUN="$(WASM_RUN_WASI)" WASM_NOWASI=1
test-wasm-selfhost-wasi: $(WASM_RTS) $(WASM_JS_STUB)
	$(MAKE) test-wasm-selfhost WASM_RUN="$(WASM_RUN_WASI)"

# test/testGo.lns (+ test/subGoPkg) を、Lua 版と wasm 版で比べる。
# "_section(); // nowasm" の section は、wasm で未対応の機能を使う場合の印で、
# 行番号を保ったまま空行にして除く (今は全ての section が対応済みで、印は無い)。
# "_section(); // nowasi" の section は、WASI のランタイムでは動かない機能 (io.popen) を使う場合の印で、
# WASM_NOWASI=1 (test-wasm-wasi) の場合に除く。
# 実行時間などの "DIFF_IGNORE" を含む行は比べない。
# import したモジュールも wasm にまとめるので、--uptodate forceAll で AST を作る。
WASM_TESTGO_DIR=$(WASM_WORK_DIR)/testgo
WASM_NOWASI=0
test-wasm-testgo: $(WASM_RTS)
	rm -rf $(WASM_TESTGO_DIR)
	mkdir -p $(WASM_TESTGO_DIR)/test
	rsync -a --include='*.lua' --include='*.lns' --exclude='*.*' lune $(WASM_TESTGO_DIR)
	if [ -f lnsc ]; then cp -a lnsc $(WASM_TESTGO_DIR); fi
	cp -r test/subGoPkg $(WASM_TESTGO_DIR)/test/
	awk -v nowasi=$(WASM_NOWASI)						\
		'/^_section\(\);/ { skip = ( $$0 ~ /\/\/ *nowasm/ ) || ( nowasi && $$0 ~ /\/\/ *nowasi/ ) } \
		{ print skip ? "" : $$0 }' \
		test/testGo.lns | sed 's@/.nolua51./@//@g' > $(WASM_TESTGO_DIR)/test/testGo.lns
	(cd $(WASM_TESTGO_DIR); $(LNSC) test/testGo.lns exe $(EXE_OPT)	\
		> testGo.lua.result 2> testGo.lua.err)
	(cd $(WASM_TESTGO_DIR); $(LNSC) test/testGo.lns SAVE . -langWasm	\
		--uptodate forceAll > /dev/null 2> testGo.conv.err)
	(cd $(WASM_TESTGO_DIR); $(WASM_RUN) test/testGo.wasm				\
		> testGo.wasm.result 2>&1)
	diff -u -I DIFF_IGNORE $(WASM_TESTGO_DIR)/testGo.lua.result $(WASM_TESTGO_DIR)/testGo.wasm.result
	@echo "ok -- testGo ($$(stat -c %s $(WASM_TESTGO_DIR)/test/testGo.wasm) bytes,"	\
		"$$(wc -l < $(WASM_TESTGO_DIR)/testGo.lua.result) lines)"

# go 版のネイティブ実行と wasm 実行の速度を比べる (docs/sw_design/wasm/wasm_benchmark.md)
# WASM_OPT に Binaryen の wasm-opt を指定すると、wasm-opt -O3 した wasm も測る
bench-wasm: $(WASM_RTS)
	WASM_OPT=$(WASM_OPT) misc/wasm/bench.sh

# test/start.lns (test-comp の主なテスト) を、Lua 版と wasm 版で比べる。
# "skip-cross" (Lua 以外への変換で除く行) と "skip-wasm" (wasm で除く行) の行は、
# 行番号を保ったまま空行にして除く。
# skip-wasm は、C のグルー・__free (ファイナライザ) など、wasm では動かない部分に付けている。
WASM_START_DIR=$(WASM_WORK_DIR)/start
test-wasm-start: $(WASM_RTS)
	rm -rf $(WASM_START_DIR)
	mkdir -p $(WASM_START_DIR)
	rsync -a --include='*.lua' --include='*.lns' --exclude='*.*' lune $(WASM_START_DIR)
	if [ -f lnsc ]; then cp -a lnsc $(WASM_START_DIR); fi
	cp -a test $(WASM_START_DIR)
	sed -i -e 's/^.*skip-cross.*$$//' -e 's/^.*skip-wasm.*$$//'			\
		$(WASM_START_DIR)/test/start.lns $(WASM_START_DIR)/test/Sub11.lns
	(cd $(WASM_START_DIR); $(LNSC) test/start.lns exe $(EXE_OPT)		\
		> start.lua.result 2> start.lua.err)
	(cd $(WASM_START_DIR); $(LNSC) test/start.lns SAVE . -langWasm		\
		--uptodate forceAll > /dev/null 2> start.conv.err)
	(cd $(WASM_START_DIR); $(WASM_RUN) test/start.wasm				\
		> start.wasm.result 2>&1)
	diff -u -I DIFF_IGNORE $(WASM_START_DIR)/start.lua.result $(WASM_START_DIR)/start.wasm.result
	@echo "ok -- start ($$(stat -c %s $(WASM_START_DIR)/test/start.wasm) bytes,"	\
		"$$(wc -l < $(WASM_START_DIR)/start.lua.result) lines)"

# wasm 版の lnsc のセルフホストの確認 (M4)。
# 1. Lua 版 (または go 版) の lnsc で、lnsc (lune/base/front.lns) を wasm に変換する (第 1 世代)
# 2. 第 1 世代の wasm 版 lnsc で、lnsc を wasm に変換し、第 1 世代と一致することを確認する (不動点)
# 3. lnsc の全モジュールを、Lua 版の lnsc と wasm 版の lnsc で変換 (.lua / .meta / .go) し、一致することを確認する
# 4. test/start.lns を、Lua 版の lnsc と wasm 版の lnsc で実行 (exe) し、結果が一致することを確認する。
#    wasm 版の Lua は C のモジュールを読み込めないので、skip-wasm の行 (C のグルーなど) は空行にする
# マクロの実行に使う Lua 版の lnsc は wasm 版の lnsc に埋め込んでいるので、LUA_PATH は指定しない
# (LUA_PATH は、lnsc がモジュールの .lns を探すのにも使うので、Lua の標準のパスにする)。
# 2 の gen2 のディレクトリには lune/base/*.lua が無いので、埋め込んだ Lua 版の lnsc だけでマクロを実行する。
WASM_SELFHOST_DIR=$(WASM_WORK_DIR)/selfhost
WASM_LNSC=env -u LUA_PATH $(WASM_RUN)
WASM_SELFHOST_SRCS=$(addsuffix .lns,$(LNS_SRCS))
test-wasm-selfhost: $(WASM_RTS)
	rm -rf $(WASM_SELFHOST_DIR)
	mkdir -p $(WASM_SELFHOST_DIR)/gen2
	rsync -a --include='*.lua' --include='*.lns' --exclude='*.*' lune $(WASM_SELFHOST_DIR)
	if [ -f lnsc ]; then cp -a lnsc $(WASM_SELFHOST_DIR); fi
	(cd $(WASM_SELFHOST_DIR); $(LNSC) lune/base/front.lns SAVE . -langWasm		\
		--uptodate forceAll > /dev/null 2> gen1.conv.err)
	mv $(WASM_SELFHOST_DIR)/lune/base/front.wasm $(WASM_SELFHOST_DIR)/lnsc1.wasm
	rsync -a --include='*.lns' --include='*/' --exclude='*' lune $(WASM_SELFHOST_DIR)/gen2
	(cd $(WASM_SELFHOST_DIR)/gen2; $(WASM_LNSC) ../lnsc1.wasm			\
		lune/base/front.lns SAVE . -langWasm --uptodate forceAll > /dev/null 2> conv.err)
	cmp $(WASM_SELFHOST_DIR)/lnsc1.wasm $(WASM_SELFHOST_DIR)/gen2/lune/base/front.wasm
	@echo "ok -- fixed point ($$(stat -c %s $(WASM_SELFHOST_DIR)/lnsc1.wasm) bytes)"
	mkdir -p $(WASM_SELFHOST_DIR)/out-lua/lune/base $(WASM_SELFHOST_DIR)/out-wasm/lune/base
	(cd $(WASM_SELFHOST_DIR); echo $(WASM_SELFHOST_SRCS) | tr ' ' '\n' |	\
		$(LNSC) @- SAVE out-lua --package lnsc $(COM_OPT) > out-lua.log 2>&1)
	(cd $(WASM_SELFHOST_DIR); echo $(WASM_SELFHOST_SRCS) | tr ' ' '\n' |	\
		$(WASM_LNSC) lnsc1.wasm @- SAVE out-wasm --package lnsc $(COM_OPT)	\
		> out-wasm.log 2>&1)
	diff -r $(WASM_SELFHOST_DIR)/out-lua $(WASM_SELFHOST_DIR)/out-wasm
	@echo "ok -- save ($$(find $(WASM_SELFHOST_DIR)/out-wasm -type f | wc -l) files)"
	cp -a test $(WASM_SELFHOST_DIR)
	sed -i -e 's/^.*skip-wasm.*$$//' $(WASM_SELFHOST_DIR)/test/start.lns $(WASM_SELFHOST_DIR)/test/Sub11.lns
	(cd $(WASM_SELFHOST_DIR); $(LNSC) test/start.lns exe $(EXE_OPT)	\
		> start.lua.result 2>&1)
	(cd $(WASM_SELFHOST_DIR); $(WASM_LNSC) lnsc1.wasm test/start.lns exe $(EXE_OPT)	\
		> start.wasm.result 2>&1)
	diff -u $(WASM_SELFHOST_DIR)/start.lua.result $(WASM_SELFHOST_DIR)/start.wasm.result
	@echo "ok -- start ($$(wc -l < $(WASM_SELFHOST_DIR)/start.lua.result) lines)"
