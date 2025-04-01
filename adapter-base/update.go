package adapter_base

import "fmt"

func Update(tableName string) string {
	return fmt.Sprintf("UPDATE %s", tableName)
}
