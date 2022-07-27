package service

import (
	"github.com/bedoke/go-scaffold/internal/service/admin"
	"github.com/bedoke/go-scaffold/internal/service/web"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewWebsocketService, admin.NewAdminService, web.NewWebService)
