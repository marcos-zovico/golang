package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := multiplex(write4("Olá mundo!"), write4("Programando em Go!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

func multiplex(inChannel1, inChannel2 <-chan string) <-chan string {
	outChannel := make(chan string)

	go func() {
		for {
			select {
			case message := <- inChannel1:
				outChannel <- message
			case message := <- inChannel2:
				outChannel <- message
			}
		}
	}()
	return outChannel
}

func write4(text string) <-chan string {
	channel := make(chan string)
	go func() {
		for {
			channel <- fmt.Sprintf("Valor recebido: %s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return channel
}
