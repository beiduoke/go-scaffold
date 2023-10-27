package biz

import (
	"context"
	"fmt"
	"time"

	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
)

// Post is a Post model.
type Post struct {
	CreatedAt time.Time `json:"createdAt,omitempty" form:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" form:"updatedAt,omitempty"`
	ID        uint      `json:"id,omitempty" form:"id,omitempty"`
	Code      string    `json:"code,omitempty" form:"code,omitempty"`
	Remarks   string    `json:"remarks,omitempty" form:"remarks,omitempty"`
	Name      string    `json:"name,omitempty" form:"name,omitempty"`
	Sort      int32     `json:"sort,omitempty" form:"sort,omitempty"`
	State     int32     `json:"state,omitempty" form:"state,omitempty"`
}

// PostRepo is a Greater repo.
type PostRepo interface {
	Save(context.Context, *Post) (*Post, error)
	Update(context.Context, *Post) (*Post, error)
	FindByID(context.Context, uint) (*Post, error)
	FindByCode(context.Context, string) (*Post, error)
	FindByName(context.Context, string) (*Post, error)
	ListByIDs(context.Context, ...uint) ([]*Post, error)
	ListByName(context.Context, string) ([]*Post, error)
	Delete(context.Context, *Post) error
	ListAll(context.Context) ([]*Post, error)
	ListPage(context.Context, *pagination.Pagination) ([]*Post, int64)
}

// PostUsecase is a Post usecase.
type PostUsecase struct {
	biz  *Biz
	log  *log.Helper
	repo PostRepo
}

// NewPostUsecase new a Post usecase.
func NewPostUsecase(logger log.Logger, biz *Biz, repo PostRepo) *PostUsecase {
	return &PostUsecase{biz: biz, log: log.NewHelper(logger), repo: repo}
}

// Create creates a Post, and returns the new Post.
func (uc *PostUsecase) Create(ctx context.Context, g *Post) (*Post, error) {
	uc.log.WithContext(ctx).Debugf("Create: %v", g.Name)
	err := uc.biz.tm.InTx(ctx, func(ctx context.Context) error {
		_, err := uc.repo.Save(ctx, g)
		return err
	})
	return g, err
}

// ListByIDs 获取指定岗位ID集合
func (uc *PostUsecase) ListByIDs(ctx context.Context, id ...uint) (roles []*Post, err error) {
	// roles, _ = uc.repo.ListPage(ctx, noop.NewPagination(noop.WithNopaging(), noop.WithCondition("id in ?", id)))
	return
}

// Update 修改岗位
func (uc *PostUsecase) Update(ctx context.Context, g *Post) error {
	uc.log.WithContext(ctx).Debugf("UpdatePost: %v", g)

	post, _ := uc.repo.FindByID(ctx, g.ID)
	if post == nil {
		return errors.New("岗位未创建")
	}

	if post.Name != g.Name && g.Name != "" {
		name, _ := uc.repo.FindByName(ctx, g.Name)
		if name != nil {
			return errors.New("岗位名已存在")
		}
	}

	fmt.Printf("%s", g.Remarks)
	_, err := uc.repo.Update(ctx, g)
	return err
}

// UpdateState 修改岗位状态
func (uc *PostUsecase) UpdateState(ctx context.Context, g *Post) error {
	uc.log.WithContext(ctx).Debugf("UpdatePostState: %v", g)

	post, _ := uc.repo.FindByID(ctx, g.ID)
	if post == nil {
		return errors.New("岗位不存在")
	}

	post.State = g.State
	_, err := uc.repo.Update(ctx, post)
	return err
}

// List 岗位列表全部
func (uc *PostUsecase) ListAll(ctx context.Context) ([]*Post, int64) {
	uc.log.WithContext(ctx).Debugf("PostList")
	return uc.repo.ListPage(ctx, &pagination.Pagination{Nopaging: true, OrderBy: map[string]bool{"sort": true}})
}

// List 岗位列表分页
func (uc *PostUsecase) ListPage(ctx context.Context, paging *pagination.Pagination) ([]*Post, int64) {
	uc.log.WithContext(ctx).Debugf("PostPage")
	return uc.repo.ListPage(ctx, paging)
}

// GetID 根据角色ID岗位
func (uc *PostUsecase) GetID(ctx context.Context, g *Post) (*Post, error) {
	uc.log.WithContext(ctx).Debugf("GetPostID: %v", g)
	return uc.repo.FindByID(ctx, g.ID)
}

// Delete 根据角色ID删除岗位
func (uc *PostUsecase) Delete(ctx context.Context, g *Post) error {
	uc.log.WithContext(ctx).Debugf("DeletePost: %v", g)
	return uc.biz.tm.InTx(ctx, func(ctx context.Context) error {
		return uc.repo.Delete(ctx, g)
	})
}
