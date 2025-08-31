package main

import (
	"fmt"
	"time"
)

func main() {
	ch := write("meu generator")
	
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}

func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Valor recebido: %s", text)
			time.Sleep(time.Microsecond * 500)
		}
	}()

	return channel
}