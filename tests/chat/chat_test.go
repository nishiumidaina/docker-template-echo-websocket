package chat

import (
    "fmt"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gorilla/websocket"
)

func TestChat(t *testing.T) {
    go startServer()

    conn, err := connect()

    if err != nil {
        t.Fatalf("failed to connect to WebSocket: %s", err)
    }

    defer conn.Close()

    message := "Hello, world!"
    err = conn.WriteMessage(websocket.TextMessage, []byte(message))
    if err != nil {
        t.Fatalf("failed to write message to WebSocket: %s", err)
    }

    _, p, err := conn.ReadMessage()
    if err != nil {
        t.Fatalf("failed to read message from WebSocket: %s", err)
    }

    if string(p) != message {
        t.Errorf("expected message %q, but got %q", message, string(p))
    }
}

func startServer() {
    err := http.ListenAndServe(":8000", nil)
    if err != nil {
        fmt.Printf("failed to start server: %s\n", err)
    }
}

func connect() (*websocket.Conn, error) {
    u := url.URL{Scheme: "ws", Host: "localhost:8000", Path: "/chat"}
    conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
    if err != nil {
        return nil, fmt.Errorf("failed to connect to WebSocket: %s", err)
    }

    return conn, nil
}