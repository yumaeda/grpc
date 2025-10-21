package server

import (
	"context"

	pb "github.com/yumaeda/grpc/internal/proto/photo"
	"github.com/yumaeda/grpc/internal/service"
)

type PhotoServer struct {
	pb.UnimplementedPhotoServiceServer
	photoService service.PhotoService
}

func NewPhotoServer(photoService service.PhotoService) *PhotoServer {
	return &PhotoServer{photoService: photoService}
}

func (s *PhotoServer) GetPhoto(ctx context.Context, req *pb.GetPhotoRequest) (*pb.GetPhotoResponse, error) {
	photo, err := s.photoService.GetPhoto(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetPhotoResponse{
		Photo: &pb.Photo{
			Id:            photo.ID,
			RestaurantId:  photo.RestaurantID,
			Name:          photo.Name,
			Image:         photo.Image,
			ImageWebp:     photo.ImageWebp,
			Thumbnail:     photo.Thumbnail,
			ThumbnailWebp: photo.ThumbnailWebp,
		},
	}, nil
}
