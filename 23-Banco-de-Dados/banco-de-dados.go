package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		log.Fatal(erro.Error())
	}
	defer db.Close()
	if erro = db.Ping(); erro != nil {
		log.Fatal(erro.Error())
	}

	fmt.Println("Conexão está aberta!")
	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		log.Fatal(erro)
	}

	defer linhas.Close()
	//for linhas.Next() {
	//	var usuario Usuario
	//	erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email, &usuario.Senha)
	//	if erro != nil {
	//		log.Fatal(erro)
	//	}
	//	fmt.Println(usuario)
	//}
}
