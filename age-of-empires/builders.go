package main

import (
	"ageofempires/websockets"
	"sync"
	"time"
  "math/rand"
)

// Busca recursos en el warehouse y construye cosas
// Si no puede construir por falta de recursos y los recolectores terminaron, termina
func build(warehouse *Warehouse, wg *sync.WaitGroup, weapon Weapon, id int) {
	go func() {
		defer wg.Done()
		for {
			time.Sleep(time.Duration(rand.Intn(5)) * time.Second) // Working...
			websockets.ShowMessage("START_BUILD %v", weapon.Name)
			ok := warehouse.Use(weapon.Materials)
			if ok {
				time.Sleep(time.Duration(rand.Intn(5)) * time.Second) // Working...
				websockets.ShowMessage("FINISHED_BUILD %v %v", weapon.Name, weapon.Materials)
				warehouse.Add(weapon.Name, 1)
			} else {
				websockets.ShowMessage("FAIL_BUILD %v", weapon.Name)
				time.Sleep(1 * time.Second) // Waiting for gatherers to finish...
				if warehouse.done {
					return
				}
			}
		}
	}()
}
