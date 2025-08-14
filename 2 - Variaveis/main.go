package main

import "fmt"

func main() {
	var i int32 // Declaração
	i = 2

	y := 2 // inferência

	var t, z int32 // 2 dados int32
	t, z = 2, 2    // valores na respectiva ordem
	t = 4
	z = 5
	// Separados

	fmt.Println(i, y)
	fmt.Println(t, z)
}
