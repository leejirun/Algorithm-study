# 💡Github URL

: 

# 💡**문제 분석**

퇴사까지 N일 남음

- 매일 컨설팅 일을 선택할 수 있음
- 낮에는 아래 일을 처리함
1. T[i] → 완료하는 데 며칠이 걸림
2. P[i] → 완료되면 돈을 지불
- N일, N일 이전에 끝나는 일자리만 수락할 수 있음
- 퇴사 전 총 이익을 극대화하는 게 목적

---

백준이가 얻을 수 있는 최대 수익 구하기
-입력: T[i] - 몇일, P[i] - 수익

- 제약 조건
1 ≤ N ≤ 15, 1 ≤ T[i] ≤ 5, 1 ≤ P[i] ≤ 1,000
- 매일 두 가지 선택이 가능

# 💡**알고리즘 설계**

1.DP 적용
2.일(DAY)로 시작하여 벌 수 있는 최대 이익
3.해당 일자리 맡는 것이 가능 => i + T[i] <= N
1)일자리 잡기: P[i] + dp[i + T[i]]
2)일 skip: dp[i+1]
3)직무 수행 불가능한 경우 => i + T[i] > N
4.퇴사 후 수익 창출 없는 경우: dp[N]=0

# 💡코드

```go
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

```

# 💡틀린 이유

1. 거꾸로 계산함: 뒤에 일수부터 계산해야한다. 7~1 day 까지 ⇒ DP 방식

```go
for i := N - 1; i >= 0; i-- {
                if i+T[i] <= N {
                        dp[i] = max(P[i]+dp[i+T[i]], dp[i+1])
                } else {
                        dp[i] = dp[i+1]
                }
        }
```

⇒ 

dp[i]는 i번째날부터 퇴사일까지 얻을 수 있는 최대 이익을 의미한다. 그래서 미래의 최대 이익을 알아야 현재 결정을 할 수 있기 때문에 뒤에서부터 채워나가야한다. 그래서 dp[N-1]부터 dp[0]까지 거꾸로 계산해나감

1. 조건문 실수

```go
 for i := N - 1; i >= 0; i-- {
                //상담이 퇴사일 전에 끝나는 경우
                if i+T[i] <= N {
                        //상담할 지 말지 결정
                        dp[i] = max(P[i]+dp[i+T[i]], dp[i+1])
                } else {
                        //퇴사일을 넘기는 경우: 그 다음날의 이익 계산
                        dp[i] = dp[i+1]
                }
        }

        //얻을 수 있는 최대 이익
        return dp[0]

```

⇒ 

1.상담 할 수 있는 경우

- 상담 수락 시 :
P[i] (현재 상담의 보수) + dp[i+T[i]] (상담 끝난 후 얻을 수 있는 최대 이익)
- 상담 거절 시 :
dp[i+1] (그냥 다음 날의 최대 이익을 가져감)

2.상담을 할 수 없는 경우

- 상담이 퇴사일을 초과하면 이 상담 안함
- 다음날 dp[i+1]의 최대 이익 가져옴

1. 디버깅

```go
dp[7]: 0
dp[6]: 0
P[i]:  15
dp[i+T[i]]:  0
dp[i+1]:  0
dp[5]:  15

P[i]:  20
dp[i+T[i]]:  15
dp[i+1]:  15
dp[4]:  35

P[i]:  10
dp[i+T[i]]:  35
dp[i+1]:  35
dp[3]:  45

P[i]:  20
dp[i+T[i]]:  0
dp[i+1]:  45
dp[2]:  45

P[i]:  10
dp[i+T[i]]:  35
dp[i+1]:  45
dp[1]:  45

```

# 💡 적용 개념

- 재귀적 접근 방식
- DP + 최적의 O(N)

# 💡 다른 풀이