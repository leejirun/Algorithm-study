## 💡문제 분석 요약
-5×5 크기의 빙고판이 주어지며, 각 칸에는 1~25까지의 숫자가 중복 없이 배치됨
-사회자가 부르는 숫자 25개가 순서대로 주어짐
-사회자가 숫자를 몇 번째 부른 후 처음으로 빙고(Bingo)가 되는지를 출력

조건
-가로, 세로, 대각선 5줄이 모두 지워지면 빙고
-빙고가 3줄 이상 되면 종료

## 💡알고리즘 설계
-5×5 빙고판을 map[int][2]int 자료구조를 사용하여 숫자의 위치를 저장
-사회자가 부르는 숫자 25개를 배열에 저장
-숫자가 불릴 때마다 해당 위치를 visited 배열에 표시
-가로, 세로, 두 개의 대각선을 확인하여 몇 줄이 완성되었는지 체크
-빙고가 3개 이상이면 정답 출력
-시간복잡도: 숫자를 찾는 연산은 O(1), 줄 확인은 O(5), 전체 반복은 최대 O(25), 따라서 O(25) ≈ O(1).

## 💡코드
package main

import (
	"bufio"
	"fmt"
	"os"
)

var board [5][5]int       // 빙고판
var visited [5][5]bool    // 체크 여부
var position map[int][2]int // 숫자의 위치 저장

func checkBingo() int {
	count := 0

	// 가로, 세로 체크
	for i := 0; i < 5; i++ {
		row, col := true, true
		for j := 0; j < 5; j++ {
			if !visited[i][j] {
				row = false
			}
			if !visited[j][i] {
				col = false
			}
		}
		if row {
			count++
		}
		if col {
			count++
		}
	}

	// 대각선 체크
	diag1, diag2 := true, true
	for i := 0; i < 5; i++ {
		if !visited[i][i] {
			diag1 = false
		}
		if !visited[i][4-i] {
			diag2 = false
		}
	}
	if diag1 {
		count++
	}
	if diag2 {
		count++
	}

	return count
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	position = make(map[int][2]int)

	// 빙고판 입력
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Scan(&board[i][j])
			position[board[i][j]] = [2]int{i, j}
		}
	}

	// 사회자가 부르는 숫자
	count := 0
	for i := 0; i < 25; i++ {
		var num int
		fmt.Scan(&num)
		count++

		// 숫자 찾고 체크
		if pos, found := position[num]; found {
			visited[pos[0]][pos[1]] = true
		}

		// 빙고 개수 확인
		if checkBingo() >= 3 {
			fmt.Println(count)
			break
		}
	}
}


## 💡개념
1.해시맵이란?
키(Key)와 값(Value)을 저장하는 자료구조로, 빠른 데이터 검색과 삽입이 가능한 구조

배열이나 리스트를 사용하면 값을 찾기 위해 모든 요소를 확인해야 하지만,
해시맵을 사용하면 키를 이용해 바로 값을 찾을 수 있기 때문에 속도가 훨씬 빠름


2.go에서 해시맵 활용
숫자의 위치를 빠르게 찾기 위해 map[int][2]int를 활용
`
package main

import "fmt"

func main() {
    // 1. 해시맵 생성 (키: string, 값: int)
    age := make(map[string]int)

    // 2. 데이터 추가
    age["Alice"] = 25
    age["Bob"] = 30
    age["Charlie"] = 22

    // 3. 데이터 조회
    fmt.Println("Alice's age:", age["Alice"]) // 출력: Alice's age: 25

    // 4. 데이터 삭제
    delete(age, "Bob")

    // 5. 키 존재 여부 확인
    value, exists := age["Bob"]
    if exists {
        fmt.Println("Bob's age:", value)
    } else {
        fmt.Println("Bob is not found!") // 출력: Bob is not found!
    }

    // 6. 전체 데이터 출력
    for key, value := range age {
        fmt.Println(key, ":", value)
    }
}
`

3.빙고 문제에서 해시맵 사용 이유

빙고 문제에서 해시맵을 사용한 이유는 각 숫자의 위치를 빠르게 찾기 위해서다.
5×5 빙고판에 숫자가 랜덤하게 배치되기 때문에,
해시맵을 사용하지 않으면 불린 숫자의 위치를 찾는 데 O(25) 시간이 걸린다.

하지만 해시맵을 사용하면 O(1) 만에 숫자의 위치를 찾을 수 있다.
빙고 문제에서 사용한 해시맵 코드:

`
position = make(map[int][2]int) // 숫자의 위치 저장

// 빙고판 입력 시 해시맵에 숫자의 위치 저장
for i := 0; i < 5; i++ {
    for j := 0; j < 5; j++ {
        fmt.Scan(&board[i][j])
        position[board[i][j]] = [2]int{i, j} // 숫자 → (행, 열) 저장
    }
}
`

빙고판이 다음과 같다면

11 12  2 24 10
16  1 13  3 25
 6 20  5 21 17
19  4  8 14  9
22 15  7 23 18

position[12] → [0,1] (12는 (0,1) 위치에 있음)
position[21] → [2,3] (21은 (2,3) 위치에 있음)

이제 사회자가 숫자를 부르면 해시맵을 이용해 O(1) 만에 위치를 찾고,
바로 visited 배열을 업데이트할 수 있다.


## 💡틀린 이유
시간을 최대한 많이 주어서 풀려고 많이 시도했는데, 잘 풀리지 않았다.
그래서 gpt의 도움을 받아서 이해하는 게 현명하다고 생각했다.

막혔던 부분은 

1)'대각선'체크이다.
왼쪽 오른쪽 오른쪽 왼쪽 대각선 체크를 하는데, 체크를 빠뜨렸다.

`
diag1, diag2 := true, true
for i := 0; i < 5; i++ {
    if !visited[i][i] {  //왼쪽 위 -> 오른쪽 아래 대각선 확인
        diag1 = false
    }
    if !visited[i][4-i] { //오른쪽 위 -> 왼쪽 아래 대각선 확인
        diag2 = false
    }
}
if diag1 {
    count++
}
if diag2 {
    count++
}
`
=> visited[i][4-i] 체크 하지 않았다.

2) 빙고 3개 이상 시 체크: 3빙고
25개 숫자를 전부 다 받고 나서 빙고 개수를 계산해버렸다.
`
if checkBingo() >= 3 {
    fmt.Println(count)
    break
}
`
=> 3빙고 이상이 되면 즉시 종료하게 break를 써야 한다.

## 느낀점
처음에 문제를 읽고 풀어야 하는 걸 생각할 때 머릿속으로 정리가 잘 안되는 것같다.

신경 써야 할 부분들이 많다보니 하나씩 놓치는 것 같다.
특히, 대각선의 경우가 그렇다.

나는 시간을 재고 풀기보다는 푸는 데 집중해야할 듯하다.
