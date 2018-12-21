PROJ_ROOT=test/proj


define comp
	@echo "$1 -> $2"
	@(cd ../..; lua5.3 lune/base/base.lua $(PROJ_ROOT)/$1 --depends $(PROJ_ROOT)/depends/$1.d.tmp SAVE)
	@cat depends/$1.d.tmp | sed 's@$(PROJ_ROOT)/@@g' > depends/$1.d
	@rm depends/$1.d.tmp
endef


%.meta: %.lns
	$(call comp,$<,$@)

SRCS =
SRCS += Mod1.lns
SRCS += Mod2.lns
SRCS += Mod3.lns
SRCS += Mod4.lns
#SRCS := $(addprefix $(PROJ_ROOT)/,$(SRCS))

META_LIST=$(SRCS:.lns=.meta)
LUA_LIST=$(SRCS:.lns=.lua)

-include depends/*.d


all:
	@echo make setup
	@echo make build

setup:
	mkdir -p depends

build: $(META_LIST)


clean:
	@rm $(META_LIST) $(LUA_LIST) depends/*


test:
	@echo =======================
	touch Mod4.lns
	@$(MAKE) build
	@echo =======================
	touch Mod3.lns
	@$(MAKE) build
	@echo =======================
	touch Mod2.lns
	@$(MAKE) build
	@echo =======================
	touch Mod1.lns
	@$(MAKE) build
