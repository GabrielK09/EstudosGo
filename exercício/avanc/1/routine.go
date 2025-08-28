package main

import (
	"fmt"
	"sync"
	"time"
)

func count(name string, wg *sync.WaitGroup) {
	fmt.Println(name)
	wg.Done()
	time.Sleep(500 * time.Millisecond)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(5)

	go func() {
		count("Gabriel", nil)
	}()

	wg.Wait()
}
