package controller

import (
	"docker-echo-template/controllers/chat"

	"github.com/labstack/echo"
)

func Router() *echo.Echo {
	e := echo.New()

	e.GET("/open_chat", chat.OpenChatHandler)

	e.Start(":8000")

	return e
}