package chat_tutorial

/*
Hub maintains the set of active clients and broadcasts messages to the
connection.
*/
type Hub struct {
	rooms map[string]map[*connection]bool
	// Registered clients.
	clients map[*Client]bool
	// Inbound messages from the clients.
	broadcast chan []byte
	// Register requests from the clients.
	register chan *Client
	// Unregister requests from clients.
	unregister chan *Client
}

type subscription struct {
	conn *connection
	room string
}

type message struct {
	data []byte
	room string
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		rooms:      make(map[string]map[*connection]bool),
	}
}

// monitor h.register, h.unregister, h.broadcast whether have value
func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			// if register set client to true
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					// no message
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
