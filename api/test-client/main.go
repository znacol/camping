package main

import (
	"context"
	"fmt"
	"os"

	pb "github.com/znacol/camping/api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	fmt.Printf("in main\n")
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	args := os.Args
	conn, err := grpc.Dial("localhost:30251", opts...)
	if err != nil {
		grpclog.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewCampingServiceClient(conn)
	request := &pb.Request{
		Message: args[1],
	}

	fmt.Printf("request: %+v\n", request)

	response, err := client.Do(context.Background(), request)
	if err != nil {
		grpclog.Fatalf("failed to call endpoint: %v", err)
	}

	fmt.Println(response.Message)
}
