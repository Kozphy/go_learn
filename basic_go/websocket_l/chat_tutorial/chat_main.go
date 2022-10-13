// source: https://github.com/gorilla/websocket/tree/master/examples/chat
package chat_tutorial

import (
	"flag"
	"log"
	"net/http"
	"os"
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
