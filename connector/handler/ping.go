package handler

import (
	"encoding/json"
	"log"
	pb "ws-server/proto"
	"ws-server/server"

	"google.golang.org/protobuf/proto"
)

func StartPing(s *server.WSServer, worker int) {
	for i := 0; i < worker; i++ {
		go func() {
			for msg := range s.Ack {
				var pingValue Ping
				if err := json.Unmarshal([]byte(msg.GetPayload()), &pingValue); err != nil {
					log.Printf("[ERROR] Unmarshal failed: %v, Payload: %s\n", err, msg.GetPayload())
					continue
				}

				ackJson, err := json.Marshal(Ack{Value: pingValue.Value})
				if err != nil {
					log.Printf("[ERROR] Marshal JSON failed: %v\n", err)
					continue
				}

				rsEvent, err := proto.Marshal(&pb.ResponseEvent{
					Type: pb.EventType_ACK,
					Content: &pb.ResponseEvent_Payload{
						Payload: string(ackJson),
					},
				})
				if err != nil {
					log.Printf("[ERROR] Marshal Protobuf failed: %v\n", err)
					continue
				}

				// Lock chỉ khi cần gửi dữ liệu
				s.Mu.RLock()
				clients := make([]*server.Connection, 0, len(s.Clients))
				for _, client := range s.Clients {
					clients = append(clients, client)
				}
				s.Mu.RUnlock()

				// Gửi dữ liệu ra ngoài vòng lock
				for _, client := range clients {
					select {
					case client.Send <- rsEvent:
					default:
						log.Println("[WARNING] Client buffer full, dropping message")
					}
				}
			}
		}()
	}
}

type Ping struct {
	Value string `json:"value"`
}

type Ack struct {
	Value string `json:"value"`
}
