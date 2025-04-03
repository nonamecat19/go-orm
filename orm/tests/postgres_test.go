package tests

import (
	_ "github.com/lib/pq"
	"testing"
)

func TestFindUsersWithOrdersAndRole(t *testing.T) {
	client := GetPostgresTestClient()
	FindUsersWithOrdersAndRole(t, client)
}

func TestFindOrdersWithUsers(t *testing.T) {
	client := GetPostgresTestClient()
	FindOrdersWithUsers(t, client)
}
