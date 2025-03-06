package handler

import (
	"log"
	"time"
	"ws-server/server"

	"github.com/gorilla/websocket"
)

func StartPing(s *server.WSServer, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		<-ticker.C
		s.Mu.RLock()
		for _, client := range s.Clients {
			if err := client.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Println("Ping error:", err)
				s.RemoveClient(client.Conn)
			}
		}
		s.Mu.RUnlock()
	}
}
