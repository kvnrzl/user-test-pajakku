package main

import "strings"

func isPalindrome(text string) bool {
	lowercaseText := strings.ToLower(text)

	reversedText := ""
	for i := len(lowercaseText) - 1; i >= 0; i-- {
		reversedText += string(lowercaseText[i])
	}

	return lowercaseText == reversedText
}
