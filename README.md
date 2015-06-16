# bronzedagger
Simple load tester and page testing

This is still very much beta and under active development. But most of the basic features I want are already implemented.

Logging on bronzedagger can also be imported into firesword (https://github.com/lmorg/firesword) for analytics.

### How to compile:

If you don't already have Go installed, then that can be downloaded from https://golang.org/ or installed via your OS's package manager (if available). Make sure you follow the guide to set your GOPATH environmental variables, eg `export GOPATH=~/go; mkdir -p $GOPATH/src $GOPATH/bin`

Then download this repo and install

    go get github.com/lmorg/bronzedagger
    go install github.com/lmorg/bronzedagger
