package radom

import (
	crand "crypto/rand"
	"fmt"
	"math/big"
	"math/rand/v2"
	"time"
	"unsafe"
)

const (
	letterBytes       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterNumberBytes = "abcdefghijklmnopqrstuvwxyz0123456789"
	letterIdxBits     = 6                    // 6 bits to represent a letter index
	letterIdxMask     = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax      = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// Alphanumeric 生成随机字母 + 数字字符串
func Alphanumeric(count int) string {
	return radom(count, letterNumberBytes)
}

// Alphabetic 生成随机字母字符串
func Alphabetic(count int) string {
	return radom(count, letterBytes)
}

// String 生成随机字符串
func String(n int) string {
	b := make([]byte, n)
	for i := range b {
		num, _ := crand.Int(crand.Reader, big.NewInt(int64(len(letterBytes))))
		b[i] = letterBytes[num.Int64()]
	}
	return string(b)
}

// TimeID 生成时间戳 + 随机字符串
func TimeId() string {
	now := time.Now()
	// 获取 Unix 时间戳和纳秒部分，生成基础字符串
	baseStr := fmt.Sprintf("%d%d", now.Unix(), now.Nanosecond())

	// 生成随机字母字符串
	randomStr := String(16)

	// 拼接并截取前 16 位
	id := (baseStr + randomStr)[:16]

	return id
}

// radom 生成随机字符串
func radom(count int, str string) string {
	b := make([]byte, count)
	for i, cache, remain := count-1, rand.Int64(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int64(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(str) {
			b[i] = str[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}
