package entities

import "github.com/nonamecat19/go-orm/core/lib/entities"

type Task struct {
	entities.Model
	Title       string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
}

func (user *Task) TableName() string {
	return "tasks"
}
