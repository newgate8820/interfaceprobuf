all: clean build

build:
	cd ../../../
	pwd
	protoc --gogofast_out=plugins=grpc:. ./gitlab.chatserver.im/interfaceprobuf/pbimofflineserver/imofflineserver.proto
	cd gitlab.chatserver.im/interfaceprobuf/pbimofflineserver

clean:
	rm -rf ./imofflineserver.pb.go
