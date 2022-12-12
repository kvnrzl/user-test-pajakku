package main

import "fmt"

func main() {
	// for i := -1; i <= 105; i++ {
	// 	// fmt.Println(i, weirdNotWeird(i))
	// 	x := generateRandomOneToHundred()
	// 	fmt.Println(x, grade(x))
	// }
	// input := "Selamat Pagi Semuanya"
	// output := generateOutputFormat(input)

	// fmt.Println("Format Satu :", output.FormatSatu)
	// fmt.Println("Format Dua :", output.FormatDua)
	// fmt.Println("Format Tiga :", output.FormatTiga)
	// fmt.Println("Format Empat :", output.FormatEmpat)

	input := 1225441
	output := generateOutputArrayString(input)

	fmt.Println("OutputArrayString:")
	for _, value := range output.Values {
		fmt.Println(value)
	}

	// cetakMitraPajakku()

	// fmt.Println(isValid("katak"))
	// fmt.Println(isValid("aku"))

	// array := []int{20, 50, 60, 50, 70, 80, 81, 20, 21}
	// sortedArray := sorting(array)

	// // menampilkan data yang telah diurutkan
	// for _, value := range sortedArray {
	// 	fmt.Println(value)
	// }
}
