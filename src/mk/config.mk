LUA=lua5.3
#LUA=lua5.2


LUACHECK=~/.luarocks/bin/luacheck

PKG_NAME=lua-5.3
LUA_LIB=$(shell pkg-config --libs $(PKG_NAME))
LUA_INC_DIR=$(shell pkg-config --cflags-only-I $(PKG_NAME) | sed 's/^-I//')
ifneq ($(LUA_INC_DIR),)
LUA_CFLAGS=-I$(LUA_INC_DIR)
endif

PKG_NAME_51=lua-5.1
LUA51=lua5.1
LUA51_INC=$(shell pkg-config --cflags-only-I $(PKG_NAME_51) | sed 's/^-I//')


LUAROCKS=luarocks

ifdef LOG
LOG_OPT=--log info
endif

LNSC_LUA=$(LUA) lune/base/base.lua
LNSC_GO=./lnsc
LNSC_LUAROCKS=~/.luarocks/bin/lnsc

LNSC_LUA51=$(LUA51) lune/base/base.lua

# make のパラメータで言語を切り替える場合、
# LNSC_CMD や LNSC を直接変更するのではなく LANG_KIND を変更する。
LANG_KIND=lua

ifeq ($(LANG_KIND),lua)
LNSC_CMD=$(LNSC_LUA)
endif
ifeq ($(LANG_KIND),go)
LNSC_CMD=$(LNSC_GO)
endif
ifeq ($(LANG_KIND),luarocks)
LNSC_CMD=$(LNSC_LUAROCKS)
endif

ifeq ($(LNSC_CMD),)
$(error "LANG_KIND is illegal -- " $(LNSC_CMD) )
endif

LNSC=LUA_PATH=./?.lua $(LNSC_CMD)

ifdef ENABLE_LUAROCKS
ifneq ($(LUA51),)
ENABLE_LUA51:=$(shell which $(LUA51))
endif
ifndef NO_LUAROCKS
LUAROCKS_VER:=$(shell $(LUAROCKS) | grep 'Lua version' | awk '//{print $$3}')
endif
endif

ifneq ($(findstring $(MAKECMDGOALS),"install" "uninstall"),)
# LUA_DIR_SCRIPT='for path in string.gmatch( package.path, "[^;]+" ) do if not path:find( "^%./" ) then print( (path:gsub( "/%?%.lua", "" )) ); break end end'
# LUA_MOD_DIR=$(shell echo $(LUA_DIR_SCRIPT) | $(LUA) 2> /dev/null)
ifneq "$(wildcard lune.mk)" ""
include lune.mk
endif
endif

INSTALL=/usr/bin/install -c


# $(PWD) だと make -C dir で実行した時に、
# $(PWD) した結果は dir ではなく、 make を実行したディレクトリになるので
# $(shell pwd) で取得する。
CUR_DIR:=$(shell pwd)

LNSC_OP=

CONVGO_OP=
COMMON_OP=
EXE_OPT := --valid-luaval
COM_OPT := --use-ipairs -Werror --valid-luaval $(LNSC_OP)
ifeq ($(WITH_GO),y)
COM_OPT += -langGo
endif

# 標準動作だと diff が出てしまうので、 diff が出ないようにオプションを指定する
TEST_OPT=--threadPerUnitThread 0 --convGoRunnerNum 4



ifdef depend_on
DEPENDS := $(shell find lune -iname '*.d')
DEPENDS_META := $(patsubst %.d,%.meta,$(DEPENDS))

include $(DEPENDS)

$(DEPENDS_META):

%.meta: %.lns
	(cd $(FROM_DIR); $(LNSC) $< SAVE $(TO_DIR) \
			$(BYTECOMP) $(COM_OPT) $(LOG_OPT) $(LNS_OPT))

build_from_d: $(DEPENDS_META)
ifeq ($(WITH_GO),y)
	$(call binding-lns-go,.)
ifndef SKIP_GO_BUILD
	$(call go-build-lns,.)
endif
endif

endif


# sort コマンドに環境依存があるので、 sort コマンドを使わずにソートするスクリプト
SORT:=python3 -c 'import sys; print( "".join( sorted( sys.stdin.readlines() ) ) )'


ifneq "$(wildcard Makefile.local)" ""
include Makefile.local
endif

.PHONY: help exec build build_from_d luarocks-check luarocks-check-sub
.PHONY: release-check build test-save build-countup test-comp
.PHONY: test-registered-rockspec test-etc test-error-all test-error test-org
.PHONY: test-new test-parallel test-comp2 check-rockspec diff-result
.PHONY: test-lua51 test-lua51-sub accept-package check-runtime-format
.PHONY: check-meta-format build-glue check_lua_mod_dir install uninstall
.PHONY: test-ccc clean-ccc build-miniC test-miniC test-subC build-cc test-formatter
.PHONY: test-installed-cmd test-on-docker


