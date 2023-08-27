package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor que 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Maior que 0")
	} else if numero < -10 {
		fmt.Println("Menor que -10")
	} else {
		fmt.Println("Entre 0 e -10")
	}
}
