#!/usr/bin/env bash

CURDIR=`pwd`
OLDGOPATH="$GOPATH"
export GOPATH="$CURDIR"

gofmt -w src/

go install server

export GOPATH="$OLDGOPATH"

echo 'finished'