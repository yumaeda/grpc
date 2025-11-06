package server

import (
	"context"

	"github.com/yumaeda/sakabas-grpc/internal/service"
	pb "github.com/yumaeda/sakabas-grpc/proto/genre"
)

type GenreServer struct {
	pb.UnimplementedGenreServiceServer
	genreService service.GenreService
}

func NewGenreServer(genreService service.GenreService) *GenreServer {
	return &GenreServer{genreService: genreService}
}

func (s *GenreServer) GetGenre(ctx context.Context, req *pb.GetGenreRequest) (*pb.GetGenreResponse, error) {
	genre, err := s.genreService.GetGenre(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetGenreResponse{
		Genre: &pb.Genre{
			Id:   genre.ID,
			Name: genre.Name,
		},
	}, nil
}
