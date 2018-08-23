LUA=lua5.3
#LUA=lua5.2

NDIR=$(shell pwd)

.PHONY: help exec

help:
	@echo usage:
	@echo "  " make test-comp


exec:
	$(LUA) lune/base/base.lua test/test.lns token
	$(LUA) lune/base/base.lua test/test.lns ast
	$(LUA) lune/base/base.lua test/test.lns LUA


define trans
	$(LUA) lune/base/base.lua lune/base/$1.lns SAVE work/1
endef

define cutMeta
	(cd work/$2/lune/base; cat $1.lua | awk '/^----- meta/{META++;}; //{if (META != 1) {print $$0}} ' > $1.cut.lua )
endef

define transwork
	(cd work/$1; $(LUA) lune/base/base.lua lune/base/$3.lns SAVE ../$2)
endef

define exeLune
	(cd $1; $(LUA) lune/base/base.lua $2 $3 $4)
endef

define exeComp
	echo lune | cat $3 - | $(call exeLune,$1,$2,comp -i,$4) | \
			json_pp --json_opt=canonical,pretty | sort
endef

test-comp:
	rm -rf work
	mkdir -p work/1/lune/base
	mkdir -p work/2/lune/base
	mkdir -p work/3/lune/base
	@cp lune/base/base*.lua lune/base/*.lns work/1/lune/base
	@cp lune/base/base*.lua lune/base/*.lns work/2/lune/base
	@cp lune/base/base*.lua lune/base/*.lns work/3/lune/base
	$(call trans,Depend)
	$(call trans,Util)
	$(call trans,Writer)
	$(call trans,Parser)
	$(call trans,Option)
	$(call trans,Ast)
	$(call trans,TransUnit)
	$(call trans,convLua)
	$(call trans,dumpNode)
	$(MAKE) test-comp2

test-error:
	-$(call exeLune,$(TEST_DIR),test/error/nilable1.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/access1.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/access2.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/access3.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/access4.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/access5.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/list.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/map.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/scope1.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/scope2.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/func.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/class.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/provide1.lns,exe,2>&1)
	-$(call exeLune,$(TEST_DIR),test/error/provide2.lns,exe,2>&1)


test-comp2:
	$(call exeLune,.,test/start.lns,exe,) > work/test.result


	cp work/test.result work/test1.result
	$(call exeComp,.,test/comp.lns,test/comp.lns,test.comp 15 7) >> work/test1.result
	$(call exeComp,.,test/comp2.lns,test/compSub.lns,test.compSub 5 6) >> work/test1.result
	$(MAKE) -s TEST_DIR=. test-error | grep 'error:' >> work/test1.result

#		sed 's/.*error:/error:/' >> work/test1.result

	$(call transwork,1,2,Depend)
	$(call transwork,1,2,Util)
	$(call transwork,1,2,Writer)
	$(call transwork,1,2,Parser)
	$(call transwork,1,2,Option)
	$(call transwork,1,2,Ast)
	$(call transwork,1,2,TransUnit)
	$(call transwork,1,2,convLua)
	$(call transwork,1,2,dumpNode)


	$(call transwork,2,3,Depend)
	$(call transwork,2,3,Util)
	$(call transwork,2,3,Writer)
	$(call transwork,2,3,Parser)
	$(call transwork,2,3,Option)
	$(call transwork,2,3,Ast)
	$(call transwork,2,3,TransUnit)
	$(call transwork,2,3,convLua)
	$(call transwork,2,3,dumpNode)

	diff -r work/2 work/3

	cp -a test work/3
	$(call exeLune,work/3,test/start.lns,ast,)
	$(call exeLune,work/3,test/start.lns,lua,)
	$(call exeLune,work/3,test/start.lns,exe,) | tee work/test2.result
	$(call exeComp,work/3,test/comp.lns,test/comp.lns,test.comp 15 7) | tee -a work/test2.result
	$(call exeComp,work/3,test/comp2.lns,test/compSub.lns,test.compSub 5 6) | tee -a work/test2.result
	$(MAKE) -s TEST_DIR=work/3 test-error | grep 'error:' | \
				tee -a work/test2.result
#sed 's/.*error:/error:/' | tee -a work/test2.result


## テストスクリプトを lua 化して、 lunescript なしで動かすテスト。
	$(call exeLune,work/3,test/Sub2.lns,save,) > /dev/null
	$(call exeLune,work/3,test/Sub3.lns,save,) > /dev/null
	$(call exeLune,work/3,test/Sub4.lns,save,) > /dev/null
	$(call exeLune,work/3,test/simple.lns,save,) > /dev/null
	$(call exeLune,work/3,test/start.lns,save,) > /dev/null
	(cd work/3; $(LUA) test/start.lua) > work/test3.result


	diff work/test1.result work/test2.result
	diff work/test.result work/test3.result

	$(MAKE) accept-package ACCEPT=y

accept-package:
ifndef ACCEPT
	@read -p "accept? $1 (y/n): " ANS; \
	if [ $${ANS}x != yx ]; then exit 1; fi;
endif
	cp work/3/lune/base/*.lua lune/base/
	cp work/3/lune/base/*.meta lune/base/
