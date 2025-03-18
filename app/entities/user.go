package entities

import "github.com/nonamecat19/go-orm/core/lib/entities"

type User struct {
	entities.Model
	Name      string     `db:"name" json:"name"`
	Email     string     `db:"email" json:"email"`
	Gender    string     `db:"gender" json:"gender"`
	Orders    []Order    `db:"orders" json:"orders"`
	Favorites []Favorite `db:"favorites" json:"favorites"`
}

func (user User) Info() string {
	return "users"
}
