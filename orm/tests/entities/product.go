package entities

import "github.com/nonamecat19/go-orm/core/lib/entities"

type Product struct {
	entities.Model
	Name  string  `db:"name" json:"name"`
	Price float64 `db:"price" json:"price"`
}

func (user Product) Info() string {
	return "products"
}
