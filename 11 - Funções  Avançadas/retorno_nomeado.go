package main

import "fmt"

func calculos(n1 int, n2 int) (soma int, subtracao int, divisao float64) {
	soma = n1 + n2
	subtracao = n1 - n2
	divisao = float64(n1 / n2)

	return
}

func main() {

	soma, subtracao, divisao := calculos(15, 6)

	fmt.Println("Soma", soma, "\nSubtração:", subtracao, "\nDivisão:", divisao)

}
