package main

import (
	"context"
	"log"

	"github.com/claudioemmanuel/grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer connection.Close()

	client := pb.NewHelloServiceClient(connection)

	Hello(err, client)
}

func Hello(err error, client pb.HelloServiceClient) {
	if err != nil {
		log.Fatalf("failed to greet: %v", err)
	}

	response, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "World"})
	if err != nil {
		log.Fatalf("failed to greet: %v", err)
	}

	log.Printf("Greeting: %s", response.Message)
}
