package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func main() {
	templates = template.Must(template.ParseGlob("*.html"))

	homeHandler := func(w http.ResponseWriter, r *http.Request) {
		u := usuario{"João", "a@a.com"}
		templates.ExecuteTemplate(w, "home.html", u)
	}

	http.HandleFunc("/", homeHandler)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
