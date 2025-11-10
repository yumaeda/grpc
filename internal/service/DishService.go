package service

import (
	"context"

	"github.com/yumaeda/sakabas-grpc/internal/model"
	"github.com/yumaeda/sakabas-grpc/internal/repository"
)

type DishService interface {
	GetDish(ctx context.Context, id int64) (*model.Dish, error)
}

type dishService struct {
	dishRepository repository.DishRepository
}

func NewDishService(dishRepository repository.DishRepository) DishService {
	return &dishService{dishRepository: dishRepository}
}

func (s *dishService) GetDish(ctx context.Context, id int64) (*model.Dish, error) {
	return s.dishRepository.GetDish(ctx, id)
}
