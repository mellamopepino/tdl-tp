package main

import (
	"fmt"
	"sync"
)

// Producer consumer + shared state. Concurrency is handled using a stateful channel
// Multiple producers, multiple consumers

type readOp struct {
	key    string
	amount int
	resp   chan bool
}

type writeOp struct {
	key string
	val int
}

type printOp struct {
	resources chan map[string]int
}

func share(reads <-chan readOp, writes <-chan writeOp, prints <-chan printOp) {
	var resources = make(map[string]int)
	for {
		select {
		case read := <-reads:
			if resources[read.key] >= read.amount {
				resources[read.key] -= read.amount
				read.resp <- true
				resources["houses"]++
				fmt.Printf("Usando 15 de madera. Total: %v\n", resources[read.key])
			} else {
				read.resp <- false
			}
		case write := <-writes:
			resources[write.key] += write.val
			fmt.Printf("Agregando %v de madera. Total: %v\n", write.val, resources[write.key])
		case print := <-prints:
			print.resources <- resources
			return
		}
	}
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

// Read data from channel and send it to the stateful channel
func consume(link <-chan int, wg *sync.WaitGroup, writes chan<- writeOp) {
	defer wg.Done()
	for amount := range link {
		write := writeOp{
			key: "wood",
			val: amount}
		writes <- write
	}
}

// Read from the stateful channel and build houses
// If there are no resources left and all consumers have finished, it stops
func build(wg *sync.WaitGroup, reads chan<- readOp, producersFinished *bool) {
	defer wg.Done()
	for {
		read := readOp{
			key:    "wood",
			amount: 15,
			resp:   make(chan bool)}
		reads <- read
		built := <-read.resp
		if !built {
			fmt.Println("No pude construir >:c")
			if *producersFinished {
				fmt.Println("Terminaron los productores y no hay madera, nos re vimos")
				return
			}
		}
	}
}

func main() {
	reads := make(chan readOp)
	writes := make(chan writeOp)
	prints := make(chan printOp)

	link := make(chan int, 100)
	wp := &sync.WaitGroup{}
	wc := &sync.WaitGroup{}
	wb := &sync.WaitGroup{}

	wp.Add(producerCount)
	wc.Add(consumerCount)
	wb.Add(builderCount)

	go share(reads, writes, prints)

	for i := 0; i < producerCount; i++ {
		go produce(link, i, wp)
	}

	for i := 0; i < consumerCount; i++ {
		go consume(link, wc, writes)
	}

	producersFinished := false
	for i := 0; i < builderCount; i++ {
		go build(wb, reads, &producersFinished)
	}

	wp.Wait()
	close(link)
	wc.Wait()
	producersFinished = true
	wb.Wait()

	print := printOp{
		resources: make(chan map[string]int)}
	prints <- print
	fmt.Println("Resources: ", <-print.resources)
}
