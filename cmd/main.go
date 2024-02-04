package main

import (
	"fmt"

	"github.com/ricsjs/password-generator/password"
	"github.com/ricsjs/password-generator/utils"
)

func main() {
	fmt.Println("Welcome to Password Generator")

	length := 32
	includeDigits := utils.AskForYesOrNo("Do you want to include numbers in your password?")
	includeUppercase := utils.AskForYesOrNo("Do you want to include capital letters in your password?")
	includeSpecialChars := utils.AskForYesOrNo("Do you want to include special chars in your password?")

	newPass := password.GeneratePassword(length, includeDigits, includeUppercase, includeSpecialChars)

	fmt.Println(newPass)
}
