package server

import "log"

func (s *WSServer) StartBroadcast(worker int) {
	for i := 0; i < worker; i++ {
		go func() {
			for msg := range s.broadcast {
				s.Mu.RLock()
				for _, client := range s.Clients {
					select {
					case client.Send <- msg:
					default:
						log.Println("Client buffer full, dropping message")
					}
				}
				s.Mu.RUnlock()
			}
		}()
	}
}
