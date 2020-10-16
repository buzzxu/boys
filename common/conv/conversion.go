package conv

import (
	"errors"
	"fmt"
	"github.com/buzzxu/boys/common/bytess"
	"reflect"
	"strconv"
	"time"
)

// []byte to string
func String(bytes *[]byte) *string {
	return bytess.String(bytes)
}

func IntPtrTo64(ptr interface{}) (value int64) {
	if v := reflect.ValueOf(ptr); v.Kind() == reflect.Ptr {
		p := v.Elem()
		switch p.Kind() {
		case reflect.Int:
			value = int64(*ptr.(*int))
		case reflect.Int8:
			value = int64(*ptr.(*int8))
		case reflect.Int16:
			value = int64(*ptr.(*int16))
		case reflect.Int32:
			value = int64(*ptr.(*int32))
		case reflect.Int64:
			value = *ptr.(*int64)
		}
	}
	return
}

func UintPtrTo64(ptr interface{}) (value uint64) {
	if v := reflect.ValueOf(ptr); v.Kind() == reflect.Ptr {
		p := v.Elem()
		switch p.Kind() {
		case reflect.Uint:
			value = uint64(*ptr.(*uint))
		case reflect.Uint8:
			value = uint64(*ptr.(*uint8))
		case reflect.Uint16:
			value = uint64(*ptr.(*uint16))
		case reflect.Uint32:
			value = uint64(*ptr.(*uint32))
		case reflect.Uint64:
			value = *ptr.(*uint64)
		}
	}
	return
}

//StructToMap struct 转Map
func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if len(f.Tag.Get("json")) == 0 {
			data[f.Name] = v.Field(i).Interface()
			continue
		}
		data[f.Tag.Get("json")] = v.Field(i).Interface()
	}
	return data
}

//MapToStruct map转Struct
func MapToStruct(data map[string]interface{}, obj interface{}) error {
	for k, v := range data {
		err := setField(obj, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

//setField 用map的值替换结构的值
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

//typeConversion 类型转换
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
