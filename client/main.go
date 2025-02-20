package main

import (
	"fmt"
	"log"
	"time"

	pb "client/proto"

	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

const serverURL = "ws://localhost:8080/ws"

func main() {
	// Kết nối đến WebSocket server
	conn, _, err := websocket.DefaultDialer.Dial(serverURL, nil)
	if err != nil {
		log.Fatalf("WebSocket connection error: %v", err)
	}
	defer conn.Close()

	// Gửi yêu cầu tham gia room chat
	chatReq := &pb.ChatRequest{RoomId: "12345"}
	data, err := proto.Marshal(chatReq)
	if err != nil {
		log.Fatalf("Failed to marshal request: %v", err)
	}

	if err := conn.WriteMessage(websocket.BinaryMessage, data); err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	// Lắng nghe tin nhắn từ server
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		var chatResp pb.ChatResponse
		if err := proto.Unmarshal(msg, &chatResp); err != nil {
			log.Println("Failed to unmarshal message:", err)
			continue
		}

		fmt.Printf("[%s] %s: %s\n",
			time.Unix(chatResp.Message.Timestamp, 0).Format(time.RFC3339),
			chatResp.Message.Sender,
			chatResp.Message.Content)
	}
}
