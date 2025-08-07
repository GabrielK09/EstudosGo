package main

import (
	"fmt"
)

// switch vai possuir 2 meios para validação
/*
	Subondo que switch numero tenha o case 1, 2, 3
	Se o numero = 2, vai ser a condição 2, sendo retornado a Segunda
	Formas de otimizar o return é inicializar a váriavel:
	var day string

	switch numero {
	case 1:
		//return "Domingo"
		day = "Domingo"

	case 2:
		//return "Segunda"
		day = "Segunda"

	case 3:
		//return "Terça"
		day = "Terça"

	default:
		//return "Inválido"
		day = "Inválido"
	}

	E retornar o dia diretamente
	return day
	E o outro meio para validação do case, é o case numero == 1, 2, 3, vai ter o mesmo resultado

	O fallthrough vai simplesmente fazer com que o resultado seja "passado", de forma que,
	Dado o numero = 1, o mesmo vai ser retornado Domingo,
	var day string
	case 1:
		day = "Domingo"
		fallthrough // Com o fallthrough aqui, vai passar para o case de baixo independente do que seja

	case 2:
		day = "Segunda"

	Retornando "Segunda"
*/

func weekDay(numero int) string {
	switch numero {
	case 1:
		return "Domingo"

	case 2:
		return "Segunda"

	default:
		return "Dia inválido"
	}
}

func secondWeekDay(numero int) string {
	var day string
	switch {
	case numero == 1:
		day = "Domingo"
		fallthrough

	default:
		day = "Dia inválido"
	}

	return day
}

func main() {
	day := weekDay(2)
	secondDay := secondWeekDay(1)

	fmt.Println(day)
	fmt.Println(secondDay)
}
