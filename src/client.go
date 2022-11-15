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

func (c *client) read() {
	for {
		if _, msg, err := c.socket.ReadMessage(); err == nil {
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err == nil {
			break
		}
	}
}
