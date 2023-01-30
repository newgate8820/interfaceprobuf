#!/bin/bash
cd ../../..
protoc --gogofast_out=plugins=grpc:. ./gitlab.chatserver.im/interfaceprobuf/botserver/botserver.proto
