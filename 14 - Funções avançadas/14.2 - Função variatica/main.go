package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, v := range numeros {
		total += v
	}
	return total
}

func main() {
	fmt.Println(soma(1 ,2 , 3, 4, 4,5,6, 14))
}