package main

import (
	pb "local.com/grpc-example/proto/token" // 引入proto包

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials" // 引入grpc认证包
	"google.golang.org/grpc/grpclog"

	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"google.golang.org/grpc/metadata"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"

	// OpenTLS 是否开启TLS认证
	OpenTLS = false
)

//token list
var (
	testPairs = []string{"authorization","bearer 1234","singlekey", "uno", "multikey", "one", "multikey", "two", "multikey", "three"}
	parentCtx = context.WithValue(context.TODO(), "parentKey", "parentValue")
)

// customCredential 自定义认证
type customCredential struct{}

// GetRequestMetadata 实现自定义认证接口
//func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
//	return map[string]string{
//		"appid":  "101010",
//		"appkey": "i am key",
//		"authorization": "bearer 123",
//	}, nil
//}

// RequireTransportSecurity 自定义认证是否开启TLS
func (c customCredential) RequireTransportSecurity() bool {
	return OpenTLS
}

func main() {
	var err error
	var opts []grpc.DialOption

	if OpenTLS {
		// TLS连接
		creds, err := credentials.NewClientTLSFromFile("../../keys/server.pem", "minicool")
		if err != nil {
			grpclog.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	// 使用自定义认证
	//opts = append(opts, grpc.WithPerRPCCredentials(new(customCredential)))

	conn, err := grpc.Dial(Address, opts...)

	if err != nil {
		grpclog.Fatalln(err)
	}

	defer conn.Close()

	// 初始化客户端
	c := pb.NewHelloClient(conn)

	//nmd := metautils.NiceMD(metadata.Pairs(testPairs...))
	nCtx := metautils.NiceMD(metadata.Pairs(testPairs...)).ToOutgoing(context.Background())

	// 调用方法
	req := &pb.HelloRequest{Name: "gRPC"}
	res, err := c.SayHello(/*context.Background()*/nCtx, req)
	if err != nil {
		grpclog.Fatalln(err)
	}

	grpclog.Println(res.Message)
}
