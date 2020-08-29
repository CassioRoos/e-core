FROM golang:1.14

COPY . $GOPATH/src/github.com/CassioRoos/e-core
COPY docker/e2e/scripts /scripts

WORKDIR $GOPATH/src/github.com/CassioRoos/e-core

ENTRYPOINT ["bash", "/scripts/tests_e2e.sh"]
