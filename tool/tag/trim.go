package tag

import (
	"reflect"
	"strings"
)

const TrimTag = "trim"

func TrimFields(data any) {
	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return
	}

	v = v.Elem()
	if v.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.String {
			_, ok := v.Type().Field(i).Tag.Lookup(TrimTag)
			if ok {
				field.SetString(strings.TrimSpace(field.String()))
			}
		}
	}
}
