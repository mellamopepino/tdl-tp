package main

import (
	"fmt"
	"sync"
	"time"
)

// START OMIT

func trabajador(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Trabajador %d empezando\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Trabajador %d terminó\n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go trabajador(i, &wg)
	}
	wg.Wait()
}

// END OMIT
