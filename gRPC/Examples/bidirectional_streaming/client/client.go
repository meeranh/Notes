package main

import (
	"context"
	"log"

	pb "github.com/meeranh/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const HOST = "localhost:9001"

func checkErr(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func main() {
	conn, err := grpc.NewClient(HOST, grpc.WithTransportCredentials(insecure.NewCredentials()))
	checkErr(err)

	client := pb.NewHelloServiceClient(conn)
	stream, err := client.Hello(context.Background())
	checkErr(err)

	for {
		err := stream.Send(&pb.HelloRequest{Name: "A"})
		checkErr(err)
		res, err := stream.Recv()
		checkErr(err)
		log.Printf(res.NameEdited)
	}

}
