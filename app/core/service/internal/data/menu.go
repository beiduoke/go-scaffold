package data

import (
	"context"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/beiduoke/go-scaffold/api/common/pagination"
	v1 "github.com/beiduoke/go-scaffold/api/core/service/v1"
	"github.com/beiduoke/go-scaffold/app/core/service/internal/data/ent"
	"github.com/beiduoke/go-scaffold/app/core/service/internal/data/ent/user"
	"github.com/beiduoke/go-scaffold/pkg/util/convert"
	"github.com/beiduoke/go-scaffold/pkg/util/trans"
	"github.com/go-kratos/kratos/v2/log"
	entgo "github.com/tx7do/go-utils/entgo/query"
)

func (r *MenuRepo) toProto(in *ent.Menu) *v1.Menu {
	if in == nil {
		return nil
	}
	return &v1.Menu{
		CreatedAt:         convert.TimeValueToString(in.CreatedAt, time.DateTime),
		UpdatedAt:         convert.TimeValueToString(in.UpdatedAt, time.DateTime),
		Id:                in.ID,
		Name:              in.Name,
		Sort:              in.Sort,
		State:             in.State,
		Remark:            in.Remark,
		ParentId:          in.ParentID,
		Title:             in.Title,
		Type:              in.Type,
		Path:              in.Path,
		Component:         in.Component,
		Icon:              in.Icon,
		IsExt:             in.IsExt,
		ExtUrl:            in.ExtURL,
		Permissions:       trans.String(strings.Join(in.Permissions, ",")),
		Redirect:          in.Redirect,
		CurrentActiveMenu: in.CurrentActiveMenu,
		KeepAlive:         in.KeepAlive,
		Visible:           in.Visible,
		HideTab:           in.HideTab,
		HideMenu:          in.HideMenu,
		HideBreadcrumb:    in.HideBreadcrumb,
	}
}

type MenuRepo struct {
	data *Data
	log  *log.Helper
}

// NewMenuRepo .
func NewMenuRepo(data *Data, logger log.Logger) *MenuRepo {
	return &MenuRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *MenuRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Menu.Query()
	if len(whereCond) > 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
	}

	return count, err
}

func (r *MenuRepo) CreateMenu(ctx context.Context, req *v1.CreateMenuRequest) (*v1.CreateMenuResponse, error) {
	builder := r.data.db.Menu.Create().SetCreatedAt(time.Now())
	builder = builder.SetName(req.Menu.GetName()).
		SetNillableState(req.Menu.State).
		SetNillableRemark(req.Menu.Remark).
		SetNillableSort(req.Menu.Sort).
		SetNillableComponent(req.Menu.Component).
		SetNillableCurrentActiveMenu(req.Menu.CurrentActiveMenu).
		SetNillableExtURL(req.Menu.ExtUrl).
		SetNillableHideBreadcrumb(req.Menu.HideBreadcrumb).
		SetNillableHideMenu(req.Menu.HideMenu).
		SetNillableHideTab(req.Menu.HideTab).
		SetNillableIcon(req.Menu.Icon).
		SetNillablePath(req.Menu.Path).
		SetNillableRedirect(req.Menu.Redirect).
		SetNillableParentID(req.Menu.ParentId).
		SetNillableTitle(req.Menu.Title).
		SetNillableKeepAlive(req.Menu.KeepAlive)

	_, err := builder.Save(ctx)
	if err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return nil, err
	}

	return &v1.CreateMenuResponse{}, err
}

func (r *MenuRepo) UpdateMenu(ctx context.Context, req *v1.UpdateMenuRequest) (*v1.UpdateMenuResponse, error) {
	builder := r.data.db.Menu.UpdateOneID(req.Id)
	builder = builder.SetName(req.Menu.GetName()).
		SetNillableState(req.Menu.State).
		SetNillableRemark(req.Menu.Remark).
		SetNillableSort(req.Menu.Sort).
		SetNillableComponent(req.Menu.Component).
		SetNillableCurrentActiveMenu(req.Menu.CurrentActiveMenu).
		SetNillableExtURL(req.Menu.ExtUrl).
		SetNillableHideBreadcrumb(req.Menu.HideBreadcrumb).
		SetNillableHideMenu(req.Menu.HideMenu).
		SetNillableHideTab(req.Menu.HideTab).
		SetNillableIcon(req.Menu.Icon).
		SetNillablePath(req.Menu.Path).
		SetNillableRedirect(req.Menu.Redirect).
		SetNillableParentID(req.Menu.ParentId).
		SetNillableTitle(req.Menu.Title).
		SetNillableKeepAlive(req.Menu.KeepAlive)

	_, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateMenuResponse{}, err
}

func (r *MenuRepo) DeleteMenu(ctx context.Context, req *v1.DeleteMenuRequest) (*v1.DeleteMenuResponse, error) {
	err := r.data.db.Menu.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteMenuResponse{}, nil
}

func (r *MenuRepo) GetMenu(ctx context.Context, req *v1.GetMenuRequest) (*v1.Menu, error) {
	ret, err := r.data.db.Menu.Get(ctx, req.GetId())
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}
	return r.toProto(ret), err
}

func (r *MenuRepo) ListMenu(ctx context.Context, req *pagination.PagingRequest) (*v1.ListMenuResponse, error) {
	builder := r.data.db.Debug().Menu.Query()
	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), user.FieldCreatedAt,
		req.GetFieldMask().GetPaths(),
	)
	if err != nil {
		r.log.Errorf("解析条件发生错误[%s]", err.Error())
		return nil, err
	}

	if querySelectors != nil {
		builder.Modify(querySelectors...)
	}

	results, err := builder.All(ctx)
	if err != nil {
		r.log.Errorf("query list failed: %s", err.Error())
		return nil, err
	}

	items := make([]*v1.Menu, 0, len(results))
	for _, res := range results {
		items = append(items, r.toProto(res))
	}
	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &v1.ListMenuResponse{
		Total: int32(count),
		Items: items,
	}, nil
}
