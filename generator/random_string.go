package generator

import (
	"time"

	"math/rand"
)

func GenerateRandomString(n int) string {
	var randomizer = rand.New(rand.NewSource(time.Now().UnixNano()))
	var rn = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

	str := make([]rune, n)
	for i := range str {
		str[i] = rn[randomizer.Intn(len(rn)-1)]
	}

	return string(str)
}
