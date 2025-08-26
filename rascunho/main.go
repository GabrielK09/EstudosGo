package main

import "fmt"

func main() {
	ch := make(chan string)
	go myChannel(ch)

	msg := <-ch
	fmt.Println(msg)
}

func myChannel(channel chan string) {
	channel <- "Meu canal"
}
