package types

import "reflect"

// Optional 定义一个 Optional 类型。
type Optional[T any] struct {
	value T
}

// NewOptional 创建一个新的 Optional 对象。
func NewOptional[T any](value T) Optional[T] {
	return Optional[T]{value}
}

// IsEmpty 检查 Optional 对象是否为空。
func (o Optional[T]) IsEmpty() bool {
	v := reflect.ValueOf(o.value)
	switch v.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Func, reflect.Chan, reflect.Interface:
		return v.IsNil()
	default:
		return false
	}
}

// Get 获取 Optional 对象的值。
func (o Optional[T]) Get() T {
	if o.IsEmpty() {
		panic("Optional is empty")
	}
	return o.value
}

// OrElseGet 如果 Optional 对象不为空，则返回其值，否则返回给定值。
func (o Optional[T]) OrElseGet(value T) T {
	if o.IsEmpty() {
		return value
	}
	return o.value
}

// IfPresent 如果 Optional 对象不为空，则执行给定的消费者。
func (o Optional[T]) IfPresent(consumer func(T)) {
	if !o.IsEmpty() {
		consumer(o.value)
	}
}
