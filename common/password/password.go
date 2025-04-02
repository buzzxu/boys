package password

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
)

// MD5Hash returns the MD5 hash of the input string.
func MD5Hash(data string) string {
	hasher := md5.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}

// SHA1Hash returns the SHA-1 hash of the input string.
func SHA1Hash(data string) string {
	hasher := sha1.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}

// SHA256Hash returns the SHA-256 hash of the input string.
func SHA256Hash(data string) string {
	hasher := sha256.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}

// SHA384Hash returns the SHA-384 hash of the input string.
func SHA384Hash(data string) string {
	hasher := sha512.New384()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}

// SHA512Hash returns the SHA-512 hash of the input string.
func SHA512Hash(data string) string {
	hasher := sha512.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}

// HashPassword hashes the given password with bcrypt.
func HashPassword(password string, cost int) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPassword verifies if the provided password matches the hashed password.
func CheckPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
