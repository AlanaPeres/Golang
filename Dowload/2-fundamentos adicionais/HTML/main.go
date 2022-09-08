package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

func main() {
	templates = template.Must(template.ParseGlob("*.html")) // aqui eu estou pedindo que todos os meus documentos que tiverem .html recebam esses templates
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Olá"))
		templates.ExecuteTemplate(w, "home.html", nil)
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}