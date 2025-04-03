package tests

import (
	"github.com/nonamecat19/go-orm/core/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func CompareTestOutput(t *testing.T, slice any, fileName string) {
	output := utils.GetStructJSON(slice)
	fileContent, err := os.ReadFile(fileName)
	assert.NoError(t, err, "Expected no error when reading the JSON file")

	assert.JSONEq(t, string(fileContent), output, "Expected JSON output to match contents of the file")
}
