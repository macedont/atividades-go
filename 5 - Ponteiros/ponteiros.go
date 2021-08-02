package main

import "fmt"


func main(){
	fmt.Println("Ponteiros")

	var variavel int = 10
	var variavel1 int = variavel

	fmt.Println(variavel, variavel1)

	variavel++
	fmt.Println(variavel, variavel1)

	//PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA
	var variavel2 int
	var ponteiro *int

	variavel2 = 100
	ponteiro = &variavel2
	
	fmt.Println(variavel2, *ponteiro) //desreferenciação 

	variavel2 = 150

	fmt.Println(variavel2, *ponteiro) //desreferenciação 
}