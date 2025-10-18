package server

import (
	"context"

	area_proto "github.com/yumaeda/grpc/demo-grpc/internal/proto"
	"github.com/yumaeda/grpc/demo-grpc/internal/service"
)

type AreaServer struct {
	area_proto.UnimplementedAreaServiceServer
	areaService service.AreaService
}

func NewAreaServer(areaService service.AreaService) *AreaServer {
	return &AreaServer{areaService: areaService}
}

func (s *AreaServer) GetArea(ctx context.Context, req *area_proto.GetAreaRequest) (*area_proto.GetAreaResponse, error) {
	area, err := s.areaService.GetArea(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &area_proto.GetAreaResponse{
		Area: &area_proto.Area{
			Id:    area.ID,
			Name:  area.Name,
			Value: area.Value,
		},
	}, nil
}
