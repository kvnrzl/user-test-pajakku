package main

import (
	"math/rand"
)

func generateRandomOneToHundred() int {
	randomNumber := rand.Intn(100) + 1
	return randomNumber
}

func grade(score int) string {
	if score <= 60 {
		return "Kurang"
	} else if score > 60 && score <= 70 {
		return "Cukup"
	} else if score > 70 && score <= 80 {
		return "Baik"
	} else {
		return "Luar Biasa"
	}
}
