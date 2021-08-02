package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, num := range numeros {
		total += num
	}

	return total
}

func numString(text string, nummeros ...int) {
	for num := range nummeros {
		fmt.Println(text, num)
	}
}

func main() {
	resultado := soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println(resultado)

	numString("Olá mundo", 1, 2, 3, 4, 5, 6)
}
