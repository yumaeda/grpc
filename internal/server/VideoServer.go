package server

import (
	"context"

	pb "github.com/yumaeda/grpc/internal/proto/video"
	"github.com/yumaeda/grpc/internal/service"
)

type VideoServer struct {
	pb.UnimplementedVideoServiceServer
	videoService service.VideoService
}

func NewVideoServer(videoService service.VideoService) *VideoServer {
	return &VideoServer{videoService: videoService}
}

func (s *VideoServer) GetVideo(ctx context.Context, req *pb.GetVideoRequest) (*pb.GetVideoResponse, error) {
	video, err := s.videoService.GetVideo(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetVideoResponse{
		Video: &pb.Video{
			Id:           video.ID,
			RestaurantId: video.RestaurantID,
			Name:         video.Name,
			Url:          video.URL,
		},
	}, nil
}
