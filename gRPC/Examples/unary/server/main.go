package main

import (
	"context"
	"log"
	"net"

	pb "github.com/meeranh/protobuf/proto"
	"google.golang.org/grpc"
)

const HOST = "0.0.0.0:50051"

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalf(msg, err)
	}
}

type Server struct {
	pb.UnimplementedHelloServiceServer
}

func (s *Server) Hello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.MyName)
	return &pb.HelloResponse{
		Greeting: "Hello " + in.MyName,
	}, nil
}

func main() {
	// Setup TCP Listener
	lis, err := net.Listen("tcp", HOST)
	checkErr(err, "Failed to listen: %v")
	log.Printf("Server listening on %s", HOST)

	// Setup the gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, &Server{})

	// Serve the connection over TCP
	srvErr := grpcServer.Serve(lis)
	checkErr(srvErr, "Failed to serve: %v")
}
