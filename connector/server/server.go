package server

import (
	"fmt"
	"log"
	"net/http"
	pb "ws-server/proto"

	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

type Server struct {
	upgrader websocket.Upgrader
}

func NewServer() *Server {
	return &Server{
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

func (s *Server) HandleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		log.Println("Received:", string(msg))

		// Convert msg to string before using it in fmt.Sprintf

		resp := &pb.ChatResponse{Message: &pb.Message{
			Id:        "id1",
			Sender:    "toanmai1",
			Content:   fmt.Sprintf("Rep: %s", string(msg)),
			Timestamp: 0,
		}}
		data, err := proto.Marshal(resp)
		if err != nil {
			log.Println("Failed to marshal response:", err)
			continue
		}

		err = conn.WriteMessage(websocket.BinaryMessage, data)
		if err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}
