package strs

import (
	"crypto/sha1"
	"encoding/hex"
	"hash/fnv"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
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

func IsLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch >= 0x80 && unicode.IsLetter(ch)
}

func TruncateString(str string, limit int) string {
	if len(str) < limit {
		return str
	}
	return str[:limit]
}

func EllipsisString(str string, length int) string {
	if len(str) < length {
		return str
	}
	return str[:length-3] + "..."
}

// IsAllLowerCase checks if the string contains only lowercase characters.
func IsAllLowerCase(str string) bool {
	if IsEmpty(str) {
		return false
	}
	for _, c := range str {
		if !unicode.IsLower(c) {
			return false
		}
	}
	return true
}

// IsAllUpperCase checks if the string contains only uppercase characters.
func IsAllUpperCase(str string) bool {
	if IsEmpty(str) {
		return false
	}
	for _, c := range str {
		if !unicode.IsUpper(c) {
			return false
		}
	}
	return true
}

// IsAlpha checks if the string contains only Unicode letters.
func IsAlpha(str string) bool {
	if IsEmpty(str) {
		return false
	}
	for _, c := range str {
		if !unicode.IsLetter(c) {
			return false
		}
	}
	return true
}

// IsAlphanumeric checks if the string contains only Unicode letters and digits.
func IsAlphanumeric(str string) bool {
	if IsEmpty(str) {
		return false
	}
	for _, c := range str {
		if !unicode.IsLetter(c) && !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

// IsAlphaSpace checks if the string contains only Unicode letters and spaces.
func IsAlphaSpace(str string) bool {
	if IsEmpty(str) {
		return false
	}
	for _, c := range str {
		if !unicode.IsLetter(c) && !unicode.IsSpace(c) {
			return false
		}
	}
	return true
}

// IsAlphanumericSpace checks if the string contains only Unicode letters, digits and spaces.
func IsAlphanumericSpace(str string) bool {
	if IsEmpty(str) {
		return false
	}
	for _, c := range str {
		if !unicode.IsLetter(c) && !unicode.IsDigit(c) && !unicode.IsSpace(c) {
			return false
		}
	}
	return true
}

// IsEmpty checks if a string is empty.
func IsEmpty(s string) bool {
	if s == "" {
		return true
	}
	return false
}

// IsNotEmpty checks if a string is not empty.
func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}

// IsAnyEmpty checks if any one of the given strings are empty.
func IsAnyEmpty(strings ...string) bool {
	for _, s := range strings {
		if IsEmpty(s) {
			return true
		}
	}
	return false
}

// IsNoneEmpty checks if none of the strings are empty.
func IsNoneEmpty(strings ...string) bool {
	for _, s := range strings {
		if IsEmpty(s) {
			return false
		}
	}
	return true
}

// IsBlank checks if a string is whitespace or empty
func IsBlank(s string) bool {
	if s == "" {
		return true
	}
	if regexp.MustCompile(`^\s+$`).MatchString(s) {
		return true
	}
	return false
}

// IsNotBlank checks if a string is not empty or containing only whitespaces.
func IsNotBlank(s string) bool {
	return !IsBlank(s)
}

// IsAnyBlank checks if any one of the strings are empty or containing only whitespaces.
func IsAnyBlank(strings ...string) bool {
	for _, s := range strings {
		if IsBlank(s) {
			return true
		}
	}
	return false
}

// IsNoneBlank checks if none of the strings are empty or containing only whitespaces.
func IsNoneBlank(strings ...string) bool {
	for _, s := range strings {
		if IsBlank(s) {
			return false
		}
	}
	return true
}

// IsNumeric checks if the string contains only digits.
func IsNumeric(str string) bool {
	for _, c := range str {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

// IsNumericSpace checks if the string contains only digits and whitespace.
func IsNumericSpace(str string) bool {
	for _, c := range str {
		if !unicode.IsDigit(c) && !unicode.IsSpace(c) {
			return false
		}
	}
	return true
}

// IsWhitespace checks if the string contains only whitespace.
func IsWhitespace(str string) bool {
	for _, c := range str {
		if !unicode.IsSpace(c) {
			return false
		}
	}
	return true
}

// Join joins an array of strings into a string where each item of the array is separated with a separator.
func Join(a []string, sep string) string {
	return strings.Join(a, sep)
}

// JoinBool is the same as Join but joining boolean.
func JoinBool(a []bool, sep string) string {
	strs := make([]string, len(a), len(a))
	for idx, i := range a {
		strs[idx] = strconv.FormatBool(i)
	}
	return Join(strs, sep)
}

// JoinFloat64 is the same as Join but joining float64.
// The default format given to strconv.FormatFloat is 'G' and bitSize is 32.
func JoinFloat64(a []float64, sep string) string {
	return JoinFloat64WithFormatAndPrecision(a, 'G', 32, sep)
}

// JoinFloat64WithFormatAndPrecision is the same as Join but joining float64 with a custom precision (bitSize) and format.
func JoinFloat64WithFormatAndPrecision(a []float64, fmt byte, precision int, sep string) string {
	strs := make([]string, len(a), len(a))
	for idx, i := range a {
		strs[idx] = strconv.FormatFloat(i, fmt, -1, precision)
	}
	return Join(strs, sep)
}

// JoinInt is the same as Join but joining integers.
func JoinInt(a []int, sep string) string {
	strs := make([]string, len(a), len(a))
	for idx, i := range a {
		strs[idx] = strconv.Itoa(i)
	}
	return Join(strs, sep)
}

// JoinInt64 is the same as Join but joining int64.
func JoinInt64(a []int64, sep string) string {
	strs := make([]string, len(a), len(a))
	for idx, i := range a {
		strs[idx] = strconv.FormatInt(i, 10)
	}
	return Join(strs, sep)
}

// JoinUint64 is the same as Join but joining uint64.
func JoinUint64(ints []uint64, sep string) string {
	strs := make([]string, len(ints), len(ints))
	for idx, i := range ints {
		strs[idx] = strconv.FormatUint(i, 10)
	}
	return Join(strs, sep)
}

// Left gets the leftmost len characters of a string.
func Left(str string, size int) string {
	if str == "" || size < 0 {
		return ""
	}
	if len(str) <= size {
		return str
	}
	return str[0:size]
}

// Trim removes control characters from both ends of this string.
func Trim(str string) string {
	return strings.Trim(str, " ")
}

// internalStartsWith internal method to check if a string starts with a specified prefix ignoring case or not.
func internalStartsWith(str string, prefix string, ignoreCase bool) bool {
	if str == "" || prefix == "" {
		return str == "" && prefix == ""
	}
	if utf8.RuneCountInString(prefix) > utf8.RuneCountInString(str) {
		return false
	}
	if ignoreCase {
		return strings.HasPrefix(strings.ToLower(str), strings.ToLower(prefix))
	}
	return strings.HasPrefix(str, prefix)
}

// StartsWith check if a string starts with a specified prefix.
func StartsWith(str string, prefix string) bool {
	return internalStartsWith(str, prefix, false)
}

// StartsWithIgnoreCase case insensitive check if a string starts with a specified prefix.
func StartsWithIgnoreCase(str string, prefix string) bool {
	return internalStartsWith(str, prefix, true)
}

// StartsWithAny check if a string starts with any of an array of specified strings.
func StartsWithAny(str string, prefixes ...string) bool {
	for _, prefix := range prefixes {
		if internalStartsWith(str, (string)(prefix), false) {
			return true
		}
	}
	return false
}

// StartsWithAnyIgnoreCase check if a string starts with any of an array of specified strings (ignoring case).
func StartsWithAnyIgnoreCase(str string, prefixes ...string) bool {
	for _, prefix := range prefixes {
		if internalStartsWith(str, (string)(prefix), true) {
			return true
		}
	}
	return false
}

// Internal method to check if a string ends with a specified suffix ignoring case or not.
func internalEndsWith(str string, suffix string, ignoreCase bool) bool {
	if str == "" || suffix == "" {
		return (str == "" && suffix == "")
	}
	if utf8.RuneCountInString(suffix) > utf8.RuneCountInString(str) {
		return false
	}
	if ignoreCase {
		return strings.HasSuffix(strings.ToLower(str), strings.ToLower(suffix))
	}
	return strings.HasSuffix(str, suffix)
}

// EndsWith check if a string ends with a specified suffix.
func EndsWith(str string, suffix string) bool {
	return internalEndsWith(str, suffix, false)
}

// EndsWithIgnoreCase case insensitive check if a string ends with a specified suffix.
func EndsWithIgnoreCase(str string, suffix string) bool {
	return internalEndsWith(str, suffix, true)
}

// EndsWithAny check if a string ends with any of an array of specified strings.
func EndsWithAny(str string, suffixes ...string) bool {
	for _, suffix := range suffixes {
		if internalEndsWith(str, (string)(suffix), false) {
			return true
		}
	}
	return false
}

// EndsWithAnyIgnoreCase check if a string ends with any of an array of specified strings (ignoring case).
func EndsWithAnyIgnoreCase(str string, suffixes ...string) bool {
	for _, suffix := range suffixes {
		if internalEndsWith(str, (string)(suffix), true) {
			return true
		}
	}
	return false
}

// Concat
func Concat(strs ...string) string {
	var b strings.Builder
	for _, v := range strs {
		b.WriteString(v)
	}
	return b.String()
}
