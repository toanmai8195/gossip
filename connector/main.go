package main

import (
	"log"
	"net/http"
	"ws-server/server"
)

func main() {
	srv := server.NewServer()

	http.HandleFunc("/ws", srv.HandleWS)

	log.Println("WebSocket server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
