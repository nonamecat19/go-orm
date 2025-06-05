package adapter_base

import (
	"fmt"
)

func DeleteFromTable(tableName string) string {
	return fmt.Sprintf("DELETE FROM %s", tableName)
}
