package main

import (
	"fmt"
)

func main() {
	user := map[string]string{
		"name":    "Gabriel",
		"surname": "Kochem",
	}

	fmt.Println(user)

	// Vai ser um map dentro de um map
	// Definindo o primeiro map que tem a sua chave de string: map[string] = "nome"
	// E o seu valor, que é outro map de chaves e valores de strings, "name": e "surname":

	disgraca := map[string]map[string]string{
		"nome": {
			"name":    "Gabriel",
			"surname": "Kochem",
		},

		"address": {
			"rua":    "Alexandre Lorezent",
			"bairro": "Parque",
		},
	}

	fmt.Println(disgraca)
	// Remover uma chave
	delete(disgraca, "nome")

	fmt.Println(disgraca) // Apenas o address

	// Adicionar é um nova chave, desde que atenda aos critérios do map, chave string, valor string, e valor map[string]

	disgraca["outra"] = map[string]string{
		"outra": "a",
	}

	fmt.Println(disgraca)
}
