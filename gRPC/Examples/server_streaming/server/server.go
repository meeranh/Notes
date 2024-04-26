package main

import (
	"fmt"
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

func (s *Server) Hello(in *pb.HelloRequest, stream pb.HelloService_HelloServer) error {
	log.Printf("Connection received from %s", in.Name)
	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s: %d", in.Name, i)
		stream.Send(&pb.HelloResponse{
			Result: res,
		})
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", HOST)
	checkErr(err)
	log.Printf("Server listening on %s", HOST)

	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, &Server{})
	srvErr := grpcServer.Serve(lis)
	checkErr(srvErr)
}
