LNSC=lnsc
LNSTAGS=lnstags

SRCS=
SRCS += $(shell find lns -iname '*.lns')
GO_SRCS += $(patsubst %.lns,%.go,$(SRCS))

ifneq "$(wildcard Makefile.local)" ""
include Makefile.local
endif


.PHONY: build


help:
	@echo make build [LNS_OPT=opt]
	@echo make build-go
	@echo make update_lnstags
	@echo make clean
	@echo make view-goprof

update_lns:
	go get -u github.com/ifritJP/LuneScript/src@latest

build-go:
	(cd entry; go build -tags gopherlua -o ../formatter)

build:
	@echo $(SRCS) | sed 's/ /\n/g' |	\
		$(LNSC) @- save -langGo -noLua --main lns.main $(LNS_OPT)
	$(LNSC) lns/main.lns mkmain entry/entry.go
	$(MAKE) build-go

goprof:
	./formatter -prof ../../src/lune/base/Ast.lns dummy.lns

view-goprof:
	go tool pprof -http=0.0.0.0:8100 -no_browser formatter prof


clean:
	-rm $(GO_SRCS)

update_lnstags:
	$(LNSTAGS) build lns/main.lns --log log
