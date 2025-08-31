package main

import (
	"fmt"
	"time"
)

func main() {
	mainChannel := multiplexar(write("vai ser no canala 1"), write("Vai ser no canal 2"))

	for i := 0; i < 20; i++ {
		fmt.Println(<-mainChannel)
	}
}

func multiplexar(canal1, canal2 <-chan string) <-chan string {
	canalSaida := make(chan string)

	go func() {
		for {
			select {
			case msg1 := <-canal1:
				canalSaida <- msg1

			case msg2 := <-canal2:
				canalSaida <- msg2

			}
		}
	}()

	return canalSaida
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
