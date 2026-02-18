package models

import (
	"fmt"
	"time"
)

type Product struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewProductModel(name string, price float64) (*Product, error) {
	product := &Product{
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	isValid, err := product.isValid()
	if !isValid {
		return nil, err
	}

	return product, nil
}

func (p *Product) isValid() (bool, error) {
	if len(p.Name) < 1 || len(p.Name) > 100 {
		return false, fmt.Errorf("invalid name: must be between 1 and 100 characters")
	}

	if p.Price < 0 {
		return false, fmt.Errorf("invalid price: must be greater than or equal to 0")
	}

	return true, nil
}
