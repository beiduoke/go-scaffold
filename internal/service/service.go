package service

import (
	"github.com/beiduoke/go-scaffold/internal/pkg/websocket"
	"github.com/beiduoke/go-scaffold/internal/service/api"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(api.NewApiService, websocket.NewWebsocketService)
