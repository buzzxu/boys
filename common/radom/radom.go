package radom

import (
	"math/rand"
	"time"
	"unsafe"
)

var (
	randSrc = rand.NewSource(time.Now().UnixNano())
)

const (
	letterBytes       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterNumberBytes = "abcdefghijklmnopqrstuvwxyz0123456789"
	letterIdxBits     = 6                    // 6 bits to represent a letter index
	letterIdxMask     = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax      = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func Numeric(count int) string {
	return ""
}

func Alphanumeric(count int) string {
	return radom(count, letterNumberBytes)
}

func Alphabetic(count int) string {
	return radom(count, letterBytes)
}

func radom(count int, str string) string {
	b := make([]byte, count)
	for i, cache, remain := count-1, randSrc.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = randSrc.Int63(), letterIdxMax
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
