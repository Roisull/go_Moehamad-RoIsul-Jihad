package main

import (
	"fmt"
)

// Fungsi untuk mengkonversi angka normal ke angka Romawi
func intToRoman(num int) string {
	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	romanStr := ""
	for i := 0; i < len(vals); i++ {
		for num >= vals[i] {
			romanStr += romans[i]
			num -= vals[i]
		}
	}

	return romanStr
}

func main() {
	// Mengkonversi angka-angka yang diberikan ke angka Romawi
	fmt.Println(intToRoman(4))    // IV
	fmt.Println(intToRoman(9))    // IX
	fmt.Println(intToRoman(23))   // XXIII
	fmt.Println(intToRoman(2021)) // MMXXI
	fmt.Println(intToRoman(1646)) // MDCXLVI
}