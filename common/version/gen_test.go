package version

import "testing"

func TestGenUpgrade(t *testing.T) {
	println(GenUpgrade("v0.0.10", VERSION_UPGRADE_TYPE_MINOR))
}

func TestGenNext(t *testing.T) {
	println(GenNext("v0.0.10"))
}
