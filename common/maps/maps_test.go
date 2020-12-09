package maps

import "testing"

func TestFilter(t *testing.T) {

	m := make(map[string]interface{})
	m["a"] = "a1"
	m["b"] = "b1"
	m["c"] = "c1"
	m["d"] = "d1"
	m["e"] = "e1"
	m["f"] = "f1"

	for k, v := range m {
		println(k + "===" + v.(string))
	}
	println("**************")
	arrays := Filter(m, func(K string) bool {
		return K != "a"
	})
	for k, v := range arrays {
		println(k + "===" + v.(string))
	}
}

func TestFilterToArrays(t *testing.T) {

	m := make(map[string]interface{})
	m["a"] = "a1"
	m["b"] = "b1"
	m["c"] = "c1"
	m["d"] = "d1"
	m["e"] = "e1"
	m["f"] = "f1"

	for k, v := range m {
		println(k + "===" + v.(string))
	}
	println("**************")
	arrays := FilterToArrays(m, func(K string) bool {
		return K != "a"
	})
	for _, v := range arrays {
		println(v.(string))
	}
}
