package service

import (
	"context"

	"github.com/yumaeda/sakabas-grpc/internal/model"
	"github.com/yumaeda/sakabas-grpc/internal/repository"
)

type RestaurantGenreService interface {
	GetRestaurantGenre(ctx context.Context, restaurantID string, genreID int64) (*model.RestaurantGenre, error)
}

type restaurantGenreService struct {
	restaurantGenreRepository repository.RestaurantGenreRepository
}

func NewRestaurantGenreService(restaurantGenreRepository repository.RestaurantGenreRepository) RestaurantGenreService {
	return &restaurantGenreService{restaurantGenreRepository: restaurantGenreRepository}
}

func (s *restaurantGenreService) GetRestaurantGenre(ctx context.Context, restaurantID string, genreID int64) (*model.RestaurantGenre, error) {
	return s.restaurantGenreRepository.GetRestaurantGenre(ctx, restaurantID, genreID)
}
