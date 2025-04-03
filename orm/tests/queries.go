package tests

import (
	"github.com/nonamecat19/go-orm/app/entities"
	client2 "github.com/nonamecat19/go-orm/orm/lib/client"
	"github.com/nonamecat19/go-orm/orm/lib/querybuilder"
	"github.com/stretchr/testify/assert"
	"testing"
)

func FindUsersWithOrdersAndRole(t *testing.T, client client2.DbClient) {
	var users []entities.User

	err := querybuilder.CreateQueryBuilder(client).
		Where("name <> ? OR name <> ?", "test1", "User 200").
		AndWhere("name <> '2'").
		AndWhere("name <> ?", '3').
		Preload("orders").
		Preload("role").
		OrderBy("id DESC").
		Limit(8).
		Offset(2).
		FindMany(&users)

	CompareTestOutput(t, users, "./outputs/FindUsersWithOrdersAndRole.json")
	assert.NoError(t, err, "Expected no error")
}

func FindOrdersWithUsers(t *testing.T, client client2.DbClient) {
	var orders []entities.Order

	err := querybuilder.CreateQueryBuilder(client).
		Where("id <> ?", 8).
		AndWhere("count <> ?", 7).
		OrderBy("id ASC").
		Preload("user").
		Limit(15).
		Offset(1).
		FindMany(&orders)

	CompareTestOutput(t, orders, "./outputs/FindOrdersWithUsers.json")
	assert.NoError(t, err, "Expected no error")
}
