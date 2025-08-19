package tag

import (
	"fmt"
	"reflect"

	"github.com/llxxbb/platform-common/def"
	"go.uber.org/zap"
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
		if v.IsNil() {
			return nil, nil
		}
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
		setting := getMapSetting(tag)
		fV := v.Field(i)
		fT := fV.Type().Name()
		if setting.drillSub {
			subMap, err := ToMapN(fV.Interface())
			if err != nil {
				return nil, err
			}
			if len(subMap) > 0 {
				for k, v := range subMap {
					rtn[k] = v
				}
			}
			continue
		}
		switch fT {
		case "int":
			rValue := int(fV.Int())
			if setting.omit && rValue == 0 {
				continue
			}
			rtn[setting.key] = rValue
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
	if ins == nil || mVal == nil || len(mVal) == 0 {
		return nil
	}
	// process ------------------------------
	v := reflect.ValueOf(ins)
	t := v.Type()
	if t.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil
		}
		v = v.Elem()
		t = t.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get(mapNTag)
		if tag == "" {
			continue
		}
		setting := getMapSetting(tag)
		if setting.drillSub {
			subValue := v.Field(i)
			if subValue.Kind() == reflect.Ptr {
				if subValue.IsNil() {
					subValue.Set(reflect.New(subValue.Type().Elem()))
				}
			} else if subValue.Kind() == reflect.Struct {
				subValue = subValue.Addr()
			}
			err := FromMapN(mVal, subValue.Interface(), remove)
			if err != nil {
				return err
			}
			continue
		}
		mV, ok := mVal[setting.key]
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
			delete(mVal, setting.key)
		}
	}
	return nil
}
