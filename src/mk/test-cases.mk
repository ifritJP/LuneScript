test-etc:
# 存在しないファイルを指定した時に、不正なファイルが作成されないことを確認する
	-@$(call exeLune,$(LNSC),$(TEST_DIR),test/error/__.lns,SAVE,2>&1) > /dev/null
	@echo -- not exist file access --
	@(cd $(TEST_DIR); find test/error -iname '__.*')


TEST_CASE += nilable1.lns
TEST_CASE += access1.lns
TEST_CASE += access2.lns
TEST_CASE += access3.lns
TEST_CASE += access4.lns
TEST_CASE += access5.lns
TEST_CASE += list.lns
TEST_CASE += map.lns
TEST_CASE += set.lns
TEST_CASE += scope1.lns
TEST_CASE += scope2.lns
TEST_CASE += scope3.lns
TEST_CASE += func.lns
TEST_CASE += class.lns
TEST_CASE += class2.lns
TEST_CASE += class3.lns
TEST_CASE += class4.lns
TEST_CASE += class5.lns
TEST_CASE += class6.lns
TEST_CASE += class7.lns
TEST_CASE += class8.lns
TEST_CASE += class9.lns
TEST_CASE += class10.lns
TEST_CASE += class11.lns
TEST_CASE += depend.lns
TEST_CASE += provide1.lns
TEST_CASE += provide2.lns
TEST_CASE += provide3.lns
TEST_CASE += shadowing.lns
TEST_CASE += enum.lns
TEST_CASE += eof.lns
TEST_CASE += mutable.lns
TEST_CASE += var.lns
TEST_CASE += var2.lns
TEST_CASE += var3.lns
TEST_CASE += arg.lns
TEST_CASE += arg2.lns
TEST_CASE += unwrap.lns
TEST_CASE += operand.lns
TEST_CASE += operand2.lns
TEST_CASE += loop.lns
TEST_CASE += cast.lns
TEST_CASE += prototype.lns
TEST_CASE += statement.lns
TEST_CASE += flow.lns
TEST_CASE += flow_switch.lns
TEST_CASE += flow_match.lns
TEST_CASE += import.lns
TEST_CASE += import2.lns
TEST_CASE += import3.lns
TEST_CASE += import4.lns
TEST_CASE += import5.lns
TEST_CASE += import6.lns
TEST_CASE += import7.lns
TEST_CASE += strformat.lns
TEST_CASE += Alge.lns
TEST_CASE += Alge2.lns
TEST_CASE += Alge3.lns
TEST_CASE += Alge4.lns
TEST_CASE += exp.lns
TEST_CASE += symbol.lns
TEST_CASE += symbol2.lns
TEST_CASE += symbolEmpty.lns
TEST_CASE += static.lns
TEST_CASE += generics.lns
TEST_CASE += generics2.lns
TEST_CASE += generics3.lns
TEST_CASE += generics4.lns
TEST_CASE += box.lns
TEST_CASE += box2.lns
TEST_CASE += form.lns
TEST_CASE += ddd.lns
TEST_CASE += immediate.lns
TEST_CASE += macro.lns
TEST_CASE += macro2.lns
TEST_CASE += macro3.lns
TEST_CASE += macro4.lns
TEST_CASE += macro5.lns
TEST_CASE += macro6.lns
TEST_CASE += macro7.lns
TEST_CASE += macro8.lns
TEST_CASE += apply.lns
TEST_CASE += multiple-class.lns
TEST_CASE += multiple-if.lns
TEST_CASE += multiple.lns
TEST_CASE += switch.lns
TEST_CASE += lua.lns
TEST_CASE += async1.lns
TEST_CASE += async3.lns
TEST_CASE += async4.lns
TEST_CASE += async5.lns
TEST_CASE += tuple.lns
TEST_CASE += tuple2.lns
TEST_CASE += tuple3.lns
TEST_CASE += condRet.lns
TEST_CASE_SRCS := $(addprefix test/error/,$(TEST_CASE))
TEST_CASE_LOGS := $(addprefix $(LOG_DIR)/error/,$(addsuffix .log,$(TEST_CASE)))



# テストケースを実行して、テストケース毎にログをログディレクトリに出力する
$(LOG_DIR)/error/%.log: test/error/%
	echo $@ > $@
	-$(call exeLune,$(LNSC),$(TEST_DIR),$<,diag --disableMultiPhaseAst --disableMultiThreadAst,$(EXE_OPT) 2>&1)  >> $@

test-error-all: $(TEST_CASE_LOGS)
	cat $(TEST_CASE_LOGS) | grep -e 'error:' -e 'warning:'

