package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		write("Linha 14")
		waitGroup.Done()
	}()

	go func() {
		write("Linha 19")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func write(text string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Texto: %s\n", text)
		time.Sleep(time.Second)
	}
}
