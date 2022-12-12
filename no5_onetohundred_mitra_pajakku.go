package main

import "fmt"

func cetakMitraPajakku() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Mitra Pajakku")
		} else if i%3 == 0 {
			fmt.Println("Mitra")
		} else if i%5 == 0 {
			fmt.Println("Pajakku")
		} else {
			fmt.Println(i)
		}
	}
}
