package main

import "fmt"

func main() {
	// TEST NO.1
	fmt.Println("=============TEST NO.1================")
	fmt.Println(weirdNotWeird(1))
	fmt.Println(weirdNotWeird(98))

	// TEST NO.2
	fmt.Println("=============TEST NO.2================")
	fmt.Println(grade(generateRandomOneToHundred()))

	// TEST NO.4
	fmt.Println("=============TEST NO.4================")
	inputSoal4 := "Selamat Pagi Semuanya"
	output := generateOutputFormat(inputSoal4)
	fmt.Println("Format Satu :", output.FormatSatu)
	fmt.Println("Format Dua :", output.FormatDua)
	fmt.Println("Format Tiga :", output.FormatTiga)
	fmt.Println("Format Empat :", output.FormatEmpat)

	// TEST NO.5 - Output Struct/Object with Array String
	fmt.Println("=============TEST NO.5================")
	inputSoal5 := 1225441
	outputSoal5 := generateOutputArrayString(inputSoal5)
	for _, value := range outputSoal5.Values {
		fmt.Println(value)
	}

	// TEST NO.5 - Print 1 - 100 and Mitra Pajakku
	fmt.Println("=============TEST NO.5================")
	cetakMitraPajakku()

	// TEST NO.5 - Palindrome
	fmt.Println("=============TEST NO.5================")
	fmt.Println(isPalindrome("MADAM"))
	fmt.Println(isPalindrome("KATAK"))
	fmt.Println(isPalindrome("AKU"))

	// TEST NO.5 - Remove duplicates and sorting
	fmt.Println("=============TEST NO.5================")
	listOfInt := []int{20, 50, 60, 50, 70, 80, 81, 20, 21}
	sortedListOfInt := removeDuplicatesAndSorting(listOfInt)
	fmt.Println(sortedListOfInt)
	fmt.Println("=====================================")
}
