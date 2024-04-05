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

	con, err := grpc.NewClient(HOST, grpc.WithTransportCredentials(insecure.NewCredentials()))
	checkErr(err)

	srvClient := pb.NewHelloServiceClient(con)
	client, err := srvClient.Hello(context.Background())

	data := []*pb.HelloRequest{
		{Name: "Meeran"},
		{Name: "Is"},
		{Name: "The"},
		{Name: "Name"},
	}

	for _, v := range data {
		print(v.Name)
		client.Send(&pb.HelloRequest{Name: v.Name})
	}

	res, err := client.CloseAndRecv()
	checkErr(err)

	log.Println(res.Result)

}
