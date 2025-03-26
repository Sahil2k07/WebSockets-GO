package interfaces

import "github.com/gorilla/websocket"

type IWebsocketConnections interface {
	RegisterChatSession(userID string, conn *websocket.Conn)
	TerminateChatSession(userID string)
	GetUserSession(userID string) *websocket.Conn
	RegisterGroupSession(groupID string, conn *websocket.Conn)
	TerminateGroupSession(groupID string, conn *websocket.Conn)
	GetGroupSessions(groupID string) []*websocket.Conn
}
