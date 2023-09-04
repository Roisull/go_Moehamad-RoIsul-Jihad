package main

import (
	"fmt"
)

var memo map[int]int

func fibonacci(n int) int {
	// Jika n <= 1, maka Fibonacci(n) = n (untuk kondisi awal, biar ada batas minimalnya)
	if n <= 1 {
		return n
	}

	// menghitung fibonaci dengan rekursi
	result := fibonacci(n-1) + fibonacci(n-2)

	// Simpan hasil perhitungan dalam memo
	memo[n] = result

	return result
}

func main() {
	memo = make(map[int]int) // Inisialisasi memo
	fmt.Println(fibonacci(0))  // 0
	fmt.Println(fibonacci(2))  // 1
	fmt.Println(fibonacci(9))  // 34
	fmt.Println(fibonacci(10)) // 55
	fmt.Println(fibonacci(12)) // 144
}