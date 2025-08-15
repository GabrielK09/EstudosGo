package main

import "fmt"

func tentativaRecover() {
	if r := recover(); r != nil {
		fmt.Println("Recuperção com sucesso!")
	}
}

func calculateAvg(n1, n2 float64) bool {
	defer tentativaRecover()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MÉDIA É EXATAMENTE 6")
	// Vai chamar tudo o que estiver no defer, nesse caso, uma tentativa de recuperar o sistema, se não conseguir, F
}

func main() {
	fmt.Println(calculateAvg(6, 6))
	fmt.Println("Pós")
}
