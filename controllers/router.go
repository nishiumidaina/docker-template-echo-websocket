package controller

import (
	"docker-echo-template/controllers/chat"

	"github.com/labstack/echo"
)

func Router() *echo.Echo {
	e := echo.New()

	e.GET("/chat", chat.ChatHandler)

	e.Start(":8000")

	return e
}