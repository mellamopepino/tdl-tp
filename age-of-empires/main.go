package main

import (
	"fmt"
	"sync"
)

const builders = 2

func produceAndConsumeResource(name string, workers int, warehohuse *Warehouse) *sync.WaitGroup {
	resourceChannel := make(chan int, 20)

	wg := &sync.WaitGroup{}
	wg.Add(workers)

	produce(resourceChannel, name)
	for i := 0; i < workers; i++ {
		consume(resourceChannel, wg, warehohuse, name, i+1)
	}

	return wg
}

func main() {
	var resources []string
	var gatherers []int
	var err bool

	resources, gatherers, err = ReadConfig("resources.csv")
	if err {
		fmt.Println("Irrecoverable error, exiting...")
		return
	}
	var resourcesQ int = len(resources)
	var gatherersWaitGroups []*sync.WaitGroup

	// Warehouse guarda los recursos ya listos para usar
	warehouse := MakeWarehouse()
	warehouse.Listen()

	// Por cada recurso generamos un "producer" y múltiples "consumers".
	// Los producers envian por el canal del recurso los recursos disponibles "en el mapa"
	// Los consumers "cosechan" esos recursos y los agregan al warehouse
	for i := 0; i < resourcesQ; i++ {
		resourceWaitGroup := produceAndConsumeResource(resources[i], gatherers[i], warehouse)
		gatherersWaitGroups = append(gatherersWaitGroups, resourceWaitGroup)
	}

	buildersWaitGroup := &sync.WaitGroup{}
	buildersWaitGroup.Add(builders)

	// Generamos constructores que toman recursos del warehouse y los transforman en escudos y espadas
	for i := 0; i < builders; i++ {
		if i < builders/2 {
			build(warehouse, buildersWaitGroup, "shield", resources[:resourcesQ/2+1], 10, i+1)
		} else {
			build(warehouse, buildersWaitGroup, "sword", resources[resourcesQ/2:], 10, i+1)
		}
	}

	// Esperamos que los consumers (gatherers) terminen de cosechar recursos y les avisamos a los builders
	for _, wg := range gatherersWaitGroups {
		wg.Wait()
	}
	showMessage("All gatherers finished")
	warehouse.done = true
	// Esperamos que los builders terminen y mostramos los recursos finales
	buildersWaitGroup.Wait()

	fmt.Println(warehouse.GetAll())
}
