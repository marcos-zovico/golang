package main

import (
	"fmt"
	"time"
)

func main() {
	channel := write3("Olá mundo!")

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

func write3(text string) <-chan string {
	channel := make(chan string)
	go func() {
		for {
			channel <- fmt.Sprintf("Valor recebido: %s", text)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return channel
}
