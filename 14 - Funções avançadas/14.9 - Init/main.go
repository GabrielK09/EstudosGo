package main

import "fmt"

// Init é rodado antes da main
func init() {
	fmt.Println("Init")
}

func main() {
	fmt.Println("Main")
}
