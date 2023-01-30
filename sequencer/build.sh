#!/bin/bash
protoc --gogofast_out=plugins=grpc:. ./sequencer.proto
