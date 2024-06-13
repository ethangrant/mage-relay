package socket

import "github.com/gorilla/websocket"

type Relay struct {
	Clients map[int]*websocket.Conn
}

func NewRelay() Relay {
	return Relay{Clients: make(map[int]*websocket.Conn)}
}