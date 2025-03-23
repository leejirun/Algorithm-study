# 💡Github URL

: 

# 💡**문제에서 구해야 할 것**

문제 조건)

정수: N(0≤𝑁≤90), 64비트 정수형

# 💡**알고리즘 설계**

- N=0 ⇒ 반품
- N=1 ⇒ 반품
- 반복 계산, 2부터 시작하기

# 💡코드
package main

import "fmt"

func fibonacci(n int) uint64 {
    if n == 0 {
        return 0
    } else if n == 1 {
        return 1
    }

    var a, b uint64 = 0, 1
    var result uint64

    for i := 2; i <= n; i++ {
        result = a + b
        a, b = b, result
    }

    return result
}

func main() {
    var n int
    fmt.Scan(&n)
    fmt.Println(fibonacci(n))
}


 💡 개념: 피보나치 수열

정의)

각 숫자가 이전 두 숫자의 합이 되는 수열기를 의미한다.
=> ( n ) 에프 ( n )=F ( 엔−1 )+F ( 엔−2 )

예시)
0, 1, 1, 2, 3, 5, 8, 13, 21, 34, ...

- `F(0) = 0`
- `F(1) = 1`
- `F(2) = 1`→`0 + 1`
- `F(3) = 2`→`1 + 1`
- `F(4) = 3`→`1 + 2`
- `F(5) = 5`→`2 + 3`

사용한 사례)

- 빠른 검색 알고리즘
- 동적 계획법(DP)
- 식별할 가능성이 있는 배열
- 주식 시장 이용