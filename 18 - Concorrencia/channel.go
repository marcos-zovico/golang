package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	go write2("Hello World", channel)

	fmt.Println("Depois da funcão write começara aser executada")

	// modo roots de receber valo do canal
	//for {
	//	message, channelOpen := <- channel // recebe um valor do canal
	//	if !channelOpen {
	//		break
	//	}
	//	fmt.Println(message)
	//}

	// maneira correta
	for message := range channel {
		fmt.Println(message)
	}

	fmt.Println("Fim do programa")
}

func write2(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- text // envia um valor para um canal
		time.Sleep(time.Second)
	}

	close(channel)
}
