package data

import (
	"context"

	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/beiduoke/go-scaffold/pkg/util/convert"
	"github.com/beiduoke/go-scaffold/pkg/util/pagination"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm/clause"
)

type UserRepo struct {
	data   *Data
	log    *log.Helper
	domain DomainRepo
	role   RoleRepo
}

// NewUserRepo .
func NewUserRepo(logger log.Logger, data *Data) biz.UserRepo {
	return &UserRepo{
		data: data,
		log:  log.NewHelper(logger),
		role: RoleRepo{},
	}
}

func (r *UserRepo) toModel(d *biz.User) *SysUser {
	if d == nil {
		return nil
	}
	roles := []SysRole{}
	for _, v := range d.Roles {
		roles = append(roles, *r.role.toModel(v))
	}
	return &SysUser{
		DomainModel: DomainModel{
			ID:        d.ID,
			CreatedAt: d.CreatedAt,
			UpdatedAt: d.UpdatedAt,
		},
		Name:     d.Name,
		Avatar:   d.Avatar,
		NickName: d.NickName,
		RealName: d.RealName,
		Password: d.Password,
		Birthday: d.Birthday,
		Gender:   d.Gender,
		Phone:    d.Phone,
		Email:    d.Email,
		State:    d.State,
		Roles:    roles,
	}
}

func (r *UserRepo) toBiz(d *SysUser) *biz.User {
	if d == nil {
		return nil
	}
	roles := []*biz.Role{}
	for _, v := range d.Roles {
		roles = append(roles, r.role.toBiz(&v))
	}
	return &biz.User{
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
		ID:        d.ID,
		Avatar:    d.Avatar,
		Name:      d.Name,
		NickName:  d.NickName,
		RealName:  d.RealName,
		Password:  d.Password,
		Birthday:  d.Birthday,
		Gender:    d.Gender,
		Phone:     d.Phone,
		Email:     d.Email,
		State:     d.State,
		Roles:     roles,
	}
}

func (r *UserRepo) Save(ctx context.Context, g *biz.User) (*biz.User, error) {
	d := r.toModel(g)
	d.DomainID = r.data.DomainID(ctx)
	result := r.data.DB(ctx).Omit(clause.Associations).Create(d).Error
	return r.toBiz(d), result
}

func (r *UserRepo) Update(ctx context.Context, g *biz.User) (*biz.User, error) {
	d := r.toModel(g)
	result := r.data.DBD(ctx).Model(d).Select("*").Omit("CreatedAt").Updates(d)
	return r.toBiz(d), result.Error
}

func (r *UserRepo) FindByID(ctx context.Context, id uint) (*biz.User, error) {
	user := SysUser{}
	result := r.data.DBD(ctx).Last(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return r.toBiz(&user), nil
}

func (r *UserRepo) ListAll(ctx context.Context) ([]*biz.User, error) {
	return nil, nil
}

func (r *UserRepo) Delete(ctx context.Context, g *biz.User) error {
	return r.data.DBD(ctx).Delete(r.toModel(g)).Error
}

func (r *UserRepo) ListPage(ctx context.Context, handler pagination.PaginationHandler) (users []*biz.User, total int64) {
	db := r.data.DBD(ctx).Model(&SysUser{})
	sysUsers := []*SysUser{}
	// 查询条件
	for _, v := range handler.GetConditions() {
		db = db.Where(v.Query, v.Args...)
	}
	// 排序
	for _, v := range handler.GetOrders() {
		db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: v.Column}, Desc: v.Desc})
	}
	if !handler.GetNopaging() {
		db = db.Count(&total).Offset(handler.GetPageOffset())
	}
	result := db.Limit(int(handler.GetPageSize())).Find(&sysUsers)
	if result.Error != nil {
		return nil, 0
	}

	for _, v := range sysUsers {
		users = append(users, r.toBiz(v))
	}

	if handler.GetNopaging() {
		total = int64(len(users))
	}

	return users, total
}

func (r *UserRepo) FindByName(ctx context.Context, s string) (*biz.User, error) {
	user := SysUser{}
	result := r.data.DBD(ctx).Debug().Last(&user, "name = ?", s)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.toBiz(&user), nil
}

func (r *UserRepo) FindByPhone(ctx context.Context, s string) (*biz.User, error) {
	user := SysUser{}
	result := r.data.DBD(ctx).Last(&user, "phone = ?", s)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.toBiz(&user), nil
}

func (r *UserRepo) FindByEmail(ctx context.Context, s string) (*biz.User, error) {
	user := SysUser{}
	result := r.data.DBD(ctx).Last(&user, "email = ?", s)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.toBiz(&user), nil
}

func (r *UserRepo) ListByName(ctx context.Context, s string) ([]*biz.User, error) {
	sysUsers, bizUsers := []*SysUser{}, []*biz.User{}
	result := r.data.DBD(ctx).Find(&sysUsers, "name LIKE ?", "%"+s)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, v := range sysUsers {
		bizUsers = append(bizUsers, r.toBiz(v))
	}
	return bizUsers, nil
}

func (r *UserRepo) ListByPhone(ctx context.Context, s string) ([]*biz.User, error) {
	sysUsers, bizUsers := []*SysUser{}, []*biz.User{}
	result := r.data.DBD(ctx).Find(&sysUsers, "phone LIKE ?", "%"+s)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, v := range sysUsers {
		bizUsers = append(bizUsers, r.toBiz(v))
	}
	return bizUsers, nil
}

func (r *UserRepo) ListByEmail(ctx context.Context, s string) ([]*biz.User, error) {
	sysUsers, bizUsers := []*SysUser{}, []*biz.User{}
	result := r.data.DBD(ctx).Find(&sysUsers, "email LIKE ?", "%"+s)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, v := range sysUsers {
		bizUsers = append(bizUsers, r.toBiz(v))
	}
	return bizUsers, nil
}

// HandleDomainRole 绑定权限
func (r *UserRepo) HandleRole(ctx context.Context, g *biz.User) error {
	domainId := r.data.Domain(ctx)
	rules := make([][]string, 0, len(g.Roles))
	for _, v := range g.Roles {
		rules = append(rules, []string{convert.UnitToString(g.ID), convert.UnitToString(v.ID), domainId, "0"})
		// if _, err := r.data.enforcer.AddRoleForUserInDomain(convert.UnitToString(g.ID), convert.UnitToString(v.ID), domainId); err != nil {
		// 	r.log.Errorf("领域权限绑定失败 %v", err)
		// }
	}
	_, err := r.data.enforcer.AddGroupingPolicies(rules)
	// r.log.Debugf("策略添加 %t %v", success, err)
	return err
}
