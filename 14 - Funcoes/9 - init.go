package main

import "fmt"

var n int

func main() {
	fmt.Println("Executando a função main")
}

// init pode ter uma por arquivo dentro do mesmo pacote
// main so pode ter uma por pacote
func init()  {
	fmt.Println("Executando a função init")
	n = 10
}
