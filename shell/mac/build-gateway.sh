#!/usr/bin/env bash

echo "====build gateway===="
#define
PROTO_PATH="../../proto"
GRPC_EXAMPLE_GATEWAY="../../example/gateway"
BIN_EXAMPLE_GATEWAY="../../bin/example/gateway"
PROTO_EXAMPLE_GATEWAY="gateway.proto"

#build
#echo "====rm grpc-http example===="
rm ${BIN_EXAMPLE_GATEWAY}/*

#echo "====protoc grpc-http example===="
protoc --proto_path=${PROTO_PATH} \
--go_out=plugins=grpc:${PROTO_PATH}/gateway \
${PROTO_PATH}/gateway/${PROTO_EXAMPLE_GATEWAY}

#echo "====build grpc-http gateway===="
protoc --proto_path=${GOPATH}/src/ \
--proto_path=${PROTO_PATH} \
--grpc-gateway_out=logtostderr=true:${PROTO_PATH}/gateway \
${PROTO_PATH}/gateway/${PROTO_EXAMPLE_GATEWAY}

#echo "====build grpc-http example===="
#go build -o ${BIN_EXAMPLE_GATEWAY}/client ${GRPC_EXAMPLE_GATEWAY}/client/client.go
go build -o ${BIN_EXAMPLE_GATEWAY}/server ${GRPC_EXAMPLE_GATEWAY}/server/main.go


# 编译google.api
#protoc --proto_path=${PROTO_PATH}/gateway --go_out=plugins=grpc,Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor:. google/api/*.proto

# 编译gateway.proto
#protoc --proto_path=${PROTO_PATH}/gateway --go_out=plugins=grpc,Mgoogle/api/annotations.proto=git.vodjk.com/go-grpc/example/proto/google/api:. ./*.proto

