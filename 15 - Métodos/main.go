package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Println("Salvando usuário ....")
}

func main() {
	usuario.salvar(usuario{nome: "Teste", idade: 12})
}
