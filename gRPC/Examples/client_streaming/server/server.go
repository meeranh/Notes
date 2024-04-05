package main

import (
	"io"
	"log"
	"net"

	pb "github.com/meeranh/protobuf/proto"
	"google.golang.org/grpc"
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
	concat := "Hello "

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			stream.SendAndClose(&pb.HelloResponse{Result: concat})
			break
		}
		checkErr(err)

		concat += res.Name
	}

	log.Println(concat)

	return nil
}

func main() {
	lis, err := net.Listen("tcp", HOST)
	checkErr(err)
	log.Printf("Listening on %v", HOST)

	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, &Server{})
	grpcServer.Serve(lis)
}
