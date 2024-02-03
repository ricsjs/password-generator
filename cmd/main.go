package main

import (
	"fmt"

	"github.com/ricsjs/password-generator/password"
)

func main() {
	fmt.Println("Welcome to Password Generator")

	length := 32
	includeDigits := true
	includeUppercase := true
	includeSpecialChars := true

	newPass := password.GeneratePassword(length, includeDigits, includeUppercase, includeSpecialChars)

	fmt.Println(newPass)
}
