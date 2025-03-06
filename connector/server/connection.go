package server

import (
	"log"
	"net/http"
	bp "ws-server/proto"

	"google.golang.org/protobuf/proto"
)

func (s *WSServer) HandleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Websocket upgrade error:", err)
		return
	}
	s.AddClient(conn)
	defer s.RemoveClient(conn)

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		event := &bp.RequestEvent{}

		errParse := proto.Unmarshal(msg, event)
		if errParse != nil {
			log.Println("Event convert failed", err)
		}

		switch event.Type {
		case bp.EventType_PING:
			{
				s.Ack <- event
			}
		default:
			{
				log.Printf("No handler event %v!", event.Type)

			}
		}

		s.broadcast <- msg
	}
}
