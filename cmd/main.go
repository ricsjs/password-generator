package main

import (
	"fmt"
	"strconv"

	"github.com/ricsjs/password-generator/password"
	"github.com/ricsjs/password-generator/utils"
)

func main() {
	fmt.Println("Welcome to Password Generator")

	length, err := AskForPassLength()
	if err != nil {
		fmt.Println("Error getting password length.", err)
		return
	}

	includeDigits := utils.AskForYesOrNo("Do you want to include numbers in your password? ")
	includeUppercase := utils.AskForYesOrNo("Do you want to include capital letters in your password? ")
	includeSpecialChars := utils.AskForYesOrNo("Do you want to include special chars in your password? ")

	newPass := password.GeneratePassword(length, includeDigits, includeUppercase, includeSpecialChars)

	fmt.Println("\nGenerated password:", newPass)
}

func AskForPassLength() (int, error) {
	fmt.Println("Enter the length of the desired password: ")
	var input string
	fmt.Scanln(&input)

	// Converte a entrada para um número inteiro
	length, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}

	return length, nil
}
