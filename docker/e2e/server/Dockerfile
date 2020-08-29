FROM golang:1.14

COPY . $GOPATH/src/github.com/CassioRoos/e-core

WORKDIR $GOPATH/src/github.com/CassioRoos/e-core

RUN go install github.com/CassioRoos/e-core

ENTRYPOINT sleep 90; $GOPATH/bin/e-core
