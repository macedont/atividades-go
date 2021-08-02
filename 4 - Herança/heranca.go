package main

import "fmt"

type pessoa struct{
	nome string
	sobrenome string
	idade int
	altura int
}

type estudante struct{
	pessoa
	curso string
	faculdade string
}


func main(){
	fmt.Println("Herança")

	p1 := pessoa{"Jaum", "Pedro", 22, 170}
	fmt.Println(p1)

	e1 := estudante{p1, "ciência da computação", "estácio de sá"}

	fmt.Println(e1) //ver campos especificos e1.nome
}