package main

import (
	"fmt"
)

func munculSekali(angka string) []int {
	counts := make(map[rune]int)
	var result []int

	for _, digit := range angka {
		counts[digit]++
	}

	for _, digit := range angka {
		if counts[digit] == 1 {
			result = append(result, int(digit-'0'))
			counts[digit] = -1 // Menandakan angka ini sudah ditambahkan ke hasil
		}
	}

	return result
}

func soalSlice() {
	fmt.Println(munculSekali("1234123"))    // [4]
	fmt.Println(munculSekali("76523752"))   // [6 3]
	fmt.Println(munculSekali("12345"))      // [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504"))    // [8 7 2 5 4]
}