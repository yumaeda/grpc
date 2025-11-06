package service

import (
	"context"

	"github.com/yumaeda/sakabas-grpc/internal/model"
	"github.com/yumaeda/sakabas-grpc/internal/repository"
)

type MenuService interface {
	GetMenu(ctx context.Context, id string) (*model.Menu, error)
}

type menuService struct {
	menuRepository repository.MenuRepository
}

func NewMenuService(menuRepository repository.MenuRepository) MenuService {
	return &menuService{menuRepository: menuRepository}
}

func (s *menuService) GetMenu(ctx context.Context, id string) (*model.Menu, error) {
	return s.menuRepository.GetMenu(ctx, id)
}
