package main

import "fmt"

//"time"

func main() {
	/* x := 0

	for x <= 10 {
		time.Sleep(time.Second)
		fmt.Println(x)
		x++
	}

	for y := 0; y < 10; y++ {
		time.Sleep(time.Second)
		fmt.Println(y)
	} */

	nomes := [3]string{"Macedo", "Juaum", "Pedru"}

	for _, nome := range nomes { //retorna indice e valor
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra)) //sem a conversão de string() a var letra retorna os número dos caracteres de acordo com a tabela ascii
	}

	var nomesw = map[string]string{
		"primeiro nome": "Macedo",
		"segundo nome":  "Chan",
	}

	for x, nome := range nomesw {
		fmt.Println(x, nome)
	}

	/* for { loop infinito
		fmt.Println(x)
		x++
	} */

	//não da pra usar o range em um struct
}
