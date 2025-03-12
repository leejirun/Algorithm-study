package main

import (
	"bufio"
	"fmt"
	"os"
)

var board [5][5]int         // 빙고판
var visited [5][5]bool      // 체크 여부
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
