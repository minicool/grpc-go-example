package main

import (
	"net"

	pb "github.com/minicool/grpc-go-example/proto/example"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"encoding/json"
	"io/ioutil"
	"strconv"
	"github.com/golang/glog"
	"os"
	"flag"
	"sync"
	"io"
	"path/filepath"
	"strings"
)

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func getParentDirectory(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		glog.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

type Config struct{
	Host string
	Port int
}

var json_config = map[string]Config{}

func readFile(filename string) (map[string]Config, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		grpclog.Error("ReadFile: ", err.Error())
		return nil, err
	}

	if err := json.Unmarshal(bytes, &json_config); err != nil {
		grpclog.Error("Unmarshal: ", err.Error())
		return json_config, err
	}

	return json_config, nil
}

type Emplyer struct{
	EmployerId int32
	Age int32
	Name string
	PicName string
}

//type EmplyerList map[string][]Emplyer
//var json_data = map[string][]Emplyer{}

type EmplyerList map[string][]*pb.Employer
var json_data = map[string][]*pb.Employer{}

func readData(filename string) (EmplyerList, error) {
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

// 定义helloService并实现约定的接口
type employerService struct{
	employers []*pb.Employer
	m      sync.RWMutex
	//emplyer data
	address string
}

// HelloService Hello服务
var EmployerService = employerService{}

func newServer() *employerService {
	s := new(employerService)
//	s.loadFeatures(*jsonDBFile)
//	s.employers = make(map[uint64][]*pb.Employer)
//	s.employers = make([]*pb.Employer,4)
	s.loadConfig()
	return s
}

//load config
func (service *employerService)loadConfig() bool {
	//data read file
	data,err := readData("../../data/employer.json")
	if err != nil {
		glog.Error("employer.json readFile: ", err.Error())
		return  false
	}

	service.employers = data["employees"]

	//config read file
	config, err := readFile("../../config/config.json")
	if err != nil {
		glog.Error("config readFile: ", err.Error())
		return false
	}
	address_port := strconv.Itoa(config["client_conf"].Port)
	service.address = config["client_conf"].Host + ":" + address_port

	return true
}

// SayHello 实现Hello服务接口
func (service *employerService) GetEmployer(ctx context.Context, in *pb.EmployerRequest) (*pb.EmployerResponse, error) {
	glog.Errorf("EmployerRequest %s",in)
	//Employer := &pb.Employer{
	//	EmployerId:2,
	//	Age:18,
	//	Name:"test",
	//	PicName:"1.jpeg",
	//}

	var resp = &pb.EmployerResponse{}
	for _,value := range service.employers{
		if value.EmployerId == in.EmployerId {
			glog.Errorf("found %d %s",value.EmployerId,value)
			resp = &pb.EmployerResponse{value}
		}
	}

	return resp, nil
}

func (service *employerService) GetEmployerList(ctx context.Context, in *pb.EmployerListRequest) (*pb.EmployerListResponse, error) {
	glog.Errorf("EmployerListRequest %s",in)

	var resp *pb.EmployerListResponse
	if uint32(len(service.employers)) < (in.EmployerIndex + in.EmployerCount) {
		resp = &pb.EmployerListResponse{}
		return resp,nil
	}

	resp = &pb.EmployerListResponse{service.employers[in.EmployerIndex:in.EmployerCount]}

/*	resp := &pb.EmployerListResponse{}
	resp.Employers = make([]*pb.Employer,2)
	resp.Employers[0] = &pb.Employer{
		EmployerId:0,
		Age:18,
		Name:"test",
		PicName:"1.jpeg",
	}
	resp.Employers[1] = &pb.Employer{
		EmployerId:1,
		Age:12,
		Name:"test02",
		PicName:"12.jpeg",
	}*/

	return resp, nil
}

func (service *employerService) GetEmployerMap(ctx context.Context, in *pb.EmployerMapRequest) (*pb.EmployerMapResponse, error){
	glog.Errorf("GetEmployerMap %s",in)

	var data map[string]*pb.Employer
	data = make(map[string]*pb.Employer)

	for _,value := range service.employers{
		data[value.Name] = value
	}
	resp := &pb.EmployerMapResponse{data}
	return resp, nil
}

func (service *employerService)GetEmployerAll(in *pb.EmployerAllRequest,stream pb.EmployerService_GetEmployerAllServer) error{
	glog.Errorf("GetEmployerAll %s",in)
	//	for _,value := range service.employers{
		if err := stream.Send( &pb.EmployerAllResponse{service.employers}); err != nil {
			return err
		}
	//	}
	return nil
}

func (service *employerService)AddEmployerImage(stream pb.EmployerService_AddEmployerImageServer) error{
	glog.Error("AddEmployerImage")
	var fp *os.File
//	startTime := time.Now()
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			fp.Close()
//			endTime := time.Now()
			return stream.SendAndClose(&pb.EmployerImageResponse{(true)})
		}
		if request.PicName != ""{
			fp,err = os.OpenFile(request.PicName, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				glog.Fatal(err)
			}
			glog.Errorf("employerId %d filename %s",request.EmployerId,request.PicName)
		}
		fp.Write(request.PicData)
		glog.Error("WriteFile")
		if err != nil {
			return err
		}
	}
	return nil
}

func (service *employerService)GetEmployerImage(in *pb.GetEmployerImageRequest,stream pb.EmployerService_GetEmployerImageServer) error{
	glog.Error("GetEmployerImage %d",in.EmployerId)
	var filename string
	for _,value := range service.employers{
		if value.EmployerId == in.EmployerId {
			glog.Errorf("found %d %s",value.EmployerId,value)
			filename = value.PicName
		}
	}
	fp,err := os.Open(getCurrentDirectory()+"/../../image/"+filename)
	glog.Error(getCurrentDirectory()+"/../../image/"+filename)
	if err != nil {
		glog.Error(err)
	}

	var fileData []byte
	//use 10240
	fileData = make([]byte,10240)
	fp.Read(fileData)

	if err := stream.Send( &pb.GetEmployerImageResponse{filename,fileData}); err != nil {
		return err
	}
	return nil
}
//func (h employerServiceServer)GetEmployerAllImage(ctx context.Context) (pb.EmployerService_GetEmployerAllImageClient, error){
//	resp := new(pb.EmployerService_GetEmployerAllImageClient)
//	return resp, nil
//}

func main() {
	//Init the command-line flags.
	flag.Parse()

	// Flushes all pending log I/O.
	defer glog.Flush()


	//config read file
	config, err := readFile("../../config/config.json")
	if err != nil {
		glog.Error("config readFile: ", err.Error())
	}
	address_port := strconv.Itoa(config["client_conf"].Port)
	address := config["client_conf"].Host + ":" + address_port

	listen, err := net.Listen("tcp", address)
	if err != nil {
		glog.Fatalf("Failed to listen: %v", err)
	}

	// TLS认证
	//creds, err := credentials.NewServerTLSFromFile("../../key/server.pem", "../../key/server.key")
	//if err != nil {
	//	grpclog.Fatalf("Failed to generate credentials %v", err)
	//}

	// 实例化grpc Server, 并开启TLS认证
	s := grpc.NewServer(/*grpc.Creds(creds)*/)
	// 注册HelloService
	pb.RegisterEmployerServiceServer(s, newServer())


	glog.Errorf("Listen on " + address + " with TLS")
	glog.Infof("Temp folder for log files: %s\n", os.TempDir())
//	grpclog.Infoln("Listen on " + address + " with TLS")

	s.Serve(listen)
}