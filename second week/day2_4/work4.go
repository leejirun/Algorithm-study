package main

import (
	"fmt"
)

// 최대 상담 수익 계산 (DP 방식)
func maxProfit(N int, T []int, P []int) int {
	dp := make([]int, N+1) // dp[i]: i일부터 퇴사일까지 얻을 수 있는 최대 이익

	// 뒤에서부터 탐색
	for i := N - 1; i >= 0; i-- {
		if i+T[i] > N { // 상담 기간이 넘어가면
			dp[i] = dp[i+1]
		} else {
			dp[i] = max(dp[i+1], P[i]+dp[i+T[i]])
		}
	}

	return dp[0] // 첫날부터 시작했을 때의 최대 이익
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var N int
	fmt.Scan(&N)

	T := make([]int, N)
	P := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&T[i], &P[i])
	}

	fmt.Println(maxProfit(N, T, P))
}
