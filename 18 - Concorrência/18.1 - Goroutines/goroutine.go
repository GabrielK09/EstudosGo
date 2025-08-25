package main

import (
	"fmt"
	"time"
)

func main() {
	write("Linha 9")
	write("Linha 10")

}

func write(text string) {
	for {
		fmt.Printf("Texto: %s\n", text)
		time.Sleep(time.Second)
	}
}
