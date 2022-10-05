package chat_tutorial

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = 60 * time.Second
	maxMessageSize = 512
)

// Upgrader specifies parameters for upgrading an HTTP connection to WebSocket connection.
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Client is middleman between the websocket and the hub.
type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
}

func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	/*
		A server application calls the upgrader.Upgrade method from an HTTP
		request handler to get a *Conn, which represents a Websocket connection:
	*/
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client
}

// readPump pumps messages from the websocket connection to the hub.
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

}
