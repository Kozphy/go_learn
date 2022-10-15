package pingpong

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func serveWs(w http.ResponseWriter, r *http.Request) *websocket.Conn {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return ws
}

func Execute_ping_pong_client() {
	conn1, _, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:8080/room1", nil)
	if err != nil {
		log.Fatalf("client room 1 err: %v\n", err)
	}
	_, msg, err := conn1.ReadMessage()
	if err != nil {
		log.Fatalf("client read message from room1 err: %v", err)
	}
	fmt.Printf("client get: %v\n", string(msg))
	time.Sleep(time.Millisecond * 100)
	conn2, _, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:8080/room2", nil)
	if err != nil {
		log.Fatalf("client room2 connect err: %v", err)
	}
	err = conn2.WriteMessage(websocket.TextMessage, []byte("client pong"))
	if err != nil {
		log.Fatalf("client send pong to room2 err: %v", err)
	}

	_, msg, err = conn1.ReadMessage()
	if err != nil {
		log.Fatalf("client read message from room1 err second time: %v", err)
	}
	fmt.Printf("get from room1 second time: %v\n", string(msg))
}
