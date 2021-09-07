package main

import "fmt"

func main() {

	func(texto string) {
		fmt.Println(texto)
	}("Parametro do texto")

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Parametro do texto")

	fmt.Println(retorno)
}
