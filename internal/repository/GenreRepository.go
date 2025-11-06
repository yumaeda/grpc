package repository

import (
	"context"
	"errors"

	"github.com/yumaeda/sakabas-grpc/internal/model"
	"gorm.io/gorm"
)

type GenreRepository interface {
	GetGenre(ctx context.Context, id int64) (*model.Genre, error)
}

type genreRepository struct {
	DB *gorm.DB
}

func NewGenreRepository(DB *gorm.DB) GenreRepository {
	return &genreRepository{DB: DB}
}

func (r *genreRepository) GetGenre(ctx context.Context, id int64) (*model.Genre, error) {
	var genre model.Genre
	if err := r.DB.WithContext(ctx).Where("id = ?", id).First(&genre).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("genre not found")
		}
		return nil, err
	}

	return &genre, nil
}
