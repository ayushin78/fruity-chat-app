package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/ws", http.HandlerFunc(wsEndpoint))
	fmt.Println("Server started at port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("ListenAndServe : ", err)
	}

}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// func wsEndpoint(w http.ResponseWriter, r *http.Request) {
// 	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

// 	// upgrade this connection to a WebSocket
// 	// connection
// 	ws, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	// helpful log statement to show connections
// 	log.Println("Client Connected")

// 	reader(ws)
// }
