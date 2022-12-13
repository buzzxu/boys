package aksk

import (
	"fmt"
	"testing"
)

func TestGenApp(t *testing.T) {
	appKey := GenAppKey()
	println(appKey)
	println(GenAppSecretBy(appKey))
	println(GenAppSecret())
}

func TestSignature(t *testing.T) {
	appkey := GenAppKey()
	fmt.Println("appkey=", appkey)
	secret := GenAppSecretBy(appkey)
	fmt.Println("secret=", secret)

	content := "xxxxx"

	sign, err := SHA256.Signature(content, secret)
	if err != nil {
		fmt.Errorf("签名错误: %s", sign)
	}
	fmt.Println("签名=", sign)
	if SHA256.Verify(content, sign, secret) != nil {
		fmt.Errorf("签名错误: %s", sign)
	}
	fmt.Println("签名OK")
}
