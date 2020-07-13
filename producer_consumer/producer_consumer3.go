package main

import (
	"fmt"
	"sync"
)

// Producer consumer + mutex example
// Multiple producers, multiple consumers
// Consumers share one SafeCounter, which acts as monitor for a map and a counter, using a mutex
// TODO: try to implement the same example using a channel instead of a mutex

type SafeCounter struct {
	resources     map[string]int
	mux           sync.Mutex
	houses        int
	finishCounter int
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
		return true
	}
	return false
}

func (counter *SafeCounter) Sum() {
	counter.mux.Lock()
	counter.houses++
	counter.mux.Unlock()
}

func (counter *SafeCounter) Finish() {
	counter.mux.Lock()
	counter.finishCounter++
	counter.mux.Unlock()
}

var messages = [][]int{
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

const producerCount int = 4
const consumerCount int = 3
const builderCount int = 2

// Send data from messages to channel
func produce(link chan<- int, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, amount := range messages[id] {
		link <- amount
	}
}

// Read data from channel and add it to the counter
func consume(link <-chan int, wg *sync.WaitGroup, counter *SafeCounter) {
	defer wg.Done()
	for amount := range link {
		counter.Add("wood", amount)
	}
	counter.Finish()
}

// Read from the counter and build houses
// If there are no resources left and all consumers have finished, it stops
func build(wg *sync.WaitGroup, counter *SafeCounter) {
	defer wg.Done()
	for {
		built := counter.Substract("wood", 15)
		if built {
			counter.Sum()
		} else {
			if counter.finishCounter == consumerCount {
				return
			}
		}
	}
}

func main() {
	counter := SafeCounter{resources: make(map[string]int), houses: 0, finishCounter: 0}
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
		go consume(link, wc, &counter)
	}

	for i := 0; i < builderCount; i++ {
		go build(wb, &counter)
	}

	wp.Wait()
	close(link)
	wc.Wait()
	wb.Wait()

	fmt.Println("Resources left: ", counter.resources)
	fmt.Println("Number of houses built:", counter.houses)
	fmt.Println("Finished consumers", counter.finishCounter)
}
