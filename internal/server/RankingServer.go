package server

import (
	"context"

	"github.com/yumaeda/sakabas-grpc/internal/service"
	pb "github.com/yumaeda/sakabas-grpc/proto/ranking"
)

type RankingServer struct {
	pb.UnimplementedRankingServiceServer
	rankingService service.RankingService
}

func NewRankingServer(rankingService service.RankingService) *RankingServer {
	return &RankingServer{rankingService: rankingService}
}

func (s *RankingServer) GetRanking(ctx context.Context, req *pb.GetRankingRequest) (*pb.GetRankingResponse, error) {
	ranking, err := s.rankingService.GetRanking(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetRankingResponse{
		Ranking: &pb.Ranking{
			Id:      ranking.ID,
			Rank:    ranking.Rank,
			DishId:  ranking.DishID,
			PhotoId: ranking.PhotoID,
		},
	}, nil
}
