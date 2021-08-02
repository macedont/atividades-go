package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct{
	Nome string
	Email string
}

func index(w http.ResponseWriter, r *http.Request){
	u := usuario{"Macedo chan", "macedochan@gmail.com"}
	templates.ExecuteTemplate(w, "index.html", u)
}

func main(){

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/index", index)

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
