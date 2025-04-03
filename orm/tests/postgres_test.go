package tests

import (
	_ "github.com/lib/pq"
	"testing"
)

func TestPostgresFindUsersWithOrdersAndRole(t *testing.T) {
	client := GetPostgresTestClient()
	FindUsersWithOrdersAndRole(t, client)
}

func TestPostgresFindOrdersWithUsers(t *testing.T) {
	client := GetPostgresTestClient()
	FindOrdersWithUsers(t, client)
}
