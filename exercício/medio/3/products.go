package main

import "fmt"

type product struct {
	Nome  string
	Preco float64
}

type productsMethod interface {
	desconto(perc float64) float64
}

func (p product) desconto(perc float64) float64 {
	return (p.Preco * perc) / 100

}

func main() {
	p := product{Nome: "Teste", Preco: 100.00}

	applyDescont(p)
}

func applyDescont(p product) {
	fmt.Printf("Produto %s :", p.Nome)
	fmt.Printf("Produto com desconto R$ %.2f: ", p.desconto(50))
}
