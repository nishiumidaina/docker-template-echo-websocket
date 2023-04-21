package chat

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
)

func TestRoomChat(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(roomChatHandle))
	defer server.Close()

	url := "ws" + strings.TrimPrefix(server.URL, "http") + "/room_chat/0000"
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("Failed to dial WebSocket: %v", err)
	}
	defer ws.Close()

	message := "test message"
	err = ws.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	_, p, err := ws.ReadMessage()
	if err != nil {
		t.Fatalf("Failed to read message: %v", err)
	}
	received := string(p)

	if received != message {
		t.Errorf("Received message does not match sent message: expected=%q, actual=%q", message, received)
	}
}

func roomChatHandle(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		return
	}
	defer conn.Close()

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			return
		}
		err = conn.WriteMessage(websocket.TextMessage, p)
		if err != nil {
			return
		}
	}
}