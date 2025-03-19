package entities

import (
	"github.com/nonamecat19/go-orm/core/lib/entities"
	"time"
)

type Order struct {
	entities.Model
	Count     int       `db:"count" json:"count,omitempty"`
	UserId    int64     `db:"user_id" json:"userId,omitempty"`
	User      *User     `db:"user" json:"user,omitempty" relation:"foreign-key:user_id"`
	OrderDate time.Time `db:"order_date" json:"orderDate,omitempty"`
}

func (user Order) Info() string {
	return "orders"
}
