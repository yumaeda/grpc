package repository

import (
	"context"
	"errors"

	"github.com/yumaeda/sakabas-grpc/internal/model"
	"gorm.io/gorm"
)

type RankingRepository interface {
	GetRanking(ctx context.Context, id int64) (*model.Ranking, error)
}

type rankingRepository struct {
	DB *gorm.DB
}

func NewRankingRepository(DB *gorm.DB) RankingRepository {
	return &rankingRepository{DB: DB}
}

func (r *rankingRepository) GetRanking(ctx context.Context, id int64) (*model.Ranking, error) {
	var ranking model.Ranking
	if err := r.DB.WithContext(ctx).
		Raw(`
			SELECT 
				id,
				`+"`rank`"+`,
				dish_id,
				photo_id
			FROM rankings
			WHERE id = ?
		`, id).
		Scan(&ranking).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("ranking not found")
		}
		return nil, err
	}

	return &ranking, nil
}
