package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/transport/websocket"

	"github.com/beiduoke/go-scaffold/internal/conf"
	ws "github.com/beiduoke/go-scaffold/internal/pkg/websocket"
)

// NewWebsocketServer create a websocket server.
func NewWebsocketServer(c *conf.Server, _ log.Logger, ws *ws.WebsocketService) *websocket.Server {
	var opts = []websocket.ServerOption{
		websocket.ConnectHandle(ws.OnWebsocketConnect),
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
		opts = append(opts, websocket.ReadHandle(c.Websocket.Path, ws.OnWebsocketMessage))
	}
	srv := websocket.NewServer(opts...)
	ws.SetWebsocketServer(srv)
	return srv
}
