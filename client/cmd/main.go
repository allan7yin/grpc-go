package main

import (
	"context"
	"log"

	pb "github.com/allan7yin/grpc-go-client/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error: Failed to connect: %v", err)
	}
	// defer conn.Close()
	log.Printf("Connected to grpc server")
	client := pb.NewUserServiceClient(conn)

	// Let's test our grpc server and client by making a user, and then retreiving the created user
	createUserRequest := &pb.CreateUserRequest{
		Name:  "Allan Yin",          // Replace with the actual name
		Email: "a7yin@uwaterloo.ca", // Replace with the actual email
	}

	// NOTE: The fields from pb.something, is always capital first letter

	newUser, err := client.Create(context.Background(), createUserRequest)
	if err != nil {
		log.Printf("UserProfileResponse: %v\n", newUser)
		log.Fatalf("Error: Failed to create user, error: %v", err)
	} else {
		log.Printf("Created single user: %v", newUser)
	}

	// Note, this is how we are calling the `SingleUserRequest` method on the server (get user request)
	singleUserRequest := &pb.SingleUserRequest{
		Id: newUser.Id,
	}
	user, err := client.Read(context.Background(), singleUserRequest)
	if err != nil {
		log.Fatalf("Error: Failed to fetch user, error: %v", err)
	} else {
		log.Printf("Retrieved single user: %v", user)
	}

	// Now, make call to delete this user
	if response, err := client.Delete(context.Background(), singleUserRequest); err != nil {
		log.Fatalf("Error: Failed to delete user, error: %v", err)
	} else {
		log.Printf("Deleted single user, respone code: %v", response)
	}
}
