package service

import (
	"context"

	"github.com/yumaeda/sakabas-grpc/internal/model"
	"github.com/yumaeda/sakabas-grpc/internal/repository"
)

type GenreService interface {
	GetGenre(ctx context.Context, id int64) (*model.Genre, error)
}

type genreService struct {
	genreRepository repository.GenreRepository
}

func NewGenreService(genreRepository repository.GenreRepository) GenreService {
	return &genreService{genreRepository: genreRepository}
}

func (s *genreService) GetGenre(ctx context.Context, id int64) (*model.Genre, error) {
	return s.genreRepository.GetGenre(ctx, id)
}
