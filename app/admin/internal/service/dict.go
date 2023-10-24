package service

import (
	"context"

	v1 "github.com/beiduoke/go-scaffold/api/admin/v1"
	"github.com/beiduoke/go-scaffold/app/admin/internal/biz"
	"github.com/beiduoke/go-scaffold/app/admin/internal/pkg/constant"
	"github.com/beiduoke/go-scaffold/pkg/util/convert"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ v1.AdminServiceServer = (*AdminService)(nil)

func TransformDict(data *biz.Dict) *v1.Dict {
	return &v1.Dict{
		CreatedAt: timestamppb.New(data.CreatedAt),
		UpdatedAt: timestamppb.New(data.UpdatedAt),
		Id:        uint64(data.ID),
		Name:      data.Name,
		Type:      data.Type,
		Sort:      &data.Sort,
		State:     &data.State,
		Remarks:   &data.Remarks,
	}
}

// ListDict 列表-字典
func (s *AdminService) ListDict(ctx context.Context, in *v1.ListDictReq) (*v1.ListDictReply, error) {
	results, total := s.dictCase.ListPage(ctx, pagination.NewPagination(pagination.WithPage(in.GetPage()), pagination.WithPageSize(in.GetPageSize())))
	return &v1.ListDictReply{
		Total: total,
		Items: convert.ArrayToAny(results, func(v *biz.Dict) *v1.Dict {
			return TransformDict(v)
		}),
	}, nil
}

// CreateDict 创建字典
func (s *AdminService) CreateDict(ctx context.Context, in *v1.CreateDictReq) (*v1.CreateDictReply, error) {
	role, err := s.dictCase.Create(ctx, &biz.Dict{
		Name:    in.GetName(),
		Type:    in.GetType(),
		Sort:    in.GetSort(),
		State:   int32(in.GetState()),
		Remarks: in.GetRemarks(),
	})
	if err != nil {
		return nil, v1.ErrorDictCreateFail("字典创建失败: %v", err.Error())
	}

	data, _ := anypb.New(&v1.Result{
		Id: uint64(role.ID),
	})
	return &v1.CreateDictReply{
		Type:    constant.HandleType_success.String(),
		Message: "创建成功",
		Result:  data,
	}, nil
}

// UpdateDict 修改字典
func (s *AdminService) UpdateDict(ctx context.Context, in *v1.UpdateDictReq) (*v1.UpdateDictReply, error) {
	v := in.GetData()
	err := s.dictCase.Update(ctx, &biz.Dict{
		ID:      uint(in.GetId()),
		Name:    v.GetName(),
		Sort:    v.GetSort(),
		State:   int32(v.GetState()),
		Remarks: v.GetRemarks(),
	})
	if err != nil {
		return nil, v1.ErrorDictUpdateFail("字典修改失败: %v", err.Error())
	}
	return &v1.UpdateDictReply{
		Type:    constant.HandleType_success.String(),
		Message: "修改成功",
	}, nil
}

// UpdateDictState 修改字典-状态
func (s *AdminService) UpdateDictState(ctx context.Context, in *v1.UpdateDictStateReq) (*v1.UpdateDictStateReply, error) {
	v := in.GetData()
	err := s.dictCase.UpdateState(ctx, &biz.Dict{
		ID:    uint(in.GetId()),
		State: int32(v.GetState()),
	})
	if err != nil {
		return nil, v1.ErrorDictUpdateFail("领域状态修改失败: %v", err.Error())
	}
	return &v1.UpdateDictStateReply{
		Message: "修改成功",
		Type:    constant.HandleType_success.String(),
	}, nil
}

// GetDict 获取字典
func (s *AdminService) GetDict(ctx context.Context, in *v1.GetDictReq) (*v1.Dict, error) {
	role, err := s.dictCase.GetID(ctx, &biz.Dict{ID: uint(in.GetId())})
	if err != nil {
		return nil, v1.ErrorDictNotFound("字典未找到")
	}
	return TransformDict(role), nil
}

// DeleteDict 删除字典
func (s *AdminService) DeleteDict(ctx context.Context, in *v1.DeleteDictReq) (*v1.DeleteDictReply, error) {
	if err := s.dictCase.Delete(ctx, &biz.Dict{ID: uint(in.GetId())}); err != nil {
		return nil, v1.ErrorDictDeleteFail("字典删除失败：%v", err)
	}
	return &v1.DeleteDictReply{
		Type:    constant.HandleType_success.String(),
		Message: "删除成功",
	}, nil
}

func TransformDictData(data *biz.DictData) *v1.DictData {
	return &v1.DictData{
		CreatedAt: timestamppb.New(data.CreatedAt),
		UpdatedAt: timestamppb.New(data.UpdatedAt),
		Id:        uint64(data.ID),
		Label:     data.Label,
		Value:     data.Value,
		DictType:  data.DictType,
		ColorType: new(string),
		CssClass:  new(string),
		Sort:      &data.Sort,
		Remarks:   &data.Remarks,
		State:     &data.State,
	}
}

// ListDictData 列表-字典数据
func (s *AdminService) ListDictData(ctx context.Context, in *v1.ListDictDataReq) (*v1.ListDictDataReply, error) {
	results, total := s.dictCase.DataListPage(ctx, pagination.NewPagination(pagination.WithPage(in.GetPage()), pagination.WithPageSize(in.GetPageSize()), pagination.WithQuery(map[string]interface{}{"dictType": in.GetDictType()})))
	return &v1.ListDictDataReply{
		Total: total,
		Items: convert.ArrayToAny(results, func(v *biz.DictData) *v1.DictData {
			return TransformDictData(v)
		}),
	}, nil
}

// CreateDictData 创建字典数据
func (s *AdminService) CreateDictData(ctx context.Context, in *v1.CreateDictDataReq) (*v1.CreateDictDataReply, error) {
	role, err := s.dictCase.DataCreate(ctx, &biz.DictData{
		Label:     in.GetLabel(),
		Value:     in.GetValue(),
		ColorType: in.GetColorType(),
		CssClass:  in.GetCssClass(),
		DictType:  in.GetDictType(),
		Remarks:   in.GetRemarks(),
		Sort:      in.GetSort(),
		State:     int32(in.GetState()),
	})
	if err != nil {
		return nil, v1.ErrorDictDataCreateFail("字典创建失败: %v", err.Error())
	}

	data, _ := anypb.New(&v1.Result{
		Id: uint64(role.ID),
	})
	return &v1.CreateDictDataReply{
		Type:    constant.HandleType_success.String(),
		Message: "创建成功",
		Result:  data,
	}, nil
}

// UpdateDictData 修改字典数据
func (s *AdminService) UpdateDictData(ctx context.Context, in *v1.UpdateDictDataReq) (*v1.UpdateDictDataReply, error) {
	v := in.GetData()
	err := s.dictCase.DataUpdate(ctx, &biz.DictData{
		ID:        uint(in.GetId()),
		Label:     v.GetLabel(),
		Value:     v.GetValue(),
		CssClass:  v.GetCssClass(),
		ColorType: v.GetColorType(),
		Remarks:   v.GetRemarks(),
		Sort:      v.GetSort(),
		State:     int32(v.GetState()),
	})
	if err != nil {
		return nil, v1.ErrorDictDataUpdateFail("字典修改失败: %v", err.Error())
	}
	return &v1.UpdateDictDataReply{
		Type:    constant.HandleType_success.String(),
		Message: "修改成功",
	}, nil
}

// UpdateDictDataState 修改字典数据-状态
func (s *AdminService) UpdateDictDataState(ctx context.Context, in *v1.UpdateDictDataStateReq) (*v1.UpdateDictDataStateReply, error) {
	v := in.GetData()
	err := s.dictCase.DataUpdateState(ctx, &biz.DictData{
		ID:    uint(in.GetId()),
		State: int32(v.GetState()),
	})
	if err != nil {
		return nil, v1.ErrorDictDataUpdateFail("字典数据状态修改失败: %v", err.Error())
	}
	return &v1.UpdateDictDataStateReply{
		Message: "修改成功",
		Type:    constant.HandleType_success.String(),
	}, nil
}

// GetDictData 获取字典数据
func (s *AdminService) GetDictData(ctx context.Context, in *v1.GetDictDataReq) (*v1.DictData, error) {
	v, err := s.dictCase.DataGetID(ctx, &biz.DictData{ID: uint(in.GetId())})
	if err != nil {
		return nil, v1.ErrorDictDataNotFound("字典未找到")
	}
	return TransformDictData(v), nil
}

// DeleteDictData 删除字典数据
func (s *AdminService) DeleteDictData(ctx context.Context, in *v1.DeleteDictDataReq) (*v1.DeleteDictDataReply, error) {
	if err := s.dictCase.DataDelete(ctx, &biz.DictData{ID: uint(in.GetId())}); err != nil {
		return nil, v1.ErrorDictDeleteFail("字典删除失败：%v", err)
	}
	return &v1.DeleteDictDataReply{
		Type:    constant.HandleType_success.String(),
		Message: "删除成功",
	}, nil
}
