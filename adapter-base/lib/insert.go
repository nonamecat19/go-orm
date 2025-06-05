package adapter_base

import (
	"fmt"
	"github.com/nonamecat19/go-orm/core/utils"
)

func Insert(tableName string, fieldNames []string, values []any, args []any) (string, []any) {
	paramsSlice := utils.GenerateParamsSlice(len(values))
	valueSlices := utils.Chunk(paramsSlice, len(fieldNames))
	stringRecords := utils.Map(valueSlices, func(valueSlice []string) string {
		return fmt.Sprintf("(%s)", JoinFields(valueSlice))
	})

	insertQuery := fmt.Sprintf("INSERT INTO %s (%s) VALUES %s;", tableName, JoinFields(fieldNames), JoinFields(stringRecords))
	newArgs := append(args, values...)
	return NormalizeSqlWithArgs(insertQuery, args), newArgs
}
