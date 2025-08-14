package main

import "fmt"

func main() {
	var p *int

	i := 10

	p = &i

	fmt.Println(*p, i)
	i++
	fmt.Println(*p, i)
}
