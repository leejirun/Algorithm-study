package main

import (
	"fmt"
)

func maxProfit(N int, T []int, P []int) int {
	dp := make([]int, N+1) // dp[i] stores max profit starting from day i

	for i := N - 1; i >= 0; i-- {
		if i+T[i] <= N { // If job can be completed before retirement
			dp[i] = max(P[i]+dp[i+T[i]], dp[i+1])
		} else { // Job cannot be taken
			dp[i] = dp[i+1]
		}
	}

	return dp[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var N int
	fmt.Scan(&N) // Read number of days

	T := make([]int, N)
	P := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&T[i], &P[i])
	}

	fmt.Println(maxProfit(N, T, P))
}
