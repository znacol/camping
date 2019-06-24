package main

import (
	"net"

	pb "github.com/znacol/camping/backend/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

// Start service
func main() {
	listener, err := net.Listen("tcp", ":5300")
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterCampingServiceServer(grpcServer, &server{})

	grpcServer.Serve(listener)
}

type server struct{}

func (s *server) Do(c context.Context, request *pb.Request) (response *pb.Response, err error) {
	response = &pb.Response{
		Message: "Hello world",
	}

	return response, nil
}
