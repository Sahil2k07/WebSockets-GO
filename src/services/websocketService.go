package services

import (
	"sync"

	"github.com/Sahil2k07/WebSockets-GO/src/interfaces"
	"github.com/gorilla/websocket"
)

type websocketConnections struct {
	chatConnections  map[string]*websocket.Conn
	groupConnections map[string][]*websocket.Conn
	mu               sync.RWMutex
}

func newWebsocketConnections() *websocketConnections {
	return &websocketConnections{
		chatConnections:  make(map[string]*websocket.Conn),
		groupConnections: make(map[string][]*websocket.Conn),
	}
}

// Singletion Registration
var WSManager interfaces.IWebsocketConnections = newWebsocketConnections()

func (wc *websocketConnections) RegisterChatSession(userID string, conn *websocket.Conn) {
	wc.mu.Lock()
	defer wc.mu.Unlock()

	wc.chatConnections[userID] = conn
}

func (wc *websocketConnections) TerminateChatSession(userID string) {
	wc.mu.Lock()
	defer wc.mu.Unlock()

	delete(wc.chatConnections, userID)
}

func (wc *websocketConnections) GetUserSession(userID string) *websocket.Conn {
	wc.mu.RLock()
	defer wc.mu.RUnlock()

	return wc.chatConnections[userID]
}

func (wc *websocketConnections) RegisterGroupSession(groupID string, conn *websocket.Conn) {
	wc.mu.Lock()
	defer wc.mu.Unlock()

	wc.groupConnections[groupID] = append(wc.groupConnections[groupID], conn)
}

func (wc *websocketConnections) GetGroupSessions(groupID string) []*websocket.Conn {
	wc.mu.RLock()
	defer wc.mu.RUnlock()

	return wc.groupConnections[groupID]
}

func (wc *websocketConnections) TerminateGroupSession(groupID string, conn *websocket.Conn) {
	wc.mu.Lock()
	defer wc.mu.Unlock()

	if conns, exists := wc.groupConnections[groupID]; exists {
		for i, c := range conns {
			if c == conn {
				wc.groupConnections[groupID] = append(conns[:i], conns[i+1:]...)
				conn.Close()
				break
			}
		}

		if len(wc.groupConnections[groupID]) == 0 {
			delete(wc.groupConnections, groupID)
		}
	}
}
