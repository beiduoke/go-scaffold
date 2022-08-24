package biz

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"time"

	pb "github.com/beiduoke/go-scaffold/api/protobuf"
	"github.com/beiduoke/go-scaffold/internal/conf"
	myAuthz "github.com/beiduoke/go-scaffold/internal/pkg/authz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/zzsds/go-tools/pkg/password"
)

// User is a User model.
type User struct {
	CreatedAt            time.Time
	UpdatedAt            time.Time
	ID                   uint
	Name                 string
	NickName             string
	RealName             string
	Password             string
	Birthday             *time.Time
	Gender               int32
	Mobile               string
	Email                string
	State                int32
	Domains              []Domain
	Authorities          []Authority
	DomainAuthorityUsers []DomainAuthorityUser
}

// UserRepo is a Greater repo.
type UserRepo interface {
	// 基准操作
	Save(context.Context, *User) (*User, error)
	Update(context.Context, *User) (*User, error)
	FindByID(context.Context, int64) (*User, error)
	ListAll(context.Context) ([]*User, error)
	// 自定义操作
	FindByName(context.Context, string) (*User, error)
	FindByMobile(context.Context, string) (*User, error)
	FindByEmail(context.Context, string) (*User, error)
	ListByName(context.Context, string) ([]*User, error)
	ListByMobile(context.Context, string) ([]*User, error)
	ListByEmail(context.Context, string) ([]*User, error)
}

// UserUsecase is a User usecase.
type UserUsecase struct {
	ac         *conf.Auth
	repo       UserRepo
	domainRepo DomainRepo
	log        *log.Helper
	tm         Transaction
}

// NewUserUsecase new a User usecase.
func NewUserUsecase(ac *conf.Auth, repo UserRepo, tm Transaction, logger log.Logger, domainRepo DomainRepo) *UserUsecase {
	return &UserUsecase{ac: ac, repo: repo, tm: tm, log: log.NewHelper(logger), domainRepo: domainRepo}
}

func (uc *UserUsecase) GenerateToken(g *User) (token string, expiresAt time.Time) {
	authorityId := []string{}
	for _, v := range g.Authorities {
		authorityId = append(authorityId, strconv.Itoa(int(v.ID)))
	}
	expiresAt = time.Now().Add(time.Hour * 24)
	securityUser := myAuthz.NewSecurityUserData(
		myAuthz.WithID(strconv.Itoa(int(g.ID))),
		myAuthz.WithExpiresAt(expiresAt.Unix()),
		myAuthz.WithDomain(strconv.Itoa(int(g.Domains[0].ID))),
		myAuthz.WithAuthorityId(strings.Join(authorityId, "-")),
	)
	token = securityUser.CreateAccessJwtToken([]byte(uc.ac.ApiKey))
	return
}

// NamePasswordLogin 用户密码登录
func (uc *UserUsecase) NamePasswordLogin(ctx context.Context, domainId string, g *User) (*User, error) {
	domain, err := uc.domainRepo.FindByDomainID(ctx, domainId)
	if err != nil {
		return nil, errors.New("Domain查询失败")
	}
	u, err := uc.repo.FindByName(ctx, g.Name)
	if err != nil {
		return nil, err
	}
	err = password.Verify(u.Password, g.Password)
	if err != nil {
		return nil, err
	}
	u.Domains = []Domain{*domain}
	return u, nil
}

// MobileSmsLogin 手机验证码登录
func (uc *UserUsecase) MobileSmsLogin(ctx context.Context, domainId string, g *User) (*User, error) {
	uc.log.WithContext(ctx).Infof("MobileSmsLogin: %v", g)
	return uc.repo.FindByMobile(ctx, g.Mobile)
}

// NamePasswordRegister 用户密码注册
func (uc *UserUsecase) NamePasswordRegister(ctx context.Context, domainId string, g *User) (*User, error) {
	uc.log.WithContext(ctx).Infof("NamePasswordRegister: %v", g.Name)
	domain, err := uc.domainRepo.FindByDomainID(ctx, domainId)
	if err != nil {
		return nil, errors.New("Domain查询失败")
	}
	user, _ := uc.repo.FindByName(ctx, g.Name)
	if user != nil && user.Name != "" {
		return nil, errors.New("用户已注册")
	}
	password, err := password.Encryption(g.Password)
	if err != nil {
		return nil, errors.New("密码加密失败")
	}
	g.Password = password
	if g.State <= 0 {
		g.State = int32(pb.UserState_ACTIVE)
	}

	err = uc.tm.InTx(ctx, func(ctx context.Context) error {
		g, err = uc.repo.Save(ctx, g)
		if err != nil {
			return err
		}
		_, err := uc.domainRepo.AuthorityUserSave(ctx, &DomainAuthorityUser{
			UserID:      g.ID,
			AuthorityID: domain.DefaultAuthorityID,
			DomainID:    domain.ID,
		})
		return err
	})
	return g, err
}
