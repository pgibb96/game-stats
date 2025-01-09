package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/pgibb96/game-stats/proto" // Import the generated proto package
	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedMyServiceServer // Embed the generated server interface
}

func (s *server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	// Implement the SayHello method from the service
	return &proto.HelloResponse{
		Message: fmt.Sprintf("Hello, %s!", req.GetName()),
	}, nil
}

func main() {
	// Start the gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterMyServiceServer(s, &server{}) // Register the service with the gRPC server

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
