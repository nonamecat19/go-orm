package entities

import . "github.com/nonamecat19/go-orm/core/lib/entities"

type List struct {
	Model
	Name  string `db:"name" json:"name"`
	Items []Item `db:"items" relation:"list_id" json:"items,omitempty"`
}

func (list List) Info() string {
	return "lists"
}
