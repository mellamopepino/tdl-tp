package main

import (
	"fmt"
)

// Dead simple producer consumer example
// https://medium.com/better-programming/hands-on-go-concurrency-the-producer-consumer-pattern-c42aab4e3bd2

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
