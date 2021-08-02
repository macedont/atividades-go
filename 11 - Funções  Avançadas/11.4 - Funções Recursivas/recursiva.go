package main

import "fmt"

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {

	posicao := 12

	for x := 1; x < posicao; x++ {
		fmt.Println(fibonacci(x))
	}

}
