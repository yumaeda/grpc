package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/yumaeda/sakabas-grpc/internal/infrastructure"
	"github.com/yumaeda/sakabas-grpc/internal/repository"
	"github.com/yumaeda/sakabas-grpc/internal/server"
	"github.com/yumaeda/sakabas-grpc/internal/service"

	pb "github.com/yumaeda/sakabas-grpc/swapi/video/video"
)

func main() {
	db, dbCloser, dbErr := infrastructure.ConnectToDB()
	if dbErr != nil {
		panic(dbErr.Error())
	}
	defer dbCloser()

	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer listener.Close()

	grpcServer := grpc.NewServer()

	videoRepository := repository.NewVideoRepository(db)
	videoService := service.NewVideoService(videoRepository)
	videoServer := server.NewVideoServer(videoService)

	pb.RegisterVideoServiceServer(grpcServer, videoServer)
	reflection.Register(grpcServer)

	log.Println("gRPC server listening on :50052")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
