package main

import (
	"fmt"
)

func main() {
	i := 0

	//for i < 10 {
	//	i++
	//	fmt.Println("Valor i é ", i)
	//	time.Sleep(time.Second)
	//}

	fmt.Println(i)

	//for j := 0; j < 10; j++ {
	//	fmt.Println("Valor j é ", j)
	//	time.Sleep(time.Second)
	//}

	nomes := [3]string{"João", "Davi", "Lucas"}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for _, letra := range "ANDRE" {
		fmt.Println(letra, string(letra))
	}

	usuario := map[string]string{
		"nome":      "André",
		"sobrenome": "Santos",
	}

	for key, value := range usuario {
		fmt.Println(key, value)
	}

	for {
		fmt.Println("Loop infinito")
	}
}
