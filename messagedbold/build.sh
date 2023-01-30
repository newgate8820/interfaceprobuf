#!/bin/bash
protoc --gogofast_out=plugins=grpc:. ./hbase.proto
protoc  --java_out=. ./hbase.proto
