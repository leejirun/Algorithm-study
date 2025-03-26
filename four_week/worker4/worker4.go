package main

import "fmt"

func tile2xN(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	const MOD = 10007
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % MOD
	}

	return dp[n]
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(tile2xN(n))
}
