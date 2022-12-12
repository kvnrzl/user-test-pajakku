package main

import (
	"strconv"
	"strings"
)

type OutputArrayString struct {
	Values []string
}

func generateOutputArrayString(input int) *OutputArrayString {
	inputStr := strconv.Itoa(input)

	outputValues := make([]string, 0)
	for i := 0; i < len(inputStr); i++ {
		a := string(inputStr[i]) + strings.Repeat("0", len(inputStr)-i-1)
		outputValues = append(outputValues, a)
	}

	return &OutputArrayString{
		Values: outputValues,
	}
}

// 5. Buatlah sebuah fungsi dengan contoh input dan output seperti berikut :
// •	Input : 1225441
// •	Output :
// 1000000
// 200000
// 20000
// 5000
// 400
// 40
// 1
