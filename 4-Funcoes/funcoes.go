package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	sub := n1 - n2
	return soma, sub
}

func main() {
	soma := somar(10, 20)
	println(soma)

	var f = func() {
		fmt.Println("Função f")
	}

	f()

	resultadoSoma, resultadoSub := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSub)
}
