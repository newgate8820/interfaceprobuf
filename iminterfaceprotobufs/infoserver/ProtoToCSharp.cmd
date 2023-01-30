@echo off
protoc --gogofast_out=plugins=grpc:. .\gitlab.potato.im\iminterfaceprotobufs\infoserver\*.proto
protoc --gogofast_out=plugins=grpc:. ./gitlab.potato.im/iminterfaceprotobufs/infoserver/*.proto

@echo All Done!
@pause
