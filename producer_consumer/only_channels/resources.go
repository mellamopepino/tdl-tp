package main

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

func MakeResources() Resources {
	return Resources{
		reads: make(chan readResource),
		writes: make(chan writeResource),
		gets: make(chan getResource),
	}
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
