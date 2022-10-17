package chat_tutorial

import (
	"time"

	"github.com/gorilla/websocket"
)

// connection is an middleman between the websocket connection and the hub
type connection struct {
	ws   *websocket.Conn
	send chan []byte
}

// write writes a message with the given message type and payload
func (c *connection) write(mt int, payload []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteMessage(mt, payload)
}

func (s subscription) readPump() {
	c := s.conn
	defer func() {
		h.unregister <- s
	}()
}
