package tests

import (
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestMySQLFindUsersWithOrdersAndRole(t *testing.T) {
	client := GetMySQLTestClient()
	FindUsersWithOrdersAndRole(t, client)
}

func TestMySQLFindOrdersWithUsers(t *testing.T) {
	client := GetMySQLTestClient()
	FindOrdersWithUsers(t, client)
}
