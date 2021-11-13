package main

import (
	"github.com/gorilla/websocket"
)

type client struct {
	// this socket is websocket for client
	socket *websocket.Conn
	// The send is channel sent message
	send chan []byte
	// The room is chat room which client is joining
	room *room
}
