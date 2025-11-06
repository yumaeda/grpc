package server

import (
	"context"

	"github.com/yumaeda/sakabas-grpc/internal/service"
	pb "github.com/yumaeda/sakabas-grpc/proto/admin_user"
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
			Id:    adminUser.ID,
			Email: adminUser.Email,
		},
	}, nil
}
