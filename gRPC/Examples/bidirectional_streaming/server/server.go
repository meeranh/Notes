package main

import (
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "github.com/meeranh/protobuf/proto"
)

const HOST = "0.0.0.0:9001"

func checkErr(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

type Server struct {
	pb.UnimplementedHelloServiceServer
}

func (s *Server) Hello(stream pb.HelloService_HelloServer) error {
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		checkErr(err)

		log.Printf("Request received: %v", res.Name)
		stream.Send(&pb.HelloResponse{NameEdited: res.Name})
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", HOST)
	checkErr(err)
	log.Printf("Server started at %s", HOST)

	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, &Server{})
	grpcServer.Serve(lis)
}
