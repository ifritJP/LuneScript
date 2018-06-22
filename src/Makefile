LUA=lua5.3

.PHONY: test

exec:
	$(LUA) base.lua test/test.ls

test:
	$(MAKE) exec | tee result
	diff result.expect result 
