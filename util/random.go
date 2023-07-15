package util

import (
	"fmt"
	"math/rand"
	"time"
)

const charSet = "abcdefghijklmnopqrstuvwxyz"


func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max - min + 1)
}


func RandomString(n int) string {
	// Create a byte slice of length n
	randomBytes := make([]byte, n)

	// Fill the byte slice with random characters from the character set
	for i := 0; i < n; i++ {
		randomBytes[i] = charSet[rand.Intn(len(charSet))]
	}

	// Convert the byte slice to a string and return it
	return string(randomBytes)
}


func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0,10000)
}

func RandomCurrency() string{
	currencies := []string {MAD,USD,CAD,EUR}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}