package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/meeranh/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const SERVER = "localhost:50051"

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalf(msg, err)
	}
}

func main() {
	// Ask the user for user input
	fmt.Println("What is your name?")
	fmt.Printf("Name: ")
	var name string
	fmt.Scanln(&name)

	// Initialize a gRPC connection
	conn, err := grpc.Dial(SERVER, grpc.WithTransportCredentials(insecure.NewCredentials()))
	checkErr(err, "Failed to connect to server: %v")

	// Creating a context with 1 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Create a client for our gRPC server
	client := pb.NewHelloServiceClient(conn)
	res, err := client.Hello(ctx, &pb.HelloRequest{MyName: name})
	checkErr(err, "Failed to call Hello: %v")

	// Print the response
	fmt.Println(res.Greeting)
}
