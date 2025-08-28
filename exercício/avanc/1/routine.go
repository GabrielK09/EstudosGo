package main

import (
	"fmt"
	"sync"
	"time"
)

func count(name string) {
	for i := 0; i < 5; i++ {
		fmt.Println(name)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(5)

	go func() {
		count("Gabriel")
		wg.Done()

	}()

	go func() {
		count("adsadasdasds")
		wg.Done()

	}()

	wg.Wait()
	fmt.Println("Fim ...")
}
