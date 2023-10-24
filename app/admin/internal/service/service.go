package service

import (
	v1 "github.com/beiduoke/go-scaffold/api/admin/v1"
	"github.com/beiduoke/go-scaffold/app/admin/internal/biz"
	"github.com/beiduoke/go-scaffold/app/admin/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewAdminService)

var _ v1.AdminServiceServer = (*AdminService)(nil)

// 使用i18n包进行国际化
// localizer := localize.FromContext(ctx)
// fmt.Println(localizer)
// helloMsg, err := localizer.Localize(&i18n.LocalizeConfig{
// 	DefaultMessage: loginMessage,
// 	TemplateData: map[string]interface{}{
// 		"Name":     in.Name,
// 		"Password": in.Password,
// 	},
// })
// println(helloMsg, 11111)
// if err != nil {
// 	return nil, err
// }

// Service is a  service.
type AdminService struct {
	v1.UnimplementedAdminServiceServer
	ac  *conf.Auth
	log *log.Helper
	// dig        *dig.Container
	authCase   *biz.AuthUsecase
	userCase   *biz.UserUsecase
	domainCase *biz.DomainUsecase
	roleCase   *biz.RoleUsecase
	menuCase   *biz.MenuUsecase
	deptCase   *biz.DeptUsecase
	postCase   *biz.PostUsecase
	dictCase   *biz.DictUsecase
}

// NewService new a  service.
func NewAdminService(
	logger log.Logger,
	authCase *biz.AuthUsecase,
	userCase *biz.UserUsecase,
	domainCase *biz.DomainUsecase,
	roleCase *biz.RoleUsecase,
	menuCase *biz.MenuUsecase,
	deptCase *biz.DeptUsecase,
	postCase *biz.PostUsecase,
	dictCase *biz.DictUsecase,
) *AdminService {
	l := log.NewHelper(log.With(logger, "module", "service"))
	return &AdminService{
		log:        l,
		authCase:   authCase,
		userCase:   userCase,
		domainCase: domainCase,
		roleCase:   roleCase,
		menuCase:   menuCase,
		deptCase:   deptCase,
		postCase:   postCase,
		dictCase:   dictCase,
	}
}
