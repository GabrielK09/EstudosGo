package main

import "fmt"

func main() {
	alunos := make(map[string]int)

	alunos["Pedro"] = 7
	alunos["Bruno"] = 10
	alunos["Julia"] = 4

	for x, y := range alunos {
		fmt.Printf("Aluno: %s - Nota: %d\n", x, y)
	}
}
