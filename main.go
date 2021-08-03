package main

import (
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("oi"))
}

func main() {

	http.HandleFunc("/index", index)

	log.Fatalln(http.ListenAndServe(":5000", nil))
}
