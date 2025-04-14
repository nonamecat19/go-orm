package test_utils

import (
	"fmt"
	"github.com/nonamecat19/go-orm/core/utils"
	"github.com/nonamecat19/go-orm/orm/lib/client"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func CompareTestOutput(t *testing.T, slice any, fileName string) {
	output := utils.GetStructJSON(slice)
	fileContent, err := os.ReadFile(fileName + ".json")
	assert.NoError(t, err, "Expected no error when reading the JSON file")

	assert.JSONEq(t, string(fileContent), output, "Expected JSON output to match contents of the file")
}

func PrepareDB(t *testing.T, client client.DbClient) {
	dbDriver := client.GetAdapter().GetDbDriver()
	fileContent, err := os.ReadFile(fmt.Sprintf("../dumps/%s.sql", dbDriver))
	if err != nil {
		assert.NoError(t, err, "Expected no error")
	}

	queries := strings.Split(string(fileContent), ";")

	for _, query := range queries {
		query = strings.TrimSpace(query)
		if query == "" {
			continue
		}
		_, err := client.GetDb().Exec(query)
		assert.NoError(t, err, "Error executing query: %s", query)
	}

	assert.NoError(t, err, "Expected no error")
}
