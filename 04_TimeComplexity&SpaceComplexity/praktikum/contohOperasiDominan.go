package main

import "fmt"

func main() {
	fmt.Print(dominant(10))
}

// contoh constant time
func dominant(n int) int {
	var result int = 0

	for i := 0; i < n; i++ {
		// pada line 14 adalah operasi dominant, karena dia yang paling banyak di eksekusi sebanyak kita menginputkan jumlah n
		result += 1
	}
	result = result + 10

	return result
}