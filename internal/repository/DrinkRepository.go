package repository

import (
	"context"
	"errors"

	"github.com/yumaeda/sakabas-grpc/internal/model"
	"gorm.io/gorm"
)

type DrinkRepository interface {
	GetDrink(ctx context.Context, id int64) (*model.Drink, error)
}

type drinkRepository struct {
	DB *gorm.DB
}

func NewDrinkRepository(DB *gorm.DB) DrinkRepository {
	return &drinkRepository{DB: DB}
}

func (r *drinkRepository) GetDrink(ctx context.Context, id int64) (*model.Drink, error) {
	var drink model.Drink
	if err := r.DB.WithContext(ctx).Where("id = ?", id).First(&drink).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("drink not found")
		}
		return nil, err
	}

	return &drink, nil
}
