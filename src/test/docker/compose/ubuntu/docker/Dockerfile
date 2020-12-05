FROM ubuntu:latest

RUN apt update
RUN apt-get install -y \
	git \
	make \
	diffutils \
	gcc \
	curl \
	libreadline-dev \
	rsync \
	time \
	software-properties-common
RUN add-apt-repository -y ppa:longsleep/golang-backports
RUN apt update
RUN apt-get install -y golang-go

RUN git clone --depth 1 https://github.com/ifritJP/LuneScript.git
WORKDIR LuneScript/src
RUN git reset --hard HEAD
RUN git pull
RUN git log
RUN make setup-pkg-lua && make go-build LUA=/LuneScript/src/work/lua/src/lua
RUN install lnsc /usr/bin
RUN lnsc --version
