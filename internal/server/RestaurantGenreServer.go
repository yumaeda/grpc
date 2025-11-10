package server

import (
	"context"

	"github.com/yumaeda/sakabas-grpc/internal/service"
	pb "github.com/yumaeda/sakabas-grpc/proto/restaurant_genre"
)

type RestaurantGenreServer struct {
	pb.UnimplementedRestaurantGenreServiceServer
	restaurantGenreService service.RestaurantGenreService
}

func NewRestaurantGenreServer(restaurantGenreService service.RestaurantGenreService) *RestaurantGenreServer {
	return &RestaurantGenreServer{restaurantGenreService: restaurantGenreService}
}

func (s *RestaurantGenreServer) GetRestaurantGenre(ctx context.Context, req *pb.GetRestaurantGenreRequest) (*pb.GetRestaurantGenreResponse, error) {
	restaurantGenre, err := s.restaurantGenreService.GetRestaurantGenre(ctx, req.RestaurantId, req.GenreId)
	if err != nil {
		return nil, err
	}

	return &pb.GetRestaurantGenreResponse{
		RestaurantGenre: &pb.RestaurantGenre{
			RestaurantId: restaurantGenre.RestaurantID,
			GenreId:      restaurantGenre.GenreID,
		},
	}, nil
}
