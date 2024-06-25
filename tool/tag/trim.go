package tag

import (
	"reflect"
	"strings"
)

const TrimTag = "trim"
const CmdSub = "sub"

func TrimFields(data any) {
	v := reflect.ValueOf(data)
	t := v.Type()
	if t.Kind() == reflect.Ptr {
		if v.IsNil() {
			return
		}
		v = v.Elem()
		t = t.Elem()
	} else {
		// Do not support non-pointer struct
		return
	}

	for i := 0; i < v.NumField(); i++ {
		vField := v.Field(i)
		tField := t.Field(i)
		fKind := vField.Kind()
		if fKind == reflect.Ptr && vField.IsNil() {
			continue
		}
		cmd, ok := tField.Tag.Lookup(TrimTag)
		if !ok {
			continue
		}
		if fKind == reflect.String {
			vField.SetString(strings.TrimSpace(vField.String()))
		} else if fKind == reflect.Ptr {
			if cmd == CmdSub {
				TrimFields(vField.Interface())
			}
		} else if fKind == reflect.Struct {
			if cmd == CmdSub {
				TrimFields(vField.Addr().Interface())
			}
		}
	}
}
