package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Oi"))
}

func main(){

	http.HandleFunc("/home", home)

	http.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Carregando a tela de usuários"))
	})

	log.Fatalln(http.ListenAndServe(":5000",nil))
}
