package main

import "fmt"

type Animal interface {
	Onomatopeia() string
}

type Urso struct{}

func (u Urso) Onomatopeia() string {
	return "AAAAAA"
}

type Passaro struct{}

func (p Passaro) Onomatopeia() string {
	return "BBBBBBBBBb"
}

func Play(a Animal) {
	fmt.Println(a.Onomatopeia())
}

func main() {
	p := Passaro{}
	fmt.Println(p.Onomatopeia())
}
