#!/usr/bin/env bash

echo "====build example===="
PROTO_PATH="../../proto"
GRPC_EXAMPLE_NORMAL="../../example/normal/"
BIN_EXAMPLE_NORMAL="../../bin/example/normal/"
PROTO_EXAMPLE_NORMAL="example.proto"

rm ${BIN_EXAMPLE_NORMAL}/*

protoc --proto_path=${PROTO_PATH}/example \
--go_out=plugins=grpc:${PROTO_PATH}/example \
${PROTO_PATH}/example/${PROTO_EXAMPLE_NORMAL}

go build -o ${BIN_EXAMPLE_NORMAL}/client ${GRPC_EXAMPLE_NORMAL}/client/client.go
go build -o ${BIN_EXAMPLE_NORMAL}/server ${GRPC_EXAMPLE_NORMAL}/server/server.go