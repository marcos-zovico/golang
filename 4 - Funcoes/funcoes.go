package main

import "fmt"

func somar(n1 int8, n2 int8) int8  {
	return n1 + n2
}

func calculos(n1, n2 int8) (int8, int8)  {
	soma := n1 + n2
	sub := n1 - n2
	return soma, sub
}

func main() {
	soma := somar(10, 10)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("Texto da função")
	fmt.Println(resultado)

	sum, sub := calculos(10, 15)
	fmt.Println(sum, sub)

	_, sub2 := calculos(10, 15)
	fmt.Println(sub2)
}
