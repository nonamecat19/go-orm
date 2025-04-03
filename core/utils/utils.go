package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

func If[T any](cond bool, vTrue, vFalse T) T {
	if cond {
		return vTrue
	}
	return vFalse
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func PrintStructSlice(value any) {
	fmt.Println(GetStructJSON(value))
}

func GetStructJSON(value any) string {
	jsonData, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return ""
	}

	return string(jsonData)
}

func StringsIntersection(slice1, slice2 []string) []string {
	var result []string
	set := make(map[string]struct{})

	for _, v := range slice2 {
		set[v] = struct{}{}
	}

	for _, v := range slice1 {
		if _, found := set[v]; found {
			result = append(result, v)
		}
	}

	return result
}

func GetFieldNameByTagValue(val reflect.Type, dbTagValue string) (string, error) {
	valType := val

	if valType.Kind() == reflect.Ptr {
		valType = valType.Elem()
	}

	if valType.Kind() != reflect.Struct {
		return "", errors.New("provided value is not a struct or a pointer to a struct")
	}

	for i := 0; i < valType.NumField(); i++ {
		field := valType.Field(i)
		dbTag := field.Tag.Get("db")

		if dbTag == dbTagValue {
			return field.Name, nil
		}
	}

	return "", fmt.Errorf("no field with the specified db tag value found: %s", dbTagValue)
}

func GenerateParamsSlice(n int) []string {
	if n <= 0 {
		return []string{}
	}
	result := make([]string, n)
	for i := 0; i < n; i++ {
		result[i] = "?"
	}
	return result
}

func Map[T any, U any](input []T, transformFunc func(T) U) []U {
	result := make([]U, len(input))
	for i, v := range input {
		result[i] = transformFunc(v)
	}
	return result
}

func MapWithIndex[T any, U any](input []T, transformFunc func(T, int) U) []U {
	result := make([]U, len(input))
	for i, v := range input {
		result[i] = transformFunc(v, i)
	}
	return result
}

func Chunk[T any](input []T, size int) [][]T {
	if size <= 0 {
		panic("Chunk size must be greater than 0")
	}

	var result [][]T
	for i := 0; i < len(input); i += size {
		end := i + size
		if end > len(input) {
			end = len(input)
		}
		result = append(result, input[i:end])
	}
	return result
}

func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func GetModelFields(model any) map[string]any {
	v := reflect.ValueOf(model).Elem()
	t := v.Type()

	fields := make(map[string]any)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		dbTag := field.Tag.Get("db")

		fieldPtr := v.Field(i).Addr().Interface()

		if dbTag != "" {
			fields[dbTag] = fieldPtr
		}
	}

	return fields
}

// AddPrefix adds a prefix to each string in the slice
func AddPrefix(prefix string, slice []string) []string {
	result := make([]string, len(slice))
	for i, s := range slice {
		result[i] = fmt.Sprintf("%s.%s", prefix, s)
	}
	return result
}

// ExtractFields extract all field names from entity
func ExtractFields(entity reflect.Type) []string {
	var fieldNames []string

	for i := 0; i < entity.NumField(); i++ {
		fieldTags := entity.Field(i).Tag
		dbTag := fieldTags.Get("db")
		relationTag := fieldTags.Get("relation")

		if dbTag != "" && relationTag == "" {
			fieldNames = append(fieldNames, dbTag)
		}
	}

	return fieldNames
}
