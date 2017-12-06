#!/usr/bin/env bash

echo "====build TOKEN===="
PROTO_PATH="../../proto"
GRPC_EXAMPLE_TOKEN="../../example/token/"
BIN_EXAMPLE_TOKEN="../../bin/example/token/"
PROTO_EXAMPLE_TOKEN="hello.proto"

protoc --proto_path=${PROTO_PATH}/token \
--go_out=plugins=grpc:${PROTO_PATH}/token \
${PROTO_PATH}/token/${PROTO_EXAMPLE_TOKEN}

go build -o ${BIN_EXAMPLE_TOKEN}/client ${GRPC_EXAMPLE_TOKEN}/client/client.go
go build -o ${BIN_EXAMPLE_TOKEN}/server ${GRPC_EXAMPLE_TOKEN}/server/server.go