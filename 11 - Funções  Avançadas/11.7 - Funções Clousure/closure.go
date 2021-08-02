package main

import "fmt"

/*
	são funções que referenciam variáveis de fora do scope
*/

func closure() func() {
	txt := "Dentro de uma clousure"

	funcao := func() {
		fmt.Println(txt)
	}

	return funcao
}

func main() {
	text := "dentro da func main"
	fmt.Println(text)

	funcNova := closure()
	funcNova()
}
