package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/nonamecat19/go-orm/orm/tests/test-utils"
	"testing"
)

func TestMySQLFindUsersWithOrdersAndRole(t *testing.T) {
	client := test_utils.GetMySQLTestClient()
	test_utils.FindUsersWithOrdersAndRole(t, client)
}

func TestMySQLFindOrdersWithUsers(t *testing.T) {
	client := test_utils.GetMySQLTestClient()
	test_utils.FindOrdersWithUsers(t, client)
}

func TestMySQLDeleteUsers(t *testing.T) {
	client := test_utils.GetMySQLTestClient()
	test_utils.DeleteUsers(t, client)
}

func TestMySQLUpdateUsers(t *testing.T) {
	client := test_utils.GetMySQLTestClient()
	test_utils.UpdateUsers(t, client)
}
