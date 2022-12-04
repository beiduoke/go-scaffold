package admin

import (
	"context"

	v1 "github.com/beiduoke/go-scaffold/api/admin/v1"
	"github.com/beiduoke/go-scaffold/api/protobuf"
	"github.com/beiduoke/go-scaffold/internal/biz"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ v1.AdminServer = (*AdminService)(nil)

func TransformDepartment(data *biz.Department) *v1.Department {
	return &v1.Department{
		CreatedAt: timestamppb.New(data.CreatedAt),
		UpdatedAt: timestamppb.New(data.UpdatedAt),
		Id:        uint64(data.ID),
		Name:      data.Name,
		ParentId:  uint64(data.ParentID),
		Sort:      data.Sort,
		State:     protobuf.DepartmentState(data.State),
		Remarks:   data.Remarks,
	}
}

// TreeMenu 部门树形
func TreeDepartment(departments []*biz.Department, pid uint) []*v1.Department {
	list := make([]*v1.Department, 0)
	for _, dept := range departments {
		if dept.ParentID == pid {
			m := TransformDepartment(dept)
			m.Children = append(m.Children, TreeDepartment(departments, dept.ID)...)
			list = append(list, m)
		}
	}
	return list
}

// GetTreeDepartment 列表部门-树形
func (s *AdminService) ListDepartmentTree(ctx context.Context, in *v1.ListDepartmentTreeReq) (*v1.ListDepartmentTreeReply, error) {
	results := s.departmentCase.GetTree(ctx, uint(in.GetParentId()))
	return &v1.ListDepartmentTreeReply{
		Items: TreeDepartment(results, uint(in.GetParentId())),
	}, nil
}

// ListDepartment 列表-部门角色
func (s *AdminService) ListDepartment(ctx context.Context, in *protobuf.PagingReq) (*protobuf.PagingReply, error) {
	results, total := s.departmentCase.ListPage(ctx, in.GetPage(), in.GetPageSize(), in.GetQuery(), in.GetOrderBy())
	items := make([]*anypb.Any, 0, len(results))
	for _, v := range results {
		item, _ := anypb.New(TransformDepartment(v))
		items = append(items, item)
	}
	return &protobuf.PagingReply{
		Total: int32(total),
		Items: items,
	}, nil
}

// CreateDepartment 创建部门角色
func (s *AdminService) CreateDepartment(ctx context.Context, in *v1.CreateDepartmentReq) (*v1.CreateDepartmentReply, error) {
	user, err := s.departmentCase.Create(ctx, &biz.Department{
		Name:     in.GetName(),
		ParentID: uint(in.GetParentId()),
		State:    int32(in.GetState()),
		Remarks:  in.GetRemarks(),
	})
	if err != nil {
		return nil, v1.ErrorDepartmentCreateFail("部门角色创建失败: %v", err.Error())
	}
	data, _ := anypb.New(&protobuf.DataProto{
		Id: uint64(user.ID),
	})
	return &v1.CreateDepartmentReply{
		Success: true,
		Message: "创建成功",
		Data:    data,
	}, nil
}

// UpdateDepartment 创建部门角色
func (s *AdminService) UpdateDepartment(ctx context.Context, in *v1.UpdateDepartmentReq) (*v1.UpdateDepartmentReply, error) {
	v := in.GetData()
	err := s.departmentCase.Update(ctx, &biz.Department{
		ID:       uint(in.GetId()),
		Name:     v.GetName(),
		ParentID: uint(v.GetParentId()),
		Sort:     v.GetSort(),
		State:    int32(v.GetState()),
		Remarks:  v.GetRemarks(),
	})
	if err != nil {
		return nil, v1.ErrorDepartmentUpdateFail("部门角色创建失败: %v", err.Error())
	}
	return &v1.UpdateDepartmentReply{
		Success: true,
		Message: "修改成功",
	}, nil
}

// GetDepartment 获取部门角色
func (s *AdminService) GetDepartment(ctx context.Context, in *v1.GetDepartmentReq) (*v1.Department, error) {
	department, err := s.departmentCase.GetID(ctx, &biz.Department{ID: uint(in.GetId())})
	if err != nil {
		return nil, v1.ErrorDepartmentNotFound("部门角色未找到")
	}
	return TransformDepartment(department), nil
}

// DeleteDepartment 删除部门角色
func (s *AdminService) DeleteDepartment(ctx context.Context, in *v1.DeleteDepartmentReq) (*v1.DeleteDepartmentReply, error) {
	if err := s.departmentCase.Delete(ctx, &biz.Department{ID: uint(in.GetId())}); err != nil {
		return nil, v1.ErrorDepartmentDeleteFail("部门角色删除失败：%v", err)
	}
	return &v1.DeleteDepartmentReply{
		Success: true,
		Message: "删除成功",
	}, nil
}
