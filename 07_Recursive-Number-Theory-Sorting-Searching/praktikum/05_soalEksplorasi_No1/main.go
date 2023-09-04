package main

import (
	"fmt"
	"math"
)

func MaxSequence(arr []int) int {
	maxEndingHere := 0
	maxSoFar := math.MinInt32

	for _, num := range arr {
		maxEndingHere = int(math.Max(float64(num), float64(maxEndingHere+num)))
		maxSoFar = int(math.Max(float64(maxSoFar), float64(maxEndingHere)))
	}

	return maxSoFar
}

func main() {
	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))   // 7
	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))   // 7
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))   // 8
	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))    // 12
}