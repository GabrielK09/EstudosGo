package main

import "fmt"

func closure() func() {
	text := "AAAAA"

	funcao := func() {
		fmt.Println(text)
	}

	return  funcao
}

func main() {
	text := "BBBB"
	fmt.Println(text)

	newFunc := closure()
	newFunc()

}