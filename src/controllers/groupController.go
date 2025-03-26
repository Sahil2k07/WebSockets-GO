package controllers

import (
	"github.com/Sahil2k07/WebSockets-GO/src/services"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

func GroupController(c echo.Context) error {
	ws, err := Upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	groupID := c.QueryParam("groupID")

	go services.WSManager.RegisterGroupSession(groupID, ws)
	defer services.WSManager.TerminateGroupSession(groupID, ws)

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

		gConn := services.WSManager.GetGroupSessions(groupID)
		for _, conn := range gConn {
			if conn == ws {
				continue
			}

			err = conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				c.Logger().Error(err)
			}
		}

		// Save the MSG in the DB
	}
}
