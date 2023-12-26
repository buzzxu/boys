package conv

import (
	"fmt"
	"testing"
)

func TestStructToMap(t *testing.T) {
	obj := &struct {
		Name    string
		Age     int
		Address struct {
			City string
			Code string `json:"code1"`
		}
	}{
		Name: "tom",
		Age:  18,
		Address: struct {
			City string
			Code string `json:"code1"`
		}{
			City: "shanghai",
			Code: "010",
		},
	}

	_map := StructToMap(*obj)
	println(_map)
}

func TestMapToStruct(t *testing.T) {
	obj := &struct {
		Name    string
		Age     int
		Address struct {
			City string
			Code string `json:"code1"`
		}
	}{
		Name: "tom",
		Age:  18,
		Address: struct {
			City string
			Code string `json:"code1"`
		}{
			City: "shanghai",
			Code: "010",
		},
	}

	_map := StructToMap(*obj)
	var newObj struct {
		Name    string
		Age     int
		Address struct {
			City string
			Code string `json:"code1"`
		}
	}
	err := MapToStruct(_map, &newObj)
	if err != nil {
		return
	}
	fmt.Print("$v", newObj)
}

func TestStructToMapStr(t *testing.T) {
	obj := &struct {
		Name    string
		Age     int
		Address struct {
			City string
			Code string `json:"code1"`
		}
	}{
		Name: "tom",
		Age:  18,
		Address: struct {
			City string
			Code string `json:"code1"`
		}{
			City: "shanghai",
		},
	}
	_map := StructToMapStr(*obj)
	println(_map)
}
