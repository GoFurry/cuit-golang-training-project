package main

import (
	"context"
	"log"
	"net"
	"strings"
	"time"

	pb "example.com/cuit-go-frameworks/grpc-demo/proto"
	"google.golang.org/grpc"
)

type greeterServer struct {
	pb.UnimplementedGreeterServer
}

func (greeterServer) SayHello(_ context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	name := strings.TrimSpace(req.GetName())
	if name == "" {
		name = "gRPC learner"
	}

	return &pb.HelloReply{
		Message:    "你好, " + name + "! 欢迎来到 gRPC Demo。",
		ServerTime: time.Now().Unix(),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, greeterServer{})

	log.Println("gRPC server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
