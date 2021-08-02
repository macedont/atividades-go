package main

import "fmt"

func generica(f interface{}) {
	fmt.Println(f)
}

func main() {
	generica(false)

	mapa := map[interface{}]interface{}{
		0:      "oi",
		"nome": 1,
		true:   "false",
	}

	fmt.Println(mapa)
}
