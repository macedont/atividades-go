package main

import "fmt"

func main() {
	numero := 30

	switch numero {
	case 1:
		fmt.Println("caiu no 1")
		break
	case 10:
		fmt.Println("caiu no 10")
		break
	default:
		fmt.Println("caiu no default")
	}
}
