# sudo docker build --rm --build-arg PROXY=    Dockerfile

#### buster 環境で lnsc をビルド
FROM golang:1.16-buster

ARG PROXY=""

# apt-get で入力待ちにならないようにするためセット
ENV DEBIAN_FRONTEND=noninteractive
# proxy 設定
ENV HTTP_PROXY=$PROXY
ENV http_proxy=$PROXY
ENV https_proxy=$PROXY
ENV HTTPS_PROXY=$PROXY
RUN if [ "$PROXY" != "" ]; then						\
	echo set $PROXY;						\
	echo "Acquire::http::Proxy \"$PROXY\";" >> /etc/apt/apt.conf;	\
	echo "Acquire::https::Proxy \"$PROXY\";" >> /etc/apt/apt.conf;	\
fi
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
RUN if [ "$PROXY" != "" ]; then			\
    git config --global http.proxy $PROXY;	\
fi
RUN apt update

WORKDIR /
RUN git clone --depth 1 https://github.com/ifritJP/LuneScript.git
WORKDIR /LuneScript/src
RUN git reset --hard HEAD
RUN git pull
RUN git log
RUN set
RUN make setup-pkg-lua && make go-build LUA=/LuneScript/src/work/lua/src/lua
RUN install lnsc /usr/bin
RUN lnsc --version

#### centos で動作チェック

FROM centos:7

ARG PROXY=""

# multistage でビルドしたものをコピー
COPY --from=0 /usr/bin/lnsc /usr/bin/
#COPY lnsc /usr/bin/

ENV HTTP_PROXY=$PROXY
ENV http_proxy=$PROXY
ENV https_proxy=$PROXY
ENV HTTPS_PROXY=$PROXY

RUN yum -y update && yum clean all
RUN yum install -y epel-release

RUN yum install -y \
      git \
      make \
      diffutils \
      golang \
      time
RUN git clone --depth 1 https://github.com/ifritJP/LuneScript.git
WORKDIR LuneScript/src
RUN git pull
RUN make test-go LNSC=lnsc GOTAG=gopherlua
