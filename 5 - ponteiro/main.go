package main

import "fmt"

func main() {
	x := 10
	y := &x
	*y = 20

	fmt.Println(x)
}
