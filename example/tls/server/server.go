package main

import (
	"fmt"
	"net"

	pb "local.com/grpc-example/proto/tls"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials" // 引入grpc认证包
	"google.golang.org/grpc/grpclog"
	"github.com/golang/glog"
	"flag"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

// 定义helloService并实现约定的接口
type helloService struct{}

// HelloService Hello服务
var HelloService = helloService{}

// SayHello 实现Hello服务接口
func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	resp := new(pb.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s.", in.Name)
	glog.Errorf("hello  %s.",in.Name)
	return resp, nil
}

func main() {
	//Init the command-line flags.
	flag.Parse()

	// Will be ignored as the program has exited in Fatal().
	defer func() {
		fmt.Println("Message in defer")
	}()

	// Flushes all pending log I/O.
	defer glog.Flush()

	listen, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
	}

	// TLS认证
	creds, err := credentials.NewServerTLSFromFile("../../key/server.pem", "../../key/server.key")
	if err != nil {
		glog.Errorf("Failed to generate credentials %v", err)
	}

	// 实例化grpc Server, 并开启TLS认证
	s := grpc.NewServer(grpc.Creds(creds))

	// 注册HelloService
	pb.RegisterHelloServer(s, HelloService)

	glog.Error("Listen on " + Address + " with TLS")

	s.Serve(listen)
}