# エラー系テストケースを実行し、結果を stdout に出力する
define test-error-all-case
	mkdir -p $1
	$(MAKE) test-error-all LOG_DIR=$1 TEST_DIR=$2
endef

test-error:
	-rm $(LOG_DIR)/error/*
	mkdir -p $(LOG_DIR)/error
	$(call test-error-all-case,$(LOG_DIR),$(TEST_DIR))

	-$(call exeLune,$(LNSC),$(TEST_DIR),test/error/noexist.lns,exe,2>&1)
	-$(call exeLune,$(LNSC),$(TEST_DIR),test/error/abbr.lns,exe $(COM_OPT),2>&1) && echo error:NG

	cat test/error/mutable.lns | sed 's@//_lune@_lune@' > $(TEST_DIR)/test/error/dis_mutable.lns
	-$(call exeLune,$(LNSC),$(TEST_DIR),test/error/dis_mutable.lns,exe,2>&1)
	rm $(TEST_DIR)/test/error/dis_mutable.lns


define exeComp
	echo lune | cat $3 - | $(call exeLune,$(LNSC),$1,$2,comp -i,$4) | \
			json_pp --json_opt=canonical,pretty | $(SORT)
endef


define test-shebang
	cat test/Shebang.lns | \
		sed 's@^#!.*@#! ./lnsc -shebang@g' > $1/test/Shebang2.lns
	chmod +x $1/test/Shebang2.lns
	echo "test ./test/Shebang2" >> $3
	(cd $1; ./test/Shebang2.lns $2 && echo OK || echo __ok) >> $3
	(cd $1; `pwd`/test/Shebang2.lns $2 && echo OK || echo __ok) >> $3
endef

#>>>>>>>>>>>>>>>>>>>>
# 補完系テストを並列実行するための設定郡

define path2mod
	$(shell echo $1 | sed 's:\.lns$$::g' | sed 's:/:.:g')
endef

define exec_compl-wrap2
	$(call exeComp,$(TEST_DIR),$1,$2,$(call path2mod,$2) $3)
endef

# ダミーのターゲットから目的の文字列を切り出す処理
define strip_1
	$(shell echo $1 | sed 's:@1_.*::')
endef
define strip_2
	$(shell echo $1 | sed 's:.*@1_\(.*\)@2_.*:\1:')
endef
define strip_3
	$(shell echo $1 | sed 's:.*@2_\(.*\).comp:\1:' | sed 's:\.: :' )
endef

%.comp:
	$(call exec_compl-wrap2,$(call strip_1,$@),$(call strip_2,$@),$(call strip_3,$@)) > \
		$(LOG_DIR)/$(notdir $(call strip_1,$@))
#<<<<<<<<<<<<<<<<<<<<

# 補完にかけるファイルのパターン。
# PATH1@1_PATH2@2_POS.comp
#  PATH1: 解析にかける lns ファイルパス
#  PATH2: 補完対象の lns ファイルパス
#  POS: 補完対象の位置 (line.column)
COMP_TARGETS += test/comp.lns@1_test/comp.lns@2_17.7.comp
COMP_TARGETS += test/comp-2.lns@1_test/comp-2.lns@2_17.7.comp
COMP_TARGETS += test/comp-3.lns@1_test/comp-3.lns@2_17.7.comp
COMP_TARGETS += test/comp-4.lns@1_test/comp-4.lns@2_17.12.comp
COMP_TARGETS += test/comp-5.lns@1_test/comp-5.lns@2_17.15.comp
COMP_TARGETS += test/comp-6.lns@1_test/comp-6.lns@2_2.6.comp
COMP_TARGETS += test/comp2.lns@1_test/compSub.lns@2_5.6.comp
COMP_TARGETS += test/comp4.lns@1_test/comp4.lns@2_9.8.comp
COMP_TARGETS += test/comp5.lns@1_test/comp5.lns@2_7.7.comp
COMP_TARGETS += test/comp6.lns@1_test/comp6.lns@2_3.2.comp
COMP_TARGETS += test/comp7.lns@1_test/comp7.lns@2_6.7.comp
COMP_TARGETS += test/comp8.lns@1_test/comp8.lns@2_4.15.comp
COMP_TARGETS += test/comp9.lns@1_test/comp9.lns@2_4.3.comp



# コード補完のテスト
test-cmpl: $(COMP_TARGETS)
	(find $(LOG_DIR) -type f | $(SORT) | \
		xargs -n 1 bash -c 'echo $$0 | sed "s@$(LOG_DIR)@@g"; cat $$0') >> $(LOG_FILE)


