package main

import "fmt"

func funcao() {
	fmt.Println("Executando func1")
}

func funcao1() {
	fmt.Println("Executando func2")
}

func aluno(n1, n2 float64) (bool, float64) {
	defer fmt.Println("A média foi calculada! O resultado será retornado.") //função defer faz com oq a linha que esteja com essa função seja a ultima operação a ser feita / caso tenha return na func ela vai ser executada antes de qualquer return

	media := (n1 + n2) / 2

	if media >= 6 {
		return true, media
	}

	return false, media
}

func main() {
	/* defer funcao()
	funcao1() */

	fmt.Println(aluno(6, 7))
}
