package tag

import (
	"fmt"
	"gitlab.cdel.local/platform/go/platform-common/def"
	"go.uber.org/zap"
	"reflect"
)

var mapNTag = "mapN"

// ToMapN 用于处理数值标签
func ToMapN(obj any) (map[string]int, *def.CustomError) {
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
	var rtn = make(map[string]int, cnt)
	for i := 0; i < cnt; i++ {
		field := t.Field(i)
		tag := field.Tag.Get(mapNTag)
		if tag == "" {
			continue
		}
		key, omitEmpty := isOmitEmpty(tag)
		fV := v.Field(i)
		fT := fV.Type().Name()
		switch fT {
		case "int":
			rValue := int(fV.Int())
			if omitEmpty && rValue == 0 {
				continue
			}
			rtn[key] = rValue
		default:
			msg := fmt.Sprintf("%s unsupported type: %s", def.SYS_M, fT)
			zap.L().Error(msg)
			return nil, def.NewCustomError(def.ET_SYS, def.SYS_C, msg, nil)
		}
	}
	return rtn, nil
}

// FromMapN notice: para `remove` will remove from `mVal` after restored
func FromMapN(mVal map[string]int, ins any, remove bool) *def.CustomError {
	// check before process
	if mVal == nil || len(mVal) == 0 {
		return nil
	}
	// process ------------------------------
	v := reflect.ValueOf(ins).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get(mapNTag)
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
		case "int":
			parseInt := mV
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
