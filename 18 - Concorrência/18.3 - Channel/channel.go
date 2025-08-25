package main

import (
	"fmt"
	"time"
)

func main() {
	// O canal vai trasportar dados, assim que receber o seu dado, vai parar a execução, então assim que o dado for retornado, linha 22, vai acabar ali

	channel := make(chan string)

	go write("Pegerfgref", channel) // Esse cara vai executar e ir embora, e vai deixar ali, rodando, o canal vai ter o seu inicio e o seu fim, aqui sendo o inicio

	// Aqui vai ser printado apenas um dos textos
	// Se for feito um for de forma inifinita, vai retornar deadlock, que é quando o canal está aberto esperando uma resposta, porém nunca vai ser retornado nada, pois vai ficar preso dentro do próprio for
	/*
	for {
		msg := <-channel 
		fmt.Println(msg)

	} == deadlockl*/

	msg := <-channel // Aqui é o fim
	fmt.Println(msg)

}

func write(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- text // Aqui é o 
		time.Sleep(time.Second)
	}
}
