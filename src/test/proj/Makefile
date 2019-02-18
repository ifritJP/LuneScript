PROJ_DIR=test/proj
MKFILE=$(PROJ_DIR)/Makefile
SRC_DIR=$(PROJ_DIR)/



.PHONY: test all build setup exe

# define comp
# 	@echo "$1 -> $2"
# 	lnsc $1 --depends depends/$(shell echo $1 | sed 's@/@.@g').d SAVE
# endef

define comp
	@echo "$1 -> $2"
	@(cd ../..; lua5.3 lune/base/base.lua $(PROJ_DIR)/$1 --depends $(PROJ_DIR)/depends/$1.d.tmp SAVE)
	@cat depends/$1.d.tmp | sed 's@$(PROJ_DIR)/@@g' > depends/$1.d
	@rm depends/$1.d.tmp
endef



%.meta: %.lns
	$(call comp,$<,$@)

SRCS =
SRCS += Mod1.lns
SRCS += Mod2.lns
SRCS += Mod3.lns
SRCS += Mod4.lns

SRCS := $(addprefix $(SRC_DIR),$(SRCS))


META_LIST=$(SRCS:.lns=.meta)
LUA_LIST=$(SRCS:.lns=.lua)

-include depends/*.d


all:
	@echo make setup
	@echo make build
	@echo make test

setup:
	mkdir -p depends

build: $(META_LIST)

exe: build
	@(cd ../..; lua5.3 $(PROJ_DIR)/Mod1.lua)


clean:
	@rm $(META_LIST) $(LUA_LIST) depends/*


test:
	@$(MAKE) -f $(MKFILE) setup
	@echo =======================
	touch $(SRC_DIR)Mod4.lns
	@$(MAKE) -f $(MKFILE) exe
	@echo =======================
	@sleep 0.5
	touch $(SRC_DIR)Mod3.lns
	@$(MAKE) -f $(MKFILE) exe
	@echo =======================
	@sleep 0.5
	touch $(SRC_DIR)Mod2.lns
	@$(MAKE) -f $(MKFILE) exe
	@echo =======================
	@sleep 0.5
	touch $(SRC_DIR)Mod1.lns
	@$(MAKE) -f $(MKFILE) exe
