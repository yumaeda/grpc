package repository

import (
	"context"
	"errors"

	"github.com/yumaeda/grpc/internal/model"
	"gorm.io/gorm"
)

type PhotoRepository interface {
	GetPhoto(ctx context.Context, id int64) (*model.Photo, error)
}

type photoRepository struct {
	DB *gorm.DB
}

func NewPhotoRepository(DB *gorm.DB) PhotoRepository {
	return &photoRepository{DB: DB}
}

func (r *photoRepository) GetPhoto(ctx context.Context, id int64) (*model.Photo, error) {
	var photo model.Photo
	if err := r.DB.WithContext(ctx).
		Raw(`
			SELECT 
				id,
				BIN_TO_UUID(restaurant_id, 1) as restaurant_id,
				name,
				CONCAT(name, '.jpg') AS image,
				CONCAT(name, '.webp') AS image_webp,
				CONCAT(name, '_thumbnail.jpg') AS thumbnail,
				CONCAT(name, '_thumbnail.webp') AS thumbnail_webp
			FROM photos
			WHERE id = ?
		`, id).
		Scan(&photo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("photo not found")
		}
		return nil, err
	}

	return &photo, nil
}
