package main

import (
	"github.com/Sahil2k07/WebSockets-GO/src/controllers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	controllers.WebSocketControllers(e)

	e.Logger.Fatal(e.Start(":1323"))
}
