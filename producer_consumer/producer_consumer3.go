package main

import (
	"fmt"
	"sync"
)

// Producer consumer + shared state. Explicit concurrency handling using a mutex
// Multiple producers, multiple consumers
// Consumers share one SafeCounter, which acts as monitor for a map and a counter, using a mutex

// Nota: No se puede hacer esto mismo con sync.Map (mapa thread safe) porque no provee la interfaz que necesitamos
// En el metodo Substract estamos chequeando que el valor guardado sea mayor a amount y si es mayor, restamos amount al valor guardado
// El tema es que sync.Map no provee una operacion atomica para chequear+guardar, lo cual hace posible una race condition

type SafeCounter struct {
	resources map[string]int
	mux       sync.Mutex
}

func (counter *SafeCounter) Add(resource string, amount int) {
	counter.mux.Lock()
	counter.resources[resource] += amount
	counter.mux.Unlock()
}

func (counter *SafeCounter) Substract(resource string, amount int) bool {
	counter.mux.Lock()
	defer counter.mux.Unlock()
	if counter.resources[resource] >= amount {
		counter.resources[resource] -= amount
		counter.resources["houses"]++
		return true
	}
	return false
}

var messages = [][]int{
	{
		10, 10, 10,
	},
	{
		20, 20, 20,
	},
	{
		30, 35, 30,
	},
	{
		40, 40, 40,
	},
}

// Send data from messages to channel
func produce(link chan<- int, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, amount := range messages[id] {
		link <- amount
	}
}

// Read data from channel and add it to the counter
func consume(link <-chan int, wg *sync.WaitGroup, counter *SafeCounter, resource string) {
	defer wg.Done()
	for amount := range link {
		counter.Add(resource, amount)
	}
}

// Read from the counter and build houses
// If there are no resources left and all consumers have finished, it stops
func build(wg *sync.WaitGroup, counter *SafeCounter, producersFinished *bool, resource string, amount int) {
	defer wg.Done()
	for {
		built := counter.Substract(resource, amount)
		if !built {
			if *producersFinished {
				return
			}
		}
	}
}

func main() {
	const producerCount int = 4
	const consumerCount int = 3
	const builderCount int = 2

	counter := SafeCounter{resources: make(map[string]int)}
	link := make(chan int, 100)

	wp := &sync.WaitGroup{}
	wc := &sync.WaitGroup{}
	wb := &sync.WaitGroup{}

	wp.Add(producerCount)
	wc.Add(consumerCount)
	wb.Add(builderCount)

	for i := 0; i < producerCount; i++ {
		go produce(link, i, wp)
	}

	for i := 0; i < consumerCount; i++ {
		go consume(link, wc, &counter, "wood")
	}

	producersFinished := false
	for i := 0; i < builderCount; i++ {
		go build(wb, &counter, &producersFinished, "wood", 15)
	}

	wp.Wait()
	close(link)
	wc.Wait()
	producersFinished = true
	wb.Wait()

	fmt.Println("Resources: ", counter.resources)
}
