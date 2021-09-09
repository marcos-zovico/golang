package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)

	go func() {
		write("GoRoutine 1")
		waitGroup.Done() // -1
	}()

	go func() {
		write("GoRoutine 2")
		waitGroup.Done() // -1
	}()

	waitGroup.Wait()
}

func write(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
