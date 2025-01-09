package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/pgibb96/game_stats/proto"
	"google.golang.org/grpc"
)

// Define the server that implements the gRPC service
type server struct {
	pb.UnimplementedMyServiceServer
}

// Implement the SayHello RPC method
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: fmt.Sprintf("Hello, %s!", req.Name)}, nil
}

func main() {
	// Create a listener on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register the service with the server
	pb.RegisterMyServiceServer(grpcServer, &server{})

	// Start the server
	log.Println("gRPC server listening on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
