package main

import "fmt"

func main() {

	numero := 15

	if numero >= 10 {
		fmt.Println("Aprovado")
	} else {
		fmt.Println("Reprovado")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("É um número positivo")
	} else {
		fmt.Println("É um número negativo")
	}
}
