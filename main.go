package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	pb "github.com/BiradarSandeep8/prac/train"
	"google.golang.org/grpc"
)

type trainServer struct {
	pb.UnimplementedTrainServiceServer // Embed the UnimplementedTrainServiceServer type

	tickets map[string]*pb.Ticket
	mu      sync.Mutex
}

func (s *trainServer) PurchaseTicket(ctx context.Context, req *pb.Ticket) (*pb.Ticket, error) {
	// Your implementation...
	return req, nil
}

// Rest of your trainServer implementation...

func main() {
	server := grpc.NewServer()
	pb.RegisterTrainServiceServer(server, &trainServer{
		UnimplementedTrainServiceServer: pb.UnimplementedTrainServiceServer{},
		tickets:                         make(map[string]*pb.Ticket),
	})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Println("Server is running on port 50051...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
