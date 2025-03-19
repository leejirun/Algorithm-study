package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	N       int
	grid    [][]int
	visited [][]bool
	dx      = []int{-1, 1, 0, 0} // 상, 하, 좌, 우 방향 벡터
	dy      = []int{0, 0, -1, 1}
)

// DFS 함수 정의
func dfs(x, y int) int {
	visited[x][y] = true
	count := 1 // 현재 위치도 집이므로 1부터 시작

	// 네 방향으로 이동
	for i := 0; i < 4; i++ {
		nx, ny := x+dx[i], y+dy[i]

		// 유효 범위 체크 & 방문 여부 체크 & 집이 있는지 확인
		if nx >= 0 && nx < N && ny >= 0 && ny < N {
			if !visited[nx][ny] && grid[nx][ny] == 1 {
				count += dfs(nx, ny) // 재귀적으로 탐색
			}
		}
	}
	return count
}

func main() {
	// 입력 처리
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &N)

	grid = make([][]int, N)
	visited = make([][]bool, N)

	for i := 0; i < N; i++ {
		grid[i] = make([]int, N)
		visited[i] = make([]bool, N)
		var line string
		fmt.Fscan(reader, &line)
		for j := 0; j < N; j++ {
			grid[i][j] = int(line[j] - '0') // '0'을 빼서 정수 변환
		}
	}

	// 단지 개수 및 각 단지 크기 저장
	var complexes []int

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] == 1 && !visited[i][j] {
				complexes = append(complexes, dfs(i, j))
			}
		}
	}

	// 정렬 후 출력
	sort.Ints(complexes)

	fmt.Println(len(complexes)) // 총 단지 수
	for _, size := range complexes {
		fmt.Println(size) // 각 단지 내 집의 개수
	}
}
