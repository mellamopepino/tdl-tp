package main

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

func MakeProducers() Producers {
	return Producers{
		dones: make(chan doneProducer),
		allDone: make(chan allDone),
		quit: make(chan quit),
	}
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
