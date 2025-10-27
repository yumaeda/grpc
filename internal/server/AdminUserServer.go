package server

import (
	"context"

	pb "github.com/yumaeda/grpc/internal/proto/admin_user"
	"github.com/yumaeda/grpc/internal/service"
)

type AdminUserServer struct {
	pb.UnimplementedAdminUserServiceServer
	adminUserService service.AdminUserService
}

func NewAdminUserServer(adminUserService service.AdminUserService) *AdminUserServer {
	return &AdminUserServer{adminUserService: adminUserService}
}

func (s *AdminUserServer) GetAdminUser(ctx context.Context, req *pb.GetAdminUserRequest) (*pb.GetAdminUserResponse, error) {
	adminUser, err := s.adminUserService.GetAdminUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetAdminUserResponse{
		AdminUser: &pb.AdminUser{
			Id:       adminUser.ID,
			Email:    adminUser.Email,
			Password: adminUser.Password,
		},
	}, nil
}
