package main

import "fmt"

//buffer especifica a capacidade do canal e ele só bloqueia quando atinge a capacidade máxima no segundo param
func main() {
	canal := make(chan string, 2)

	canal <- "olá canal"
	canal <- "olá canal"

	menssagem := <-canal
	menssagem1 := <-canal

	fmt.Println(menssagem)
	fmt.Println(menssagem1)

}
