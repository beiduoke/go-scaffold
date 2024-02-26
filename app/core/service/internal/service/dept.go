package service

import (
	"context"

	"github.com/beiduoke/go-scaffold/api/common/pagination"
	v1 "github.com/beiduoke/go-scaffold/api/core/service/v1"
	"github.com/beiduoke/go-scaffold/app/core/service/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

type DeptService struct {
	v1.UnimplementedDeptServiceServer
	log *log.Helper
	ac  *data.DeptRepo
}

func NewDeptService(logger log.Logger, ac *data.DeptRepo) *DeptService {
	l := log.NewHelper(log.With(logger, "module", "dept/service"))
	return &DeptService{
		log: l,
		ac:  ac,
	}
}

func (s *DeptService) CreateDept(ctx context.Context, req *v1.CreateDeptRequest) (*v1.CreateDeptResponse, error) {
	return s.ac.CreateDept(ctx, req)
}

func (s *DeptService) UpdateDept(ctx context.Context, req *v1.UpdateDeptRequest) (*v1.UpdateDeptResponse, error) {
	return s.ac.UpdateDept(ctx, req)
}

func (s *DeptService) DeleteDept(ctx context.Context, req *v1.DeleteDeptRequest) (*v1.DeleteDeptResponse, error) {
	return s.ac.DeleteDept(ctx, req)
}

func (s *DeptService) GetDept(ctx context.Context, req *v1.GetDeptRequest) (*v1.Dept, error) {
	return s.ac.GetDept(ctx, req)
}

func (s *DeptService) ListDept(ctx context.Context, req *pagination.PagingRequest) (*v1.ListDeptResponse, error) {
	return s.ac.ListDept(ctx, req)
}
