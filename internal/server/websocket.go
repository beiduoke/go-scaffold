package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/transport/websocket"

	"github.com/bedoke/go-scaffold/internal/conf"
	"github.com/bedoke/go-scaffold/internal/service"
)

// NewWebsocketServer create a websocket server.
func NewWebsocketServer(c *conf.Server, _ log.Logger, svc *service.WebsocketService) *websocket.Server {
	var opts = []websocket.ServerOption{
		websocket.ConnectHandle(svc.OnWebsocketConnect),
	}
	if c.Websocket.Network != "" {
		opts = append(opts, websocket.Network(c.Websocket.Network))
	}
	if c.Websocket.Addr != "" {
		opts = append(opts, websocket.Address(c.Websocket.Addr))
	}
	if c.Websocket.Timeout != nil {
		opts = append(opts, websocket.Timeout(c.Websocket.Timeout.AsDuration()))
	}
	if c.Websocket.Path != "" {
		opts = append(opts, websocket.ReadHandle(c.Websocket.Path, svc.OnWebsocketMessage))
	}
	srv := websocket.NewServer(opts...)
	svc.SetWebsocketServer(srv)
	return srv
}
