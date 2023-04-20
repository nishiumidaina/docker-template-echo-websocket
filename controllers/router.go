package controller

import (
	"docker-echo-template/controllers/websocket"

	"github.com/labstack/echo"
)

func Router() *echo.Echo {
	e := echo.New()

	e.GET("/websocket", websocket.WebsocketHandler)

	e.Start(":8000")

	return e
}