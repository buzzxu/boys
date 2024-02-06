package radom

import "testing"

func TestLength(t *testing.T) {
	println(len(letterNumberBytes))
}

func TestAlphanumeric(t *testing.T) {

	println(Alphanumeric(16))
}

func TestAlphabetic(t *testing.T) {
	println(Alphabetic(32))
}

func TestTimeID(t *testing.T) {
	println(TimeId())
}
