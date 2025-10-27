package main

import (
	"fmt"
	"log"
	"net"

	"github.com/yumaeda/grpc/internal/infrastructure"
	admin_user_pb "github.com/yumaeda/grpc/internal/proto/admin_user"
	area_pb "github.com/yumaeda/grpc/internal/proto/area"
	category_pb "github.com/yumaeda/grpc/internal/proto/category"
	menu_pb "github.com/yumaeda/grpc/internal/proto/menu"
	photo_pb "github.com/yumaeda/grpc/internal/proto/photo"
	restaurant_pb "github.com/yumaeda/grpc/internal/proto/restaurant"
	video_pb "github.com/yumaeda/grpc/internal/proto/video"
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

	// Photo service
	photoRepository := repository.NewPhotoRepository(db)
	photoService := service.NewPhotoService(photoRepository)
	photoServer := server.NewPhotoServer(photoService)
	photo_pb.RegisterPhotoServiceServer(grpcServer, photoServer)

	// Video service
	videoRepository := repository.NewVideoRepository(db)
	videoService := service.NewVideoService(videoRepository)
	videoServer := server.NewVideoServer(videoService)
	video_pb.RegisterVideoServiceServer(grpcServer, videoServer)

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
}
