package service

import (
	"context"

	"github.com/yumaeda/sakabas-grpc/internal/model"
	"github.com/yumaeda/sakabas-grpc/internal/repository"
)

type PhotoService interface {
	GetPhoto(ctx context.Context, id int64) (*model.Photo, error)
	ListPhotos(ctx context.Context, restaurantId string) ([]*model.Photo, error)
}

type photoService struct {
	photoRepository repository.PhotoRepository
}

func NewPhotoService(photoRepository repository.PhotoRepository) PhotoService {
	return &photoService{photoRepository: photoRepository}
}

func (s *photoService) GetPhoto(ctx context.Context, id int64) (*model.Photo, error) {
	return s.photoRepository.GetPhoto(ctx, id)
}

func (s *photoService) ListPhotos(ctx context.Context, restaurantId string) ([]*model.Photo, error) {
	return s.photoRepository.ListPhotos(ctx, restaurantId)
}
