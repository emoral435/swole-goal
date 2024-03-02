package util

import (
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// RandomInt returns a random integer between the minimum value and the maximum value
func RandomInt(min int64, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
// @param n the length of the string
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomEmail generates a random email of length n + "@gmail.com"
// @param n the length of the string before the email address suffix
func RandomEmail(n int) string {
	return RandomString(n) + "@gmail.com"
}

// RandomEmail generates a random email of length n + "@gmail.com"
// @param n the length of the string before the email address suffix
func RandomPassword(n int) string {
	return RandomString(n)
}

// RandomEmail generates a random email of length n + "@gmail.com"
// @param n the length of the string before the email address suffix
func RandomUsername(n int) string {
	return RandomString(n)
}
