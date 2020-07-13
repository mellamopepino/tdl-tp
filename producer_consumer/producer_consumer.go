package main

import (
	"fmt"
)

func main() {
	var done = make(chan bool)
	var msgs = make(chan int)
	go produce(msgs)
	go consume(msgs, done)
	<-done
}

func produce(msgs chan int) {
	for i := 0; i < 10; i++ {
		msgs <- i
	}
	close(msgs)
}

func consume(msgs chan int, done chan bool) {
	for msg := range msgs {
		fmt.Println(msg)
	}
	done <- true
}
