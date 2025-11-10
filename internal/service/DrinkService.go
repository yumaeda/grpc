package service

import (
	"context"

	"github.com/yumaeda/sakabas-grpc/internal/model"
	"github.com/yumaeda/sakabas-grpc/internal/repository"
)

type DrinkService interface {
	GetDrink(ctx context.Context, id int64) (*model.Drink, error)
}

type drinkService struct {
	drinkRepository repository.DrinkRepository
}

func NewDrinkService(drinkRepository repository.DrinkRepository) DrinkService {
	return &drinkService{drinkRepository: drinkRepository}
}

func (s *drinkService) GetDrink(ctx context.Context, id int64) (*model.Drink, error) {
	return s.drinkRepository.GetDrink(ctx, id)
}
