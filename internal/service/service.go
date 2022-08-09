package service

import (
	"github.com/beiduoke/go-scaffold/internal/pkg/websocket"
	"github.com/beiduoke/go-scaffold/internal/service/admin"
	"github.com/beiduoke/go-scaffold/internal/service/web"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(admin.NewAdminService, web.NewWebService, websocket.NewWebsocketService)
