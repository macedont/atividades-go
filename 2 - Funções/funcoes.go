package main

import(
	"fmt"
)

func somar(n1 int, n2 int) int { 
	return n1 + n2
}

func calculos(n1, n2 int) (int, int){
	somar := n1 + n2
	sub := n1 - n2
	return somar, sub
}


func main(){
	resultado := somar(5, 5)
	fmt.Println("O resultado da soma é:", resultado)

	var f = func (txt string){
		fmt.Println(txt)
	}

	f("oi eu sou o goku")

	resultadoSoma, _ := calculos(10, 15) //ignorando o segundo retorno
	fmt.Println(resultadoSoma)
}