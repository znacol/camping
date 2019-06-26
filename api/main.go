package main

import (
	"fmt"
	"net"

	pb "github.com/znacol/camping/api/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

// Start service
func main() {
	// log := slf.WithContext("camping-service")
	listener, err := net.Listen("tcp", ":30251")
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterCampingServiceServer(grpcServer, &server{})

	fmt.Print("started service")

	grpcServer.Serve(listener)
}

type server struct{}

func (s *server) Do(c context.Context, request *pb.Request) (response *pb.Response, err error) {
	response = &pb.Response{
		Message: "Hello world",
	}

	return response, nil
}
