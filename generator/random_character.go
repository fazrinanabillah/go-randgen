package generator

import (
	"math/rand"
	"time"
)

func GenerateRandomCharacter(n int) string {
	var randomizer = rand.New(rand.NewSource(time.Now().UnixNano()))
	var characters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+-=[]{}|;:,.<>/?~")

	str := make([]rune, n)
	for i := range str {
		str[i] = characters[randomizer.Intn(len(characters) - 1)]
	}

	return string(str)
}