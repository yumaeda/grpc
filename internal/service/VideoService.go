package service

import (
	"context"

	"github.com/yumaeda/grpc/internal/model"
	"github.com/yumaeda/grpc/internal/repository"
)

type VideoService interface {
	GetVideo(ctx context.Context, id int64) (*model.Video, error)
}

type videoService struct {
	videoRepository repository.VideoRepository
}

func NewVideoService(videoRepository repository.VideoRepository) VideoService {
	return &videoService{videoRepository: videoRepository}
}

func (s *videoService) GetVideo(ctx context.Context, id int64) (*model.Video, error) {
	return s.videoRepository.GetVideo(ctx, id)
}
