#!/bin/bash
protoc --gogofast_out=plugins=grpc:. ./*.proto