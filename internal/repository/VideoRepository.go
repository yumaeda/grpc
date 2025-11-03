package repository

import (
	"context"
	"errors"

	"github.com/yumaeda/grpc/internal/model"
	"gorm.io/gorm"
)

type VideoRepository interface {
	GetVideo(ctx context.Context, id int64) (*model.Video, error)
	ListVideos(ctx context.Context, restaurantId string) ([]*model.Video, error)
}

type videoRepository struct {
	DB *gorm.DB
}

func NewVideoRepository(DB *gorm.DB) VideoRepository {
	return &videoRepository{DB: DB}
}

func (r *videoRepository) GetVideo(ctx context.Context, id int64) (*model.Video, error) {
	var video model.Video
	if err := r.DB.WithContext(ctx).
		Raw(`
			SELECT 
				id,
				BIN_TO_UUID(restaurant_id, 1) as restaurant_id,
				name,
				url
			FROM videos
			WHERE id = ?
		`, id).
		Scan(&video).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("video not found")
		}
		return nil, err
	}

	return &video, nil
}

func (r *videoRepository) ListVideos(ctx context.Context, restaurantId string) ([]*model.Video, error) {
	var videos []*model.Video
	if err := r.DB.WithContext(ctx).
		Raw(`
			SELECT 
				id,
				BIN_TO_UUID(restaurant_id, 1) as restaurant_id,
				name,
				url
			FROM videos
			WHERE BIN_TO_UUID(restaurant_id, 1) = ?
		`, restaurantId).Scan(&videos).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("videos for the restaurant not found")
		}
		return nil, err
	}

	return videos, nil
}
