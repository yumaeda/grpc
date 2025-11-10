package server

import (
	"context"

	"github.com/yumaeda/sakabas-grpc/internal/service"
	pb "github.com/yumaeda/sakabas-grpc/swapi/photo/photo"
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

func (s *PhotoServer) ListPhotos(ctx context.Context, req *pb.ListPhotosRequest) (*pb.ListPhotosResponse, error) {
	photos, err := s.photoService.ListPhotos(ctx, req.RestaurantId)
	if err != nil {
		return nil, err
	}

	pbPhotos := make([]*pb.Photo, len(photos))
	for i, photo := range photos {
		pbPhotos[i] = &pb.Photo{
			Id:            photo.ID,
			RestaurantId:  photo.RestaurantID,
			Name:          photo.Name,
			Image:         photo.Image,
			ImageWebp:     photo.ImageWebp,
			Thumbnail:     photo.Thumbnail,
			ThumbnailWebp: photo.ThumbnailWebp,
		}
	}

	return &pb.ListPhotosResponse{
		Photos: pbPhotos,
	}, nil
}
