package main

import (
	"fmt"
	"sync"
)

// Producer consumer with multiple producers and multiple consumers
// https://medium.com/better-programming/hands-on-go-concurrency-the-producer-consumer-pattern-c42aab4e3bd2

var messages = [][]string{
	{
		"The world itself's",
		"just one big hoax.",
		"Spamming each other with our",
		"running commentary of bullshit,",
	},
	{
		"but with our things, our property, our money.",
		"I'm not saying anything new.",
		"We all know why we do this,",
		"not because Hunger Games",
		"books make us happy,",
	},
	{
		"masquerading as insight, our social media",
		"faking as intimacy.",
		"Or is it that we voted for this?",
		"Not with our rigged elections,",
	},
	{
		"but because we wanna be sedated.",
		"Because it's painful not to pretend,",
		"because we're cowards.",
		"- Elliot Alderson",
		"Mr. Robot",
	},
}

const producerCount int = 4
const consumerCount int = 3

func produce(link chan<- string, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, msg := range messages[id] {
		link <- msg
	}
}

func consume(link <-chan string, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for msg := range link {
		fmt.Printf("Message \"%v\" is consumed by consumer %v\n", msg, id)
	}
}

func main() {
	link := make(chan string)
	wp := &sync.WaitGroup{}
	wc := &sync.WaitGroup{}

	wp.Add(producerCount)
	wc.Add(consumerCount)

	for i := 0; i < producerCount; i++ {
		go produce(link, i, wp)
	}

	for i := 0; i < consumerCount; i++ {
		go consume(link, i, wc)
	}

	wp.Wait()
	close(link)
	wc.Wait()
}
