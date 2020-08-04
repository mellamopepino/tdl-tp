package main

import (
	"ageofempires/websockets"
	"math/rand"
	"sync"
	"time"
)

// Busca recursos en el warehouse y construye cosas
// Si no puede construir por falta de recursos y los recolectores terminaron, termina
func build(warehouse *Warehouse, wg *sync.WaitGroup, weapon Weapon, id int) {
	go func() {
		defer wg.Done()
		for {
			ok := warehouse.Use(weapon.Materials)
			if ok {
				websockets.ShowMessage("START_BUILD %v %v", weapon.Name, weapon.Materials)
				time.Sleep(time.Duration(rand.Intn(5)+5) * time.Second) // Working...
				websockets.ShowMessage("FINISHED_BUILD %v %v", weapon.Name, weapon.Materials)
				warehouse.Add(weapon.Name, 1)
			} else {
				websockets.ShowMessage("FAIL_BUILD %v", weapon.Name)
				if warehouse.done {
					return
				}
				time.Sleep(1 * time.Second) // Waiting for gatherers to finish...
			}
		}
	}()
}
