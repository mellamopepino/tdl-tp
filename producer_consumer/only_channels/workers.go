package main

type doneProducer struct {
	id int
}

type producersDone struct {
	resp chan bool
}

type doneConsumer struct {
	id int
}

type WorkersHandler struct {
	doneProducerOp chan doneProducer
	producersDoneOp chan producersDone
	doneConsumerOp chan doneConsumer
}

func MakeWorkersHandler() WorkersHandler {
	return WorkersHandler{
		doneProducerOp: make(chan doneProducer),
		producersDoneOp: make(chan producersDone),
		doneConsumerOp: make(chan doneConsumer),
	}
}

func (p WorkersHandler) Listen(totalP int, totalC int, wait chan bool) {
	go func() {
		var producers = 0
		var consumers = 0
		for {
			select {
			case <-p.doneProducerOp:
				producers++
			case pd := <-p.producersDoneOp:
				pd.resp <- (producers == totalP)
			case <-p.doneConsumerOp:
				consumers++
				if(consumers == totalC) {
					wait <- true
				}
			}
		}
	}()
}

func (wh WorkersHandler) DoneProducer(id int) {
	d := doneProducer{ id: id }
	wh.doneProducerOp <- d
}

func (wh WorkersHandler) ProducersDone() (bool) {
	pd := producersDone{ resp: make(chan bool) }
	wh.producersDoneOp <- pd
	return <-pd.resp
}

func (wh WorkersHandler) DoneConsumer(id int) {
	d := doneConsumer{ id: id }
	wh.doneConsumerOp <- d
}
