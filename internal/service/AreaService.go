package service

import (
	"context"

	"github.com/yumaeda/grpc/internal/model"
	"github.com/yumaeda/grpc/internal/repository"
)

type AreaService interface {
	GetArea(ctx context.Context, id int64) (*model.Area, error)
}

type areaService struct {
	areaRepository repository.AreaRepository
}

func NewAreaService(areaRepository repository.AreaRepository) AreaService {
	return &areaService{areaRepository: areaRepository}
}

func (s *areaService) GetArea(ctx context.Context, id int64) (*model.Area, error) {
	return s.areaRepository.GetArea(ctx, id)
}
