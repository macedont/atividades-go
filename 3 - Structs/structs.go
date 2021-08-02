package main

import(
	"fmt"
)

type usuario struct{
	nome string
	idade int
	endereco endereco 
}

type endereco struct{
	logradouro string
	numero int
}

func main(){
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Macedo chan"
	u.idade = 21
	fmt.Println(u);

	enderecoEx := endereco{"Rua dos bobos", 0}

	u2 := usuario {"Macedu", 20, enderecoEx} 
	fmt.Println(u2)

	u3 := usuario {nome: "oioi"}

	fmt.Println(u3)
}