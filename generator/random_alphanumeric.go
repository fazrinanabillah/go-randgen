package generator

import (
	"math/rand"
	"time"
)

func GenerateRandomAlphanumeric(n int) string {
	var randomizer = rand.New(rand.NewSource(time.Now().UnixNano()))
	var rn = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	str := make([]rune, n)

	for i := range str {
		str[i] = rn[randomizer.Intn(len(rn) - 1)]
	}

	return string(str)
}
