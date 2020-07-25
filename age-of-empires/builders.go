package main

import (
	"sync"
	"time"
)

// Busca recursos en el warehouse y construye cosas
// Si no puede construir por falta de recursos y los recolectores terminaron, termina
func build(warehouse *Warehouse, wg *sync.WaitGroup, thing string, materials []string, amount int, id int) {
	go func() {
		for {
			showMessage("Builder number %v is trying to build %v", id, thing)
			ok := warehouse.Use(amount, materials)
			if ok {
				time.Sleep(3 * time.Second) // Working...
				showMessage("Builder number %v finished building %v", id, thing)
				warehouse.Add(thing, 1)
			} else {
				showMessage("Builder number %v couldn't build %v", id, thing)
				time.Sleep(1 * time.Second) // Waiting for gatherers to finish...
				if warehouse.done {
					break
				}
			}
		}
		wg.Done()
	}()
}
