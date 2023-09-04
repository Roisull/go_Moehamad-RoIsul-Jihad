package main

import (
	"fmt"
	"math"
)

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n)

	dp[0], dp[1] = cost[0], cost[1]

	for i := 2; i < n; i++ {
		dp[i] = cost[i] + int(math.Min(float64(dp[i-1]), float64(dp[i-2])))
	}

	return int(math.Min(float64(dp[n-1]), float64(dp[n-2])))
}

func Frog(jumps []int) int {
	return minCostClimbingStairs(jumps)
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))                // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50}))       // 40
	fmt.Println(Frog([]int{1, 100, 1, 1, 1, 100, 1, 1, 100})) // 6
}