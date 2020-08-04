package main

import "ageofempires/websockets"

var materials = []int{1, 4, 1, 3, 1, 2, 1, 5, 1, 2, 1, 1}

// Recibir recursos y mandar esos recursos a un channel
// Ahora mismo los saca de un array, pero se puede cambiar por cualquier cosa (por ejemplo, web sockets)
func produce(rawResource chan<- int, material string) {
	go func() {
		for _, amount := range materials {
			for i := 0; i < amount; i++ {
				websockets.ShowMessage("NEW_RESOURCES %v %v", material, 1)
				rawResource <- 1
			}
		}
		close(rawResource)
	}()
}
