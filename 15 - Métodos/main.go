package main

import "fmt"

<<<<<<< HEAD
type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Println("Salvando usuário ....")
}

func main() {
	usuario.salvar(usuario{nome: "Teste", idade: 12})
=======
type myMethod struct {
	v float64
}

func (v myMethod) eMair() bool {
	return v.v >= 10

}

func main() {
	var i myMethod
	i.v = 10
	fmt.Println(i.eMair())
>>>>>>> 686038d95814ea60fe1645f6ced9004428a9a07e
}
