package randoms

import (
	"math"
	"math/rand"
	"time"
)

var (
	rnd *rand.Rand
	ch  chan int64
)

func init() {
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	ch = make(chan int64, 1000)
	go randomBase(ch)
}

// Seed inits the random seed using the current UTC Unix Nano time.
func Seed() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// NextFloat64 returns the next float64.
func NextFloat64() float64 {
	return rand.Float64()
}

// NextFloat64Bounded returns the next float64 bounded by start and end.
func NextFloat64Bounded(start float64, end float64) float64 {
	return rand.Float64()*(end-start) + start
}

// NextFloat32 returns the next float32.
func NextFloat32() float32 {
	return rand.Float32()
}

// NextFloat32Bounded returns the next float32 bounded by start and end.
func NextFloat32Bounded(start float32, end float32) float32 {
	return rand.Float32()*(end-start) + start
}

// NextInt returns the next int.
func NextInt() int {
	return rand.Int()
}

// NextIntBounded returns the next int bounded by start and end.
func NextIntBounded(start int, end int) int {
	return start + rand.Intn(end)
}

// NextIntUpperBounded returns the next int bounded by a maximum.
func NextIntUpperBounded(end int) int {
	return rand.Intn(end)
}

// NextBytes creates an array of random bytes.
func NextBytes(count int) []byte {
	a := make([]byte, count)
	for i := range a {
		a[i] = (byte)(NextInt())
	}
	return a
}

func randomBase(c chan int64) {
	for {
		c <- rnd.Int63()
	}
}

//Uniqid 随机数，length是需要返回的长度，只支持10~19位
func Uniqid(length int) int64 {
	if length < 10 || length > 19 {
		return 0
	}
	prefix := (time.Now().UnixNano() / 100000000) & 0x3fffffff
	cut := int64(math.Pow10(length - 9))
	suffix := <-ch % cut
	return prefix*cut + suffix
}
