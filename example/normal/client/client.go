package main

import (
	pb "github.com/minicool/grpc-go-example/proto/example"

	"io/ioutil"
	"encoding/json"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"strconv"
	"context"
	"github.com/golang/glog"
	"flag"
	"fmt"
)

type Emplyer struct{
	Host string
	Port int
}

type EmplyerList struct{
	emplyers []Emplyer
}

var json_data = map[string]Emplyer{}

func readFile(filename string) (map[string]Emplyer, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		grpclog.Error("ReadFile: ", err.Error())
		return nil, err
	}

	if err := json.Unmarshal(bytes, &json_data); err != nil {
		grpclog.Error("Unmarshal: ", err.Error())
		return json_data, err
	}

	return json_data, nil
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

	config, err := readFile("../../config/config.json")
	if err != nil {
		glog.Error("readFile: ", err.Error())
	}

	//grpclog.Fatal(config)
	//fmt.Println(xxxMap)

	// TLS连接
	//creds, err := credentials.NewClientTLSFromFile("../../key/server.pem", "minicool")
	//if err != nil {
	//	glog.Errorf("Failed to create TLS credentials %v", err)
	//}

	address_port := strconv.Itoa(config["client_conf"].Port)
	address := config["client_conf"].Host + ":" + address_port
	glog.Errorf("connnet to address %s",address)

	conn, err := grpc.Dial(address, grpc.WithInsecure()/*grpc.WithTransportCredentials(creds)*/)
	if err != nil {
		glog.Errorln("Faild to Dial grpc",err)
	}
	defer conn.Close()

	// 初始化客户端
	c := pb.NewEmployerServiceClient(conn)

	// 调用方法
	req := &pb.EmployerRequest{EmployerId: 1}
	res, err := c.GetEmployer(context.Background(), req)
	if err != nil {
		glog.Errorln(err)
	}
	glog.Errorln(res.Employer.GetPicName())
}
