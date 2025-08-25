package main

import (
	"fmt"
	"time"
)

func main() {
	// Resumindo isto vai rodar a próxima linha mesmo que aonde foi definido go, não tenha 'terminado', onde o write é um texto infinito, mas assim, vai ser feito os 2 da mesma forma, e assim vai levar em conta também ao número de nú
	go write("Linha 9") // Vai fazer uma linha, faz a outra, faz a outra,
	write("Linha 10")

}

func write(text string) {
	for {
		fmt.Printf("Texto: %s\n", text)
		time.Sleep(time.Second)
	}
}
