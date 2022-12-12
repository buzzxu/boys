package strs

import (
	"strconv"
	"testing"
)

func TestHash32(t *testing.T) {
	println(strconv.Itoa(int(Hash32("32323"))) + "")
	println(HashSHA1("32323"))
}
