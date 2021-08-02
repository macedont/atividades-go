package main

import (
	"fmt"
	"time"
)

func main() {
	//CONCORRÊNCIA != PARALELISMO

	go escrever("olá go! tudo bem?") //goroutine
	escrever("Programando em go")
}

func escrever(x interface{}) {
	for {
		fmt.Println(x)
		time.Sleep(time.Second)
	}
}
