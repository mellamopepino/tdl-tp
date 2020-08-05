package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"ageofempires/websockets"
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
	websockets.Init(startGame)
}

func startGame() {
	start := time.Now()

	// Cargamos configuracion
	resources, weapons, err := loadConfig()
	if err {
		fmt.Println("Fatal error. Exiting...")
		return
	}

	// Creamos warehouse
	warehouse := MakeWarehouse()

	// Generamos recursos y recolectores
	var gatherersWaitGroups []*sync.WaitGroup
	for _, resource := range resources {
		websockets.SendMessage("NEW_GATHERERS %v", resource.Gatherers)
		resourceWaitGroup := produceAndConsumeResource(resource, warehouse)
		gatherersWaitGroups = append(gatherersWaitGroups, resourceWaitGroup)
	}

	buildersWaitGroup := &sync.WaitGroup{}

	// Generamos constructores
	for _, weapon := range weapons {
		builders := weapon.Builders
		websockets.SendMessage("NEW_BUILDERS %v", builders)
		for i := 0; i < builders; i++ {
			buildersWaitGroup.Add(1)
			build(warehouse, buildersWaitGroup, weapon)
		}
	}

	// Esperamos que los recolectores terminen
	for _, wg := range gatherersWaitGroups {
		wg.Wait()
	}
	websockets.SendMessage("FINISH_ALL_GATHERERS")
	warehouse.done = true

	// Esperramos que los constructores terminen
	buildersWaitGroup.Wait()
	websockets.SendMessage("FINISH_ALL_BUILDERS")

	elapsed := time.Since(start)
	websockets.SendMessage("TOTAL_TIME %v", elapsed)

	time.Sleep(5 * time.Second)
	os.Exit(0)
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

// Por cada recurso se crea un canal por el cual los recolectores obtienen los recursos
func produceAndConsumeResource(resource Resource, warehohuse *Warehouse) *sync.WaitGroup {
	resourceChannel := make(chan int, 20)

	wg := &sync.WaitGroup{}
	wg.Add(resource.Gatherers)

	produce(resourceChannel, resource.Name)
	for i := 0; i < resource.Gatherers; i++ {
		consume(resourceChannel, wg, warehohuse, resource.Name)
	}

	return wg
}
