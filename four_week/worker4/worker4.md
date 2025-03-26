# 💡Github URL

: 

# 💡**문제 분석**

- 너비가 2이고 높이가 N인 직사각형을 1×2 또는 2×1 타일로 채우는 방법의 수를 구하는 문제
- 동적 계획법(DP, Dynamic Programming)을 활용하여 해결하기

# 💡**알고리즘 설계**

1. 식 정리
- N = 1일 때: |
    
    → 1가지 경우 (1×2 타일 1개)
    
- N = 2일 때: ||, =
    
    → 2가지 경우 (1×2 타일 2개, 2×1 타일 1개)
    

.2. 점화식

N 크기의 직사각형을 채우는 방법을 **dp[N]**이라고 하면,

⇒ 맨 마지막에 1×2 타일을 놓는 경우 :  dp[N-1]

⇒ 맨 마지막에 2×1 타일을 두 개 쌓는 경우 :  dp[N-2]

⇒ 점화식: dp[N]=dp[N−1]+dp[N−2]

1. 초기값
- dp[1] = 1
- dp[2] = 2

1. 최적화
- DP 배열을 활용하면 O(N) 시간 복잡도로 해결 가능
- 단, N이 커질 경우 큰 숫자를 다루기 때문에 **mod 10007** 연산이 필요

# 💡코드

```go
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

```

# 💡설명

- n이 1이면 1을 반환, n이 2이면 2를 반환

`=>> dp[1] = 1`, `dp[2] = 2`로 설정

- 반복문을 통해 점화식 적용: dp[i] = (dp[i-1] + dp[i-2]) % 10007
- 최종 결과 출력:  `tile2xN(n)`의 결과를 출력

# 💡 이유