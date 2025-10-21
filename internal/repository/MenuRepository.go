package repository

import (
	"context"
	"errors"

	"github.com/yumaeda/grpc/internal/model"
	"gorm.io/gorm"
)

type MenuRepository interface {
	GetMenu(ctx context.Context, id string) (*model.Menu, error)
}

type menuRepository struct {
	DB *gorm.DB
}

func NewMenuRepository(DB *gorm.DB) MenuRepository {
	return &menuRepository{DB: DB}
}

func (r *menuRepository) GetMenu(ctx context.Context, id string) (*model.Menu, error) {
	var menu model.Menu
	if err := r.DB.WithContext(ctx).
		Raw(`
			SELECT 
				BIN_TO_UUID(id, 1) as id,
				sort_order,
				category,
				sub_category,
				region,
				name,
				name_jpn,
				price,
				is_min_price,
				is_hidden
			FROM menus
			WHERE BIN_TO_UUID(id, 1) = ?
		`, id).
		Scan(&menu).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("menu not found")
		}
		return nil, err
	}

	return &menu, nil
}
