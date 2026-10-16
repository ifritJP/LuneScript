accept-package:
ifndef ACCEPT
	$(MAKE) -i diff-result
	@read -p "accept? $1 (y/n): " ANS; \
	if [ $${ANS}x != yx ]; then exit 1; fi;
endif
#	rm lune/base/*.lua
	cp work/test2.result test/result/test1.result
	cp work/3/lune/base/*.lua $(ORG_DIR)lune/base/
	cp work/3/lune/base/*.meta $(ORG_DIR)lune/base/
#	cp work/3/lune/*.* $(ORG_DIR)lune/
	if [ -f work/3/$(LNSC_GO) ]; then			\
		cp work/3/$(LNSC_GO) $(ORG_DIR).;		\
		cp work/3/lune/base/*.go $(ORG_DIR)lune/base/;	\
	fi
ifneq ($(ENABLE_LUA51),)
	cp work/6/lune/base/*.lns $(ORG_DIR)lune/base/runtime_go/legacy/lua51/lune/base
	cp work/6/lune/base/*.lua $(ORG_DIR)lune/base/runtime_go/legacy/lua51/lune/base
	cat work/6/lune/base/binding.go | sed 's@!gopherlua@gopherlua@g' |	\
		sed 's/package lnsc/package runtimelns/g'		|	\
		sed 's/func InitBinding/func init_bind/g'			|	\
		sed 's@import \.@//import .@g'				|	\
		sed 's@:embed @:embed legacy/lua51/lune/base/@g' >		\
		$(ORG_DIR)lune/base/runtime_go/binding51.go
#	cp work/6/lune/*.* $(ORG_DIR)lune/legacy/lua51/lune/ 
endif


test-installed-cmd:
	$(MAKE) ENABLE_LUAROCKS=y LANG_KIND=luarocks test-comp -j

test-on-docker:
	sudo docker build -t testlns -f test/docker/test_install test/docker
	sudo docker run --name test -it testlns \
		/bin/ash -c "make test-installed-cmd NO_LUAROCKS=y SKIP_PROJ_TEST=y"
	sudo docker container rm test
	sudo docker image rm testlns

test-on-docker_luarocks:
	sudo docker build -t testlns -f test/docker/test_install_with_luarock test/docker
	sudo docker run --name test -it testlns \
		/bin/ash -c 'PATH=/usr/local/bin:${PATH} make test-installed-cmd SKIP_PROJ_TEST=y LUAROCKS=/usr/bin/luarocks-5.3'
	sudo docker container rm test
	sudo docker image rm testlns

VER_PATH=lune/base/Ver.lua
RUNTIME_VER_PICK=grep 'local luaModVersion' | sed 's/[^0-9]//g'
ifneq ($(findstring $(MAKECMDGOALS),"check-runtime-format" "diff-result"),)
RUNTIME_ORG_VER=$(shell cat $(VER_PATH) | $(RUNTIME_VER_PICK))
RUNTIME_NEW_VER=$(shell cat work/1/$(VER_PATH) | $(RUNTIME_VER_PICK))
endif

check-runtime-format:
# ランタイムバージョンが同じ場合は、ランタイムが一致しているかどうか比較
ifeq ($(RUNTIME_ORG_VER),$(RUNTIME_NEW_VER))
	@diff lune/base/runtime.lua work/1/lune/base/runtime.lua > /dev/null || \
	(echo need to change the runtime version.; exit 1;)
else
	@echo runtime version ok.
endif



ifneq ($(findstring $(MAKECMDGOALS),"check-meta-format" ),)
VER_PICK=grep 'local metaVersion' | awk '//{print $$4}'
ORG_VER=$(shell cat $(VER_PATH) | $(VER_PICK))
NEW_VER=$(shell cat work/3/$(VER_PATH) | $(VER_PICK))
endif

check-meta-format:
# メタフォーマットバージョンが同じ場合は、メタ情報が一致しているかどうか比較
ifeq ($(ORG_VER),$(NEW_VER))
	cat work/1/test/Sub2.meta | grep -v -e '^_moduleObj.__version' > work/1/test/Sub2.meta2
	cat work/4/test/Sub2.meta | grep -v -e '^_moduleObj.__version' > work/4/test/Sub2.meta2
	@diff -I '__formatVersion' -I 'buildId = ' \
		work/1/test/Sub2.meta2 work/4/test/Sub2.meta2 > /dev/null || \
	(echo need to change the meta format version.; exit 1;)
else
	@echo meta format version ok.
endif



GLUE_DIR=work/glue
build-glue:
	gcc $(GLUE_DIR)/test_glueTest_glue.c test/glue.c -I$(GLUE_DIR) \
		-std=c99 -fPIC -shared -o test/glueTest$(LUA_GLUE).so $(LUA_CFLAGS)


check_lua_mod_dir:
ifeq ($(LUA_MOD_DIR),)
	@echo not found lua command.
	@echo please retry following command.
	@echo 
	@echo "    make install LUA='your lua command'".
	@echo 
	@exit 1
endif

install: check_lua_mod_dir
	$(MAKE) build LUA=$(LUA) LNSDIR=$(LNSDIR)
	@mkdir -p $(LUA_MOD_DIR)/lune/base
	@cat lnsc.lua | sed "s@#!.*@#! /usr/bin/env $(abspath $(LUA))@g" > /tmp/lnsc
	@chmod +x /tmp/lnsc
	@$(INSTALL) /tmp/lnsc /usr/bin
	@echo installed -- /usr/bin/lnsc
	@rm /tmp/lnsc
	@$(INSTALL) $(LNSDIR)/lune/base/*.* $(LUA_MOD_DIR)/lune/base
	@$(INSTALL) $(LNSDIR)/lune/*.* $(LUA_MOD_DIR)/lune
	@echo installed -- $(LUA_MOD_DIR)/lune

uninstall: check_lua_mod_dir
	rm /usr/bin/lnsc
	rm -rf $(LUA_MOD_DIR)/lune


# テストケースだけ修正した時に、通常のビルドチェックは重すぎるので、
# Lua に変換した結果が同じになっているかどうかを確認し、同じなら OK にする
build-testcase:
	$(call cp_src_set,work/1,.)
	$(call selfhost_build,$(LNSC),.,work/1,--testing)




