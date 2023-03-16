LNSC=lnsc

SRCS=
SRCS += $(shell find lns -iname '*.lns')

ifneq "$(wildcard Makefile.local)" ""
include Makefile.local
endif


.PHONY: lnstags


help:
	@echo make build [LNS_OPT=opt]

