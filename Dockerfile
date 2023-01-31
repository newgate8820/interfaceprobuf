FROM golang:1.19.2

WORKDIR /src

RUN apt-get update
RUN apt install -y protobuf-compiler
RUN protoc --version
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2
#RUN go install github.com/gogo/protobuf/protoc-gen-gofast@latest
#
#RUN go install github.com/gogo/protobuf/protoc-gen-gogofast@latest
#RUN go install github.com/gogo/protobuf/protoc-gen-gogofaster@latest
#RUN go install github.com/gogo/protobuf/protoc-gen-gogoslick@latest

CMD ["bash"]

