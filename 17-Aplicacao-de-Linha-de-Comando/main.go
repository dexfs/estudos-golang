package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de partida")
	aplication := app.Gerar()
	erro := aplication.Run(os.Args)
	if erro != nil {
		log.Fatal(erro)
	}
}
