package main

import "fmt"

func somar(a, b float64) (bool, error) {
	soma := a + b
	fmt.Println(soma)
	return true, nil

}

func main() {
	soma, err := somar(1, 2)

	fmt.Println(soma, err)

}
