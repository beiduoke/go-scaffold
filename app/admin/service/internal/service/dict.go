package service

import (
	"context"

	v1 "github.com/beiduoke/go-scaffold/api/admin/service/v1"
	"github.com/beiduoke/go-scaffold/api/common"
	"github.com/beiduoke/go-scaffold/app/admin/service/internal/biz"
	"github.com/beiduoke/go-scaffold/pkg/constant"
	"github.com/beiduoke/go-scaffold/pkg/util/convert"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ v1.DictServiceServer = (*DictService)(nil)

// Service is a  service.
type DictService struct {
	v1.UnimplementedDictServiceServer
	log      *log.Helper
	dictCase *biz.DictUsecase
}

// NewService new a  service.
func NewDictService(
	logger log.Logger,
	dictCase *biz.DictUsecase,
) *DictService {
	l := log.NewHelper(log.With(logger, "module", "dict/service/admin-service"))
	return &DictService{
		log:      l,
		dictCase: dictCase,
	}
}

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
func (s *DictService) ListDict(ctx context.Context, in *v1.ListDictRequest) (*v1.ListDictResponse, error) {
	results, total := s.dictCase.ListPage(ctx, pagination.NewPagination(pagination.WithPage(in.GetPage()), pagination.WithPageSize(in.GetPageSize())))
	return &v1.ListDictResponse{
		Total: total,
		Items: convert.ArrayToAny(results, func(v *biz.Dict) *v1.Dict {
			return TransformDict(v)
		}),
	}, nil
}

// CreateDict 创建字典
func (s *DictService) CreateDict(ctx context.Context, in *v1.CreateDictRequest) (*v1.CreateDictResponse, error) {
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

	data, _ := anypb.New(&common.Result{
		Id: uint64(role.ID),
	})
	return &v1.CreateDictResponse{
		Type:    constant.HandleType_success.String(),
		Message: "创建成功",
		Result:  data,
	}, nil
}

// UpdateDict 修改字典
func (s *DictService) UpdateDict(ctx context.Context, in *v1.UpdateDictRequest) (*v1.UpdateDictResponse, error) {
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
	return &v1.UpdateDictResponse{
		Type:    constant.HandleType_success.String(),
		Message: "修改成功",
	}, nil
}

// UpdateDictState 修改字典-状态
func (s *DictService) UpdateDictState(ctx context.Context, in *v1.UpdateDictStateRequest) (*v1.UpdateDictStateResponse, error) {
	v := in.GetData()
	err := s.dictCase.UpdateState(ctx, &biz.Dict{
		ID:    uint(in.GetId()),
		State: int32(v.GetState()),
	})
	if err != nil {
		return nil, v1.ErrorDictUpdateFail("租户状态修改失败: %v", err.Error())
	}
	return &v1.UpdateDictStateResponse{
		Message: "修改成功",
		Type:    constant.HandleType_success.String(),
	}, nil
}

// GetDict 获取字典
func (s *DictService) GetDict(ctx context.Context, in *v1.GetDictRequest) (*v1.Dict, error) {
	role, err := s.dictCase.GetID(ctx, &biz.Dict{ID: uint(in.GetId())})
	if err != nil {
		return nil, v1.ErrorDictNotFound("字典未找到")
	}
	return TransformDict(role), nil
}

// DeleteDict 删除字典
func (s *DictService) DeleteDict(ctx context.Context, in *v1.DeleteDictRequest) (*v1.DeleteDictResponse, error) {
	if err := s.dictCase.Delete(ctx, &biz.Dict{ID: uint(in.GetId())}); err != nil {
		return nil, v1.ErrorDictDeleteFail("字典删除失败：%v", err)
	}
	return &v1.DeleteDictResponse{
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
func (s *DictService) ListDictData(ctx context.Context, in *v1.ListDictDataRequest) (*v1.ListDictDataResponse, error) {
	results, total := s.dictCase.DataListPage(ctx, pagination.NewPagination(pagination.WithPage(in.GetPage()), pagination.WithPageSize(in.GetPageSize()), pagination.WithQuery(map[string]interface{}{"dictType": in.GetDictType()})))
	return &v1.ListDictDataResponse{
		Total: total,
		Items: convert.ArrayToAny(results, func(v *biz.DictData) *v1.DictData {
			return TransformDictData(v)
		}),
	}, nil
}

// CreateDictData 创建字典数据
func (s *DictService) CreateDictData(ctx context.Context, in *v1.CreateDictDataRequest) (*v1.CreateDictDataResponse, error) {
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

	data, _ := anypb.New(&common.Result{
		Id: uint64(role.ID),
	})
	return &v1.CreateDictDataResponse{
		Type:    constant.HandleType_success.String(),
		Message: "创建成功",
		Result:  data,
	}, nil
}

// UpdateDictData 修改字典数据
func (s *DictService) UpdateDictData(ctx context.Context, in *v1.UpdateDictDataRequest) (*v1.UpdateDictDataResponse, error) {
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
	return &v1.UpdateDictDataResponse{
		Type:    constant.HandleType_success.String(),
		Message: "修改成功",
	}, nil
}

// UpdateDictDataState 修改字典数据-状态
func (s *DictService) UpdateDictDataState(ctx context.Context, in *v1.UpdateDictDataStateRequest) (*v1.UpdateDictDataStateResponse, error) {
	v := in.GetData()
	err := s.dictCase.DataUpdateState(ctx, &biz.DictData{
		ID:    uint(in.GetId()),
		State: int32(v.GetState()),
	})
	if err != nil {
		return nil, v1.ErrorDictDataUpdateFail("字典数据状态修改失败: %v", err.Error())
	}
	return &v1.UpdateDictDataStateResponse{
		Message: "修改成功",
		Type:    constant.HandleType_success.String(),
	}, nil
}

// GetDictData 获取字典数据
func (s *DictService) GetDictData(ctx context.Context, in *v1.GetDictDataRequest) (*v1.DictData, error) {
	v, err := s.dictCase.DataGetID(ctx, &biz.DictData{ID: uint(in.GetId())})
	if err != nil {
		return nil, v1.ErrorDictDataNotFound("字典未找到")
	}
	return TransformDictData(v), nil
}

// DeleteDictData 删除字典数据
func (s *DictService) DeleteDictData(ctx context.Context, in *v1.DeleteDictDataRequest) (*v1.DeleteDictDataResponse, error) {
	if err := s.dictCase.DataDelete(ctx, &biz.DictData{ID: uint(in.GetId())}); err != nil {
		return nil, v1.ErrorDictDeleteFail("字典删除失败：%v", err)
	}
	return &v1.DeleteDictDataResponse{
		Type:    constant.HandleType_success.String(),
		Message: "删除成功",
	}, nil
}
