package adapter_mysql

import (
	"fmt"
	"github.com/nonamecat19/go-orm/core/utils"
)

func (ap AdapterMySQL) Insert(tableName string, fieldNames []string, values []any, args []any) (string, []any) {
	paramsSlice := utils.GenerateParamsSlice(len(values))
	valueSlices := utils.Chunk(paramsSlice, len(fieldNames))
	stringRecords := utils.Map(valueSlices, func(valueSlice []string) string {
		return fmt.Sprintf("(%s)", ap.JoinFields(valueSlice))
	})

	insertQuery := fmt.Sprintf("INSERT INTO %s (%s) VALUES %s;", tableName, ap.JoinFields(fieldNames), ap.JoinFields(stringRecords))
	newArgs := append(args, values...)
	return ap.NormalizeSqlWithArgs(insertQuery, args), newArgs
}
