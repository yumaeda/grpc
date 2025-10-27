package service

import (
	"context"

	"github.com/yumaeda/grpc/internal/model"
	"github.com/yumaeda/grpc/internal/repository"
)

type CategoryService interface {
	GetCategory(ctx context.Context, id int64) (*model.Category, error)
}

type categoryService struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryService(categoryRepository repository.CategoryRepository) CategoryService {
	return &categoryService{categoryRepository: categoryRepository}
}

func (s *categoryService) GetCategory(ctx context.Context, id int64) (*model.Category, error) {
	return s.categoryRepository.GetCategory(ctx, id)
}
