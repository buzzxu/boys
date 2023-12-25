package aksk

import (
	"crypto"
	"fmt"
	"testing"
	"time"
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

	sign, err := SHA256.Signature(appkey, secret, content, time.Now().Unix())
	if err != nil {
		fmt.Errorf("签名错误: %s", sign)
	}
	fmt.Println("签名=", sign)
	if Verify(content, sign, secret, crypto.SHA256) != nil {
		fmt.Errorf("签名错误: %s", sign)
	}
	fmt.Println("签名OK")
}
func TestSha1(t *testing.T) {
	appkey := "xwjd"
	secret := "XWJD010230441222212312313V"
	content := "c0qX08"
	timestamp := int64(1703514656)
	sign := SHA1(appkey, secret, content, timestamp)
	fmt.Println("签名=", sign)
}
