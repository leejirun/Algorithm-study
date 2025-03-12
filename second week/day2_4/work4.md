## 💡문제 분석 요약
조건)
-하루에 하나의 상담만 가능
-상담을 시작하면 해당 기간(Ti) 동안 다른 상담 진행 불가
-N+1일 이후에는 상담 불가

## 💡알고리즘 설계

1.선택지:
a)현재 상담을 진행하는 경우
-day + T[day]로 이동하여 다음 상담 탐색
-이익은 P[day] 증가
b)현재 상담을 진행하지 않는 경우
-day + 1로 이동하여 다음 상담 탐색
-이익 변동 없음

2.dp[i] = max(dp[i+1], P[i] + dp[i+T[i]])
-p[i]는 i일부터 퇴사일까지 얻을 수 있는 최대 이익
-dp[i+1]: i일의 상담을 하지 않는 경우
-P[i] + dp[i+T[i]]: i일의 상담을 수행하는 

3.설계
-dp 배열을 뒤에서부터 채우며 최댓값을 저장
-dp[i] = max(dp[i+1], P[i] + dp[i+T[i]])를 반복


## 💡코드
`
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
`

## 💡개념
1. 완전 탐색 (Brute Force, DFS)
모든 경우를 탐색하는 방식
작은 N에서는 가능하지만 N이 크면 비효율적

2.동적 계획법 (DP)
작은 문제를 해결하면서 결과를 저장하고 재활용하는 기법
dp[i]를 뒤에서부터 채우는 Bottom-Up 방식 사용
점화식: dp[i] = max(dp[i+1], P[i] + dp[i+T[i]])

3.배열과 인덱스 관리
Go에서는 배열을 슬라이스(slice)로 선언하고 사용
dp 배열을 N+1 크기로 만들어서 초과 계산 방지

## 💡다른 풀이 방법: 데이터 구조를 변경해서 struct를 활용(T, P)쌍으로 저장
- struct 사용해 데이터 저장
- DFS(재귀) + 메모이제이션으로 중복 계산을 줄이기
`
package main

import (
	"fmt"
)

// 상담 정보를 저장할 구조체
type Job struct {
	time int // 상담 소요 기간
	pay  int // 상담 수익
}

var (
	N     int
	jobs  []Job
	memo  []int
)

// DFS + 메모이제이션 방식으로 최대 이익 계산
func dfs(day int) int {
	// 퇴사일을 넘어가면 상담 불가능
	if day >= N {
		return 0
	}

	// 이미 계산된 값이면 반환
	if memo[day] != -1 {
		return memo[day]
	}

	// 현재 상담을 하지 않는 경우
	profit := dfs(day + 1)

	// 현재 상담을 하는 경우 (퇴사 전에 상담이 끝나는 경우만)
	if day+jobs[day].time <= N {
		profit = max(profit, jobs[day].pay+dfs(day+jobs[day].time))
	}

	// 메모이제이션
	memo[day] = profit
	return profit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Scan(&N)

	jobs = make([]Job, N)
	memo = make([]int, N)

	// 메모 배열을 -1로 초기화 (미방문 표시)
	for i := range memo {
		memo[i] = -1
	}

	// 상담 정보 입력
	for i := 0; i < N; i++ {
		fmt.Scan(&jobs[i].time, &jobs[i].pay)
	}

	// DFS 탐색 시작
	fmt.Println(dfs(0))
}
`
정리
1) 구조체 사용: struct job
-데이터 접근이 더 직관적이다.
2) DFS+Memoization 사용
- memo[day]를 활용하여 중복 계산을 피함
- 시간 복잡도가 O(N)수준으로 줄어든다.
3) 재귀(DFS)사용
- 퇴사일을 넘는 경우를 쉽게 처리할 수 있다.
EX) if day >= N

