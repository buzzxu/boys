package tools

import (
	"crypto/sha1"
	"encoding/hex"
	"hash/fnv"
)

func Hash32(str string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(str))
	return h.Sum32()
}
func HashSHA1(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
