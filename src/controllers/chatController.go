package controllers

import (
	"github.com/Sahil2k07/WebSockets-GO/src/services"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

func ChatController(c echo.Context) error {
	ws, err := Upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	// Ideally we extract from JWT
	userID := c.QueryParam("userID")
	receiverID := c.QueryParam("receiverID")

	go services.WSManager.RegisterChatSession(userID, ws)
	defer services.WSManager.TerminateChatSession(userID)

	for {
		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				return nil
			}

			c.Logger().Error(err)
			return err
		}

		// Write
		conn := services.WSManager.GetUserSession(receiverID)
		if conn != nil {
			err = conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				c.Logger().Error(err)
			}
		}

		// Save the MSG in the DB
	}
}
