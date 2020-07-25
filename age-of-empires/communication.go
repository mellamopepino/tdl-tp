package main

import "fmt"

// Envia mensajes a donde corresponda
// Ahora mismo es por consola, pero podrían ser web sockets
func showMessage(message string, variables ...interface{}) {
	filledMessage := fmt.Sprintf(message, variables...)
	fmt.Println(filledMessage)
}
