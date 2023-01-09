package biz

import (
	"context"
	"time"

	pb "github.com/beiduoke/go-scaffold/api/protobuf"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/imdario/mergo"
	"github.com/pkg/errors"
)

// Post is a Post model.
type Post struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	ID        uint
	Code      string
	Remarks   string
	Name      string
	Sort      int32
	State     int32
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
	ListPage(context.Context, pagination.PaginationHandler) ([]*Post, int64)
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
	uc.log.WithContext(ctx).Infof("Create: %v", g.Name)
	err := uc.biz.tm.InTx(ctx, func(ctx context.Context) error {
		_, err := uc.repo.Save(ctx, g)
		return err
	})
	return g, err
}

// ListByIDs 获取指定职位ID集合
func (uc *PostUsecase) ListByIDs(ctx context.Context, id ...uint) (roles []*Post, err error) {
	roles, _ = uc.repo.ListPage(ctx, pagination.NewPagination(pagination.WithNopaging(), pagination.WithCondition("id in ?", id)))
	return
}

// Update 修改职位
func (uc *PostUsecase) Update(ctx context.Context, g *Post) error {
	uc.log.WithContext(ctx).Infof("UpdatePost: %v", g)

	post, _ := uc.repo.FindByID(ctx, g.ID)
	if post == nil {
		return errors.New("职位未注册")
	}

	if post.Name != g.Name && g.Name != "" {
		name, _ := uc.repo.FindByName(ctx, g.Name)
		if name != nil {
			return errors.New("职位名已存在")
		}
	}

	if g.State <= 0 {
		g.State = int32(pb.PostState_POST_STATE_ACTIVE)
	}

	// 新数据合并到源数据
	if err := mergo.Merge(post, *g, mergo.WithOverride); err != nil {
		return errors.Errorf("数据合并失败：%v", err)
	}

	_, err := uc.repo.Update(ctx, post)
	return err
}

// UpdateState 修改职位状态
func (uc *PostUsecase) UpdateState(ctx context.Context, g *Post) error {
	uc.log.WithContext(ctx).Infof("UpdatePostState: %v", g)

	post, _ := uc.repo.FindByID(ctx, g.ID)
	if post == nil {
		return errors.New("职位不存在")
	}

	if g.State <= 0 {
		g.State = int32(pb.PostState_POST_STATE_ACTIVE)
	}

	post.State = g.State
	_, err := uc.repo.Update(ctx, post)
	return err
}

// List 职位列表全部
func (uc *PostUsecase) ListAll(ctx context.Context) ([]*Post, int64) {
	uc.log.WithContext(ctx).Infof("PostList")
	return uc.repo.ListPage(ctx, pagination.NewPagination())
}

// List 职位列表分页
func (uc *PostUsecase) ListPage(ctx context.Context, pageNum, pageSize int32, query map[string]string, order map[string]bool) ([]*Post, int64) {
	uc.log.WithContext(ctx).Infof("PostPage")
	conditions := []pagination.Condition{{Query: "id > 0"}}
	for k, v := range query {
		conditions = append(conditions, pagination.Condition{Query: k, Args: []interface{}{v}})
	}
	orders := []pagination.Order{}
	for k, v := range order {
		orders = append(orders, pagination.Order{Column: k, Desc: v})
	}

	page := pagination.NewPagination(
		pagination.WithPageNum(pageNum),
		pagination.WithPageSize(pageSize),
		pagination.WithConditions(conditions...),
		pagination.WithOrders(orders...),
	)
	return uc.repo.ListPage(ctx, page)
}

// GetID 根据角色ID职位
func (uc *PostUsecase) GetID(ctx context.Context, g *Post) (*Post, error) {
	uc.log.WithContext(ctx).Infof("GetPostID: %v", g)
	return uc.repo.FindByID(ctx, g.ID)
}

// Delete 根据角色ID删除职位
func (uc *PostUsecase) Delete(ctx context.Context, g *Post) error {
	uc.log.WithContext(ctx).Infof("DeletePost: %v", g)
	return uc.biz.tm.InTx(ctx, func(ctx context.Context) error {
		return uc.repo.Delete(ctx, g)
	})
}
