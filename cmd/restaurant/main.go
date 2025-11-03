package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/yumaeda/grpc/internal/infrastructure"
	"github.com/yumaeda/grpc/internal/repository"
	"github.com/yumaeda/grpc/internal/server"
	"github.com/yumaeda/grpc/internal/service"

	pb "github.com/yumaeda/grpc/swapi/restaurant/restaurant"
)

func main() {
	db, dbCloser, dbErr := infrastructure.ConnectToDB()
	if dbErr != nil {
		panic(dbErr.Error())
	}
	defer dbCloser()

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer listener.Close()

	grpcServer := grpc.NewServer()

	restaurantRepository := repository.NewRestaurantRepository(db)
	restaurantService := service.NewRestaurantService(restaurantRepository)
	restaurantServer := server.NewRestaurantServer(restaurantService)

	pb.RegisterRestaurantServiceServer(grpcServer, restaurantServer)
	reflection.Register(grpcServer)

	log.Println("gRPC server listening on :50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
