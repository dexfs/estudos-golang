package main

import "fmt"

func main() {
	canal := make(chan string, 2)
	canal <- "Olá Mundo" // deadlock
	canal <- "Programando em Go!"

	mensagem := <-canal
	fmt.Println(mensagem)
}
