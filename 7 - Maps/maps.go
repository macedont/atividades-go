package main

import (
	"fmt"
)

func main() {
	//fmt.Println("Maps")

	usuario := map[string]string{ //tipo das chaves -> tipo dos valores
		"nome":      "Macedo chan",
		"sobrenome": "kun",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Macedo",
			"ultimo":   "chan",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "são joão",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "curso") //deletando o map

	fmt.Println(usuario2)
}
