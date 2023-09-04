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
	n := 2 // Ganti nilai n sesuai yang Anda inginkan
	memo = make(map[int]int) // Inisialisasi memo
	result := fibonacci(n)
	fmt.Printf("Fibonacci ke-%d adalah %d\n", n, result)
}