package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/eduardor2m/go-mastery/src/grpc/hello_grpc/pb"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedHelloServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello " + in.GetName()}, nil
}

func main() {
	fmt.Println("Runnig gRPC Server...")

	listener, err := net.Listen("tcp", "localhost:9000")

	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	pb.RegisterHelloServer(server, &Server{})

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
