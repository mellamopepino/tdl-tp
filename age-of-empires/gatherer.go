package main

import (
	"sync"
	"time"
)

// Recoleta recursos provenientes del canal de su recurso
func consume(rawResource <-chan int, wg *sync.WaitGroup, warehouse *Warehouse, material string, id int) {
	go func() {
		defer wg.Done()
		for amount := range rawResource {
			showMessage("%v worker number %v started to gather %v of %v", material, id, amount, material)
			time.Sleep(2 * time.Second) // Working...
			showMessage("%v worker number %v finished gathering %v of %v", material, id, amount, material)

			warehouse.Add(material, amount)
		}
	}()
}
