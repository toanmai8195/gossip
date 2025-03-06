package server

import (
	"log"

	"github.com/gorilla/websocket"
)

func (s *WSServer) StartWorker() {
	for conn, client := range s.Clients {
		go func(conn *websocket.Conn, client *Connection) {
			for msg := range client.Send {
				if err := conn.WriteMessage(websocket.BinaryMessage, msg); err != nil {
					log.Println("Write error:", err)
					s.RemoveClient(conn)
					break
				}
			}
		}(conn, client)
	}
}
