# 💡Github URL

: 

# 💡**문제 분석**

1.제약 조건

- 초기 볼륨은 S.
- 각 노래의 볼륨은 주어진 범위 내에서 조절 가능하다. [0, M]

목표

- 기타리스트가 모든 노래를 연주한 후 낼 수 있는 최대 볼륨을 N 결정하는 것이다/

입력

- N(노래 수), S(시작 볼륨), M(최대 볼륨)

2.볼륨 변경 목록

산출
-노래가 끝난 후의 최대 볼륨: N

- 모든 노래를 재생할 수 없는 경우: -1

3.제약사항:
1 ≤ N ≤ 100, 1 ≤ S ≤ M ≤ 1000, 1 ≤ volume change ≤ 100

# 💡**알고리즘 설계**

각 노래에 대해 다음 을 사용하여 i가능한 볼륨을 업데이트하기

- v + volumes[i](범위 내에 있는 경우)
- v - volumes[i](범위 내에 있는 경우)

=> 답: 가장 큰 유효 볼륨 dp[N]

# 💡코드

```go
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

```

# 💡적용한 방법

DP 방식, 시간 복잡도

# 💡 다른 풀이