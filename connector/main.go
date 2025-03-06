package main

import (
	"log"
	"net/http"
	"time"
	"ws-server/handler"
	"ws-server/server"
)

func main() {
	wsServer := server.NewWSServer()

	go wsServer.StartBroadcast(10)
	go handler.StartPing(wsServer, 30*time.Second)
	go wsServer.StartWorker()

	http.HandleFunc("/ws", wsServer.HandleConnection)
	log.Println("WebSocket server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
