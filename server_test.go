package main

import (
	"context"
	"log"
	"net"
	"testing"

	pb "github.com/BiradarSandeep8/prac/train"
	"google.golang.org/grpc"
)

type testServer struct {
	pb.UnimplementedTrainServiceServer // Embed the UnimplementedTrainServiceServer type

	tickets map[string]*pb.Ticket
}

// Implement the PurchaseTicket method for the testServer
func (s *testServer) PurchaseTicket(ctx context.Context, req *pb.Ticket) (*pb.Ticket, error) {
	// Your test implementation...
	return &pb.Ticket{}, nil
}

// Helper function to set up the test environment
func setupTestServer(t *testing.T) (*grpc.Server, pb.TrainServiceClient, func()) {
	server := grpc.NewServer()
	testServer := &testServer{
		UnimplementedTrainServiceServer: pb.UnimplementedTrainServiceServer{},
		tickets:                         make(map[string]*pb.Ticket),
	}
	pb.RegisterTrainServiceServer(server, testServer)

	// Start the server in a goroutine
	go func() {
		if err := server.Serve(createTestListener()); err != nil {
			t.Fatal("Failed to start server:", err)
		}
	}()

	// Create a gRPC client
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		t.Fatal("Failed to connect to server:", err)
	}

	client := pb.NewTrainServiceClient(conn)

	// Return a cleanup function to stop the server and close the connection
	cleanup := func() {
		server.Stop()
		conn.Close()
	}

	return server, client, cleanup
}

// Helper function to create a test listener
func createTestListener() *net.TCPListener {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal("Failed to create test listener:", err)
	}
	return lis.(*net.TCPListener)
}

// Unit tests for the PurchaseTicket method
func TestPurchaseTicket(t *testing.T) {
	_, client, cleanup := setupTestServer(t)
	defer cleanup()

	// Make a PurchaseTicket request in the test
	response, err := client.PurchaseTicket(context.Background(), &pb.Ticket{
		From:      "TestFrom",
		To:        "TestTo",
		User:      &pb.User{FirstName: "Test", LastName: "User", Email: "test@example.com"},
		PricePaid: 25.0,
	})
	if err != nil {
		t.Fatal("PurchaseTicket request failed:", err)
	}

	// Assert the expected results
	if response.From != "TestFrom" {
		t.Errorf("Expected 'From' to be 'TestFrom', got '%s'", response.From)
	}
	// Add more assertions based on your actual implementation and expectations
}
