package websocket

import (
	"encoding/json"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/tx7do/kratos-transport/transport/websocket"
)

type Event struct {
	EventId string `json:"eventId,omitempty"`
	Payload string `json:"payload,omitempty"`
}

type Option func(*WebsocketService)

func WithSocket(w *websocket.Server) Option {
	return func(ws *WebsocketService) {
		ws.ws = w
	}
}

// WebsocketService is a Admin service.
type WebsocketService struct {
	log *log.Helper
	ws  *websocket.Server
}

// NewWebsocketService new a Admin service.
func NewWebsocketService(logger log.Logger) *WebsocketService {
	l := log.NewHelper(log.With(logger, "module", "websocket"))
	return &WebsocketService{log: l}
}

func (s *WebsocketService) SetWebsocketServer(ws *websocket.Server) {
	s.ws = ws
}

func (s *WebsocketService) OnWebsocketMessage(connectionId string, message *websocket.Message) error {
	s.log.Infof("[%s] Payload: %s\n", connectionId, string(message.Body))

	var proto Event

	if err := json.Unmarshal(message.Body, &proto); err != nil {
		s.log.Error("Error unmarshal json %v", err)
		return nil
	}

	switch proto.EventId {
	case "chat":
		chatMsg := proto.Payload
		fmt.Println("chat message:", chatMsg)
		_ = s.OnChatMessage(connectionId, &chatMsg)
	}

	return nil
}

func (s *WebsocketService) OnChatMessage(connectionId string, msg *string) error {
	s.BroadcastToWebsocketClient("chat", msg)
	return nil
}

func (s *WebsocketService) OnWebsocketConnect(connectionId string, register bool) {
	if register {
		fmt.Printf("%s connected\n", connectionId)
	} else {
		fmt.Printf("%s disconnect\n", connectionId)
	}
}

func (s *WebsocketService) BroadcastToWebsocketClient(eventId string, payload interface{}) {
	if payload == nil {
		return
	}

	bufPayload, _ := json.Marshal(&payload)

	var proto Event
	proto.EventId = eventId
	proto.Payload = string(bufPayload)

	bufProto, _ := json.Marshal(&proto)

	var msg websocket.Message
	msg.Body = bufProto

	s.ws.Broadcast(&msg)
}
