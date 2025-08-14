package main

import "fmt"

func main() {
	var i int16 = 3
	switch i {
	case 1:
		fmt.Println("É 1")

	case 2:
		fmt.Println("É 2")

	case 3:
		fmt.Println("É 3")

	default:
		fmt.Println("O valor não é um número")
	}
}
