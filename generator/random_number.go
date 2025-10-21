package generator

import (
	"math/rand"
	"time"
)

func GenerateRandomNumber(n int) string {
	var randomizer = rand.New(rand.NewSource(time.Now().UnixNano()))
	var rn = []rune("0123456789")
	str := make([]rune, n)

	for i := range str {
		str[i] = rn[randomizer.Intn(len(rn)-1)]
	}

	return string(str)

}