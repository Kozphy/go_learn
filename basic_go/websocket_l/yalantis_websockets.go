package websocket_l

import (
	"fmt"
	"net/http"
)

/*
	To write a simple WebSocket echo server based on the net/http library, you need to:
	- Initiale a handshake
	- Receive data frames from the client
	- Send data frames to the client
	- Close the handshake
*/

// HTTP server with WebSocket endpoint
/*
	The initial handshake request always comes from the client. Once the server
	has defined a WebSocket request, it needs to reply with a handshake response.

	Bear in mind that you can't write the response using the http.ResponseWriter, since
	it will close the underlying TCP connection once you start sending the response.

	So you need to use HTTP hijacking. Hijacking allows you to take over the underlying TCP
	connection handler and bufio.Writer. This gives you the possibility to read and write data
	without closing the TCP connection.
*/
func Server() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ws, err := NewHandler(w, r)
		if err != nil {
			// handle error
			panic(fmt.Errorf("NewHandler error %v\n", err))
		}
		if err = ws.Handshake(); err != nil {
			// handle error
			panic(fmt.Errorf("ws handshake error %v\n", err))
		}
	})

	// Handshake creates a handshake header
	/*

	 */
	// func (ws *WS) Handshake() error {
	// 	hash := func(key string) string {
	// 		h := sha.New()
	// 		h.Write([]byte(key))
	// 		h.Write([]byte("258EAFA5-E914-47DA-95CA-C5AB0DC85B11"))
	// 		return base64.StdEncoding.EncodeToString(h.Sum(nil))
	// 	}(ws.header.Get("Sec-WebSocket-Key"))
	// 	return nil
	// }

}

// NewHandler initializes a new handler
func NewHandler(w http.ResponseWriter, req *http.Request) (*WS, error) {
	hj, ok := w.(http.Hijacker)
	if !ok {
		// handle error
		http.Error(w, "webserver doesn't support hijacking", http.StatusInternalServerError)
		return
	}
	conn, bufrw, err := hj.Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	bufrw.WriteString("Now we're speaking raw TCP. Say hi: ")
	bufrw.Flush()
}
