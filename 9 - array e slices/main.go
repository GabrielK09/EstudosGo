package main

import "fmt"

func main() {
	// Preenchendo por posição
	var array [5]string

	array[1] = "position 1"

	fmt.Println(array)

	// --

	// Já definindo
	// Sem o var
	array2 := [5]string{"valor 1", "valor 2", "valor 3", "valor 4", "valor 5"}

	fmt.Println(array2)

	// Slices

	slice := []int{1, 2, 3, 4}

	fmt.Println(slice)

	slice2 := make([]int, 32)
	fmt.Println(slice2)
}
