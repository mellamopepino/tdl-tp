package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
} // use default options

func receive(w http.ResponseWriter, r *http.Request) {
	ws, _ := upgrader.Upgrade(w, r, nil)
	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("Error al leer, bye")
			break
		}
		fmt.Println(string(message))
	}
	ws.Close()
}

func send(w http.ResponseWriter, r *http.Request) {
	ws, _ := upgrader.Upgrade(w, r, nil)
	for {
		ws.WriteMessage(1, []byte("Este es un mensaje del server al front"))
		time.Sleep(5 * time.Second)
	}
}

func main() {
	http.HandleFunc("/receive", receive)
	http.HandleFunc("/send", send)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
