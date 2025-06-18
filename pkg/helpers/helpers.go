package helpers

import (
	"reflect"
	"strings"
)

func JSONFieldName(obj interface{}, fieldName string) string {
	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if f, ok := t.FieldByName(fieldName); ok {
		tag := f.Tag.Get("json")
		if tag == "" {
			return fieldName
		}
		parts := strings.Split(tag, ",")
		if parts[0] != "" && parts[0] != "-" {
			return parts[0]
		}
	}
	return fieldName
}
