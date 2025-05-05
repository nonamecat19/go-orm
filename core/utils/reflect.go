package utils

import (
	"reflect"
	"strconv"
	"time"
)

func StringifyReflectValue(v reflect.Value) string {
	if !v.IsValid() {
		return "NULL"
	}

	switch v.Kind() {
	case reflect.String:
		return v.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64)
	case reflect.Bool:
		if v.Bool() == true {
			return "true"
		}
		return "false"
	case reflect.Ptr:
		if v.IsNil() {
			return "NULL"
		}
		return StringifyReflectValue(v.Elem())
	case reflect.Interface:
		if v.IsNil() {
			return "NULL"
		}
		fallthrough
	case reflect.Struct:
		if v.Type() == reflect.TypeOf(time.Time{}) {
			t := v.Interface().(time.Time)
			return t.Format("02/01/2006 15:04:05")
		}
		fallthrough
	default:
		return v.Interface().(string)
	}
}
