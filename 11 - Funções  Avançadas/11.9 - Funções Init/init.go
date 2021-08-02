package main

import "fmt"

var n int

func main() {
	fmt.Println("Func main sendo exec")
	fmt.Println(n)
}

/*
	a função é a primeira a ser executada antes da main geralmente é usada para executar configurações
*/

func init() {
	fmt.Println("Executando a função init")
	n = 10
}
