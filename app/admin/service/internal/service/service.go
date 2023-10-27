package service

import (
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewAuthService,
	NewUserService,
	NewRoleService,
	NewDomainService,
	NewMenuService,
	NewPostService,
	NewDictService,
	NewDeptService,
)

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
