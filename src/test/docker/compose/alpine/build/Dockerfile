FROM golang:1.18-alpine

ARG PROXY=""

# proxy 設定
ENV HTTP_PROXY=$PROXY
ENV http_proxy=$PROXY
ENV https_proxy=$PROXY
ENV HTTPS_PROXY=$PROXY

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
	
WORKDIR /
RUN git clone --depth 1 https://github.com/ifritJP/LuneScript.git
WORKDIR /LuneScript/src
RUN git reset --hard HEAD
RUN git pull
RUN git log
RUN make setup-pkg-lua && make go-build LUA=/LuneScript/src/work/lua/src/lua
RUN install lnsc /usr/bin
RUN lnsc --version
