package chat

import (
	"log"
	"net/http"
	"github.com/labstack/echo"
	"github.com/gorilla/websocket"
)

var (
	rooms = make(map[string]map[*websocket.Conn]bool)
	roomChatUpgrader = websocket.Upgrader{
		ReadBufferSize: 1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool { return true },
	}
)

func RoomChatHandler(c echo.Context) error {
	room := c.Param("room")
	conn, err := roomChatUpgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println(err)
		return err
	}

	if rooms[room] == nil {
		rooms[room] = make(map[*websocket.Conn]bool)
	}

	rooms[room][conn] = true

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			delete(rooms[room], conn)
			return err
		}

		message := string(p)

		for client := range rooms[room] {
			err := client.WriteMessage(messageType, []byte(message))
			if err != nil {
				log.Println(err)
				delete(rooms[room], client)
			}
		}
	}
}