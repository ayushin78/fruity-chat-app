package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(displayHome))

	fmt.Println("Server started at port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("ListenAndServe : ", err)
	}

}

func displayHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello chat app"))
}
