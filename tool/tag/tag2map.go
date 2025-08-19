package tag

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/llxxbb/platform-common/def"
	"go.uber.org/zap"
)

var mapTag = "map"

func ToMap(obj any) (map[string]string, *def.CustomError) {
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
	var data = make(map[string]string, cnt)
	for i := 0; i < cnt; i++ {
		field := t.Field(i)
		tag := field.Tag.Get(mapTag)
		if tag == "" {
			continue
		}
		setting := getMapSetting(tag)
		fV := v.Field(i)
		fT := fV.Type().Name()
		if setting.drillSub {
			subMap, err := ToMap(fV.Interface())
			if err != nil {
				return nil, err
			}
			if len(subMap) > 0 {
				for k, v := range subMap {
					data[k] = v
				}
			}
			continue
		}
		switch fT {
		case "string":
			rValue := fV.Interface().(string)
			if setting.omit && rValue == "" {
				continue
			}
			data[setting.key] = rValue
		case "bool":
			rValue := fV.Bool()
			if setting.omit && !rValue {
				continue
			}
			data[setting.key] = strconv.FormatBool(rValue)
		case "int":
			rValue := int(fV.Int())
			if setting.omit && rValue == 0 {
				continue
			}
			data[setting.key] = strconv.Itoa(rValue)
		default:
			msg := fmt.Sprintf("%s unsupported type: %s", def.SYS_M, fT)
			zap.L().Error(msg)
			return nil, def.NewCustomError(def.ET_SYS, def.SYS_C, msg, nil)
		}
	}
	return data, nil
}

// FromMap notice: para `remove` will remove from `mVal` after restored
func FromMap(mVal map[string]string, ins any, remove bool) *def.CustomError {
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
		tag := field.Tag.Get(mapTag)
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
			err := FromMap(mVal, subValue.Interface(), remove)
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
			delete(mVal, setting.key)
		}
	}
	return nil
}
