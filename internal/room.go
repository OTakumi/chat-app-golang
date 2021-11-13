package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type room struct {
	// forward is a channel that holds messages to be forwarded to other clients
	forward chan []byte
	// join is a channel for clients who are trying to join a chat room
	join chan *client
	// leave is a channel for clients who are about to leave a chat room
	leave chan *client
	// All clients present in the room will be retained in clients.
	clients map[*client]bool
}

func (r *room) run() {
	for {
		select {
		// processing when join
		case client := <-r.join:
			r.clients[client] = true

		// processing when leave
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)

		case msg := <-r.forward:
			// send message to all client
			for client := range r.clients {
				select {
				case client.send <- msg:
					// send message
				default:
					// Message transmission failure
					delete(r.clients, client)
					close(client.send)
				}
			}
		}
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  socketBufferSize,
	WriteBufferSize: socketBufferSize,
}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}

	client := &client{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   r,
	}

	r.join <- client
	defer func() { r.leave <- client }()
	go client.write()
	client.read()
}
