package main

import (
	"context"
	"fmt"

	pb "github.com/znacol/camping/backend/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	fmt.Printf("in main\n")
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial("localhost:30251", opts...)
	if err != nil {
		grpclog.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewCampingServiceClient(conn)
	request := &pb.GetAllSitesRequest{}

	fmt.Printf("request: %+v\n", request)

	response, err := client.GetAllSites(context.Background(), request)
	if err != nil {
		grpclog.Fatalf("failed to call endpoint: %v", err)
	}

	fmt.Println("response: %+v", response)
}
