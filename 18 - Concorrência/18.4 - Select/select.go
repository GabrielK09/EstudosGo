package main

import (
	"fmt"
	"time"
)

func main() {
	channel1, channel2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			channel1 <- "Mensagem enviada para o canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			channel2 <- "Mensagem enviada para o canal 2"

		}
	}()

	/*for {
		msg1 := <-channel1
		fmt.Println("Mensagem canal 1:", msg1)

		msg2 := <-channel2
		fmt.Println("Mensagem canal 2:", msg2)
	}*/

	// Desta forma pegamos o canal que já está pronto para envio, sem depender do canal 2 finalizar para interrar no array
	
	for {
		select {
		case msg1 := <-channel1:
			fmt.Println("Mensagem canal 1:", msg1)
		case msg2 := <-channel2:
			fmt.Println("Mensagem canal 2:", msg2)
		}
	}
}
