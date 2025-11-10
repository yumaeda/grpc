package server

import (
	"context"

	"github.com/yumaeda/sakabas-grpc/internal/service"
	pb "github.com/yumaeda/sakabas-grpc/proto/dish"
)

type DishServer struct {
	pb.UnimplementedDishServiceServer
	dishService service.DishService
}

func NewDishServer(dishService service.DishService) *DishServer {
	return &DishServer{dishService: dishService}
}

func (s *DishServer) GetDish(ctx context.Context, req *pb.GetDishRequest) (*pb.GetDishResponse, error) {
	dish, err := s.dishService.GetDish(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetDishResponse{
		Dish: &pb.Dish{
			Id:   dish.ID,
			Name: dish.Name,
		},
	}, nil
}
