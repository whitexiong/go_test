package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

//导入我们在protos文件中定义的服务
import pb "github.com/whitexiong/go_grpc/protos/helloworld"

// 定义一个结构体，作用是实现helloworld中的GreeterServer
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

// 定义端口号 支持启动的时候输入端口号
var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	//解析输入的端口号 默认50051
	flag.Parse()
	//tcp协议监听指定端口号
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//实例化gRPC服务
	s := grpc.NewServer()
	//服务注册
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	//启动服务
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
