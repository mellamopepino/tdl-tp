package main

import (
	"fmt"
)

// Producer consumer, just with channels example.
// Multiple producers, multiple consumers

const producerCount int = 4
const builderCount int = 2

var materials = [][]int{
	{
		10, 10, 10,
	},
	{
		20, 20, 20,
	},
	{
		30, 30, 30,
	},
	{
		40, 40, 40,
	},
}

// Send data from messages to channel
func produce(resources ResourcesHandler, id int, workers WorkersHandler) {
	for _, amount := range materials[id] {
		res, ok := resources.Add("wood", amount)
		fmt.Println("Producing wood:", res, ok)
	}
	fmt.Println("Producer done:", id)
	workers.DoneProducer(id)
}

// Read from the counter and build houses
// If there are no resources left and all consumers have finished, it stops
func build(resources ResourcesHandler, id int, workers WorkersHandler) {
	for {
		res, ok := resources.Build("house", "wood", 15)
		fmt.Println("Building with wood:", res, ok)

		if(!ok) {
			fmt.Println("Can not build :(")
			if(workers.ProducersDone()) {
				workers.DoneConsumer(id)
				fmt.Println("Nos re vimo")
				return
			}
		}
	}
}

func main() {
	wait := make(chan bool)

	workers := MakeWorkersHandler()
	workers.Listen(producerCount, builderCount, wait)

	resources := MakeResourcesHandler()
	resources.Listen()

	for i := 0; i < producerCount; i++ {
		go produce(resources, i, workers)
	}

	for i := 0; i < builderCount; i++ {
		go build(resources, i, workers)
	}

	<-wait
	fmt.Println("Resources left and houses built:", resources.GetAll())
}
