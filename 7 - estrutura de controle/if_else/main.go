package main

import (
	"fmt"
)

func main() {
	i := 1

	if i < 2 {
		fmt.Println("Número menor que 2")
	} else {
		fmt.Println("Número maior ou igual que 2")

	}

	// if init
	// y a variável y fica apenas disponível para o escopo que foi executada
	if y := 23; y < 0 {
		fmt.Println("Número menor que 0")

	} else {
		fmt.Println("Número maior ou igual que 0")

	}
}
