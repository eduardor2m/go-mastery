package main

import (
	"context"
	"log"

	"github.com/eduardor2m/go-mastery/src/grpc/hello_grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	req := &pb.HelloRequest{Name: "World"}

	client := pb.NewHelloClient(conn)

	res, err := client.SayHello(context.Background(), req)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.GetMessage())

}
