package server

import (
	"context"

	pb "github.com/yumaeda/grpc/internal/proto/category"
	"github.com/yumaeda/grpc/internal/service"
)

type CategoryServer struct {
	pb.UnimplementedCategoryServiceServer
	categoryService service.CategoryService
}

func NewCategoryServer(categoryService service.CategoryService) *CategoryServer {
	return &CategoryServer{categoryService: categoryService}
}

func (s *CategoryServer) GetCategory(ctx context.Context, req *pb.GetCategoryRequest) (*pb.GetCategoryResponse, error) {
	category, err := s.categoryService.GetCategory(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetCategoryResponse{
		Category: &pb.Category{
			Id:       category.ID,
			ParentId: category.ParentID,
			Name:     category.Name,
		},
	}, nil
}
