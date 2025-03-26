package main

import (
	"fmt"
)

func maxVolume(N, S, M int, volumes []int) int {
	// Create a DP table initialized with false
	dp := make([][]bool, N+1)
	for i := range dp {
		dp[i] = make([]bool, M+1)
	}

	// Initialize the starting volume
	dp[0][S] = true

	// Process each song
	for i := 0; i < N; i++ {
		for v := 0; v <= M; v++ {
			if dp[i][v] { // If volume v is possible at song i
				if v+volumes[i] <= M {
					dp[i+1][v+volumes[i]] = true
				}
				if v-volumes[i] >= 0 {
					dp[i+1][v-volumes[i]] = true
				}
			}
		}
	}

	// Find the maximum possible volume at the last song
	for v := M; v >= 0; v-- {
		if dp[N][v] {
			return v
		}
	}

	return -1 // If no valid volume is found
}

func main() {
	var N, S, M int
	fmt.Scan(&N, &S, &M)
	volumes := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&volumes[i])
	}

	result := maxVolume(N, S, M, volumes)
	fmt.Println(result)
}
