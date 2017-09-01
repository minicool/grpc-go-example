#!/bin/sh

echo "-----------------------------build go proto------------------------------"
PROTO_PATH="../../proto/"

function func_proto_path(){
    return $(${PROTO_PATH}${1}/${1}.proto)
}

echo "$(func_proto_path tls)"

echo "====build example===="
GRPC_EXAMPLE_NORMAL="../../example/normal/"
BIN_EXAMPLE_NORMAL="../../bin/example/normal/"
PROTO_EXAMPLE_NORMAL="example.proto"

rm ${BIN_EXAMPLE_NORMAL}/*
protoc --proto_path=${PROTO_PATH}/example --go_out=plugins=grpc:${PROTO_PATH}/example ${PROTO_PATH}/example/${PROTO_EXAMPLE_NORMAL}
go build -o ${BIN_EXAMPLE_NORMAL}/client ${GRPC_EXAMPLE_NORMAL}/client/client.go
go build -o ${BIN_EXAMPLE_NORMAL}/server ${GRPC_EXAMPLE_NORMAL}/server/server.go

#echo "====build TLS===="
#GRPC_EXAMPLE_TLS="../../example/tls/"
#BIN_EXAMPLE_TLS="../../bin/example/tls/"
#PROTO_EXAMPLE_TLS="hello.proto"
#
#rm ${BIN_EXAMPLE_TLS}/*
#protoc --proto_path=${PROTO_PATH}/tls --go_out=plugins=grpc:${PROTO_PATH}/tls ${PROTO_PATH}/tls/${PROTO_EXAMPLE_TLS}
#go build -o ${BIN_EXAMPLE_TLS}/client ${GRPC_EXAMPLE_TLS}/client/client.go
#go build -o ${BIN_EXAMPLE_TLS}/server ${GRPC_EXAMPLE_TLS}/server/server.go
#
#echo "====build TOKEN===="
#GRPC_EXAMPLE_TOKEN="../../example/token/"
#BIN_EXAMPLE_TOKEN="../../bin/example/token/"
#PROTO_EXAMPLE_TOKEN="hello.proto"
#
#protoc --proto_path=${PROTO_PATH}/token --go_out=plugins=grpc:${PROTO_PATH}/token ${PROTO_PATH}/token/${PROTO_EXAMPLE_TOKEN}
#go build -o ${BIN_EXAMPLE_TOKEN}/client ${GRPC_EXAMPLE_TOKEN}/client/client.go
#go build -o ${BIN_EXAMPLE_TOKEN}/server ${GRPC_EXAMPLE_TOKEN}/server/server.go
#
#echo "====build interceptor===="
#GRPC_EXAMPLE_INTER="../../example/interceptor/"
#BIN_EXAMPLE_INTER="../../bin/example/interceptor/"
#PROTO_EXAMPLE_INTER="hello.proto"
#
#protoc --proto_path=${PROTO_PATH}/interceptor --go_out=plugins=grpc:${PROTO_PATH}/interceptor ${PROTO_PATH}/interceptor/${PROTO_EXAMPLE_INTER}
#go build -o ${BIN_EXAMPLE_INTER}/client ${GRPC_EXAMPLE_INTER}/client/client.go
#go build -o ${BIN_EXAMPLE_INTER}/server ${GRPC_EXAMPLE_INTER}/server/server.go