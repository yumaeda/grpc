package main

import (
	"fmt"
	"log"
	"net"

	"github.com/yumaeda/grpc/internal/infrastructure"
	area_proto "github.com/yumaeda/grpc/internal/proto"
	"github.com/yumaeda/grpc/internal/repository"
	"github.com/yumaeda/grpc/internal/server"
	"github.com/yumaeda/grpc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, dbCloser, dbErr := infrastructure.ConnectToDB()
	if dbErr != nil {
		panic(dbErr.Error())
	}
	defer dbCloser()

	areaRepository := repository.NewAreaRepository(db)
	areaService := service.NewAreaService(areaRepository)

	grpcServer := grpc.NewServer()
	areaServer := server.NewAreaServer(areaService)
	area_proto.RegisterAreaServiceServer(grpcServer, areaServer)

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
