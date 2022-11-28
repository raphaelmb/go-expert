package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data <-chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func main() {
	data := make(chan int)
	workerAmount := 10

	// initialize workers
	for i := 0; i < workerAmount; i++ {
		go worker(i, data)
	}
	// go worker(1, data)
	// go worker(2, data)
	for i := 0; i < 100; i++ {
		data <- i
	}
}
