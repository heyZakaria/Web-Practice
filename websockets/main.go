package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	setupAPI()
	fmt.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setupAPI() {
	manager := NewManager()
	http.Handle("/", http.FileServer(http.Dir("./frontend")))

	http.HandleFunc("/ws", manager.serveWS)
}
