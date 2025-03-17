package main

import (
	"fmt"
	"math"
)

var (
	n         int
	numbers   []int
	operators [4]int
	maxResult = math.MinInt32
	minResult = math.MaxInt32
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
