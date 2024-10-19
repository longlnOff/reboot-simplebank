package utils

import (
	"math/rand"
	"strings"
	"time"
)



var alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// Generate random string
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphanum)
	for i := 0; i < n; i++ {
		c := alphanum[RandomInt(0, int64(k))]
		sb.WriteByte(c)
	}
	return sb.String()
}

// Generate random number
func RandomInt(min int64, max int64) int64 {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	return rng.Int63n(max-min) + min
}

func RandomBalance() int64 {
	return RandomInt(0, 1000)
}

func RandomEmail() string {
	return RandomString(10) + "@gmail.com"
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}


// Generate random currency