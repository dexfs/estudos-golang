package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	fmt.Println("Métodos")
	usuario1 := usuario{"Usuario 1", 20}
	usuario1.salvar()
	fmt.Println(usuario1)

	usuario2 := usuario{"Usuario 2", 17}
	fmt.Println(usuario2.maiorIdade())

	usuario3 := usuario{"Usuario 3", 17}
	usuario3.fazerAniversario()
	fmt.Println(usuario3)
}
