package cryptos

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

// MD5 生成32位MD5
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}
func MD5B(bytes []byte) string {
	h := md5.New()
	h.Write(bytes)
	return hex.EncodeToString(h.Sum(nil))
}

func Sha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

func HamSha1(data string, key []byte) string {
	hmac := hmac.New(sha1.New, key)
	hmac.Write([]byte(data))
	return base64.StdEncoding.EncodeToString(hmac.Sum(nil))
}
