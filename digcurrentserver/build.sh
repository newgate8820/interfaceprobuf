#!/bin/bash
protoc --gogofast_out=plugins=grpc:. ./digcurrentserver.proto