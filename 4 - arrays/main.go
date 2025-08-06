package main

import (
	"fmt"
)

func main() {
	n := map[string]int{"t": 1}

	for key, value := range n {
		fmt.Println(key)
		fmt.Println(value)
	}
}
