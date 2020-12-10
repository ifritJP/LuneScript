FROM centos:7

RUN yum -y update && yum clean all
RUN yum install -y epel-release

RUN yum install -y \
      git \
      make \
      diffutils \
      gcc \
      time \
      readline-devel \
      golang
RUN git clone --depth 1 https://github.com/ifritJP/LuneScript.git
WORKDIR LuneScript/src
RUN git pull
RUN make setup-pkg-lua && make go-build LUA=/LuneScript/src/work/lua/src/lua
RUN install lnsc /usr/bin
RUN lnsc --version
