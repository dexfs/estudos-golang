package main

import "fmt"

func main() {
	// ARITMÉTICOS
	// + - * / %

	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int32 = 25
	fmt.Println(numero1)
	fmt.Println(numero2)
	// não pode somar int16 com int32
	//soma2 := numero1 + numero2
	//fmt.Println(soma2)

	// FIM DOS ARITMÉTICOS

	// ATRIBUIÇÃO
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)
	// FIM DOS ATRIBUIÇÃO

	// OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)
	// FIM DOS OPERADORES RELACIONAIS
	// OPERADORES LÓGICOS
	fmt.Println("OPERADORES LÓGICOS")
	fmt.Println(1 == 1 && 2 == 2)
	fmt.Println(1 == 1 || 2 == 2)
	fmt.Println(1 == 1 && 2 != 2)
	fmt.Println(!true)
	fmt.Println("FIM DOS OPERADORES LÓGICOS")
	// FIM DOS OPERADORES LÓGICOS

	// OPERADORES UNÁRIOS
	fmt.Println("OPERADORES UNÁRIOS")
	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)
	fmt.Println("FIM DOS OPERADORES UNÁRIOS")
	// FIM DOS OPERADORES UNÁRIOS
}
