package main

import "fmt"

func calculos(n1, n2 int) (soma int, sub int) {
	soma = n1 + n2
	sub = n1 - n2
	return
}

func main() {
	soma, sub := calculos(10, 20)
	fmt.Println(soma, sub)
}
