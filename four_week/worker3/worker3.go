package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func stairClimb(scores []int) int {
	n := len(scores)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return scores[0]
	}
	if n == 2 {
		return scores[0] + scores[1]
	}

	// DP 테이블 초기화
	dp := make([]int, n)
	dp[0] = scores[0]
	dp[1] = scores[0] + scores[1]
	dp[2] = max(scores[0]+scores[2], scores[1]+scores[2])

	// 점화식을 이용하여 DP 테이블 채우기
	for i := 3; i < n; i++ {
		dp[i] = max(dp[i-2]+scores[i], dp[i-3]+scores[i-1]+scores[i])
	}

	// 마지막 계단까지 갔을 때의 최대 점수 반환
	return dp[n-1]
}

func main() {
	var n int
	fmt.Scan(&n)

	scores := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&scores[i])
	}

	fmt.Println(stairClimb(scores))
}
