package main

import (
	"context"
	"log"
	"time"

	pb "github.com/rashidkalwar/demo-grpc/check_prime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Establish a connection to the gRPC server on localhost:50051 using insecure credentials.
	// Do not use insecure credentials in the production.
	connection, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed to connect to gRPC server")
	}
	defer connection.Close()

	client := pb.NewCheckPrimeClient(connection)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	output, err := client.IsPrime(ctx, &pb.IsPrimeRequest{
		Number: int64(7),
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v", output)
}
