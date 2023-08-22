package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Structs")
	var u usuario // u é do tipo usuario
	u.nome = "João"
	u.idade = 30
	fmt.Println(u)

	endereco1 := endereco{"Rua dos Bobos", 0}

	u2 := usuario{"Maria", 40, endereco1}
	fmt.Println(u2)

	u3 := usuario{nome: "José"}
	fmt.Println(u3)

	u4 := usuario{idade: 60}
	fmt.Println(u4)
}
