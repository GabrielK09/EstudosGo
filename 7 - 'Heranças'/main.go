package main

type pessoa struct {
	Nome string
	Pet  pet
}

type pet struct {
	Dono  string
	Nome  string
	Idade int
}

// Isso é o máximo de herança que temos
