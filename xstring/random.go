package xstring

import (
	"math/rand"
	"strings"
	"time"
)

// const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
// const (
// // letterIdxBits = 6                    // 6 bits to represent a letter index
// // letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
// // letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
// )

// var src = rand.NewSource(time.Now().UnixNano())

//RandomNumber  - generates a random number based on the length set
func RandomNumber(n int) string {
	rand.Seed(time.Now().UnixNano())
	var builder strings.Builder
	builder.Grow(n)
	for i := 0; i < n; i++ {
		builder.WriteByte(byte(rand.Intn(10)))
	}
	return builder.String()
}
