package server

import (
	"context"

	"github.com/yumaeda/grpc/internal/service"
	pb "github.com/yumaeda/grpc/proto/area"
)

type AreaServer struct {
	pb.UnimplementedAreaServiceServer
	areaService service.AreaService
}

func NewAreaServer(areaService service.AreaService) *AreaServer {
	return &AreaServer{areaService: areaService}
}

func (s *AreaServer) GetArea(ctx context.Context, req *pb.GetAreaRequest) (*pb.GetAreaResponse, error) {
	area, err := s.areaService.GetArea(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetAreaResponse{
		Area: &pb.Area{
			Id:    area.ID,
			Name:  area.Name,
			Value: area.Value,
		},
	}, nil
}
