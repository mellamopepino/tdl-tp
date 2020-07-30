package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var messages = make(chan string, 1000)

func showMessage(message string, variables ...interface{}) {
	filledMessage := fmt.Sprintf(message, variables...)
	fmt.Println(filledMessage)
	messages <- filledMessage
}

func sendHandler(w http.ResponseWriter, r *http.Request) {
	ws, _ := upgrader.Upgrade(w, r, nil)
	for {
		message := <-messages
		ws.WriteMessage(1, []byte(message))
	}
}
