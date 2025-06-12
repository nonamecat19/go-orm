package entities

import "github.com/nonamecat19/go-orm/core/lib/entities"

type User struct {
	entities.Model
	Name   string  `db:"name" json:"name,omitempty"`
	Email  string  `db:"email" json:"email,omitempty"`
	Gender string  `db:"gender" json:"gender,omitempty"`
	Orders []Order `db:"orders" json:"orders,omitempty" relation:"user_id"`
	RoleId *int64  `db:"role_id" json:"roleId,omitempty"`
	Role   *Role   `db:"role" relation:"role_id" json:"role,omitempty"`
}

func (user User) Info() string {
	return "users"
}
