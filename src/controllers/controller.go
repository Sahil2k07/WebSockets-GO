package controllers

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var Upgrader = websocket.Upgrader{}

func WebSocketControllers(e *echo.Echo) {
	e.GET("/chat", ChatController)

	e.GET("/group", GroupController)
}
