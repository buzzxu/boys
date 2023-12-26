package types

import "testing"

func TestOptional_IsEmpty(t *testing.T) {
	v := NewOptional(1)
	println(v.IsEmpty())
}
