package main

import "fmt"

func main() {
	dominant(10)
}

func dominant(n int) {
	var result int = 0

	for i := 0; i < n; i++ {
		// pada line 14 adalah operasi dominant, karena dia yang paling banyak di eksekusi sebanyak kita menginputkan jumlah n
		result += 1
	}

	fmt.Println(result)
}