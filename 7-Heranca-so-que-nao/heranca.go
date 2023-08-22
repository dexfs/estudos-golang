package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint
	altura    uint
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança em Go, #sqn")
	p1 := pessoa{"João", "Pedro", 20, 178}
	fmt.Println(p1)
	e1 := estudante{p1, "Engenharia", "USP"}
	fmt.Println(e1.nome, e1.curso, e1.faculdade)
}
