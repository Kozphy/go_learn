package chat_tutorial

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func serveWs_conn(hub *Hub_conn, w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	vars := mux.Vars(r)
	log.Println(vars["room1"])
	log.Println(vars["room2"])
	if err != nil {
		log.Println(err)
		return
	}
	// for router, _ := range vars {

	// }
	c := &connection{send: make(chan []byte, 256), ws: ws}
	s := subscription{c, vars["room"]}
	hub.register <- s
	go s.writePump()
	s.readPump(hub)
}

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

// readPump pumps messages from the websocket connection to the hub.
func (s subscription) readPump(hub *Hub_conn) {
	c := s.conn
	defer func() {
		hub.unregister <- s
		c.ws.Close()
	}()
	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error {
		c.ws.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, msg, err := c.ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		m := message{msg, s.room}
		hub.broadcast <- m
	}
}

// writePump pumps messages from the hub to the websocket connection.
func (s *subscription) writePump() {
	c := s.conn
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}
