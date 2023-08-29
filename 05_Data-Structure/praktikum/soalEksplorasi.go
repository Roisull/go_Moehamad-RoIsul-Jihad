package main

import (
	"fmt"
	"math"
)

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}

	diagonalLeftToRight := 0
	diagonalRightToLeft := 0

	size := len(matrix)

	for i := 0; i < size; i++ {
		diagonalLeftToRight += matrix[i][i]
		diagonalRightToLeft += matrix[i][size-i-1]
	}

	absoluteDifference := int(math.Abs(float64(diagonalLeftToRight - diagonalRightToLeft)))
	fmt.Printf("Selisih absolut antara dua diagonal: %d\n", absoluteDifference)
}