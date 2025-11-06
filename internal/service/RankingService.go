package service

import (
	"context"

	"github.com/yumaeda/sakabas-grpc/internal/model"
	"github.com/yumaeda/sakabas-grpc/internal/repository"
)

type RankingService interface {
	GetRanking(ctx context.Context, id int64) (*model.Ranking, error)
}

type rankingService struct {
	rankingRepository repository.RankingRepository
}

func NewRankingService(rankingRepository repository.RankingRepository) RankingService {
	return &rankingService{rankingRepository: rankingRepository}
}

func (s *rankingService) GetRanking(ctx context.Context, id int64) (*model.Ranking, error) {
	return s.rankingRepository.GetRanking(ctx, id)
}
