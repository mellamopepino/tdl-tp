package main

import (
	"ageofempires/websockets"
	"math/rand"
	"sync"
	"time"
)

// Recoleta recursos provenientes del canal de su recurso
func consume(rawResource <-chan int, wg *sync.WaitGroup, warehouse *Warehouse, material string, id int) {
	go func() {
		defer wg.Done()
		for amount := range rawResource {
			websockets.SendMessage("START_GATHER %v", material)
			time.Sleep(time.Duration(rand.Intn(4)+1) * time.Second) // Working...
			gatheredAmount := amount * 10
			websockets.SendMessage("FINISHED_GATHER %v %v", material, gatheredAmount)

			warehouse.Add(material, gatheredAmount)
		}
	}()
}
