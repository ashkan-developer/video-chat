package main

import (
	"log"
	"net/http"

	"github.com/ashkan-developer/video-chat/server"
)

func main() {

	http.Handle("/create", server.CreateRoomRequestHandler)
	http.Handle("/join", server.JoinRoomRequestHandler)

	log.Println("Starting server on port 8000")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
