package api

import (
	v1 "github.com/beiduoke/go-scaffold/api/server/v1"
	"github.com/beiduoke/go-scaffold/internal/biz"
	"github.com/beiduoke/go-scaffold/internal/conf"
	"github.com/beiduoke/go-scaffold/internal/pkg/websocket"
	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/dig"
)

var _ v1.ApiServer = (*ApiService)(nil)

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

// ApiService is a Api service.
type ApiService struct {
	v1.UnimplementedApiServer
	ac           *conf.Auth
	log          *log.Helper
	ws           *websocket.WebsocketService
	dig          *dig.Container
	authCase     *biz.AuthUsecase
	userCase     *biz.UserUsecase
	domainCase   *biz.DomainUsecase
	roleCase     *biz.RoleUsecase
	menuCase     *biz.MenuUsecase
	resourceCase *biz.ResourceUsecase
	deptCase     *biz.DeptUsecase
	postCase     *biz.PostUsecase
}

// NewApiService new a Api service.
func NewApiService(
	logger log.Logger,
	ac *conf.Auth,
	ws *websocket.WebsocketService,
	authCase *biz.AuthUsecase,
	userCase *biz.UserUsecase,
	domainCase *biz.DomainUsecase,
	roleCase *biz.RoleUsecase,
	menuCase *biz.MenuUsecase,
	resourceCase *biz.ResourceUsecase,
	deptCase *biz.DeptUsecase,
	postCase *biz.PostUsecase,
) *ApiService {
	l := log.NewHelper(log.With(logger, "module", "service/api"))
	return &ApiService{
		log:          l,
		ac:           ac,
		ws:           ws,
		authCase:     authCase,
		userCase:     userCase,
		domainCase:   domainCase,
		roleCase:     roleCase,
		menuCase:     menuCase,
		resourceCase: resourceCase,
		deptCase:     deptCase,
		postCase:     postCase,
	}
}
