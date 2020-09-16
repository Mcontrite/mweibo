package utils

import (
	"crypto/rand"
	"fmt"
	r "math/rand"
	"strings"
	"time"
)

var alphaNum = []byte(`0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`)

func CreateRandomBytes(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNum
	}
	bytes := make([]byte, n)
	randBy := false
	if num, err := rand.Read(bytes); num != n || err != nil {
		r.Seed(time.Now().UnixNano())
		randBy = true
	}
	for k, v := range bytes {
		if randBy {
			bytes[k] = alphabets[r.Intn(len(alphabets))]
		} else {
			bytes[k] = alphabets[v%byte(len(alphabets))]
		}
	}
	return bytes
}

func CreateRandomInt(min, max int) int {
	if min >= max || min < 0 || max == 0 {
		return max
	}
	return r.Intn(max-min) + min
}

func GenRandCode(length int) string {
	arr := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r.Seed(time.Now().UnixNano())
	var builder strings.Builder
	for i := 0; i < length; i++ {
		fmt.Fprintf(&builder, "%d", arr[r.Intn(len(arr))])
	}
	return builder.String()
}
