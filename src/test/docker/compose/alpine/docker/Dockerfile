FROM alpine:latest

RUN apk --no-cache add \
	git \
	make \
	diffutils \
	gcc \
	libc-dev \
	curl \
	readline-dev \
	rsync \
	go
RUN git clone --depth 1 https://github.com/ifritJP/LuneScript.git
WORKDIR LuneScript/src
RUN git reset --hard HEAD
RUN git pull
RUN git log
RUN make setup-pkg-lua && make go-build LUA=/LuneScript/src/work/lua/src/lua
RUN install lnsc /usr/bin
RUN lnsc --version

