package main

import (
	"context"
	"log"
	"time"

	pb "example.com/cuit-go-frameworks/grpc-demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Go learner"})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("message: %s", resp.GetMessage())
	log.Printf("server_time: %d", resp.GetServerTime())
}
