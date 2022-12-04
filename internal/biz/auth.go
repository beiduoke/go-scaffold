package biz

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"

	"github.com/beiduoke/go-scaffold/internal/conf"
	myAuthz "github.com/beiduoke/go-scaffold/internal/pkg/authz"
	"github.com/beiduoke/go-scaffold/pkg/util/convert"
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
	biz *Biz
	ac  *conf.Auth
	log *log.Helper
}

// NewAuthUsecase new a User usecase.
func NewAuthUsecase(logger log.Logger, biz *Biz, ac *conf.Auth) *AuthUsecase {
	return &AuthUsecase{log: log.NewHelper(logger), ac: ac, biz: biz}
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
		myAuthz.WithUser(strconv.Itoa(int(claims.ID))),
		myAuthz.WithDomain(strconv.Itoa(int(claims.Domain))),
		myAuthz.WithSubject(strings.Join(authorities, ",")),
		myAuthz.WithExpires(expiresAt),
	)
	claims.Token = securityUser.CreateAccessJwtToken([]byte(ac.ac.ApiKey))
	if claims.Token == "" {
		return errors.New("token生成失败")
	}
	return nil
}

// PassLogin 登录-密码登录
func (ac *AuthUsecase) PassLogin(ctx context.Context, g *User) (*AuthClaims, error) {
	u, err := ac.biz.userRepo.FindByName(ctx, g.Name)
	if err != nil {
		return nil, err
	}
	err = password.Verify(u.Password, g.Password)
	if err != nil {
		return nil, err
	}

	userUsecase := UserUsecase{log: ac.log, biz: ac.biz, ac: ac.ac}
	domain, err := userUsecase.GetLastUseDomain(ctx, u)
	if err != nil {
		return nil, err
	}
	// 默认过期时间 24小时
	var expiresAt = time.Now().Add(time.Hour * 24)
	// 配置过期时间
	if ac.ac.GetExpiresTime() != nil {
		expiresAt = time.Now().Add(ac.ac.ExpiresTime.AsDuration())
	}
	// 生成token
	authClaims := &AuthClaims{
		ID:          u.ID,
		ExpiresAt:   &expiresAt,
		Domain:      domain.ID,
		Authorities: []uint{domain.DefaultAuthorityID},
	}

	if err := ac.GetToken(authClaims); err != nil {
		return nil, err
	}

	if err := ac.biz.userRepo.SetTokenCache(ctx, *authClaims); err != nil {
		ac.log.Errorf("token 缓存失败 %v", err)
	}

	return authClaims, nil
}

// LoginNamePassword 登录-用户密码
func (ac *AuthUsecase) LoginNamePassword(ctx context.Context, domainCode string, g *User) (*AuthClaims, error) {
	domain, err := ac.biz.domainRepo.FindByCode(ctx, domainCode)
	if err != nil {
		return nil, errors.New("领域查询失败")
	}
	u, err := ac.biz.userRepo.FindByName(ctx, g.Name)
	if err != nil {
		return nil, err
	}
	err = password.Verify(u.Password, g.Password)
	if err != nil {
		return nil, err
	}

	authorities := ac.biz.enforcer.GetRolesForUserInDomain(convert.UnitToString(u.ID), convert.UnitToString(domain.ID))
	if len(authorities) == 0 {
		return nil, errors.New("权限未配置")
	}

	// 默认过期时间 24小时
	var expiresAt = time.Now().Add(time.Hour * 24)
	// 配置过期时间
	if ac.ac.GetExpiresTime() != nil {
		expiresAt = time.Now().Add(ac.ac.ExpiresTime.AsDuration())
	}
	// 生成token
	authClaims := &AuthClaims{
		ID:          u.ID,
		Domain:      domain.ID,
		Authorities: []uint{},
		ExpiresAt:   &expiresAt,
	}
	// 组装权限
	for _, v := range authorities {
		authClaims.Authorities = append(authClaims.Authorities, convert.StringToUint(v))
	}
	if err := ac.GetToken(authClaims); err != nil {
		return nil, err
	}

	if err := ac.biz.userRepo.SetTokenCache(ctx, *authClaims); err != nil {
		ac.log.Errorf("token 缓存失败 %v", err)
	}

	return authClaims, nil
}

// LoginMobileSms 登录-手机验证码
func (ac *AuthUsecase) LoginMobileSms(ctx context.Context, domainCode string, g *User) (*User, error) {
	ac.log.WithContext(ctx).Infof("mobileSmsLogin: %v", g)
	return ac.biz.userRepo.FindByMobile(ctx, g.Mobile)
}

// RegisterNamePassword 注册-用户密码
func (ac *AuthUsecase) RegisterNamePassword(ctx context.Context, domainCode string, g *User) (*User, error) {
	ac.log.WithContext(ctx).Infof("NamePasswordRegister: %v", g.Name)
	domain, err := ac.biz.domainRepo.FindByCode(ctx, domainCode)
	if err != nil {
		return nil, errors.New("领域查询失败")
	}
	userUsecase := &UserUsecase{ac.biz, ac.log, ac.ac}
	g, err = userUsecase.Create(ctx, g)
	if err != nil {
		return nil, err
	}

	g.Domains = []*Domain{domain}
	if err != userUsecase.HandleDomain(ctx, g) {
		return nil, err
	}
	return g, err
}

// Logout 退出登录
func (ac *AuthUsecase) Logout(ctx context.Context) error {

	return nil
}
