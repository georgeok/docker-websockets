package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func ws(w http.ResponseWriter, r *http.Request) {
	// Upgrade connection
	upgrader := websocket.Upgrader{}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	log.Printf("New connection from: %s", r.RemoteAddr)
	// Read messages from socket
	for {
		mt, msg, err := conn.ReadMessage()
		if err != nil {
			conn.Close()
			return
		}
		log.Printf("msg: %s", string(msg))
		log.Printf("recv: %s", string(msg))
		err = conn.WriteMessage(mt, msg)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func main() {
	log.Println("Starting ...")
	http.HandleFunc("/", ws)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
	log.Println("Stopping ...")
}
