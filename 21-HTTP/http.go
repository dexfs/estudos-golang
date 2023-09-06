package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// HTTP É UM PROTOCOLO DE COMUNICAÇÃO - BASE DA COMUNICAÇÃO WEB
	// CLIENT (FAZ REQUISIÇÃO) - SERVIDOR (PROCESSA REQUISIÇÃO E ENVIA RESPOSTA)

	// REQUEST (PEDIDO) - RESPONSE (RESPOSTA)
	// ROTAS
	// UIR - Identificador Universal de Recursos
	// Metodos mais comuns de requisição HTTP: GET, POST, PUT, DELETE, PATCH
	homeHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("Bem vindo ao site %s-%s", r.Method, r.URL.Path)))
	}

	usuariosHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("Bem vindo ao site %s%s%s", "http", r.Method, r.Host)))
	}
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/usuarios", usuariosHandler)
	log.Fatal(http.ListenAndServe(":5000", nil))

}
