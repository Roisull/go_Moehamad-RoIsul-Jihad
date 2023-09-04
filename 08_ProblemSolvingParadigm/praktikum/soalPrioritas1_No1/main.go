package main

import (
	"fmt"
)

func countBits(num int) int {
	count := 0
	for num > 0 {
		count += num & 1
		num >>= 1
	}
	return count
}

func generateBits(n int) []int {
	result := make([]int, n+1)
	for i := 0; i <= n; i++ {
		result[i] = countBits(i)
	}
	return result
}

func main() {
	n := 3
	ans := generateBits(n)
	fmt.Println(ans)
}