package service

import (
	"context"

	"github.com/yumaeda/grpc/internal/model"
	"github.com/yumaeda/grpc/internal/repository"
)

type AdminUserService interface {
	GetAdminUser(ctx context.Context, id string) (*model.AdminUser, error)
}

type adminUserService struct {
	adminUserRepository repository.AdminUserRepository
}

func NewAdminUserService(adminUserRepository repository.AdminUserRepository) AdminUserService {
	return &adminUserService{adminUserRepository: adminUserRepository}
}

func (s *adminUserService) GetAdminUser(ctx context.Context, id string) (*model.AdminUser, error) {
	return s.adminUserRepository.GetAdminUser(ctx, id)
}
