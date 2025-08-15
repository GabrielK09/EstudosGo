package main

import "fmt"

func funcao1() {
	fmt.Println("Executando funcao1")
}

func funcao2() {
	fmt.Println("Executando funcao2")
}

func studant(n ...int) bool {
	total := 0

	for _, v := range n {
		total += v

	}

	if avg := total / len(n); avg >= 6 {
		fmt.Printf("média: %d\n", avg)
		return true
	} else {
		fmt.Printf("média: %d\n", avg)
		return false
	}
}

func main() {
	defer funcao1()
	funcao2()

	// Defer vai adiar o que foi adiado, defer = adiar
	// Então o que vai ser escrito, vai ser
	// Executando funcao2
	// Executando funcao1
	// Mesmo chamando a função 1 primeiro

	fmt.Println(studant(6, 6, 6))
}
