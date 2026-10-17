# wasm (WasmGC) への直接変換のテスト
#
# test/wasm/*.lns と test/tutorial/*.lns を、Lua で実行した結果と、
# wasm に変換して Node.js で実行した結果で比べる。
# 設計は docs/sw_design/wasm_direct_design.md を参照。

NODE=node
WASM_WORK_DIR=work/wasm

# wasm にした Lua (lnslua.wasm)。
# 変換した wasm は、Lua 互換の処理 (string.format やパターンマッチなど) にこれを使う。
# Lua のソースは lua.org から取得する。ビルドには wasi-sdk が必要。
WASI_SDK ?= $(HOME)/opt/wasi-sdk-34.0-x86_64-linux
WASM_LUA_VER=5.3.6
WASM_LUA_SHA256=fc5fd69bb8736323f026672b1b7235da613d7177e72558893a0bdcd320466d60
WASM_LUA_DIR=work/wasm-lua
WASM_LUA_SRC_DIR=$(WASM_LUA_DIR)/lua-$(WASM_LUA_VER)/src
WASM_LUA=$(WASM_LUA_DIR)/lnslua.wasm
# io・os・package などのライブラリは除く (ホスト側で提供する)
WASM_LUA_MODS=lapi lcode lctype ldebug ldo ldump lfunc lgc llex lmem lobject	\
	lopcodes lparser lstate lstring ltable ltm lundump lvm lzio		\
	lauxlib lbaselib ltablib lstrlib lmathlib lutf8lib
# l_signalT: wasm には signal.h が無い
# l_randomizePivot: table.sort の pivot の乱数化に clock() を使わない
WASM_LUA_CFLAGS=-Os -DLUA_COMPAT_5_2 -Dl_signalT=int '-Dl_randomizePivot()=0'	\
	-mllvm -wasm-enable-sjlj -mexec-model=reactor -Wl,--strip-all

$(WASM_LUA_SRC_DIR)/lua.h:
	mkdir -p $(WASM_LUA_DIR)
	curl -sSL -o $(WASM_LUA_DIR)/lua-$(WASM_LUA_VER).tar.gz \
		https://www.lua.org/ftp/lua-$(WASM_LUA_VER).tar.gz
	echo "$(WASM_LUA_SHA256)  $(WASM_LUA_DIR)/lua-$(WASM_LUA_VER).tar.gz" | sha256sum -c
	tar xzf $(WASM_LUA_DIR)/lua-$(WASM_LUA_VER).tar.gz -C $(WASM_LUA_DIR)
	touch $@

$(WASM_LUA): misc/wasm/lnslua.c $(WASM_LUA_SRC_DIR)/lua.h
	$(WASI_SDK)/bin/clang --target=wasm32-wasip1 $(WASM_LUA_CFLAGS)	\
		-I$(WASM_LUA_SRC_DIR)						\
		$(addprefix $(WASM_LUA_SRC_DIR)/,$(addsuffix .c,$(WASM_LUA_MODS)))	\
		misc/wasm/lnslua.c -lsetjmp -o $@

wasm-lua: $(WASM_LUA)
# test/tutorial/class1.lns は Lua 版でもコンパイルエラーになるテストなので除外する
WASM_TEST_SRCS:=$(wildcard test/wasm/*.lns) \
	$(filter-out test/tutorial/class1.lns,$(wildcard test/tutorial/*.lns))

test-wasm: $(WASM_LUA)
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
		$(NODE) misc/wasm/lnswasm.mjs					\
			$(WASM_WORK_DIR)/$$dir/$$name.wasm			\
			> $(WASM_WORK_DIR)/$$name.wasm.result 2>&1 || exit 1;	\
		diff -u $(WASM_WORK_DIR)/$$name.lua.result			\
			$(WASM_WORK_DIR)/$$name.wasm.result || exit 1;		\
		echo "ok -- $$name ($$(stat -c %s $(WASM_WORK_DIR)/$$dir/$$name.wasm) bytes)"; \
	done

# go 版のネイティブ実行と wasm 実行の速度を比べる (docs/sw_design/wasm_benchmark.md)
# WASM_OPT に Binaryen の wasm-opt を指定すると、wasm-opt -O3 した wasm も測る
bench-wasm: $(WASM_LUA)
	WASM_OPT=$(WASM_OPT) misc/wasm/bench.sh
