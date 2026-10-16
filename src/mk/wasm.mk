# wasm (WasmGC) への直接変換のテスト
#
# test/wasm/*.lns を、Lua で実行した結果と、wasm に変換して Node.js で実行した結果で比べる。
# 設計は docs/sw_design/wasm_direct_design.md を参照。

NODE=node
WASM_WORK_DIR=work/wasm
WASM_TEST_SRCS:=$(wildcard test/wasm/*.lns)

test-wasm:
	rm -rf $(WASM_WORK_DIR)
	mkdir -p $(WASM_WORK_DIR)
	@for src in $(WASM_TEST_SRCS); do					\
		name=$$(basename $$src .lns);					\
		mkdir -p $(WASM_WORK_DIR)/$$(dirname $$src);			\
		$(LNSC) $$src exe $(EXE_OPT)					\
			> $(WASM_WORK_DIR)/$$name.lua.result 2>&1 || exit 1;	\
		$(LNSC) $$src SAVE $(WASM_WORK_DIR) -langWasm > /dev/null || exit 1; \
		$(NODE) misc/wasm/lnswasm.mjs					\
			$(WASM_WORK_DIR)/test/wasm/$$name.wasm			\
			> $(WASM_WORK_DIR)/$$name.wasm.result 2>&1 || exit 1;	\
		diff -u $(WASM_WORK_DIR)/$$name.lua.result			\
			$(WASM_WORK_DIR)/$$name.wasm.result || exit 1;		\
		echo "ok -- $$name ($$(stat -c %s $(WASM_WORK_DIR)/test/wasm/$$name.wasm) bytes)"; \
	done
