help:
	@echo usage:
	@echo "  make install"
	@echo "  make test-comp [NO_ACCEPT=y]"
	@echo "  make test-comp-with-go [NO_ACCEPT=y]"
	@echo "  make test-global"
	@echo "  make release-check"
	@echo "  make accept-package [WITH_GO=y]"
	@echo "  make build-glue [LUA_INC_DIR=path]"
	@echo "  make build-cc"
	@echo "  make test-miniC"
	@echo "  make test-subC"
	@echo "  make test-testC"
	@echo "  make test-formatter"
	@echo "  make test-mkbuiltin"
	@echo "  make build-testcase"
	@echo "  make test-tutorial"
	@echo "  make test-installed-cmd"
	@echo "  make test-on-docker"
	@echo "  make test-reference"
	@echo "  make test-go"
	@echo "  make test-miniGo [NO_TEST=y] [NO_LUA_DIFF=y]"
	@echo "  make scratch-go"
	@echo "  make test-tutorial-go"
	@echo "  make go-build [GOTAG=gopherlua] [NO_TEST=y] [NO_ACCEPT=y]"
	@echo "  make setup-pkg-lua"
	@echo "  make test-govslua [NO_BATCH=y] [ONLY_CONV=y]"
	@echo "  make goprof [LNSC_OP=]"
	@echo "  make view-goprof [BROWSER=y]"
	@echo "  make update-lnstags"


exec:
	$(LUA) lune/base/base.lua test/test.lns token
	$(LUA) lune/base/base.lua test/test.lns ast
	$(LUA) lune/base/base.lua test/test.lns LUA



