LUA=lua5.3

NDIR=$(shell pwd)

.PHONY: test

exec:
	$(LUA) base.lua test/test.lns token
	$(LUA) base.lua test/test.lns ast
	$(LUA) base.lua test/test.lns lua

test:
	LANG= $(MAKE) exec | \
		sed "s@$(NDIR)@NDIR@g" | \
                tee result
	diff result.expect result 
