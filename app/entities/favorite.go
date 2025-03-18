package entities

import (
	"github.com/nonamecat19/go-orm/core/lib/entities"
)

type Favorite struct {
	entities.Model
	UserId    int64   `db:"user_id" json:"userId"`
	User      User    `db:"user" json:"user" relation:"foreign-key:id"`
	ProductId int64   `db:"product_id" json:"productId"`
	Product   Product `db:"product" json:"product" relation:"foreign-key:id"`
}

func (user Favorite) Info() string {
	return "favorites"
}
