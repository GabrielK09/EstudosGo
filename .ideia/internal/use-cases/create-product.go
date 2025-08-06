package use_cases

import (
	"log"

	"g2l.com/internal/entities"
)

type createProductUseCase struct {
}

func NewProductUseCase() *createProductUseCase {
	return &createProductUseCase{}

}

func (p *createProductUseCase) Execute(name string, qtde float64) error {
	product, err := entities.NewProduct(name, qtde)

	if err != nil {
		return err

	}

	// TODO: a
	log.Println(product)
	return nil
}
