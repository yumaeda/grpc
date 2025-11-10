package server

import (
	"context"

	"github.com/yumaeda/sakabas-grpc/internal/service"
	pb "github.com/yumaeda/sakabas-grpc/proto/drink"
)

type DrinkServer struct {
	pb.UnimplementedDrinkServiceServer
	drinkService service.DrinkService
}

func NewDrinkServer(drinkService service.DrinkService) *DrinkServer {
	return &DrinkServer{drinkService: drinkService}
}

func (s *DrinkServer) GetDrink(ctx context.Context, req *pb.GetDrinkRequest) (*pb.GetDrinkResponse, error) {
	drink, err := s.drinkService.GetDrink(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetDrinkResponse{
		Drink: &pb.Drink{
			Id:   drink.ID,
			Name: drink.Name,
		},
	}, nil
}
