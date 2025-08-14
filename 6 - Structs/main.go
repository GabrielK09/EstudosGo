package main

import "fmt"

type pessoaS struct {
	Nome      string
	SobreNome string
}

func main() {
	var pessoa pessoaS

	pessoa.Nome = "Teste"
	pessoa.SobreNome = "Da Silva"
	fmt.Println(pessoa)
	// { Teste Da Silva}
	// { Nome SobreNome}
}
