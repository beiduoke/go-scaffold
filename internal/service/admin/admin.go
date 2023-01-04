package admin

import (
	v1 "github.com/beiduoke/go-scaffold/api/admin/v1"
	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/beiduoke/go-scaffold/internal/conf"
	"github.com/beiduoke/go-scaffold/internal/pkg/websocket"
	"github.com/go-kratos/kratos/v2/log"
)

var _ v1.AdminServer = (*AdminService)(nil)

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

// AdminService is a Admin service.
type AdminService struct {
	v1.UnimplementedAdminServer
	ac             *conf.Auth
	log            *log.Helper
	ws             *websocket.WebsocketService
	authCase       *biz.AuthUsecase
	userCase       *biz.UserUsecase
	domainCase     *biz.DomainUsecase
	roleCase       *biz.RoleUsecase
	menuCase       *biz.MenuUsecase
	resourceCase   *biz.ResourceUsecase
	departmentCase *biz.DepartmentUsecase
}

// NewAdminService new a Admin service.
func NewAdminService(
	logger log.Logger,
	ac *conf.Auth,
	ws *websocket.WebsocketService,
	authCase *biz.AuthUsecase,
	userCase *biz.UserUsecase,
	domainCase *biz.DomainUsecase,
	roleCase *biz.RoleUsecase,
	menuCase *biz.MenuUsecase,
	resourceCase *biz.ResourceUsecase,
	departmentCase *biz.DepartmentUsecase,
) *AdminService {
	l := log.NewHelper(log.With(logger, "module", "service/admin"))
	return &AdminService{
		log:            l,
		ac:             ac,
		ws:             ws,
		authCase:       authCase,
		userCase:       userCase,
		domainCase:     domainCase,
		roleCase:       roleCase,
		menuCase:       menuCase,
		resourceCase:   resourceCase,
		departmentCase: departmentCase,
	}
}
