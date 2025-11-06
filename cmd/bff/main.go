package main

import (
	"fmt"
	"log"
	"net"

	"github.com/yumaeda/sakabas-grpc/internal/infrastructure"
	"github.com/yumaeda/sakabas-grpc/internal/repository"
	"github.com/yumaeda/sakabas-grpc/internal/server"
	"github.com/yumaeda/sakabas-grpc/internal/service"
	admin_user_pb "github.com/yumaeda/sakabas-grpc/proto/admin_user"
	area_pb "github.com/yumaeda/sakabas-grpc/proto/area"
	category_pb "github.com/yumaeda/sakabas-grpc/proto/category"
	menu_pb "github.com/yumaeda/sakabas-grpc/proto/menu"
	photo_pb "github.com/yumaeda/sakabas-grpc/proto/photo"
	ranking_pb "github.com/yumaeda/sakabas-grpc/proto/ranking"
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

	listener, err := net.Listen("tcp", ":50050")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer listener.Close()

	fmt.Println("gRPC server listening on :50050")
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

	// Menu service
	menuRepository := repository.NewMenuRepository(db)
	menuService := service.NewMenuService(menuRepository)
	menuServer := server.NewMenuServer(menuService)
	menu_pb.RegisterMenuServiceServer(grpcServer, menuServer)

	// Photo service
	photoRepository := repository.NewPhotoRepository(db)
	photoService := service.NewPhotoService(photoRepository)
	photoServer := server.NewPhotoServer(photoService)
	photo_pb.RegisterPhotoServiceServer(grpcServer, photoServer)

	// AdminUser service
	adminUserRepository := repository.NewAdminUserRepository(db)
	adminUserService := service.NewAdminUserService(adminUserRepository)
	adminUserServer := server.NewAdminUserServer(adminUserService)
	admin_user_pb.RegisterAdminUserServiceServer(grpcServer, adminUserServer)

	// Category service
	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryServer := server.NewCategoryServer(categoryService)
	category_pb.RegisterCategoryServiceServer(grpcServer, categoryServer)

	// Ranking service
	rankingRepository := repository.NewRankingRepository(db)
	rankingService := service.NewRankingService(rankingRepository)
	rankingServer := server.NewRankingServer(rankingService)
	ranking_pb.RegisterRankingServiceServer(grpcServer, rankingServer)
}
