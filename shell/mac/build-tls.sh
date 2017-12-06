#!/usr/bin/env bash

echo "====build TLS===="
PROTO_PATH="../../proto"
GRPC_EXAMPLE_TLS="../../example/tls/"
BIN_EXAMPLE_TLS="../../bin/example/tls/"
PROTO_EXAMPLE_TLS="hello.proto"

rm ${BIN_EXAMPLE_TLS}/*

protoc --proto_path=${PROTO_PATH}/tls \
--go_out=plugins=grpc:${PROTO_PATH}/tls \
${PROTO_PATH}/tls/${PROTO_EXAMPLE_TLS}

go build -o ${BIN_EXAMPLE_TLS}/client ${GRPC_EXAMPLE_TLS}/client/client.go
go build -o ${BIN_EXAMPLE_TLS}/server ${GRPC_EXAMPLE_TLS}/server/server.go