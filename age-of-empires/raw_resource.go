package main

var materials = []int{10, 10, 10, 20, 20, 20, 30, 30, 30, 40, 40, 40}

// Recibir recursos y mandar esos recursos a un channel
// Ahora mismo los saca de un array, pero se puede cambiar por cualquier cosa (por ejemplo, web sockets)
func produce(rawResource chan<- int, material string) {
	go func() {
		for _, amount := range materials {
			showMessage("New resources discovered! %v: %v", material, amount)
			rawResource <- amount
		}
		close(rawResource)
	}()
}
