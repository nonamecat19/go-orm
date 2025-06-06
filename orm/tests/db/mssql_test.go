package db

import (
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/nonamecat19/go-orm/orm/tests/test-utils"
	"testing"
)

func TestMSSQLFindUsersWithOrdersAndRole(t *testing.T) {
	client := test_utils.GetMSSQLTestClient()
	test_utils.FindUsersWithOrdersAndRole(t, client)
}

func TestMSSQLFindOrdersWithUsers(t *testing.T) {
	client := test_utils.GetMSSQLTestClient()
	test_utils.FindOrdersWithUsers(t, client)
}

func TestMSSQLDeleteUsers(t *testing.T) {
	client := test_utils.GetMSSQLTestClient()
	test_utils.DeleteUsers(t, client)
}

func TestMSSQLUpdateUsers(t *testing.T) {
	client := test_utils.GetMSSQLTestClient()
	test_utils.UpdateUsers(t, client)
}
