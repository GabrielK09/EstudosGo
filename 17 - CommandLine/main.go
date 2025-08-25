package main

import (
	"Mcli/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Inicio")

	aplication := app.Generate()
	if err := aplication.Run(os.Args); err != nil {
		log.Fatal("Error in run app:", err)
	}
}