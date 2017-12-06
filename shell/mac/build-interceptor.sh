#!/usr/bin/env bash

echo "====build interceptor===="
PROTO_PATH="../../proto"
GRPC_EXAMPLE_INTER="../../example/interceptor/"
BIN_EXAMPLE_INTER="../../bin/example/interceptor/"
PROTO_EXAMPLE_INTER="hello.proto"

protoc --proto_path=${PROTO_PATH}/interceptor \
--go_out=plugins=grpc:${PROTO_PATH}/interceptor \
${PROTO_PATH}/interceptor/${PROTO_EXAMPLE_INTER}

go build -o ${BIN_EXAMPLE_INTER}/client ${GRPC_EXAMPLE_INTER}/client/client.go
go build -o ${BIN_EXAMPLE_INTER}/server ${GRPC_EXAMPLE_INTER}/server/server.go
