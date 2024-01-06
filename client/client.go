package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	pb "github.com/BiradarSandeep8/prac/train"
	"google.golang.org/grpc"
)

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func main() {
	// Set up a connection to the server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client.
	client := pb.NewTrainServiceClient(conn)

	// Reading user input for PurchaseTicket request
	from := getUserInput("Enter 'From' location")
	to := getUserInput("Enter 'To' location: ")
	firstName := getUserInput("Enter first name: ")
	lastName := getUserInput("Enter last name: ")
	email := getUserInput("Enter email: ")
	section := getUserInput("Enter 'Section':")

	// Make a PurchaseTicket request
	response, err := client.PurchaseTicket(context.Background(), &pb.Ticket{
		To: to, From: from,
		Section: section,
		User: &pb.User{
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
		},
		PricePaid: 20.0,
	})
	if err != nil {
		log.Fatalf("train ticket purchase request failed: %v", err)
	}

	// Printing all details of the response.
	fmt.Println("train ticket purchaising request succeeded.")
	fmt.Printf("Server Response:\n")
	fmt.Printf("From: %s\n", response.From)
	fmt.Printf("To: %s\n", response.To)
	fmt.Printf("User: %+v\n", response.User)
	fmt.Printf("PricePaid: %f\n", response.PricePaid)
	fmt.Printf("Section: %s\n", response.Section)
}
