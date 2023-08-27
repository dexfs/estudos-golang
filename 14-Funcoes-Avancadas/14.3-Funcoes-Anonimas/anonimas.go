package main

import "fmt"

func main() {

	fmt.Println("Funções anônimas")
	returno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Olá mundo!")

	fmt.Println(returno)

}
