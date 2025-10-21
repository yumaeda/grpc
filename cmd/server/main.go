package main

import (
	"fmt"
	"log"
	"net"

	"github.com/yumaeda/grpc/internal/infrastructure"
	area_pb "github.com/yumaeda/grpc/internal/proto/area"
	menu_pb "github.com/yumaeda/grpc/internal/proto/menu"
	restaurant_pb "github.com/yumaeda/grpc/internal/proto/restaurant"
	"github.com/yumaeda/grpc/internal/repository"
	"github.com/yumaeda/grpc/internal/server"
	"github.com/yumaeda/grpc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

func main() {
	db, dbCloser, dbErr := infrastructure.ConnectToDB()
	if dbErr != nil {
		panic(dbErr.Error())
	}
	defer dbCloser()

	grpcServer := grpc.NewServer()
	registerServices(grpcServer, db)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer listener.Close()

	fmt.Println("gRPC server listening on :50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func registerServices(grpcServer *grpc.Server, db *gorm.DB) {
	// Area service
	areaRepository := repository.NewAreaRepository(db)
	areaService := service.NewAreaService(areaRepository)
	areaServer := server.NewAreaServer(areaService)
	area_pb.RegisterAreaServiceServer(grpcServer, areaServer)

	// Restaurant service
	restaurantRepository := repository.NewRestaurantRepository(db)
	restaurantService := service.NewRestaurantService(restaurantRepository)
	restaurantServer := server.NewRestaurantServer(restaurantService)
	restaurant_pb.RegisterRestaurantServiceServer(grpcServer, restaurantServer)

	// Menu service
	menuRepository := repository.NewMenuRepository(db)
	menuService := service.NewMenuService(menuRepository)
	menuServer := server.NewMenuServer(menuService)
	menu_pb.RegisterMenuServiceServer(grpcServer, menuServer)
}
