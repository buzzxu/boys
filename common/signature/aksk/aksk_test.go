package aksk

import "testing"

func TestGenApp(t *testing.T) {
	appKey := GenAppKey()
	println(appKey)
	println(GenAppSecretBy(appKey))
	println(GenAppSecret())
}
