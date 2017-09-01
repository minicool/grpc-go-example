package main

import (
	pb "local.com/grpc-example/proto/tls" // 引入proto包

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials" // 引入grpc认证包
	"google.golang.org/grpc/grpclog"
	"flag"
	"fmt"
	"github.com/golang/glog"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

func main() {
	//Init the command-line flags.
	flag.Parse()

	// Will be ignored as the program has exited in Fatal().
	defer func() {
		fmt.Println("Message in defer")
	}()

	// Flushes all pending log I/O.
	defer glog.Flush()

	// TLS连接
	creds, err := credentials.NewClientTLSFromFile("../../key/server.pem", "minicool")
	if err != nil {
		grpclog.Fatalf("Failed to create TLS credentials %v", err)
	}

	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(creds))
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()

	// 初始化客户端
	c := pb.NewHelloClient(conn)

	// 调用方法
	req := &pb.HelloRequest{Name: "gRPC"}
	res, err := c.SayHello(context.Background(), req)
	if err != nil {
		grpclog.Fatalln(err)
	}

	grpclog.Fatal(res.Message)
}
