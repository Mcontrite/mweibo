package utils

import (
	"crypto/rand"
	r "math/rand"
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
