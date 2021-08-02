package main

import "fmt"

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

//recebe dados / envia dados
func worker(tarefas <-chan int, resultados chan<- int) {
	for tarefa := range tarefas {
		resultados <- fibonacci(tarefa)
	}
}

func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for x := 0; x < 45; x++ {
		tarefas <- x
	}
	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}

}
