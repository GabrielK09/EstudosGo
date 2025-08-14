package main

import "fmt"

func main() {
	if i := 2; i < 2 {
		fmt.Printf("I é menor que 2: %d", i)
	} else {
		fmt.Printf("I é maior que 2: %d\n", i)
	} // Esse i não vai mais existir, apenas foi criado e usado aq e acabou

	var y int32

	if y = 3; y < 3 {
		fmt.Printf("Y é menor que 2: %d", y)
	} else {
		fmt.Printf("Y é maior que 2: %d", y)
	}
}