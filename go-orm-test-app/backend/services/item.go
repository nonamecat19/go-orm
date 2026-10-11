package services

import (
	"github.com/nonamecat19/go-orm/orm/lib/querybuilder"
	"shopping-list/backend/database"
	"shopping-list/backend/entities"
)

type ItemCreate struct {
	Name string `json:"name"`
}

func GetAllItems() (err error, items []entities.Item) {
	var records []entities.Item
	err = querybuilder.CreateQueryBuilder(database.DbClient).
		Select("items.id", "items.name", "items.bought").
		FindMany(&records)

	return err, records
}

type ItemUpdate struct {
	Name   string `json:"name,omitempty"`
	Bought bool   `json:"bought,omitempty"`
}

func UpdateItem(itemUpdate ItemUpdate, id int64) (err error) {
	return querybuilder.CreateQueryBuilder(database.DbClient).
		Where("id = ?", id).
		SetValues(map[string]any{"name": itemUpdate.Name, "bought": itemUpdate.Bought}).
		UpdateMany(&entities.Item{})
}

func DeleteItem(id int64) (err error) {
	return querybuilder.CreateQueryBuilder(database.DbClient).
		Where("id = ?", id).
		DeleteMany(&entities.Item{})
}

func CreateItem(newItem entities.Item) (err error) {
	return querybuilder.
		CreateQueryBuilder(database.DbClient).
		InsertOne(newItem)
}
