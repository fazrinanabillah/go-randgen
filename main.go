package main

import (
	"fmt"
	"time"

	"math/rand"
)

func main() {
	// scan input from user string
	fmt.Println("Enter the length of the random string:")
	var n int
	fmt.Scanln(&n)

	randomString := GenerateRandomString(n)
	fmt.Println("Random String:", randomString)
}

func GenerateRandomString(n int) string {
	var randomizer = rand.New(rand.NewSource(time.Now().UnixNano()))
	var rn = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

	str := make([]rune, n)
	for i := range str {
		str[i] = rn[randomizer.Intn(len(rn))]
	}

	return string(str)
}
