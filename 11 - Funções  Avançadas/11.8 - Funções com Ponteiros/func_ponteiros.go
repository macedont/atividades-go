package main

import "fmt"

func invertSinal(num int) int {
	return num * -1
}

func inverterSinalPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20
	inverterSinalPonteiro(&numero)
	fmt.Println(numero)
}
