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
func produce(resources Resources, id int, producers Producers) {
	for _, amount := range materials[id] {
		result, ok := resources.Add("wood", amount)
		fmt.Println("Producing wood:", result, ok)
	}
	fmt.Println("Producer done:", id)
	producers.Done(id)
}

// Read from the counter and build houses
// If there are no resources left and all consumers have finished, it stops
func build(resources Resources, id int, producers Producers) {
	for {
		amount, ok := resources.Use("wood", 15)
		fmt.Println("Building with wood:", amount, ok)

		if(!ok) {
			fmt.Println("Can not build :(")
			if(producers.AllDone()) {
				producers.Quit(id)
				fmt.Println("Nos re vimo")
				return
			}
		}
	}
}

func main() {
	wait := make(chan bool)

	producers := MakeProducers()
	producers.Listen(producerCount, builderCount, wait)

	resources := MakeResources()
	resources.Listen()

	for i := 0; i < producerCount; i++ {
		go produce(resources, i, producers)
	}

	for i := 0; i < builderCount; i++ {
		go build(resources, i, producers)
	}

	<-wait
	fmt.Println("Resources left:", resources.GetAll())
	fmt.Println("Number of houses built:", 0)
	fmt.Println("Finished consumers", 3)
}
