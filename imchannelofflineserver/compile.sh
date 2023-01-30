#!/usr/bin/env bash
protoc -I$GOPATH/src -I. --go_out=plugins=grpc:. getchanneldifference.proto