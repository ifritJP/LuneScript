LNSC=lnsc

SRCS=
SRCS += $(shell find lns -iname '*.lns')
GO_SRCS += $(patsubst %.lns,%.go,$(SRCS))

ifneq "$(wildcard Makefile.local)" ""
include Makefile.local
endif


.PHONY: build


help:
	@echo make build [LNS_OPT=opt]
	@echo make clean


build:
	@echo $(SRCS) | sed 's/ /\n/g' |	\
		$(LNSC) @- save -langGo -noLua --main lns.main $(LNS_OPT)
	$(LNSC) lns/main.lns mkmain entry.go
	go build -tags gopherlua -o formatter

clean:
	-rm $(GO_SRCS)
