package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"net"

	pb "github.com/rashidkalwar/demo-grpc/check_prime"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCheckPrimeServer // Embeds unimplemented methods to maintain forward compatibility.
}

func (s *server) IsPrime(context context.Context, isPrimeRequest *pb.IsPrimeRequest) (*pb.IsPrimeResponse, error) {
	number := isPrimeRequest.Number
	primeNumberMessage := fmt.Sprintf("%d is a prime number", number)
	nonPrimeNumberMessage := fmt.Sprintf("%d is not a prime number", number)

	// Numbers <= 1 are not prime by definition.
	if number <= 1 {
		return &pb.IsPrimeResponse{
			Result: nonPrimeNumberMessage,
		}, nil
	}

	// Loop from 2 to sqrt(number) to check for divisors.
	// No need to check beyond sqrt(number) because if number is divisible by a number
	// larger than its square root, the corresponding divisor would be smaller than sqrt(number).
	for i := int64(2); i <= int64(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return &pb.IsPrimeResponse{
				Result: nonPrimeNumberMessage,
			}, nil
		}
	}

	// If no divisors were found, number is prime.
	return &pb.IsPrimeResponse{
		Result: primeNumberMessage,
	}, nil
}

func main() {
	// Start listening on port 50051 for incoming gRPC requests.
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// Register the CheckPrime service implementation with the gRPC server.
	pb.RegisterCheckPrimeServer(grpcServer, &server{})

	// Start serving requests. Blocks and listens for incoming RPCs.
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
