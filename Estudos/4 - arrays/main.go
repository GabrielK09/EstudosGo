package main

import (
	"fmt"
)

// O array vai ser definido com o seu tamanho, [5] e o tipo de dado: string
// Definido seu tamanho, não consigo adicionar maios valores, para adicionar uma valor em posição X, pode-se ser feito assim:
// array1[5] = "adsdasd", porém aqui vai ser um erro, pois a posição passa o seu máximo

// Ideal é de se utilizar o slice e não fica preso a essa limitação de tamanho
// slice := []string, a lógica dos dados é a mesma, apenas strings para esse slice
// Para adicionar um novo elemento, podese utilizar:
// slice = append(slice, "valor")

func main() {
	array1 := [5]string{"A", "B", "C", "D"}
	fmt.Println(array1) // "", "", "", "", ""

	slice := []string{}
	slice = append(slice, "A")

	fmt.Println(slice)

	// Pegando uma parte de um array

	slice2 := array1[1:3]
	// Se algo for alterado no array original, o slice também vai ser alterado por conta do apontamento do array interno que 'gerencia' o slice
	// 1 = Inclusivo, vai ser de fato o item da posição 1
	// 3 = Exclusivo, não vai ser pego o item da posição 3, e sim o anterior do 3, sendo a posição 2

	fmt.Println(slice2)
}
