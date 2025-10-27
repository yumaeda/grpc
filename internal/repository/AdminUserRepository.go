package repository

import (
	"context"
	"errors"

	"github.com/yumaeda/grpc/internal/model"
	"gorm.io/gorm"
)

type AdminUserRepository interface {
	GetAdminUser(ctx context.Context, id string) (*model.AdminUser, error)
}

type adminUserRepository struct {
	DB *gorm.DB
}

func NewAdminUserRepository(DB *gorm.DB) AdminUserRepository {
	return &adminUserRepository{DB: DB}
}

func (r *adminUserRepository) GetAdminUser(ctx context.Context, id string) (*model.AdminUser, error) {
	var adminUser model.AdminUser
	if err := r.DB.WithContext(ctx).
		Raw(`
			SELECT 
				BIN_TO_UUID(id, 1) as id,
				email,
				password
			FROM admin_users
			WHERE BIN_TO_UUID(id, 1) = ?
		`, id).
		Scan(&adminUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("admin user not found")
		}
		return nil, err
	}

	return &adminUser, nil
}
