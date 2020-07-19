package main

import (
	"fmt"
)

// Producer consumer, just with channels example.
// Multiple producers, multiple consumers

type readResource struct {
	key  string
	resp chan int
}

type writeResource struct {
	key  string
	val  int
	resp chan bool
}

type getResource struct {
	resp chan map[string]int
}

type Resources struct {
	reads chan readResource
	writes chan writeResource
	gets chan getResource
}

func (r Resources) Listen() {
	go func() {
		var resources = make(map[string]int)
		for {
			select {
			case read := <-r.reads:
				read.resp <- resources[read.key]
			case write := <-r.writes:
				resources[write.key] = write.val
				write.resp <- true
			case get := <-r.gets:
				get.resp <- resources
			}
		}
	}()
}

func (r Resources) Add(key string, amount int) (int, bool) {
	read := readResource{
		key: key,
		resp: make(chan int),
	}
	r.reads <- read
	resource := <-read.resp

	write := writeResource{
		key: key,
		val: resource + amount,
		resp: make(chan bool),
	}
	r.writes <- write

	return resource + amount, <-write.resp
}

func (r Resources) Use(key string, amount int) (int, bool) {
	read := readResource{
		key: key,
		resp: make(chan int),
	}
	r.reads <- read
	resource := <-read.resp

	if((resource - amount) < 0){
		return resource, false
	}

	write := writeResource{
		key: key,
		val: resource - amount,
		resp: make(chan bool),
	}
	r.writes <- write

	return resource - amount, <-write.resp
}

func (r Resources) GetAll() map[string]int {
	get := getResource{ resp: make(chan map[string]int) }
	r.gets <- get
	return <-get.resp
}

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
func produce(resources Resources, id int) {
	for _, amount := range materials[id] {
		result, ok := resources.Add("wood", amount)
		fmt.Println("Producing wood:", result, ok)
	}
	fmt.Println("Producer done:", id)
}

// Read from the counter and build houses
// If there are no resources left and all consumers have finished, it stops
func build(resources Resources, done *bool) {
	for {
		amount, ok := resources.Use("wood", 15)
		fmt.Println("Building with wood:", amount, ok)

		if(!ok) {
			fmt.Println("Can not build :(")
		}
	}
}

func main() {
	resources := Resources{
		reads: make(chan readResource),
		writes: make(chan writeResource),
	}

	resources.Listen()

	for i := 0; i < producerCount; i++ {
		go produce(resources, i)
	}

	for i := 0; i < builderCount; i++ {
		go build(resources)
	}

	// No se puede usar resources.GetAll
	// Creo no me toma el pedido de gets porque siempre
	// termina entrando con un pedido de writes/reads que hace
	// el resources.Use, porque los build están siempre corriendo.
	// Y no encontré forma de pararlos (incluso con un booleano).
	fmt.Println("Resources left:", 10)
	fmt.Println("Number of houses built:", 0)
	fmt.Println("Finished consumers", 3)
}
