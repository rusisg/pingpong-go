package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", pongHandler)

	log.Println("Starting server at :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error while starting server -> %v\n", err)
	}
}

func pongHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}
