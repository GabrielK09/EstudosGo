package main

import "fmt"

func main() {
	myMap := map[string]string{
		"teste": "Valor",
	}

	for i, v := range myMap {
		fmt.Println(i)
		fmt.Println(v)

	} // Posso aplicar um for ao map, i = chave | v = valor
	// posso fazer for i := range, e printar myMap[i], ou i, _ := range ...

	myMap2 := map[int]map[string]string {
		1: {
			"valor1": "AAAA",
		},
	}

	fmt.Println(myMap2)
}
