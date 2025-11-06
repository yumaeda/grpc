package server

import (
	"context"

	"github.com/yumaeda/sakabas-grpc/internal/service"
	pb "github.com/yumaeda/sakabas-grpc/swapi/video/video"
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

func (s *VideoServer) ListVideos(ctx context.Context, req *pb.ListVideosRequest) (*pb.ListVideosResponse, error) {
	videos, err := s.videoService.ListVideos(ctx, req.RestaurantId)
	if err != nil {
		return nil, err
	}

	pbVideos := make([]*pb.Video, len(videos))
	for i, video := range videos {
		pbVideos[i] = &pb.Video{
			Id:           video.ID,
			RestaurantId: video.RestaurantID,
			Name:         video.Name,
			Url:          video.URL,
		}
	}

	return &pb.ListVideosResponse{
		Videos: pbVideos,
	}, nil
}
