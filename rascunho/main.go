package main

import (
	"fmt"
	"sync"
	"time"
)

func count(name string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Nome: %s\n", name)
		time.Sleep(500 * time.Millisecond)

	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(5)

	go func() {
		count("Chamada 1")
		wg.Done()
	}()

	go func() {
		count("Chamada 2")
		wg.Done()
	}()

	wg.Wait()
}
