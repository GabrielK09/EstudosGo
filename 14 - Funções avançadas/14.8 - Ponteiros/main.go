package main

import "fmt"

func reverseSignal(n int) int {
	return n * -1

}

func reversePointerNumber(n *int) {
	*n = *n * -1
}

func main() {
	number := 20

	revereseNumber := reverseSignal(number)

	fmt.Println(revereseNumber)
}
