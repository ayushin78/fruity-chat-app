package main

import "github.com/gorilla/websocket"

// A client represents a single chatting user
type client struct {
	conn         *websocket.Conn // websocket connection for this client
	room         *room           // room in which the client is currently chatting in
	sentMessages chan []byte     // channel for messages sent by this client
}
