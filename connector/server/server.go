package server

import (
	"net/http"
	"sync"

	pb "ws-server/proto"

	"github.com/gorilla/websocket"
)

type Connection struct {
	Conn *websocket.Conn
	Send chan []byte
}

type WSServer struct {
	upgrader  websocket.Upgrader
	Clients   map[*websocket.Conn]*Connection
	broadcast chan []byte
	Ack       chan *pb.RequestEvent
	Mu        sync.RWMutex
}

func NewWSServer() *WSServer {
	return &WSServer{
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		},
		Clients:   make(map[*websocket.Conn]*Connection),
		broadcast: make(chan []byte, 1000000),
		Ack:       make(chan *pb.RequestEvent, 1000000),
	}
}

func (s *WSServer) AddClient(conn *websocket.Conn) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	s.Clients[conn] = &Connection{
		Conn: conn,
		Send: make(chan []byte, 256),
	}
}

func (s *WSServer) RemoveClient(conn *websocket.Conn) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	delete(s.Clients, conn)
	conn.Close()
}
