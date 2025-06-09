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

func TestPostgresDeleteUsers(t *testing.T) {
	client := test_utils.GetPostgresTestClient()
	test_utils.DeleteUsers(t, client)
}

func TestPostgresInsertUser(t *testing.T) {
	client := test_utils.GetPostgresTestClient()
	test_utils.InsertUser(t, client)
}

func TestPostgresUpdateUsers(t *testing.T) {
	client := test_utils.GetPostgresTestClient()
	test_utils.UpdateUsers(t, client)
}
