package conv

import (
	"encoding/json"
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"html/template"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var errNegativeNotAllowed = errors.New("unable to cast negative value")

var (
	errorType       = reflect.TypeOf((*error)(nil)).Elem()
	fmtStringerType = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
)

// ToBool casts any type to a bool type.
func ToBool(a any) bool {
	v, _ := ToBoolλ(a)
	return v
}

// ToBoolλ casts any type to a bool type.
func ToBoolλ(a any) (bool, error) {
	a = indirect(a)

	switch b := a.(type) {
	case bool:
		return b, nil
	case nil:
		return false, nil
	case int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32, uintptr, complex64, complex128:
		return !reflect.ValueOf(a).IsZero(), nil
	case string:
		return strconv.ParseBool(a.(string))
	case time.Duration:
		return b != 0, nil
	case json.Number:
		v, err := b.Float64()
		return v != 0, err
	default:
		return false, fmt.Errorf("unable to cast %#v of type %T to bool", a, a)
	}
}

// ToInt casts any type to an int type.
func ToInt(i any) int {
	v, _ := ToIntλ(i)
	return v
}

// ToIntλ casts any type to an int type.
func ToIntλ(i any) (int, error) {
	i = indirect(i)

	intv, ok := toInt(i)
	if ok {
		return intv, nil
	}

	switch s := i.(type) {
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case int64:
		return int(s), nil
	case int32:
		return int(s), nil
	case int16:
		return int(s), nil
	case int8:
		return int(s), nil
	case uint:
		return int(s), nil
	case uint64:
		return int(s), nil
	case uint32:
		return int(s), nil
	case uint16:
		return int(s), nil
	case uint8:
		return int(s), nil
	case float64:
		return int(s), nil
	case float32:
		return int(s), nil
	case string:
		v, err := strconv.ParseInt(trimZeroDecimal(s), 0, 0)
		if err == nil {
			return int(v), nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to int64", i, i)
	case json.Number:
		v, err := s.Int64()
		return int(v), err
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int", i, i)
	}
}

// ToInt8 casts any type to an int8 type.
func ToInt8(i any) int8 {
	v, _ := ToInt8λ(i)
	return v
}

// ToInt8λ casts any type to an int8 type.
func ToInt8λ(i any) (int8, error) {
	i = indirect(i)

	intv, ok := toInt(i)
	if ok {
		return int8(intv), nil
	}

	switch s := i.(type) {
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case int64:
		return int8(s), nil
	case int32:
		return int8(s), nil
	case int16:
		return int8(s), nil
	case int8:
		return s, nil
	case uint:
		return int8(s), nil
	case uint64:
		return int8(s), nil
	case uint32:
		return int8(s), nil
	case uint16:
		return int8(s), nil
	case uint8:
		return int8(s), nil
	case float64:
		return int8(s), nil
	case float32:
		return int8(s), nil
	case string:
		v, err := strconv.ParseInt(trimZeroDecimal(s), 0, 0)
		if err == nil {
			return int8(v), nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to int8", i, i)
	case json.Number:
		v, err := s.Int64()
		return int8(v), err
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int8", i, i)
	}
}

// ToInt16 casts any type to an int16 type.
func ToInt16(i any) int16 {
	v, _ := ToInt16λ(i)
	return v
}

// ToInt16λ casts any type to an int16 type.
func ToInt16λ(i any) (int16, error) {
	i = indirect(i)

	intv, ok := toInt(i)
	if ok {
		return int16(intv), nil
	}

	switch s := i.(type) {
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case int64:
		return int16(s), nil
	case int32:
		return int16(s), nil
	case int16:
		return s, nil
	case int8:
		return int16(s), nil
	case uint:
		return int16(s), nil
	case uint64:
		return int16(s), nil
	case uint32:
		return int16(s), nil
	case uint16:
		return int16(s), nil
	case uint8:
		return int16(s), nil
	case float64:
		return int16(s), nil
	case float32:
		return int16(s), nil
	case string:
		v, err := strconv.ParseInt(trimZeroDecimal(s), 0, 0)
		if err == nil {
			return int16(v), nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to int16", i, i)
	case json.Number:
		v, err := s.Int64()
		return int16(v), err
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int16", i, i)
	}
}

// ToInt32 casts any type to an int32 type.
func ToInt32(i any) int32 {
	v, _ := ToInt32λ(i)
	return v
}

// ToInt32λ casts any type to an int32 type.
func ToInt32λ(i any) (int32, error) {
	i = indirect(i)

	intv, ok := toInt(i)
	if ok {
		return int32(intv), nil
	}

	switch s := i.(type) {
	case int64:
		return int32(s), nil
	case int32:
		return s, nil
	case int16:
		return int32(s), nil
	case int8:
		return int32(s), nil
	case uint:
		return int32(s), nil
	case uint64:
		return int32(s), nil
	case uint32:
		return int32(s), nil
	case uint16:
		return int32(s), nil
	case uint8:
		return int32(s), nil
	case float64:
		return int32(s), nil
	case float32:
		return int32(s), nil
	case string:
		v, err := strconv.ParseInt(trimZeroDecimal(s), 0, 0)
		if err == nil {
			return int32(v), nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to int32", i, i)
	case json.Number:
		v, err := s.Int64()
		return int32(v), err
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int32", i, i)
	}
}

// ToInt64 casts any type to an int64 type.
func ToInt64(i any) int64 {
	v, _ := ToInt64λ(i)
	return v
}

// ToInt64λ casts any to an int64 type.
func ToInt64λ(i any) (int64, error) {
	i = indirect(i)

	intv, ok := toInt(i)
	if ok {
		return int64(intv), nil
	}

	switch s := i.(type) {
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case int64:
		return s, nil
	case int32:
		return int64(s), nil
	case int16:
		return int64(s), nil
	case int8:
		return int64(s), nil
	case uint:
		return int64(s), nil
	case uint64:
		return int64(s), nil
	case uint32:
		return int64(s), nil
	case uint16:
		return int64(s), nil
	case uint8:
		return int64(s), nil
	case float64:
		return int64(s), nil
	case float32:
		return int64(s), nil
	case string:
		v, err := strconv.ParseInt(trimZeroDecimal(s), 0, 0)
		if err == nil {
			return v, nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to int64", i, i)
	case json.Number:
		return s.Int64()
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int64", i, i)
	}
}

// ToUint casts any type to a uint type.
func ToUint(i any) uint {
	v, _ := ToUintλ(i)
	return v
}

// ToUintλ casts any type to a uint type.
func ToUintλ(i any) (uint, error) {
	i = indirect(i)

	intv, ok := toInt(i)
	if ok {
		if intv < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint(intv), nil
	}

	switch s := i.(type) {
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case int64:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint(s), nil
	case int32:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint(s), nil
	case int16:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint(s), nil
	case int8:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint(s), nil
	case uint:
		return s, nil
	case uint64:
		return uint(s), nil
	case uint32:
		return uint(s), nil
	case uint16:
		return uint(s), nil
	case uint8:
		return uint(s), nil
	case float64:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint(s), nil
	case float32:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint(s), nil
	case string:
		v, err := strconv.ParseInt(trimZeroDecimal(s), 0, 0)
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint(v), err
	case json.Number:
		v, err := s.Int64()
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint(v), err
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint", i, i)
	}
}

// ToUint8 casts any type to a uint type.
func ToUint8(i any) uint8 {
	v, _ := ToUint8λ(i)
	return v
}

// ToUint8λ casts any type to a uint type.
func ToUint8λ(i any) (uint8, error) {
	i = indirect(i)

	intv, ok := toInt(i)
	if ok {
		if intv < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint8(intv), nil
	}

	switch s := i.(type) {
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case int64:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint8(s), nil
	case int32:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint8(s), nil
	case int16:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint8(s), nil
	case int8:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint8(s), nil
	case uint:
		return uint8(s), nil
	case uint64:
		return uint8(s), nil
	case uint32:
		return uint8(s), nil
	case uint16:
		return uint8(s), nil
	case uint8:
		return s, nil
	case float64:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint8(s), nil
	case float32:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint8(s), nil
	case string:
		v, err := strconv.ParseInt(trimZeroDecimal(s), 0, 0)
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint8(v), err
	case json.Number:
		v, err := s.Int64()
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint8(v), err
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint8", i, i)
	}
}

// ToUint16 casts any type to a uint16 type.
func ToUint16(i any) uint16 {
	v, _ := ToUint16λ(i)
	return v
}

// ToUint16λ casts any type to a uint16 type.
func ToUint16λ(i any) (uint16, error) {
	i = indirect(i)

	intv, ok := toInt(i)
	if ok {
		if intv < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint16(intv), nil
	}

	switch s := i.(type) {
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case int64:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint16(s), nil
	case int32:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint16(s), nil
	case int16:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint16(s), nil
	case int8:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint16(s), nil
	case uint:
		return uint16(s), nil
	case uint64:
		return uint16(s), nil
	case uint32:
		return uint16(s), nil
	case uint16:
		return s, nil
	case uint8:
		return uint16(s), nil
	case float64:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint16(s), nil
	case float32:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint16(s), nil
	case string:
		v, err := strconv.ParseInt(trimZeroDecimal(s), 0, 0)
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint16(v), err
	case json.Number:
		v, err := s.Int64()
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint16(v), err
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint16", i, i)
	}
}

// ToUint32 casts any type to a uint32 type.
func ToUint32(i any) uint32 {
	v, _ := ToUint32λ(i)
	return v
}

// ToUint32λ casts any type to a uint32 type.
func ToUint32λ(i any) (uint32, error) {
	i = indirect(i)

	intv, ok := toInt(i)
	if ok {
		if intv < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint32(intv), nil
	}

	switch s := i.(type) {
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case int64:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint32(s), nil
	case int32:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint32(s), nil
	case int16:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint32(s), nil
	case int8:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint32(s), nil
	case uint:
		return uint32(s), nil
	case uint64:
		return uint32(s), nil
	case uint32:
		return s, nil
	case uint16:
		return uint32(s), nil
	case uint8:
		return uint32(s), nil
	case float64:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint32(s), nil
	case float32:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint32(s), nil
	case string:
		v, err := strconv.ParseInt(trimZeroDecimal(s), 0, 0)
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint32(v), err
	case json.Number:
		v, err := s.Int64()
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint32(v), err
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint32", i, i)
	}
}

// ToUint64 casts any type to a uint64 type.
func ToUint64(i any) uint64 {
	v, _ := ToUint64λ(i)
	return v
}

// ToUint64λ casts any type to a uint64 type.
func ToUint64λ(i any) (uint64, error) {
	i = indirect(i)

	intv, ok := toInt(i)
	if ok {
		if intv < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(intv), nil
	}

	switch s := i.(type) {
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case int64:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(s), nil
	case int32:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(s), nil
	case int16:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(s), nil
	case int8:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(s), nil
	case uint:
		return uint64(s), nil
	case uint64:
		return s, nil
	case uint32:
		return uint64(s), nil
	case uint16:
		return uint64(s), nil
	case uint8:
		return uint64(s), nil
	case float32:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(s), nil
	case float64:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(s), nil
	case string:
		v, err := strconv.ParseInt(trimZeroDecimal(s), 0, 0)
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(v), err
	case json.Number:
		v, err := s.Int64()
		if v < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(v), err
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint64", i, i)
	}
}

// ToFloat32 casts any type to a float32 type.
func ToFloat32(i any) float32 {
	v, _ := ToFloat32λ(i)
	return v
}

// ToFloat32λ casts any type to a float32 type.
func ToFloat32λ(i any) (float32, error) {
	i = indirect(i)

	intv, ok := toInt(i)
	if ok {
		return float32(intv), nil
	}

	switch s := i.(type) {
	case float64:
		return float32(s), nil
	case float32:
		return s, nil
	case int64:
		return float32(s), nil
	case int32:
		return float32(s), nil
	case int16:
		return float32(s), nil
	case int8:
		return float32(s), nil
	case uint:
		return float32(s), nil
	case uint64:
		return float32(s), nil
	case uint32:
		return float32(s), nil
	case uint16:
		return float32(s), nil
	case uint8:
		return float32(s), nil
	case string:
		v, err := strconv.ParseFloat(s, 32)
		return float32(v), err
	case json.Number:
		v, err := s.Float64()
		return float32(v), err
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to float32", i, i)
	}
}

// ToFloat64 casts any type to a float64 type.
func ToFloat64(i any) float64 {
	v, _ := ToFloat64λ(i)
	return v
}

// ToFloat64λ casts any type to a float64 type.
func ToFloat64λ(i any) (float64, error) {
	i = indirect(i)

	intv, ok := toInt(i)
	if ok {
		return float64(intv), nil
	}

	switch s := i.(type) {
	case float64:
		return s, nil
	case float32:
		return float64(s), nil
	case int64:
		return float64(s), nil
	case int32:
		return float64(s), nil
	case int16:
		return float64(s), nil
	case int8:
		return float64(s), nil
	case uint:
		return float64(s), nil
	case uint64:
		return float64(s), nil
	case uint32:
		return float64(s), nil
	case uint16:
		return float64(s), nil
	case uint8:
		return float64(s), nil
	case string:
		return strconv.ParseFloat(s, 64)
	case json.Number:
		return s.Float64()
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to float64", i, i)
	}
}

// ToString casts any type to a string type.
func ToString(i any) string {
	v, _ := ToStringλ(i)
	return v
}

// ToStringλ casts any type to a string type.
func ToStringλ(i any) (string, error) {
	i = indirectToStringerOrError(i)

	switch s := i.(type) {
	case string:
		return s, nil
	case bool:
		return strconv.FormatBool(s), nil
	case int:
		return strconv.Itoa(s), nil
	case int64:
		return strconv.FormatInt(s, 10), nil
	case int32:
		return strconv.Itoa(int(s)), nil
	case int16:
		return strconv.FormatInt(int64(s), 10), nil
	case int8:
		return strconv.FormatInt(int64(s), 10), nil
	case uint:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint64:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(s), 10), nil
	case float64:
		return strconv.FormatFloat(s, 'f', -1, 64), nil
	case float32:
		return strconv.FormatFloat(float64(s), 'f', -1, 32), nil
	case json.Number:
		return s.String(), nil
	case []byte:
		return string(s), nil
	case template.HTML:
		return string(s), nil
	case template.HTMLAttr:
		return string(s), nil
	case template.URL:
		return string(s), nil
	case template.JS:
		return string(s), nil
	case template.JSStr:
		return string(s), nil
	case template.CSS:
		return string(s), nil
	case template.Srcset:
		return string(s), nil
	case nil:
		return "", nil
	case fmt.Stringer:
		return s.String(), nil
	case error:
		return s.Error(), nil
	default:
		// 如果是对象，则转为 JSON 字符串
		if reflect.TypeOf(s).Kind() == reflect.Struct {
			var json = jsoniter.ConfigCompatibleWithStandardLibrary
			b, err := json.Marshal(s)
			if err != nil {
				return fmt.Sprintf("%v", s), nil
			}
			return string(b), nil
		}
		return "", fmt.Errorf("Unable to cast %#v of type %T to string", i, i)
	}
}

// ToDuration casts any type to a time.Duration type.
func ToDuration(i any) time.Duration {
	v, _ := ToDurationλ(i)
	return v
}

// ToDurationλ casts any type to time.Duration type.
func ToDurationλ(i any) (time.Duration, error) {
	i = indirect(i)

	switch s := i.(type) {
	case time.Duration:
		return s, nil
	case int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8:
		return time.Duration(ToAny[int64](s)), nil
	case float32, float64:
		return time.Duration(ToAny[float64](s)), nil
	case string:
		if strings.ContainsAny(s, "nsuµmh") {
			return time.ParseDuration(s)
		}
		return time.ParseDuration(s + "ns")
	case json.Number:
		v, err := s.Float64()
		return time.Duration(v), err
	default:
		return time.Duration(0), fmt.Errorf("unable to cast %#v of type %T to Duration", i, i)
	}
}

// toInt returns the int value of v if v or v's underlying type is an int.
// Note that this will return false for int64 etc. types.
func toInt(v any) (int, bool) {
	switch v := v.(type) {
	case int:
		return v, true
	case time.Weekday:
		return int(v), true
	case time.Month:
		return int(v), true
	default:
		return 0, false
	}
}

// Copied from html/template/content.go.
// indirect returns the value, after dereferencing as many times
// as necessary to reach the base type (or nil).
func indirect(a any) any {
	if a == nil {
		return nil
	}
	if t := reflect.TypeOf(a); t.Kind() != reflect.Pointer {
		// Avoid creating a reflect.Value if it's not a pointer.
		return a
	}
	v := reflect.ValueOf(a)
	for v.Kind() == reflect.Pointer && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

// Copied from html/template/content.go.
// indirectToStringerOrError returns the value, after dereferencing as many times
// as necessary to reach the base type (or nil) or an implementation of fmt.Stringer
// or error.
func indirectToStringerOrError(a any) any {
	if a == nil {
		return nil
	}
	v := reflect.ValueOf(a)
	for !v.Type().Implements(fmtStringerType) && !v.Type().Implements(errorType) && v.Kind() == reflect.Pointer && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

// trimZeroDecimal trims the zero decimal.
// E.g. 12.00 to 12 while 12.01 still to be 12.01.
func trimZeroDecimal(s string) string {
	var foundZero bool
	for i := len(s); i > 0; i-- {
		switch s[i-1] {
		case '.':
			if foundZero {
				return s[:i-1]
			}
		case '0':
			foundZero = true
		default:
			return s
		}
	}
	return s
}
