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
	jsonData, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonData))
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

	return "", errors.New("no field with the specified db tag value found")

}
