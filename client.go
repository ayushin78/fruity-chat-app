package main

import (
	"log"

	"github.com/gorilla/websocket"
)

// A client represents a single chatting user
type client struct {
	conn         *websocket.Conn // websocket connection for this client
	room         *room           // room in which the client is currently chatting in
	sentMessages chan []byte     // channel for messages sent by this client
}

func (c *client) read() {
	for {
		// read in a message from client
		_, p, err := c.conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		// broadcast this message on the room
		c.room.broadcast <- p
	}

	// In case there is any error, break from the loop and close the connection without waiting for a close message
	c.conn.Close()
}
