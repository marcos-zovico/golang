package main

import "fmt"

type usuario struct {
	nome  string
	idade uint
}

func (u usuario) salvar() {
	fmt.Printf("Salvando usuário %s \n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	user1 := usuario{"Marcos", 33}
	user1.salvar()

	user2 := usuario{"José", 21}
	user2.salvar()

	maiorIdade := user2.maiorDeIdade()
	fmt.Println(maiorIdade)

	user2.fazerAniversario()
	fmt.Println(user2.idade)
}
