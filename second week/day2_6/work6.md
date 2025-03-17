## 💡문제 분석 요약
이 문제는 완전 탐색(Brute Force) 및 백트래킹(Backtracking) 기법을 활용하여 가능한 모든 연산자 조합을 탐색하는 문제이다


-> 조건
1.N(2 ≤ N ≤ 11)개의 숫자가 순서대로 주어진다.
2.N-1개의 연산자가 주어지며, 각각 덧셈(+), 뺄셈(-), 곱셈(×), 나눗셈(÷) 연산의 개수를 나타낸다.
3.숫자의 순서는 변경할 수 없으며, 주어진 연산자의 개수만큼 연산자를 배치해야 한다.
4.연산은 앞에서부터 순서대로 진행해야 한다. 즉, 연산자 우선순위를 고려하지 않는다.
5.나눗셈(÷)은 C++14 스타일의 정수 나눗셈을 따른다.
즉, 음수를 나눌 때는 양수로 변환 후 몫을 구하고, 다시 음수로 변환한다.
예) -5 ÷ 3 = -1, 5 ÷ -3 = -1

## 💡알고리즘 설계
1) 완전 탐색 (Brute Force)
가능한 모든 연산자 배치를 시도하여 최댓값과 최솟값을 찾는다.
N이 최대 11이므로, 최악의 경우 (N-1)!번 연산을 수행해야 한다.
10! = 3,628,800이므로 브루트 포스로 접근이 가능하다.

2) 백트래킹 (Backtracking)
DFS(깊이 우선 탐색)를 사용하여 연산자를 하나씩 배치하면서 모든 경우를 탐색한다.
연산자를 배치할 때마다 사용한 연산자의 개수를 감소시키고, 탐색이 끝나면 복구하는 방식으로 구현한다.

## 💡코드
`
package main

import (
	"fmt"
	"math"
)

var (
	n          int
	numbers    []int
	operators  [4]int
	maxResult  = math.MinInt32
	minResult  = math.MaxInt32
)

// 연산 수행 함수
func calculate(a, b, op int) int {
	switch op {
	case 0:
		return a + b
	case 1:
		return a - b
	case 2:
		return a * b
	case 3:
		if a < 0 {
			return -((-a) / b) // C++14 스타일 나눗셈
		}
		return a / b
	}
	return 0
}

// 백트래킹을 이용한 완전 탐색
func backtrack(index, result int) {
	if index == n { // 모든 연산자를 사용했을 경우
		if result > maxResult {
			maxResult = result
		}
		if result < minResult {
			minResult = result
		}
		return
	}

	for i := 0; i < 4; i++ {
		if operators[i] > 0 {
			operators[i]-- // 연산자 사용
			nextResult := calculate(result, numbers[index], i)
			backtrack(index+1, nextResult)
			operators[i]++ // 연산자 복구 (백트래킹)
		}
	}
}

func main() {
	// 입력 받기
	fmt.Scan(&n)
	numbers = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}
	for i := 0; i < 4; i++ {
		fmt.Scan(&operators[i])
	}

	// 백트래킹 수행
	backtrack(1, numbers[0])

	// 결과 출력
	fmt.Println(maxResult)
	fmt.Println(minResult)
}
`

## 💡 정리
1.전역 변수 선언
-numbers: 입력받은 숫자 배열
-operators: 덧셈, 뺄셈, 곱셈, 나눗셈의 개수를 저장하는 배열
-maxResult: 결과 중 최댓값을 저장 (초기값: math.MinInt32)
-minResult: 결과 중 최솟값을 저장 (초기값: math.MaxInt32)

2.연산 수행 함수 calculate()
-두 수 a와 b 및 연산자 정보를 받아 계산을 수행
-switch 문을 사용하여 연산자에 따라 연산 수행
-나눗셈(÷)은 C++14 스타일을 적용하여 처리

3.백트래킹을 이용한 완전 탐색 함수 backtrack()
-index가 n에 도달하면 최댓값/최솟값을 갱신하고 종료
-4개의 연산자 중 사용 가능한 연산자가 있다면 사용
-사용한 연산자는 개수를 감소시키고, 탐색이 끝나면 복구(백트래킹)

4.메인 함수 (main())
-n과 숫자 배열, 연산자 개수를 입력받음
-backtrack(1, numbers[0]) 호출하여 탐색 시작
-최댓값과 최솟값을 출력


## 💡틀린 부분


## 💡느낀점