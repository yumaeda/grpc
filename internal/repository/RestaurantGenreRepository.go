package repository

import (
	"context"
	"errors"

	"github.com/yumaeda/sakabas-grpc/internal/model"
	"gorm.io/gorm"
)

type RestaurantGenreRepository interface {
	GetRestaurantGenre(ctx context.Context, restaurantID string, genreID int64) (*model.RestaurantGenre, error)
}

type restaurantGenreRepository struct {
	DB *gorm.DB
}

func NewRestaurantGenreRepository(DB *gorm.DB) RestaurantGenreRepository {
	return &restaurantGenreRepository{DB: DB}
}

func (r *restaurantGenreRepository) GetRestaurantGenre(ctx context.Context, restaurantID string, genreID int64) (*model.RestaurantGenre, error) {
	var restaurantGenre model.RestaurantGenre
	if err := r.DB.WithContext(ctx).
		Raw(`
			SELECT 
				BIN_TO_UUID(restaurant_id, 1) as restaurant_id,
				genre_id
			FROM restaurant_genres
			WHERE BIN_TO_UUID(restaurant_id, 1) = ? AND genre_id = ?
		`, restaurantID, genreID).
		Scan(&restaurantGenre).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("restaurant genre not found")
		}
		return nil, err
	}

	return &restaurantGenre, nil
}
