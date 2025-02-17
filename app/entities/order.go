package entities

import (
	"github.com/nonamecat19/go-orm/core/lib/entities"
	"time"
)

type Order struct {
	entities.Model
	Count     int       `db:"count" json:"count"`
	UserId    int64     `db:"user_id" json:"userId"`
	User      User      `db:"user" json:"user" relation:"foreign-key:id"`
	OrderDate time.Time `db:"order_date" json:"orderDate"`
}

func (user Order) Info() string {
	return "orders"
}
