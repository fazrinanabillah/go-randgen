package main

import (
	"fmt"

	"github.com/fazrinanabillah/go-randgen/generator"
)

func main() {
	// scan input from user string
	fmt.Println("Enter the length of the random string:")
	var n int
	fmt.Scanln(&n)

	randomString := generator.GenerateRandomString(n)
	fmt.Println("Random String:", randomString)

	randomAlphanumeric := generator.GenerateRandomAlphanumeric(n)
	fmt.Println("Random Alphanumeric String:", randomAlphanumeric)
}
