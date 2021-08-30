package main

import (
	"errors"
	"fmt"
)

func main() {

	// NÚMERO INTEIROS
	var numero int64 = -10000000000
	fmt.Println(numero)

	// unsigned int
	// uint não aceita negativo
	var numero2 uint32 = 10000
	fmt.Println(numero2)

	// alias
	// INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	// UINT8 = BYTE
	var numero4 byte = 123
	fmt.Println(numero4)
	// FIM NÚMERO INTEIROS

	// NÚMERO REAIS
	var num float32 = 123.45
	fmt.Println(num)

	var num2 float64 = 12345454545.45
	fmt.Println(num2)

	num3 := 12345.56
	fmt.Println(num3)


	// STRING
	var str1 string = "sdsdsd"
	fmt.Println(str1)

	str2 := "gdfgdgdfg"
	fmt.Println(str2)

	// char é sempre declarado em com aspas simples
	// é sempre convertido para o numero da tabel aASCII
	char := 'A'
	fmt.Println(char)


	// BOOL
	bool1 := true
	fmt.Println(bool1)

	var bool2 bool = false
	fmt.Println(bool2)


	// ERROR
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

	erro2 := errors.New("Novo erro")
	fmt.Println(erro2)

}
