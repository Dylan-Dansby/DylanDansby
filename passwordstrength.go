package main

import (
	"fmt"
	"strings"
	"unicode"
)

var MIN_LENGTH int = 8 //minimum password length
var PASS_SCORE int = 100 //passing score

// Password Getter
func getPassword() string {
	var password string
	fmt.Println("Please enter your Password. Must be at least 8 characters, contain an uppercase, lowercase, digit, and special character.")
	fmt.Println("Please do not use common patterns such as `123` `abc` or `password`")
	fmt.Scan(&password)
	return password
}

// Uppercase Checker Function
func hasUppercase(password string) bool {
	for _, i := range password {
		if unicode.IsUpper(i) {
			return true
		}
	}
	return false
}

// Lowercase checker function
func hasLowercase(password string) bool {
	for _, i := range password {
		if unicode.IsLower(i) {
			return true
		}
	}
	return false
}

// Digit checker function
func hasDigit(password string) bool {
	for _, i := range password {
		if unicode.IsDigit(i) {
			return true
		}
	}
	return false
}

// Special Character Checker
func hasSpecialChar(password string) bool {
	for _, i := range password {
		if !unicode.IsDigit(i) && !unicode.IsLetter(i) {
			return true
		}
	}
	return false
}

// Password Length Checker
func isLongenough(password string) bool {
	if len(password) < MIN_LENGTH {
		return false
	}
	return true
}

// Checks password for common password patterns
func containsCommonPatterns(password string) bool {
	var arrCommonPhrases [3]string //common phrases array
	arrCommonPhrases[0] = "123"
	arrCommonPhrases[1] = "password"
	arrCommonPhrases[2] = "abc"

	for i := 0; i < len(arrCommonPhrases); i++ {
		if strings.Contains(strings.ToUpper(password), strings.ToUpper(arrCommonPhrases[i])) {
			return true
		}
	}
	return false
}

//password strength calculator
func calculateStrength(password string) int {

	if !isLongenough(password) {
		fmt.Println("Password is not long enough. Try again.")
		return 0
	}

	if containsCommonPatterns(password) {
		fmt.Println("Password contains a common pattern such as `abc`, `password`, or `123`. Try again")
		return 0
	}

	if !hasDigit(password) {
		fmt.Println("Password doesn`t contain a digit. Try again")
		return 0
	}

	if !hasUppercase(password) {
		fmt.Println("Password does not contain Uppercase Value. Try again.")
		return 0
	}

	if !hasLowercase(password) {
		fmt.Println("Password does not contain Lowercase Value. Try again.")
		return 0
	}

	if !hasSpecialChar(password) {
		fmt.Println("Password does not contain a spceial character. Try again")
		return 0
	}

	return PASS_SCORE
}

func main() {

	var Endscore int = 0

	for Endscore < PASS_SCORE {
		var password string = getPassword()
		Endscore = calculateStrength(password)
	}

}
