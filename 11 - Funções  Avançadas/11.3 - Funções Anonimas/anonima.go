package main

import "fmt"

func main() {

	text := "declarando o texto"

	ano := func(txt string) string {
		return fmt.Sprintf("Recebeu -> %s", txt)
	}(text)

	fmt.Println(ano)

}
