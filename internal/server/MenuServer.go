package server

import (
	"context"

	pb "github.com/yumaeda/grpc/internal/proto/menu"
	"github.com/yumaeda/grpc/internal/service"
)

type MenuServer struct {
	pb.UnimplementedMenuServiceServer
	menuService service.MenuService
}

func NewMenuServer(menuService service.MenuService) *MenuServer {
	return &MenuServer{menuService: menuService}
}

func (s *MenuServer) GetMenu(ctx context.Context, req *pb.GetMenuRequest) (*pb.GetMenuResponse, error) {
	menu, err := s.menuService.GetMenu(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetMenuResponse{
		Menu: &pb.Menu{
			Id:          menu.ID,
			SortOrder:   menu.SortOrder,
			Category:    menu.Category,
			SubCategory: menu.SubCategory,
			Region:      menu.Region,
			Name:        menu.Name,
			NameJpn:     menu.NameJpn,
			Price:       menu.Price,
			IsMinPrice:  menu.IsMinPrice,
			IsHidden:    menu.IsHidden,
		},
	}, nil
}
