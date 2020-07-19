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

type doneProducer struct {
	id int
}

type allDone struct {
	resp chan bool
}

type quit struct {
	id int
}

type Producers struct {
	dones chan doneProducer
	allDone chan allDone
	quit chan quit
}

func (p Producers) Listen(totalP int, totalC int, wait chan bool) {
	go func() {
		var producers = 0
		var consumers = 0
		for {
			select {
			case <-p.dones:
				producers++
			case allDone := <-p.allDone:
				allDone.resp <- (producers == totalP)
			case <-p.quit:
				consumers++
				if(consumers == totalC) {
					wait <- true
				}
			}
		}
	}()
}

func (p Producers) Done(id int) {
	done := doneProducer{ id: id }
	p.dones <- done
}

func (p Producers) AllDone() (bool) {
	allDone := allDone{ resp: make(chan bool) }
	p.allDone <- allDone
	return <-allDone.resp
}

func (p Producers) Quit(id int) {
	quit := quit{ id: id }
	p.quit <- quit
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

	producers := Producers{
		dones: make(chan doneProducer),
		allDone: make(chan allDone),
		quit: make(chan quit),
	}
	producers.Listen(producerCount, builderCount, wait)

	resources := Resources{
		reads: make(chan readResource),
		writes: make(chan writeResource),
		gets: make(chan getResource),
	}
	resources.Listen()

	for i := 0; i < producerCount; i++ {
		go produce(resources, i, producers)
	}

	for i := 0; i < builderCount; i++ {
		go build(resources, i, producers)
	}

	<-wait
	// No se puede usar resources.GetAll
	// Creo no me toma el pedido de gets porque siempre
	// termina entrando con un pedido de writes/reads que hace
	// el resources.Use, porque los build están siempre corriendo.
	// Y no encontré forma de pararlos (incluso con un booleano).
	fmt.Println("Resources left:", resources.GetAll())
	fmt.Println("Number of houses built:", 0)
	fmt.Println("Finished consumers", 3)
}
