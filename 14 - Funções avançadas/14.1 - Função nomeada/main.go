package main

import "fmt"

func calc(n1, n2 int) (soma int, sub int) {
	soma = n1 + n2
	sub = n1 - n2
	return
}

func main() {
	res, _ := calc(1, 2)
	_, res2 := calc(1, 2)
	fmt.Println(res)
	fmt.Println(res2)
}
