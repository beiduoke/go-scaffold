package biz

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/beiduoke/go-scaffold/internal/conf"
	myAuthz "github.com/beiduoke/go-scaffold/internal/pkg/authz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/zzsds/go-tools/pkg/password"
)

// Auth 多登录方式进行接口封装
type Auth interface {
	Login() (AuthClaims, error)
	Register() error
	Logout() error
}

type AuthClaims struct {
	ID          uint
	Domain      uint
	Authorities []uint
	ExpiresAt   *time.Time
	Token       string
}

// AuthUsecase is a User usecase.
type AuthUsecase struct {
	ac         *conf.Auth
	log        *log.Helper
	tm         Transaction
	userRepo   UserRepo
	domainRepo DomainRepo
}

// NewAuthUsecase new a User usecase.
func NewAuthUsecase(ac *conf.Auth, userRepo UserRepo, tm Transaction, logger log.Logger, domainRepo DomainRepo) *AuthUsecase {
	return &AuthUsecase{ac: ac, userRepo: userRepo, tm: tm, log: log.NewHelper(logger), domainRepo: domainRepo}
}

func (ac *AuthUsecase) GetToken(claims *AuthClaims) error {
	authorities := []string{}
	for _, v := range claims.Authorities {
		authorities = append(authorities, strconv.Itoa(int(v)))
	}
	expiresAt := time.Now().Add(time.Hour * 24)
	if claims.ExpiresAt != nil {
		expiresAt = *claims.ExpiresAt
	}
	securityUser := myAuthz.NewSecurityUserData(
		myAuthz.WithID(strconv.Itoa(int(claims.ID))),
		myAuthz.WithDomain(strconv.Itoa(int(claims.Domain))),
		myAuthz.WithAuthority(strings.Join(authorities, ",")),
		myAuthz.WithExpires(expiresAt),
	)
	claims.Token = securityUser.CreateAccessJwtToken([]byte(ac.ac.ApiKey))
	if claims.Token == "" {
		return errors.New("token生成失败")
	}
	return nil
}

// LoginNamePassword 登录-用户密码
func (ac *AuthUsecase) LoginNamePassword(ctx context.Context, domainId string, g *User) (*AuthClaims, error) {
	domain, err := ac.domainRepo.FindByDomainID(ctx, domainId)
	if err != nil {
		return nil, errors.New("domain查询失败")
	}
	u, err := ac.userRepo.FindByName(ctx, g.Name)
	if err != nil {
		return nil, err
	}
	err = password.Verify(u.Password, g.Password)
	if err != nil {
		return nil, err
	}
	authorities := ac.domainRepo.FindAuthoritiesForUserInDomain(ctx, u.ID, domain.ID)
	if len(authorities) == 0 || err != nil {
		return nil, errors.New("权限未配置")
	}

	// 生成token
	authClaims := &AuthClaims{
		ID:          u.ID,
		Domain:      domain.ID,
		Authorities: []uint{},
		ExpiresAt:   nil,
	}
	// 配置过期时间
	if ac.ac.GetExpiresTime() != nil {
		expiresAt := time.Now().Add(ac.ac.ExpiresTime.AsDuration())
		authClaims.ExpiresAt = &expiresAt
	}
	// 组装权限
	for _, v := range authorities {
		authClaims.Authorities = append(authClaims.Authorities, v.ID)
	}
	if err := ac.GetToken(authClaims); err != nil {
		return nil, err
	}

	if err := ac.userRepo.SetTokenCache(ctx, *authClaims); err != nil {
		ac.log.Errorf("token 缓存失败 %v", err)
	}

	return authClaims, nil
}

// LoginMobileSms 登录-手机验证码
func (ac *AuthUsecase) LoginMobileSms(ctx context.Context, domainId string, g *User) (*User, error) {
	ac.log.WithContext(ctx).Infof("mobileSmsLogin: %v", g)
	return ac.userRepo.FindByMobile(ctx, g.Mobile)
}

// RegisterNamePassword 注册-用户密码
func (ac *AuthUsecase) RegisterNamePassword(ctx context.Context, domainId string, g *User) (*User, error) {
	ac.log.WithContext(ctx).Infof("NamePasswordRegister: %v", g.Name)
	domain, err := ac.domainRepo.FindByDomainID(ctx, domainId)
	if err != nil {
		return nil, errors.New("领域查询失败")
	}
	g.Domains = []*Domain{domain}
	return (&UserUsecase{ac.log, ac.ac, ac.tm, ac.userRepo, ac.domainRepo}).Create(ctx, g)
}
