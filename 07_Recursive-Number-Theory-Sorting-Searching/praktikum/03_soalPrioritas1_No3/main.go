package main

import "fmt"

// Fungsi untuk memeriksa apakah sebuah bilangan adalah bilangan prima
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	i := 5
	for i*i <= n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}

// Fungsi untuk mendapatkan bilangan prima ke-n
func primeX(n int) int {
	count := 0
	num := 1

	for count < n {
		num++
		if isPrime(num) {
			count++
		}
	}

	return num
}

func main() {
	fmt.Println(primeX(1)) // 2
	fmt.Println(primeX(5)) // 11
	fmt.Println(primeX(8)) // 19
	fmt.Println(primeX(9)) // 23
	fmt.Println(primeX(10)) // 29
}
