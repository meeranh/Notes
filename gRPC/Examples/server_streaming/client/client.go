package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/meeranh/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const HOST = "localhost:9001"

func checkErr(err error) {
	if err != nil {
		log.Printf("There was an error reading the file %v", err)
	}
}

func main() {
	conn, err := grpc.NewClient(HOST, grpc.WithTransportCredentials(insecure.NewCredentials()))
	checkErr(err)
	client := pb.NewHelloServiceClient(conn)
	stream, err := client.Hello(context.Background(), &pb.HelloRequest{Name: "Meeran"})
	checkErr(err)

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		checkErr(err)
		fmt.Println("Message:", msg)
	}
}
