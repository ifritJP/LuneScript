LUA=lua5.3
#LUA=lua5.2

NDIR=$(shell pwd)

.PHONY: test help exec

help:
	@echo usage:
	@echo "  " make test-comp


exec:
	$(LUA) lune/base/base.lua test/test.lns token
	$(LUA) lune/base/base.lua test/test.lns ast
	$(LUA) lune/base/base.lua test/test.lns LUA

test:
	LANG= $(MAKE) exec | \
		sed "s@$(NDIR)@NDIR@g" | \
                tee result
	diff result.expect result 

define trans
	echo test $1
	$(LUA) lune/base/base.lua lune/base/$1.lns LUA > work/1/lune/base/$1.lua
	cp lune/base/$1.lns work/1/lune/base
endef

define cutMeta
	(cd work/$2/lune/base; cat $1.lua | awk '/^----- meta/{META++;}; //{if (META != 1) {print $$0}} ' > $1.cut.lua )
endef

define transwork
	(cd work/$1; $(LUA) lune/base/base.lua lune/base/$3.lns LUA > ../$2/lune/base/$(3).lua )
	cp lune/base/$3.lns work/$2/lune/base
endef

test-comp:
	rm -rf work
	mkdir -p work/1/lune/base
	mkdir -p work/2/lune/base
	mkdir -p work/3/lune/base
	cp lune/base/base*.lua work/1/lune/base
	cp lune/base/base*.lua work/2/lune/base
	cp lune/base/base*.lua work/3/lune/base
	$(call trans,Util)
	$(call trans,Parser)
	$(call trans,TransUnit)
	$(call trans,convLua)
	$(call trans,dumpNode)
	$(MAKE) test-comp2

test-comp2:
	$(LUA) lune/base/base.lua test/start.lns exe > test.result

	$(call transwork,1,2,Util)
	$(call transwork,1,2,Parser)
	$(call transwork,1,2,TransUnit)
	$(call transwork,1,2,convLua)
	$(call transwork,1,2,dumpNode)


	$(call transwork,2,3,Util)
	$(call transwork,2,3,Parser)
	$(call transwork,2,3,TransUnit)
	$(call transwork,2,3,convLua)
	$(call transwork,2,3,dumpNode)

	diff -r work/2 work/3

	cp -a test work/3
	(cd work/3; $(LUA) lune/base/base.lua test/start.lns ast)
	(cd work/3; $(LUA) lune/base/base.lua test/start.lns LUA)
	(cd work/3; $(LUA) lune/base/base.lua test/start.lns exe) | tee test2.result

	diff test.result test2.result

	cp work/3/lune/base/*.lua lune/base/
