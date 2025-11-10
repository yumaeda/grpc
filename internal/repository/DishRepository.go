package repository

import (
	"context"
	"errors"

	"github.com/yumaeda/sakabas-grpc/internal/model"
	"gorm.io/gorm"
)

type DishRepository interface {
	GetDish(ctx context.Context, id int64) (*model.Dish, error)
}

type dishRepository struct {
	DB *gorm.DB
}

func NewDishRepository(DB *gorm.DB) DishRepository {
	return &dishRepository{DB: DB}
}

func (r *dishRepository) GetDish(ctx context.Context, id int64) (*model.Dish, error) {
	var dish model.Dish
	if err := r.DB.WithContext(ctx).Where("id = ?", id).First(&dish).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("dish not found")
		}
		return nil, err
	}

	return &dish, nil
}
