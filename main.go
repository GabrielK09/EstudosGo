package main

import (
	"fmt"
	"log"
	db "g2l.com/crud/config"
)

// Praticamente o programa todo
func main() {
	conn, err := db.SetupDB()

	if err != nil {
		log.Fatal("Erro ao carregar o banco de dados!")
	}

}
