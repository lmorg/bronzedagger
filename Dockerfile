FROM golang:latest

RUN mkdir -p $GOPATH/src/github.com/lmorg/bronzedagger

COPY * $GOPATH/src/github.com/lmorg/bronzedagger/

RUN cd $GOPATH/src/github.com/lmorg/bronzedagger; \
    go install $GOPATH/src/github.com/lmorg/bronzedagger

ENTRYPOINT [ "/go/bin/bronzedagger" ]