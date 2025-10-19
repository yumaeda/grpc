package service

import (
	"context"

	"github.com/yumaeda/grpc/internal/model"
	"github.com/yumaeda/grpc/internal/repository"
)

type RestaurantService interface {
	GetRestaurant(ctx context.Context, id string) (*model.Restaurant, error)
}

type restaurantService struct {
	restaurantRepository repository.RestaurantRepository
}

func NewRestaurantService(restaurantRepository repository.RestaurantRepository) RestaurantService {
	return &restaurantService{restaurantRepository: restaurantRepository}
}

func (s *restaurantService) GetRestaurant(ctx context.Context, id string) (*model.Restaurant, error) {
	return s.restaurantRepository.GetRestaurant(ctx, id)
}
