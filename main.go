package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"text/template"

	"github.com/gorilla/websocket"
)

const (
	readBufferSize  = 1024
	writeBufferSize = 1024
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", ServeTemplate)
	room := &room{
		clients:    make(map[*client]bool),
		register:   make(chan *client),
		unregister: make(chan *client),
		broadcast:  make(chan []byte),
	}
	mux.HandleFunc("/room", room.roomHandler)
	go room.run()
	fmt.Println("Server started at port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("ListenAndServe : ", err)
	}

}

//ServeTemplate does this
func ServeTemplate(w http.ResponseWriter, r *http.Request) {
	lp := path.Join("static", "index.html")

	templates, err := template.ParseFiles(lp)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "500 Internal Server Error", 500)
		return
	}
	// fmt.Println(lp, fp)
	templates.ExecuteTemplate(w, "index.html", nil)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  readBufferSize,
	WriteBufferSize: writeBufferSize,
}

func (r *room) roomHandler(w http.ResponseWriter, req *http.Request) {

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// upgrade this connection to a WebSocket  connection
	ws, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println(err)
	}

	client := &client{
		conn: ws,
		room: r,
		send: make(chan []byte, writeBufferSize),
	}

	r.register <- client
	fmt.Println("Client registered")
	defer func() { r.unregister <- client }()

	log.Println("Client Connected")
	go client.write()
	client.read()

}
