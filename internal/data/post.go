package data

import (
	"context"

	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm/clause"
)

type PostRepo struct {
	data *Data
	log  *log.Helper
}

// NewPostRepo .
func NewPostRepo(logger log.Logger, data *Data) biz.PostRepo {
	return &PostRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *PostRepo) toModel(d *biz.Post) *SysPost {
	if d == nil {
		return nil
	}
	sysData := &SysPost{
		Code:    d.Code,
		Name:    d.Name,
		State:   d.State,
		Remarks: d.Remarks,
		Sort:    d.Sort,
	}
	sysData.ID = d.ID
	sysData.CreatedAt = d.CreatedAt
	sysData.CreatedAt = d.UpdatedAt
	return sysData
}

func (r *PostRepo) toBiz(d *SysPost) *biz.Post {
	if d == nil {
		return nil
	}
	return &biz.Post{
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
		ID:        d.ID,
		Code:      d.Code,
		Name:      d.Name,
		Sort:      d.Sort,
		State:     d.State,
		Remarks:   d.Remarks,
	}
}

func (r *PostRepo) Save(ctx context.Context, g *biz.Post) (*biz.Post, error) {
	d := r.toModel(g)
	sfId := r.data.sf.Generate()
	d.Code = sfId.String()
	d.DomainID = r.data.CtxDomainID(ctx)
	result := r.data.DBD(ctx).Create(d)
	return r.toBiz(d), result.Error
}

func (r *PostRepo) Update(ctx context.Context, g *biz.Post) (*biz.Post, error) {
	d := r.toModel(g)
	result := r.data.DBD(ctx).Model(d).Omit("Code").Updates(d)
	return r.toBiz(d), result.Error
}

func (r *PostRepo) FindByID(ctx context.Context, id uint) (*biz.Post, error) {
	post := SysPost{}
	result := r.data.DBD(ctx).Last(&post, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.toBiz(&post), nil
}

func (r *PostRepo) FindByName(ctx context.Context, s string) (*biz.Post, error) {
	post := SysPost{}
	result := r.data.DBD(ctx).Last(&post, "name = ?", s)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.toBiz(&post), nil
}

func (r *PostRepo) FindByCode(ctx context.Context, code string) (*biz.Post, error) {
	sysPost := &SysPost{}
	result := r.data.DBD(ctx).Last(sysPost, "code", code)
	return r.toBiz(sysPost), result.Error
}

func (r *PostRepo) ListByIDs(ctx context.Context, id ...uint) (posts []*biz.Post, err error) {
	db := r.data.DBD(ctx).Model(&SysPost{})
	sysPosts := []*SysPost{}

	err = db.Find(&sysPosts).Error
	if err != nil {
		return posts, err
	}
	for _, v := range sysPosts {
		posts = append(posts, r.toBiz(v))
	}
	return
}

func (r *PostRepo) ListByName(ctx context.Context, name string) ([]*biz.Post, error) {
	sysPosts, bizPosts := []*SysPost{}, []*biz.Post{}
	result := r.data.DBD(ctx).Find(&sysPosts, "name LIKE ?", "%"+name)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, v := range sysPosts {
		bizPosts = append(bizPosts, r.toBiz(v))
	}
	return bizPosts, nil
}

func (r *PostRepo) Delete(ctx context.Context, g *biz.Post) error {
	return r.data.DBD(ctx).Delete(r.toModel(g)).Error
}

func (r *PostRepo) ListAll(ctx context.Context) ([]*biz.Post, error) {
	return nil, nil
}

func (r *PostRepo) ListPage(ctx context.Context, paging *pagination.Pagination) (posts []*biz.Post, total int64) {
	db := r.data.DBD(ctx).Model(&SysPost{})
	sysPosts := []*SysPost{}

	// 查询条件
	if name, ok := paging.Query["name"].(string); ok {
		db = db.Where("name LIKE ?", name+"%")
	}
	if ids, ok := paging.Query["ids"]; ok {
		db = db.Where("id", ids)
	}

	// 排序

	if sortBy, ok := paging.OrderBy["sort"]; ok {
		db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: "sort"}, Desc: sortBy})
	}

	if orderBy, ok := paging.OrderBy["createdAt"]; ok {
		db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: "created_at"}, Desc: orderBy})
	}

	if idBy, ok := paging.OrderBy["id"]; ok {
		db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: "id"}, Desc: idBy})
	}

	if !paging.Nopaging {
		db = db.Count(&total).Offset(pagination.GetPageOffset(paging.Page, paging.PageSize))
	}

	result := db.Limit(int(paging.PageSize)).Find(&sysPosts)
	if result.Error != nil {
		return nil, 0
	}

	for _, v := range sysPosts {
		posts = append(posts, r.toBiz(v))
	}

	if paging.Nopaging {
		total = int64(len(posts))
	}

	return posts, total
}
