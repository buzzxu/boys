package conv

import (
	"fmt"
	"github.com/buzzxu/boys/common/bytess"
	"reflect"
)

func ToAny[T any](a any) T {
	v, _ := ToAnyλ[T](a)
	return v
}

// ToAnyλ converts one type to another and returns an error if occurred.
func ToAnyλ[T any](a any) (T, error) {
	var t T
	switch any(t).(type) {
	case bool:
		v, err := ToBoolλ(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case int:
		v, err := ToIntλ(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case int8:
		v, err := ToInt8λ(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case int16:
		v, err := ToInt16λ(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case int32:
		v, err := ToInt32λ(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case int64:
		v, err := ToInt64λ(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case uint:
		v, err := ToUintλ(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case uint8:
		v, err := ToUint8λ(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case uint16:
		v, err := ToUint16λ(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case uint32:
		v, err := ToUint32λ(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case uint64:
		v, err := ToUint64λ(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case float32:
		v, err := ToFloat32λ(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case float64:
		v, err := ToFloat64λ(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case string:
		v, err := ToStringλ(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	default:
		return t, fmt.Errorf("the type %T is not supported", t)
	}
	return t, nil
}

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
