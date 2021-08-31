package main

import "fmt"

func main() {
	var v1 int = 10
	var v2 int = v1

	fmt.Println(v1, v2)

	v1++
	fmt.Println(v1, v2)

	// PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA
	var v3 int
	var ponteiro *int
	fmt.Println(v3, ponteiro)

	v3 = 100
	ponteiro = &v3 // referencia, atribui um valor ao endereço de memoria(ponteiro)
	fmt.Println(v3, ponteiro)

	v3 = 150
	fmt.Println(v3, ponteiro)
	fmt.Println(v3, *ponteiro) // desrreferrencia do endereço de memoria
}
