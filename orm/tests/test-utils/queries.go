package test_utils

import (
	"github.com/nonamecat19/go-orm/app/entities"
	client2 "github.com/nonamecat19/go-orm/orm/lib/client"
	"github.com/nonamecat19/go-orm/orm/lib/querybuilder"
	"github.com/stretchr/testify/assert"
	"testing"
)

func FindUsersWithOrdersAndRole(t *testing.T, client client2.DbClient) {
	PrepareDB(t, client)
	var users []entities.User

	err := querybuilder.CreateQueryBuilder(client).
		Where("id > ?", 12).
		AndWhere("id < ?", 18).
		Preload("orders").
		Preload("role").
		OrderBy("id DESC").
		Limit(8).
		Offset(2).
		FindMany(&users)

	assert.NoError(t, err, "Expected no error")
	CompareTestOutput(t, users, "../outputs/FindUsersWithOrdersAndRole")
}

func FindOrdersWithUsers(t *testing.T, client client2.DbClient) {
	PrepareDB(t, client)
	var orders []entities.Order

	err := querybuilder.CreateQueryBuilder(client).
		Where("id <> ?", 8).
		AndWhere("count <> ?", 7).
		OrderBy("id ASC").
		Preload("user").
		Limit(15).
		Offset(1).
		FindMany(&orders)

	assert.NoError(t, err, "Expected no error")
	CompareTestOutput(t, orders, "../outputs/FindOrdersWithUsers")
}
