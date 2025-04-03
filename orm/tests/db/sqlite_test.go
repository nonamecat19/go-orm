package db

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/nonamecat19/go-orm/orm/tests/test-utils"
	"testing"
)

func TestSQLiteFindUsersWithOrdersAndRole(t *testing.T) {
	client := test_utils.GetSQLiteTestClient()
	test_utils.FindUsersWithOrdersAndRole(t, client)
}

func TestSQLiteFindOrdersWithUsers(t *testing.T) {
	client := test_utils.GetSQLiteTestClient()
	test_utils.FindOrdersWithUsers(t, client)
}
