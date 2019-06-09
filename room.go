package main

// A room represents a chatting room in a chat app
type room struct {
	clients    map[*client]bool // clients contains all the clients registered in this room
	register   chan client      // register channel holds clients willing to join this room
	unregister chan client      // unregister channel holds clients willing to leave this room
	message    chan []byte      // message channel holds incoming messages to be forwarded to all the clients in this room
}
