// source: https://github.com/gorilla/websocket/tree/master/examples/chat
package chat_tutorial

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

// return value of the flag
var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	log.Println(r.URL.Path == "/")
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	log.Println("Serve file")

	dir, _ := os.Getwd()
	home_lacation := dir + "/websocket_l/chat_tutorial/home.html"
	_, err := os.Stat(home_lacation)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	http.ServeFile(w, r, home_lacation)
}

func Execute_chat_websocket() {
	flag.Parse()
	hub := newHub()
	go hub.run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

func Execute_chat_websocket_conn() {
	fmt.Println("start chat websocket_conn")
	flag.Parse()
	hub := newHub_conn()
	go hub.run()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome)
	r.HandleFunc("/room1", func(w http.ResponseWriter, r *http.Request) {
		serveWs_conn(hub, w, r)
	})
	r.HandleFunc("/room2", func(w http.ResponseWriter, r *http.Request) {
		serveWs_conn(hub, w, r)
	})
	srv := &http.Server{
		Handler:      r,
		Addr:         *addr,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
