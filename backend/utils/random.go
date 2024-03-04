// Package util handles an assortment of process's within the backend process for the swole-goal module.
package util

import (
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// RandomInt: returns a random integer between the minimum value and the maximum value
//
// returns: a random integer between the minimum value and the maximum
func RandomInt(min int64, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString: generates a random string of length n
//
// returns: string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	// for i in range n, write a randome character
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	// return the string that we just built
	return sb.String()
}

// RandomEmail: generates a random email of length n + "@gmail.com"
//
// returns: string of length n + "@gmail.com"
func RandomEmail(n int) string {
	return RandomString(n) + "@gmail.com"
}

// RandomEmail: generates a random email of length n + "@gmail.com"
//
// returns: string of length n to mimic a password
func RandomPassword(n int) string {
	return RandomString(n)
}

// RandomEmail: generates a random username
//
// returns: string of length n to mimic a username
func RandomUsername(n int) string {
	return RandomString(n)
}
