package main

import "fmt"

type myMethod struct {
	v float64
}

func (v myMethod) eMair() bool {
	return v.v >= 10

}

func main() {
	var i myMethod
	i.v = 10
	fmt.Println(i.eMair())
}
