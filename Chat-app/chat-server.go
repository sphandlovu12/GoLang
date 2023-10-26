package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	clients = make(map[*websocket.Conn]bool)
)

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	clients[conn] = true

	for {
		var msg string
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error: %v", err)
			delete(clients, conn)
			break
		}

		for client := range clients {
			if client != conn {
				err := client.WriteJSON(msg)
				if err != nil {
					log.Printf("Error: %v", err)
					client.Close()
					delete(clients, client)
				}
			}
		}
	}
}

func main() {
	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "chat.html")
	})

	http.HandleFunc("/ws", handleConnections)

	fmt.Println("Server is running on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
