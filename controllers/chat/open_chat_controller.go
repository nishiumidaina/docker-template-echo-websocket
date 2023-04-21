package chat

import (
	"log"
	"github.com/labstack/echo"
	"github.com/gorilla/websocket"
)

var (
	clients = make(map[*websocket.Conn]bool)
	upgrader = websocket.Upgrader{
		ReadBufferSize: 1024,
		WriteBufferSize: 1024,
	}
)

func OpenChatHandler(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println(err)
		return err
	}
	defer conn.Close()

	clients[conn] = true
	defer delete(clients, conn)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return err
		}

		log.Printf("Received message: %s", message)

		for client := range clients {
			if err := client.WriteMessage(websocket.TextMessage, message); err != nil {
				log.Println(err)
				delete(clients, client)
			}
		}
	}
}