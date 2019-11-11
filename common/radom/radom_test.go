package radom

import "testing"

func TestLength(t *testing.T) {
	println(len(letterNumberBytes))
}

func TestAlphanumeric(t *testing.T) {

	println(Alphanumeric(4))
}

func TestAlphabetic(t *testing.T) {
	println(Alphabetic(32))
}
