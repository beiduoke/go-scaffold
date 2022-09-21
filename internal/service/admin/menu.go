package admin

import (
	"context"

	v1 "github.com/beiduoke/go-scaffold/api/admin/v1"
	"github.com/beiduoke/go-scaffold/api/protobuf"
	"google.golang.org/protobuf/types/known/anypb"
)

var _ v1.AdminServer = (*AdminService)(nil)

// ListMenu 列表菜单
func (s *AdminService) ListMenu(ctx context.Context, in *protobuf.PagingReq) (*protobuf.PagingReply, error) {
	return &protobuf.PagingReply{
		Total: 0,
		Items: []*anypb.Any{},
	}, nil
}

// CreateMenu 创建菜单
func (s *AdminService) CreateMenu(ctx context.Context, in *v1.CreateMenuReq) (*v1.CreateMenuReply, error) {
	return &v1.CreateMenuReply{
		Success: true,
	}, nil
}

// GetMenu 获取菜单
func (s *AdminService) GetMenu(ctx context.Context, in *v1.GetMenuReq) (*v1.Menu, error) {
	return &v1.Menu{}, nil
}

// UpdateMenu 修改菜单
func (s *AdminService) UpdateMenu(ctx context.Context, in *v1.UpdateMenuReq) (*v1.UpdateMenuReply, error) {
	return &v1.UpdateMenuReply{
		Success: true,
	}, nil
}

// DeleteMenu 删除菜单
func (s *AdminService) DeleteMenu(ctx context.Context, in *v1.DeleteMenuReq) (*v1.DeleteMenuReply, error) {
	return &v1.DeleteMenuReply{
		Success: true,
	}, nil
}
