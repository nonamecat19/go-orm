package entities

import "github.com/nonamecat19/go-orm/core/lib/entities"

type User struct {
	entities.Model
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
}

func (user User) TableName() string {
	return "users"
}
