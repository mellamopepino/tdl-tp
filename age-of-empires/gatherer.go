package main

import (
	"ageofempires/websockets"
	"sync"
	"time"
)

// Recoleta recursos provenientes del canal de su recurso
func consume(rawResource <-chan int, wg *sync.WaitGroup, warehouse *Warehouse, material string, id int) {
	go func() {
		defer wg.Done()
		for amount := range rawResource {
			websockets.ShowMessage("%v worker number %v started to gather %v", material, id, material)
			time.Sleep(2 * time.Second) // Working...
			gatheredAmount := amount * 10
			websockets.ShowMessage("%v worker number %v finished gathering %v of %v", material, id, gatheredAmount, material)

			warehouse.Add(material, gatheredAmount)
		}
	}()
}
