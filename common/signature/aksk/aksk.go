package aksk

import (
	"crypto"
	"crypto/hmac"
	"encoding/hex"
	"errors"
	"github.com/buzzxu/boys/common/cryptos"
	"github.com/buzzxu/boys/common/radom"
	"strings"
)

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
func Signature(content, secret string) (string, error) {
	if !crypto.SHA3_256.Available() {
		return "", errors.New("SHA3_256加密算法不可用")
	}
	hasher := hmac.New(crypto.SHA3_256.New, []byte(secret))
	hasher.Write([]byte(content))
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

// Verify 验签
func Verify(content, sign, secret string) error {
	sig, err := hex.DecodeString(sign)
	if err != nil {
		return err
	}

	if !crypto.SHA3_256.Available() {
		return errors.New("SHA3_256加密算法不可用")
	}
	hasher := hmac.New(crypto.SHA3_256.New, []byte(secret))
	hasher.Write([]byte(content))
	if !hmac.Equal(sig, hasher.Sum(nil)) {
		return errors.New("验证签名失败")
	}
	return nil
}
