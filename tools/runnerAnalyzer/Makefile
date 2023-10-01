LNSC=lnsc
LNSTAGS=lnstags
PLANTUML=plantuml.jar

SRCS=
SRCS += $(shell find lns -iname '*.lns')
GO_SRCS += $(patsubst %.lns,%.go,$(SRCS))

ifneq "$(wildcard Makefile.local)" ""
include Makefile.local
endif


.PHONY: build test


help:
	@echo make build [LNS_OPT=opt]
	@echo make build-go
	@echo make exec
	@echo make clean
	@echo make view-goprof

update_lns:
	go get -u github.com/ifritJP/LuneScript/src@latest

build-go:
	(cd entry; go build -tags gopherlua)

build:
	@echo $(SRCS) | sed 's/ /\n/g' |	\
		$(LNSC) @- save -langGo -noLua --main lns.main $(LNS_OPT)
	$(LNSC) lns/main.lns mkmain entry/entry.go
	$(MAKE) build-go

clean:
	-rm $(GO_SRCS)


exec:
	$(LNSC) lns/main.lns -shebang ../../src/@-.profirunner hoge.txt runner2.csv
	java -jar $(PLANTUML) hoge.txt -tsvg
