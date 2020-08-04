package websockets

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var messages = make(chan string, 1000)

func Init(callback func()) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		sendHandler(w, r, callback)
	}

	http.HandleFunc("/send", handler)
	http.ListenAndServe(":8080", nil)
}

func ShowMessage(message string, variables ...interface{}) {
	filledMessage := fmt.Sprintf(message, variables...)
	messages <- filledMessage
}

func sendHandler(w http.ResponseWriter, r *http.Request, callback func()) {
	ws, _ := upgrader.Upgrade(w, r, nil)
	go callback()
	for {
		message := <-messages
		ws.WriteMessage(1, []byte(message))
	}
}
