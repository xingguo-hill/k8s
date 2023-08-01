FROM golang:1.21-rc-alpine3.17
ENV VERSION=2 
WORKDIR /go/src/
ADD https://github.com/xingguo-hill/k8s/archive/refs/heads/main.zip /go/src/
RUN unzip main.zip \
  && cd k8s-main/work2/ \
  && go env -w GOPROXY=https://goproxy.cn,direct \
  && go env -w GO111MODULE=on \
  && go mod init demo \
  && go mod tidy \
  && go build -v -o /go/bin/ \
  && rm -rf /go/src/*
ENTRYPOINT /go/bin/demo -alsologtostderr true