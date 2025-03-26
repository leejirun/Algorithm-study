# 💡Github URL

: 

# 💡**문제분석**

- 계단이 N개 주어진다.
- 각 계단에는 점수가 있다.
- 계단을 오를 때의 규칙:
    1. 한 번에 한 계단 또는 두 계단을 오를 수 있다.
    2. 연속된 세 개의 계단을 밟을 수 없다. (즉, 세 개 연속해서 밟는 경우가 발생하면 안 됨)
    3. 마지막 계단은 반드시 밟아야 한다.
- 마지막 계단까지 올라갔을 때 얻을 수 있는 최대 점수를 구하자!

# 💡**알고리즘 설계**

- DP 동적 계획법 이용
- **dp[i]**: i번째 계단까지 올라갔을 때 얻을 수 있는 최대 점수
- 마지막 계단을 반드시 밟아야 하므로, 두 가지 경우가 가능
    - **(이전 계단을 밟고 올라옴)** → `dp[i] = dp[i-2] + score[i]`
    - **(이전 계단을 건너뛰고 올라옴, 하지만 전전 계단을 밟음)** → `dp[i] = dp[i-3] + score[i-1] + score[i]`
- 방식: dp[i]=max(dp[i−2]+score[i],dp[i−3]+score[i−1]+score[i])
- 초기값 설정
    - dp[1] = score[1]
    - `dp[2] = score[1] + score[2]` (두 번째 계단까지는 연속으로 밟을 수 있음)
    - dp[3] = max(score[1] + score[3], score[2] + score[3])

# 💡코드

```go
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

```

# 정리)

**1.`stairClimb(scores []int) int` 함수**

- 점수 배열을 받아서 최대 점수를 반환
- `dp` 배열을 생성하여 점화식을 이용해 최대 점수를 구함
- 계단이 1개 또는 2개인 경우를 별도로 처리
- 세 번째 계단부터 점화식을 적용하여 `dp[i]`를 채움

2.`max()` 함수

- 두 수를 비교해서 큰 값을 반환하는 함수
- 동적 계획법에서 최댓값을 찾는 데 사용됨

# 💡 개념

`O(N)`: 한 번의 반복문을 사용하여 `dp` 배열을 채우므로 선형 시간 복잡도

# 💡 틀린 이유 & 틀린 부분 수정

# 💡 다른 풀이