💡문제 분석 요약
정수 x 주어지는 경우, 1)~3)을 1로 만들어서 연산을 사용하는 횟수의 최솟값 구하기기
- 1) x가 3으로 나누어 떨얼지면, 3으로 나눈다.
- 2) x가 2로 나누어 떨어지면, 2로 나눈다.
- 3) 1을 뺀다.

💡알고리즘 설계
1) 최솟값 => 동적 계획법(DP)를 사용해야겠다.
2) 최소 연산 횟수: 문제의 1)~3)을 적용하기

💡코드
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


💡 풀이 
- 동적 계획 법 적용
1.작은 문제를 해결하면서 큰 문제를 해결하는 방식이다.
2.중복된 계산을 방지하기 위해 배열에 저장한다.

1번: 최적화하기
a[i] 구하기 위해서 => a[i-1], a[i/2], a[i/3] + 문제에서 요구한 최소값 

2번: 중복 부분 문제
같은 계산을 반복하지 않도록 배열 저장


💡느낀점
문제에 답이 있어서 풀기가 비교적 쉬웠다.
코드를 풀고, 알게 된 건 문제에서 요구한 조건이 동적 계획법이라는 것이다. 동적 개념을 이전 문제를 풀면서 알고 있기도 했기에 쉽게 풀 수 있지 않았나 쉽다.
