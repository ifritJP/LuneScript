LUA=lua5.3

NDIR=$(shell pwd)

.PHONY: test help exec

help:
	@echo usage:
	@echo "  " make test-comp


exec:
	$(LUA) lune/base/base.lua test/test.lns token
	$(LUA) lune/base/base.lua test/test.lns ast
	$(LUA) lune/base/base.lua test/test.lns lua

test:
	LANG= $(MAKE) exec | \
		sed "s@$(NDIR)@NDIR@g" | \
                tee result
	diff result.expect result 

define trans
	echo test $1
	$(LUA) primal/base.lua lune/base/$1.lns lua > work/lune/base/$1.lua
	cp lune/base/$1.lns work/lune/base
endef

define cutMeta
	(cd work/lune/base; cat $1.lua | awk '/^----- meta/{META++;}; //{if (META != 1) {print $$0}} ' > $1.cut.lua )
endef

define transwork
	(cd work; $(LUA) lune/base/base.lua lune/base/$1.lns lua > lune/base/$(1)3.lua )
	$(call cutMeta,$1)
	$(call cutMeta,$(1)3)
	diff work/lune/base/$1.cut.lua work/lune/base/$(1)3.cut.lua;
endef

test-comp:
	mkdir -p work/lune/base
	cp lune/base/base.lua work/lune/base
	$(call trans,Parser)
	$(call trans,TransUnit)
	$(call trans,convLua)
	$(call trans,dumpNode)
	$(MAKE) test-comp2

test-comp2:
	$(call transwork,Parser)
	$(call transwork,TransUnit)
	$(call transwork,convLua)
	$(call transwork,dumpNode)

	rm work/lune/base/*3.lua
	cp work/lune/base/*.lua lune/base/
