package repository

import (
	"context"
	"errors"

	"github.com/yumaeda/grpc/internal/model"
	"gorm.io/gorm"
)

type RestaurantRepository interface {
	GetRestaurant(ctx context.Context, id string) (*model.Restaurant, error)
}

type restaurantRepository struct {
	DB *gorm.DB
}

func NewRestaurantRepository(DB *gorm.DB) RestaurantRepository {
	return &restaurantRepository{DB: DB}
}

func (r *restaurantRepository) GetRestaurant(ctx context.Context, id string) (*model.Restaurant, error) {
	var restaurant model.Restaurant
	if err := r.DB.WithContext(ctx).
		Select("BIN_TO_UUID(id, 1) as id, url, name, genre, tel, business_day_info, address, latitude, longitude, area").
		Where("BIN_TO_UUID(id, 1) = ?", id).
		First(&restaurant).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("restaurant not found")
		}
		return nil, err
	}

	return &restaurant, nil
}
