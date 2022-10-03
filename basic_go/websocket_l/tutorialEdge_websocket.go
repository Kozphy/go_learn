package websocket_l

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// source: https://tutorialedge.net/golang/go-websocket-tutorial/

// WebSockets - What Are They?
/*
	WebSockets are upgraded HTTP connections that live until the connection is killed by
	either the client or the server. It’s through this WebSocket connection that we
	can perform duplex communication which is a really fancy way of saying we can
	communicate to-and-from the server from our client using this single connection.

	The real beauty of WebSockets is that they use a grand total of 1 TCP connection
	and all communication is done over this single long-lived TCP connection.

	In order to create a WebSocket endpoint, we effectively need to upgrade an incoming
	connection from a standard HTTP endpoint to a long-lasting WebSocket connection.
*/

// Upgrading a HTTP Connection
/*
	The first thing we’ll have to do is to define a websocker.Upgrader struct.
	This will hold information such as the Read and Write buffer size
	for our WebSocket connection:
*/
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Execute_simeple_http_endpoint() {
	fmt.Println("Hello World")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	/*
		This will determine whether or not an incoming request from a different
		domain is allowed to connect, and if it isn’t they’ll be hit with a CORS
		error.

		For now, we have kept it really simple and simply return true regardless of
		what host is trying to connect to our endpoint.
	*/
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	/*
		We can now start attempting to upgrade the incoming HTTP connection using
		the upgrader.Upgrade() function which will take in the Response Writer and
		the pointer to the HTTP Request and return us with a pointer to a WebSocket
		connection, or an error if it failed to upgrade.
	*/
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client Connected")
	render(ws)
}

// FIXME: can't connect to wsl with ws:// header
// will continually listen for any incoming messages sent through that WebSocket connection.
func render(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}
