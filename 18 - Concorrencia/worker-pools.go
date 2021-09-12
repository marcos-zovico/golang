package main

import (
	"fmt"
	"time"
)

const (
	workerPool  = 10
	fibPosition = 50
)

func main() {
	initTime := time.Now()
	tasks := make(chan int, fibPosition)
	results := make(chan int, fibPosition)

	for i := 0; i < workerPool; i++ {
		go worker(tasks, results)
	}

	for i := 0; i < fibPosition; i++ {
		tasks <- i
	}
	close(tasks)

	for i := 0; i < fibPosition; i++ {
		result := <-results
		fmt.Println(result)
	}
	endTime := time.Now()

	fmt.Println("===============")
	fmt.Println(endTime.Second() - initTime.Second())
	fmt.Println("===============")
}

func worker(tasks <-chan int, results chan<- int) {
	for position := range tasks {
		results <- fibonacci(position)
	}
}

func fibonacci(position int) int {
	if position <= 1 {
		return position
	}
	return fibonacci(position-2) + fibonacci(position-1)
}
