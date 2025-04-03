package db

import (
	_ "github.com/lib/pq"
	"github.com/nonamecat19/go-orm/orm/tests/test-utils"
	"testing"
)

func TestPostgresFindUsersWithOrdersAndRole(t *testing.T) {
	client := test_utils.GetPostgresTestClient()
	test_utils.FindUsersWithOrdersAndRole(t, client)
}

func TestPostgresFindOrdersWithUsers(t *testing.T) {
	client := test_utils.GetPostgresTestClient()
	test_utils.FindOrdersWithUsers(t, client)
}
