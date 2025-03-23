package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// 입력 받기
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// N 입력
	input, _ := reader.ReadString('\n')
	N, _ := strconv.Atoi(input[:len(input)-1])

	// DP 테이블 선언
	dp := make([]int, N+1)

	// DP 계산
	for i := 2; i <= N; i++ {
		// 기본적으로 1을 빼는 연산
		dp[i] = dp[i-1] + 1

		// 2로 나누어 떨어지는 경우
		if i%2 == 0 {
			dp[i] = min(dp[i], dp[i/2]+1)
		}

		// 3으로 나누어 떨어지는 경우
		if i%3 == 0 {
			dp[i] = min(dp[i], dp[i/3]+1)
		}
	}

	// 결과 출력
	fmt.Fprintln(writer, dp[N])
}
