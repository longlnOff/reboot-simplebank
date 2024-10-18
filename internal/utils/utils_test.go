package utils

import "testing"


func TestRandomString(t *testing.T) {
	for i := 0; i < 100; i++ {
		RandomString(10)
	}
}