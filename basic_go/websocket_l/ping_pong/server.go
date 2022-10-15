package pingpong

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func room1(done chan struct{}, w http.ResponseWriter, r *http.Request) {
	fmt.Println("start room1")
	ws := serveWs(w, r)
	err := ws.WriteMessage(websocket.TextMessage, []byte("room1 ping"))
	if err != nil {
		log.Fatalf("server room1 err ping first time: %v\n", err)
	}
	closeSignal := <-done
	fmt.Printf("room 1 get close Signal: %v\n", closeSignal)

	if closeSignal == struct{}{} {
		close(done)
		err = ws.WriteMessage(websocket.TextMessage, []byte("room1 ping second time"))
		if err != nil {
			log.Fatalf("server room1 err ping second time: %v\n", err)
		}
	}

}

func room2(done chan struct{}, w http.ResponseWriter, r *http.Request) {
	fmt.Println("start room2")
	ws := serveWs(w, r)
	_, payload, err := ws.ReadMessage()
	if err != nil {
		log.Fatalf("server room2 err: %v\n", err)
	}
	if payload != nil {
		fmt.Printf("get client res: %s\n", string(payload))
		done <- struct{}{}
	}
}

func Execute_ping_pong_server() {
	fmt.Println("start")
	done := make(chan struct{})
	defer close(done)
	go Execute_room1(done)
	go Execute_room2(done)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Execute_room1(done chan struct{}) {
	http.HandleFunc("/room1", func(w http.ResponseWriter, r *http.Request) {
		room1(done, w, r)
	})
}

func Execute_room2(done chan struct{}) {
	http.HandleFunc("/room2", func(w http.ResponseWriter, r *http.Request) {
		room2(done, w, r)
	})
}

// func Execute_ping_pong_websocket() {
// 	go Execute_ping_pong_websocket()
// 	time.Sleep(time.Second)
// 	Execute_ping_pong_client()
// }
