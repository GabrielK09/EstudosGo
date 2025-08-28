package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if a <= 0 || b <= 0 {
		return 0, errors.New("impossivel dividir por zero")
	}

	return a / b, nil
}

func main() {
	res, err := divide(0, 2)
	if err != nil {
		fmt.Println("Erro na divisão:", err)
		return
	}

	fmt.Println("Resultado:", res)
}
