package main

import "fmt"

func recuperarExec() {
	if r := recover(); r != nil {
		fmt.Println("A execução foi recuperada com sucesso") //recupera a aplicação caso de um panic
	}
}

func mediaArit(n1, n2 float64) bool {
	defer recuperarExec()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A média é exatamente 6") //mata a aplicação caso chegue aq e chama todas as funções defer dentro da função
}

func main() {
	fmt.Println(mediaArit(8, 7))
	fmt.Println("Pós execução")
}
