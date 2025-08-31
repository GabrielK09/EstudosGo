package main

import "fmt"

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao

	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	// fila de tarefas
	tarefas := make(chan int, 45) // Canal com buffer, 45 posições

	// Esse recebe o resultado da tarefa
	resultados := make(chan int, 45)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		tarefas <- i

	}
	close(tarefas)

	for i := 0; i < 45; i++ {
		res := <-resultados
		fmt.Println("Res:", res)
	}
}

// <-chan - Apenas recebe dados
// chan<- Penas envia dados
func worker(tarefas <-chan int, resultado chan<- int) {
	for tarefa := range tarefas {
		resultado <- fibonacci(tarefa)
	}

}