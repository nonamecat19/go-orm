package entities

import "github.com/nonamecat19/go-orm/core/lib/entities"

type Role struct {
	entities.Model
	Name  string `db:"name" json:"name,omitempty"`
	Users []User `db:"users" json:"users,omitempty" relation:"role_id"`
}

func (role Role) Info() string {
	return "roles"
}
