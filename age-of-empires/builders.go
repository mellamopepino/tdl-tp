package main

import (
	"ageofempires/websockets"
	"sync"
	"time"
)

// Busca recursos en el warehouse y construye cosas
// Si no puede construir por falta de recursos y los recolectores terminaron, termina
func build(warehouse *Warehouse, wg *sync.WaitGroup, weapon Weapon, id int) {
	go func() {
		defer wg.Done()
		for {
			websockets.ShowMessage("Builder number %v is trying to build %v", id, weapon.Name)
			ok := warehouse.Use(weapon.Materials)
			if ok {
				time.Sleep(3 * time.Second) // Working...
				websockets.ShowMessage("Builder number %v finished building %v", id, weapon.Name)
				warehouse.Add(weapon.Name, 1)
			} else {
				websockets.ShowMessage("Builder number %v couldn't build %v", id, weapon.Name)
				time.Sleep(1 * time.Second) // Waiting for gatherers to finish...
				if warehouse.done {
					return
				}
			}
		}
	}()
}
