package repository

import (
	"context"
	"errors"

	"github.com/yumaeda/grpc/demo-grpc/internal/model"
	"gorm.io/gorm"
)

type AreaRepository interface {
	GetArea(ctx context.Context, id int64) (*model.Area, error)
}

type areaRepository struct {
	DB *gorm.DB
}

func NewAreaRepository(DB *gorm.DB) AreaRepository {
	return &areaRepository{DB: DB}
}

func (r *areaRepository) GetArea(ctx context.Context, id int64) (*model.Area, error) {
	var area model.Area
	if err := r.DB.WithContext(ctx).Where("id = ?", id).First(&area).Error; err != nil {
		return nil, errors.New("area not found")
	}

	return &area, nil
}
