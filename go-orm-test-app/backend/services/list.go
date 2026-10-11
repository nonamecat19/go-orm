package services

import (
	"github.com/nonamecat19/go-orm/orm/lib/querybuilder"
	"shopping-list/backend/database"
	"shopping-list/backend/entities"
)

func GetListWithItems(id int64) (err error, listResult entities.List) {
	var list []entities.List

	err = querybuilder.
		CreateQueryBuilder(database.DbClient).
		Where("id = ?", id).
		Limit(1).
		Preload("items").
		FindMany(&list)

	if len(list) == 0 {
		return err, listResult
	}
	return err, list[0]
}

func GetAllLists() (err error, listsResult []entities.List) {
	var lists []entities.List

	err = querybuilder.
		CreateQueryBuilder(database.DbClient).
		FindMany(&lists)

	return err, lists
}

func DeleteList(id int64) (err error) {
	return querybuilder.CreateQueryBuilder(database.DbClient).
		Where("id = ?", id).
		DeleteMany(&entities.List{})
}

func RemoveItemFromList(itemId int64, listId int64) (err error) {
	return querybuilder.CreateQueryBuilder(database.DbClient).
		Where("id = ?", itemId).
		AndWhere("list_id = ?", listId).
		SetValues(map[string]any{"list_id": nil}).
		UpdateMany(&entities.Item{})
}

type AddItemToListBody struct {
	ItemId int `json:"itemId"`
	ListId int `json:"listId"`
}

func AddItemToList(addItemToList AddItemToListBody) (err error) {
	return querybuilder.CreateQueryBuilder(database.DbClient).
		Where("id = ?", addItemToList.ItemId).
		SetValues(map[string]any{"list_id": addItemToList.ListId}).
		Debug().
		UpdateMany(&entities.Item{})
}
