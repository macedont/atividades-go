package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("ola mundo", canal)

	for menssagem := range canal {
		fmt.Println(menssagem)
	}
	fmt.Println("fim do programa")
}

func escrever(x string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- x
		time.Sleep(time.Second)
	}
	close(canal) //sem isso o programa cai em deadlock e isso n é identificado na hora da compilação
}
