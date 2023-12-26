package conv

import (
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/mitchellh/mapstructure"
	"reflect"
	"strconv"
	"time"
)

// MapToStruct map转Struct
func MapToStruct(data map[string]interface{}, obj interface{}) error {
	return mapstructure.Decode(data, obj)
}

// StructToMap converts struct to map[string]any.
// Such as struct{I int, S string}{I: 1, S: "a"} to map["I":1 "S":"a"].
// Note that unexported fields of struct can't be converted.
func StructToMap(a any) map[string]any {
	v := reflect.ValueOf(a)
	if v.Kind() != reflect.Struct {
		return nil
	}

	t := reflect.TypeOf(a)
	var m = make(map[string]any)
	for i := 0; i < t.NumField(); i++ {
		// 过滤私有的,不然会G
		f := t.Field(i)
		if f.PkgPath != "" {
			continue
		}
		if f.IsExported() {
			tag, have := f.Tag.Lookup("json")
			var fieldName string
			if have {
				fieldName = tag
			} else {
				fieldName = f.Name
			}
			switch v.Field(i).Kind() {
			case reflect.Struct:
				valueValue := v.Field(i).Interface()
				m[fieldName] = StructToMap(valueValue)
			default:
				m[fieldName] = v.Field(i).Interface()
			}
		}
	}
	return m
}

// StructToMapStr converts struct to map[string]string.
// Such as struct{I int, S string}{I: 1, S: "a"} to map["I":"1" "S":"a"].
// Note that unexported fields of struct can't be converted.
func StructToMapStr(a any) map[string]string {
	v := reflect.ValueOf(a)
	if v.Kind() != reflect.Struct {
		return nil
	}

	t := reflect.TypeOf(a)
	var m = make(map[string]string)
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).IsExported() {
			m[t.Field(i).Name] = ToString(v.Field(i).Interface())
		}
	}
	return m
}

// ToMapStr converts any type to a map[string]string type.
func ToMapStr(a any) map[string]string {
	v, _ := ToMapStrλ(a)
	return v
}

// ToMapStrλ converts any type to a map[string]string type.
func ToMapStrλ(a any) (map[string]string, error) {
	var m = map[string]string{}

	switch v := a.(type) {
	case map[string]string:
		return v, nil
	case map[string]any:
		for k, val := range v {
			val, err := ToAnyλ[string](val)
			if err != nil {
				return nil, err
			}
			m[k] = val
		}
	case map[any]string:
		for k, val := range v {
			k, err := ToAnyλ[string](k)
			if err != nil {
				return nil, err
			}
			m[k] = val
		}
	case map[any]any:
		for k, val := range v {
			k, err := ToAnyλ[string](k)
			if err != nil {
				return nil, err
			}
			val, err := ToAnyλ[string](val)
			if err != nil {
				return nil, err
			}
			m[k] = val
		}
	case string:
		if err := jsonTo(v, &m); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unable to convert %#v of type %T to map[string]string", a, a)
	}
	return m, nil
}

// jsonTo attempts to unmarshall a string as JSON into
// the object passed as pointer.
func jsonTo(s string, v any) error {
	data := []byte(s)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal(data, v)
}

// setField 用map的值替换结构的值
func setField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()        //结构体属性值
	structFieldValue := structValue.FieldByName(name) //结构体单个属性值
	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}
	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}
	structFieldType := structFieldValue.Type() //结构体的类型
	val := reflect.ValueOf(value)              //map值的反射值
	var err error
	if structFieldType != val.Type() {
		val, err = typeConversion(fmt.Sprintf("%v", value), structFieldValue.Type().Name()) //类型转换
		if err != nil {
			return err
		}
	}
	structFieldValue.Set(val)
	return nil
}

// typeConversion 类型转换
func typeConversion(value string, ntype string) (reflect.Value, error) {
	if ntype == "string" {
		return reflect.ValueOf(value), nil
	} else if ntype == "time.Time" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "Time" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "int" {
		i, err := strconv.Atoi(value)
		return reflect.ValueOf(i), err
	} else if ntype == "int8" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int8(i)), err
	} else if ntype == "int32" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int64(i)), err
	} else if ntype == "int64" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(i), err
	} else if ntype == "float32" {
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(float32(i)), err
	} else if ntype == "float64" {
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(i), err
	}
	//else if .......增加其他一些类型的转换
	return reflect.ValueOf(value), errors.New("未知的类型：" + ntype)
}
