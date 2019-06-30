package main

import (
	"log"

	"github.com/gorilla/websocket"
)

// A client represents a single chatting user
type client struct {
	conn *websocket.Conn // websocket connection for this client
	room *room           // room in which the client is currently chatting in
	send chan []byte     // channel for messages sent by this client
}

/*
 *	This is a receiver function of type client. It subscribes to the web socket
 *	and read if theres is any message received from the client. Any incoming message
 *	is then sent to the broadcast channel of the room in which the client is currently
 *	chatting.
 */
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

/*
 *	This is a receiver function of type client. It subscribes on the send
 * 	of the client. Any message on the end channel is read from the channel
 *  and then written on the web socket.
 */
func (c *client) write() {
	for msg := range c.send {
		if err := c.conn.WriteMessage(1, msg); err != nil {
			log.Println(err)
			break
		}
	}
	c.conn.Close()
}
