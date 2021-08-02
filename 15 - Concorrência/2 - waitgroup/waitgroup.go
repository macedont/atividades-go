package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		escrever("olá go! tudo bem?")
		waitGroup.Done() // -1 do method add
	}()

	go func() {
		escrever("Programando em go")
		waitGroup.Done() // -1
	}()

	waitGroup.Wait() //esperar a contagem das goroutines chegar em 0
}

func escrever(x interface{}) {
	for y := 0; y < 5; y++ {
		fmt.Println(x)
		time.Sleep(time.Second)
	}
}
