package main

func weirdNotWeird(a int) string {
	if a < 1 || a > 100 {
		return "Angka harus diantara 1-100"
	}
	if a%2 == 1 {
		return "Weird"
	} else if a >= 2 && a <= 5 {
		return "Not Weird"
	} else if a >= 6 && a <= 20 {
		return "Weird"
	} else if a > 20 && a%2 == 0 {
		return "Not Weird"
	} else {
		return "Weird"
	}
}
