package main

import "fmt"


func main() {

	var variavel1 string = "Varialvel 1"
	variavel2 := "Variavel 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string =  "lalala"
		variavel4 string =  "lalala"
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "sdfsdfsdff", "xcxcxcxcxcx"
	fmt.Println(variavel5, variavel6)

	// Vale a mesma regra de declaração de variáveis

	const constante1 string = "Contante 1"
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6,  variavel5

	fmt.Println(variavel5, variavel6)

}