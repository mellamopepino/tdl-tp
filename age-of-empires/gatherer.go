package main

import (
	"ageofempires/websockets"
	"sync"
	"time"
  "math/rand"
)

// Recoleta recursos provenientes del canal de su recurso
func consume(rawResource <-chan int, wg *sync.WaitGroup, warehouse *Warehouse, material string, id int) {
	go func() {
		defer wg.Done()
		for amount := range rawResource {
			time.Sleep(time.Duration(rand.Intn(5)) * time.Second) // Working...
      websockets.ShowMessage("START_GATHER %v", material)
			time.Sleep(time.Duration(rand.Intn(5)) * time.Second) // Working...
			gatheredAmount := amount * 10
      websockets.ShowMessage("FINISHED_GATHER %v %v", material, gatheredAmount)

			warehouse.Add(material, gatheredAmount)
		}
	}()
}
