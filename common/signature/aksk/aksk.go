package aksk

import (
	"crypto"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/buzzxu/boys/common/cryptos"
	"github.com/buzzxu/boys/common/radom"
	"strings"
)

type AKSK uint

const (
	SHA256 AKSK = 1 + iota
	SHA3_256
	SHA3_512
	MD5
)

func (aksk AKSK) Hash() crypto.Hash {
	switch aksk {
	case SHA256:
		return crypto.SHA256
	case SHA3_256:
		return crypto.SHA3_256
	case SHA3_512:
		return crypto.SHA3_512
	default:
		return crypto.SHA3_256
	}
}

// GenAppKey 生成appkey
func GenAppKey() string {
	return strings.ToLower(radom.Alphabetic(16))
}

// GenAppSecret 生成默认64位AppSecret
func GenAppSecret() string {
	return radom.Alphanumeric(64)
}

// GenAppSecretBy 根据指定内容生成AppSecret
func GenAppSecretBy(appKey string) string {
	return cryptos.Sha1(appKey)
}

// Signature 生成签名
func (aksk AKSK) Signature(appId, appSecret, nonce string, timestamp int64) (string, error) {
	concatenatedString := appId + appSecret + nonce + fmt.Sprintf("%v", timestamp)
	return Signature(concatenatedString, appSecret, aksk.Hash())
}

// SHA1 sha1签名
func SHA1(appId, appSecret, nonce string, timestamp int64) string {
	message := appId + appSecret + nonce + fmt.Sprintf("%v", timestamp)
	h := sha1.New()
	h.Write([]byte(message))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// Signature 生成签名
func Signature(content, secret string, hash crypto.Hash) (string, error) {
	if !hash.Available() {
		return "", errors.New(hash.String() + "加密算法不可用")
	}
	hasher := hmac.New(hash.New, []byte(secret))
	hasher.Write([]byte(content))
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

// Verify 验签
func Verify(content, sign, secret string, hash crypto.Hash) error {
	sig, err := hex.DecodeString(sign)
	if err != nil {
		return err
	}
	if !hash.Available() {
		return errors.New(hash.String() + "加密算法不可用")
	}
	hasher := hmac.New(hash.New, []byte(secret))
	hasher.Write([]byte(content))
	if !hmac.Equal(sig, hasher.Sum(nil)) {
		return errors.New("验证签名失败")
	}
	return nil
}
