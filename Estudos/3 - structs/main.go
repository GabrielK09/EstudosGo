package main

import "fmt"

type pessoa struct {
	nome       string
	idade      uint8
	planosaude planoSaude
}

type medico struct {
	pessoa
	centroDeSaude string
}

type planoSaude struct {
	nome string
}

func main() {
	var u pessoa
	u.nome = "Teste"
	u.idade = 15

	fmt.Println(u)

	planoDeSaude := planoSaude{"Unimed"}
	u2 := pessoa{nome: "Teste 2", planosaude: planoDeSaude}

	fmt.Println(u2)
}
