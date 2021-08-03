package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando o %s no 20 - Banco de dados de dados\n", u.nome)
}

func (u usuario) maiorIdade() bool {
	return u.idade >= 18
}

func (u *usuario) aniversario() {
	*&u.idade++
}

func main() {
	u := usuario{"Macedo chan", 20}
	u.salvar()
	retorno := u.maiorIdade()
	fmt.Println(retorno)
	u.aniversario()
	fmt.Println(u)
}
