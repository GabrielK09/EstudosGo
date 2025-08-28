package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go count(ch)

	for {
		msg, isOpen := <-ch
		if !isOpen {
			break
		}
		fmt.Println(msg)
	}

	fmt.Println("Fim ...")
}

func count(ch chan int) {
	for i := 0; i <= 5; i++ {
		ch <- i
		time.Sleep(time.Second)
	}

	close(ch)
}
