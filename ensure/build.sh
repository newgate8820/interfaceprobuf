#!/bin/bash
protoc --gogofast_out=plugins=grpc:. ./ensure.proto