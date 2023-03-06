package biz

import (
	"context"
	"fmt"
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

type AuthClaims struct {
	ID        uint
	Domain    uint
	Roles     []uint
	ExpiresAt *time.Time
	Token     string
}

// AuthUsecaseOld is a User usecase.
type AuthUsecaseOld struct {
	biz *Biz
	ac  *conf.Auth
	log *log.Helper
}

// NewAuthUsecaseOld new a User usecase.
func NewAuthUsecaseOld(logger log.Logger, biz *Biz, ac *conf.Auth) *AuthUsecaseOld {
	return &AuthUsecaseOld{log: log.NewHelper(logger), ac: ac, biz: biz}
}

func (ac *AuthUsecaseOld) GetToken(claims *AuthClaims) error {
	roles := []string{}
	for _, v := range claims.Roles {
		roles = append(roles, strconv.Itoa(int(v)))
	}
	expiresAt := time.Now().Add(time.Hour * 24)
	if claims.ExpiresAt != nil {
		expiresAt = *claims.ExpiresAt
	}
	securityUser := myAuthz.NewSecurityUserData(
		myAuthz.WithUser(strconv.Itoa(int(claims.ID))),
		myAuthz.WithDomain(strconv.Itoa(int(claims.Domain))),
		myAuthz.WithSubject(strings.Join(roles, ",")),
		myAuthz.WithExpires(expiresAt),
	)
	claims.Token = securityUser.CreateAccessJwtToken([]byte(ac.ac.Jwt.GetSecretKey()))
	if claims.Token == "" {
		return errors.New("token生成失败")
	}
	return nil
}

// Login 登录-密码登录
func (ac *AuthUsecaseOld) Login(ctx context.Context, g *User) (*AuthClaims, error) {
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
	if ac.ac.Jwt.GetExpiresTime() != nil {
		expiresAt = time.Now().Add(ac.ac.Jwt.ExpiresTime.AsDuration())
	}
	// 生成token
	authClaims := &AuthClaims{
		ID:        u.ID,
		ExpiresAt: &expiresAt,
		Domain:    domain.ID,
	}

	if err := ac.GetToken(authClaims); err != nil {
		return nil, err
	}

	// if err := ac.biz.userRepo.SetTokenCache(ctx, *authClaims); err != nil {
	// 	ac.log.Errorf("token 缓存失败 %v", err)
	// }

	return authClaims, nil
}

// LoginNamePassword 登录-用户密码
func (ac *AuthUsecaseOld) LoginNamePassword(ctx context.Context, domainCode string, g *User) (*AuthClaims, error) {
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

	roles := ac.biz.enforcer.GetRolesForUserInDomain(convert.UnitToString(u.ID), convert.UnitToString(domain.ID))
	if len(roles) == 0 {
		return nil, errors.New("权限未配置")
	}

	// 默认过期时间 24小时
	var expiresAt = time.Now().Add(time.Hour * 24)
	// 配置过期时间
	if ac.ac.Jwt.GetExpiresTime() != nil {
		expiresAt = time.Now().Add(ac.ac.Jwt.ExpiresTime.AsDuration())
	}
	// 生成token
	authClaims := &AuthClaims{
		ID:        u.ID,
		Domain:    domain.ID,
		Roles:     []uint{},
		ExpiresAt: &expiresAt,
	}
	// 组装权限
	for _, v := range roles {
		authClaims.Roles = append(authClaims.Roles, convert.StringToUint(v))
	}
	if err := ac.GetToken(authClaims); err != nil {
		return nil, err
	}

	// if err := ac.biz.userRepo.SetTokenCache(ctx, *authClaims); err != nil {
	// 	ac.log.Errorf("token 缓存失败 %v", err)
	// }

	return authClaims, nil
}

// LoginPhoneSms 登录-手机验证码
func (ac *AuthUsecaseOld) LoginPhoneSms(ctx context.Context, domainCode string, g *User) (*User, error) {
	ac.log.WithContext(ctx).Infof("phoneSmsLogin: %v", g)
	return ac.biz.userRepo.FindByPhone(ctx, g.Phone)
}

// RegisterNamePassword 注册-用户密码
func (ac *AuthUsecaseOld) RegisterNamePassword(ctx context.Context, domainCode string, g *User) (*User, error) {
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
	fmt.Println(domain)
	// g.Domains = []*Domain{domain}
	// if err != userUsecase.HandleDomain(ctx, g) {
	// 	return nil, err
	// }
	return g, err
}

// Logout 退出登录
func (ac *AuthUsecaseOld) Logout(ctx context.Context) error {

	return nil
}
