package main

import (
	"fmt"
)

func main() {
	canal := make(chan string, 2) // fefine canal com tamanho limitado, para ser usado na mesma função
	canal <- "Olá Mundo!" // 1
	canal <- "Programando em Go!" // 2
	canal <- "Programando em Go De Novo!" // 3 excede o capacidade do canal e gera deadlook

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
