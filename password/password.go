package password

import (
	"math/rand"
	"time"
)

const (
	digits         = "0123456789"
	uppercaseChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialChars   = "!@#$%^&*()-_=+[]{}|;:'\",.<>/?"
)

// Ensures that the generated passwords are different each time you run
func init() {
	rand.Seed(time.Now().UnixNano())
}

func GeneratePassword(length int, includeDigits, includeUppercase, includeSpecialChars bool) string {
	allChars := "abcdefghijklmnopqrstuvwxyz"

	if includeDigits {
		allChars += digits
	}

	if includeUppercase {
		allChars += uppercaseChars
	}

	if includeSpecialChars {
		allChars += specialChars
	}

	password := make([]byte, length)
	for i := range password {
		password[i] = allChars[rand.Intn(len(allChars))]
	}

	return string(password)
}
