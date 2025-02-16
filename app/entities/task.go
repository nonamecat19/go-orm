package entities

import "github.com/nonamecat19/go-orm/core/lib/entities"

type Task struct {
	entities.Model
	Title       string `db:"title" type:"varchar(64)" json:"name"`
	Description string `db:"description" json:"description" nullable:"true"`
	User        User   `db:"user" json:"user" relation:"one-many"`
}

func (user Task) Info() string {
	// TODO: rework to Info and return all info
	return "tasks"
}
