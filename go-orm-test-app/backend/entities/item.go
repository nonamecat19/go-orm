package entities

import . "github.com/nonamecat19/go-orm/core/lib/entities"

type Item struct {
	Model
	Name   string `db:"name" json:"name"`
	Bought bool   `db:"bought" json:"bought"`
	ListId *int64 `db:"list_id" json:"listId,omitempty"`
	List   *List  `db:"list" relation:"list_id" json:"list,omitempty"`
}

func (item Item) Info() string {
	return "items"
}
