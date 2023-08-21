package main

import (
	"errors"
	"fmt"
)

func main() {
	// NUMEROS INTEIROS
	var numero int64 = -1000000
	fmt.Println(numero)

	var numero2 uint32 = 100000
	fmt.Println(numero2)

	// alias
	// INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	// BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)
	// FIM NUMEROS INTEIROS

	// NUMERO REAIS
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 12300000.45
	fmt.Println(numeroReal2)

	numerReal3 := 12345.67
	fmt.Println(numerReal3)

	// FIM NUMERO REAIS

	// STRINGS
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto 2"
	fmt.Println(str2)

	// char := 'B' // não existe char em Go
	// tabela ASCII - https://web.fe.up.pt/~ee96100/projecto/Tabela%20ascii.htm
	// resultado do char é o número 66
	// char é do tipo int
	char := 'B'
	fmt.Println(char)

	// FIM STRINGS

	text := 5
	fmt.Println(text)

	// default de booleano é false
	var booleano1 bool = true
	fmt.Println(booleano1)

	// valor é nil de zero e nao de vazio.
	var erro error = errors.New("ERROR INTERNO")
	fmt.Println(erro)
}
