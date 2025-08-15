package main

import "fmt"

func main() {
	res := func(text string) string {
		return fmt.Sprintf("Texto: %s", text)
	}("Teste")

	fmt.Println(res)

}
