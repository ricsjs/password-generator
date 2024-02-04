package utils

import (
	"fmt"
	"strings"
)

func AskForYesOrNo(prompt string) bool {
	var input string

	for {
		fmt.Print(prompt)
		fmt.Scanln(&input)

		input = strings.ToLower(strings.TrimSpace(input))

		if input == "y" || input == "yes" {
			return true
		} else if input == "n" || input == "no" {
			return false
		}
		fmt.Println("Please answer with 'y' (yes) or 'n' (no).")
	}
}
