package main

import (
	"fmt"
	"net/http"
	"sync"
)

type Resource struct {
	Name      string
	Gatherers int
}

type Weapon struct {
	Name      string
	Builders  int
	Materials map[string]int
}

func main() {
	resources, weapons, err := loadConfig()
	if err {
		fmt.Println("Fatal error. Exiting...")
		return
	}

	// Warehouse guarda los recursos ya listos para usar
	warehouse := MakeWarehouse()
	warehouse.Listen()

	// Por cada recurso generamos un "producer" y múltiples "consumers".
	// Los producers envian por el canal del recurso los recursos disponibles "en el mapa"
	// Los consumers "cosechan" esos recursos y los agregan al warehouse
	var gatherersWaitGroups []*sync.WaitGroup
	for _, resource := range resources {
		resourceWaitGroup := produceAndConsumeResource(resource, warehouse)
		gatherersWaitGroups = append(gatherersWaitGroups, resourceWaitGroup)
	}

	buildersWaitGroup := &sync.WaitGroup{}

	// Generamos constructores que toman recursos del warehouse y los transforman en escudos y espadas
	for _, weapon := range weapons {
		builders := weapon.Builders
		for i := 0; i < builders; i++ {
			buildersWaitGroup.Add(1)
			build(warehouse, buildersWaitGroup, weapon, i+1)
		}
	}

	http.HandleFunc("/send", sendHandler)
	go http.ListenAndServe(":8080", nil)

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

func loadConfig() (resources []Resource, weapons []Weapon, err bool) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("File error:", r)
			err = true
		}
	}()

	resources = ReadResourcesConfig("resources.json")
	weapons = ReadWeaponsConfig("weapons.json")
	return
}

func produceAndConsumeResource(resource Resource, warehohuse *Warehouse) *sync.WaitGroup {
	resourceChannel := make(chan int, 20)

	wg := &sync.WaitGroup{}
	wg.Add(resource.Gatherers)

	produce(resourceChannel, resource.Name)
	for i := 0; i < resource.Gatherers; i++ {
		consume(resourceChannel, wg, warehohuse, resource.Name, i+1)
	}

	return wg
}
