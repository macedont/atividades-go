package main

import (
	"fmt"
	"time"
)

func main() {
	canal, canal1 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal <- "Canal"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal1 <- "Canal 1"
		}
	}()

	for {

		select {
		case menssagem := <-canal:
			fmt.Println(menssagem)
		case menssagem1 := <-canal1:
			fmt.Println(menssagem1)
		}
	}
}
