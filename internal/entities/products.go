package entities

import (
	"fmt"
	"time"
)

type Products struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Qtde      float64   `json:"qtde"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewProduct(name string, qtde float64) (*Products, error) {
	product := &Products{
		Name: name,
		Qtde: qtde,
	}

	err := product.isValid()
	if err != nil {
		return nil, err

	}

	return product, nil
}

func (c *Products) isValid() error {
	if len(c.Name) < 3 {
		return fmt.Errorf("Nome precisa conter pelo menos 3 caracteres: %d", len(c.Name))
	}

	if c.Qtde <= 0 {
		return fmt.Errorf("A quantidade precisa ser maior que zero: %d", len(c.Name))
	}

	return nil
}
