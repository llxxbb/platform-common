package tag

import (
	"fmt"
	"gitlab.cdel.local/platform/go/platform-common/def"
	"go.uber.org/zap"
	"reflect"
	"strconv"
	"strings"
)

func ToMap(obj any) (map[string]string, *def.CustomError) {
	if obj == nil {
		return nil, nil
	}
	v := reflect.ValueOf(obj)
	t := v.Type()
	if t.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}

	cnt := t.NumField()
	var data = make(map[string]string, cnt)
	for i := 0; i < cnt; i++ {
		field := t.Field(i)
		tag := field.Tag.Get("map")
		if tag == "" {
			continue
		}
		key, omitEmpty := isOmitEmpty(tag)
		fV := v.Field(i)
		fT := fV.Type().Name()
		switch fT {
		case "string":
			rValue := fV.Interface().(string)
			if omitEmpty && rValue == "" {
				continue
			}
			data[key] = rValue
		case "bool":
			rValue := fV.Bool()
			if omitEmpty && !rValue {
				continue
			}
			data[key] = strconv.FormatBool(rValue)
		case "int":
			rValue := int(fV.Int())
			if omitEmpty && rValue == 0 {
				continue
			}
			data[key] = strconv.Itoa(rValue)
		default:
			msg := fmt.Sprintf("%s unsupported type: %s", def.SYS_M, fT)
			zap.L().Error(msg)
			return nil, def.NewCustomError(def.ET_SYS, def.SYS_C, msg, nil)
		}
	}
	return data, nil
}

// return key and omitEmpty
func isOmitEmpty(tag string) (string, bool) {
	split := strings.Split(tag, ",")
	omitEmpty := false
	if len(split) > 1 {
		if split[1] == "omitempty" {
			omitEmpty = true
		}
	}
	return split[0], omitEmpty
}

// FromMap notice: para `remove` will remove from `mVal` after restored
func FromMap(mVal map[string]string, ins any, remove bool) *def.CustomError {
	// check before process
	if mVal == nil || len(mVal) == 0 {
		return nil
	}
	// process ------------------------------
	v := reflect.ValueOf(ins).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("map")
		if tag == "" {
			continue
		}
		key, _ := isOmitEmpty(tag)
		mV, ok := mVal[key]
		if !ok {
			continue
		}
		fV := v.Field(i)
		fT := fV.Type().Name()
		switch fT {
		case "string":
			fV.SetString(mV)
		case "bool":
			parseBool, _ := strconv.ParseBool(mV)
			fV.SetBool(parseBool)
		case "int":
			parseInt, _ := strconv.Atoi(mV)
			fV.SetInt(int64(parseInt))
		default:
			msg := fmt.Sprintf("%s unsupported type: %s", def.SYS_M, fT)
			zap.L().Error(msg)
			return def.NewCustomError(def.ET_SYS, def.SYS_C, msg, nil)
		}
		if remove {
			delete(mVal, key)
		}
	}
	return nil
}
