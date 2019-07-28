package main

import (
	"fmt"
	"net/http"

	"github.com/shiv-er/go-sockets-backend/pkg/websocket"
)

func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	ws, err := websocket.Upgrade(w, r, nil)
	if err != nil {
		log.Println(w, "%+V\n", err)
	}

	go websocket.Writer(ws)
	websocket.Reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Chat App v0.01")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
