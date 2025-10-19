package server

import (
	"context"

	pb "github.com/yumaeda/grpc/internal/proto/restaurant"
	"github.com/yumaeda/grpc/internal/service"
)

type RestaurantServer struct {
	pb.UnimplementedRestaurantServiceServer
	restaurantService service.RestaurantService
}

func NewRestaurantServer(restaurantService service.RestaurantService) *RestaurantServer {
	return &RestaurantServer{restaurantService: restaurantService}
}

func (s *RestaurantServer) GetRestaurant(ctx context.Context, req *pb.GetRestaurantRequest) (*pb.GetRestaurantResponse, error) {
	restaurant, err := s.restaurantService.GetRestaurant(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetRestaurantResponse{
		Restaurant: &pb.Restaurant{
			Id:              restaurant.ID,
			Url:             restaurant.URL,
			Name:            restaurant.Name,
			Genre:           restaurant.Genre,
			Tel:             restaurant.Tel,
			BusinessDayInfo: restaurant.BusinessDayInfo,
			Address:         restaurant.Address,
			Latitude:        restaurant.Latitude,
			Longitude:       restaurant.Longitude,
			Area:            restaurant.Area,
		},
	}, nil
}
