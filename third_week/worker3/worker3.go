package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	graph   [][]int
	visited [][]bool
	dx      = []int{0, 0, -1, 1} // 상하좌우 이동
	dy      = []int{-1, 1, 0, 0}
	N, M    int
)

func dfs(x, y int) {
	visited[y][x] = true // 방문 표시

	// 상하좌우 탐색
	for i := 0; i < 4; i++ {
		nx, ny := x+dx[i], y+dy[i]

		if nx >= 0 && nx < M && ny >= 0 && ny < N { // 맵 범위 체크
			if !visited[ny][nx] && graph[ny][nx] == 1 {
				dfs(nx, ny)
			}
		}
	}
}

func main() {
	var T, K int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscan(reader, &T)

	for t := 0; t < T; t++ {
		fmt.Fscan(reader, &M, &N, &K)

		// 그래프 초기화
		graph = make([][]int, N)
		visited = make([][]bool, N)
		for i := range graph {
			graph[i] = make([]int, M)
			visited[i] = make([]bool, M)
		}

		// 배추 위치 입력받기
		for i := 0; i < K; i++ {
			var x, y int
			fmt.Fscan(reader, &x, &y)
			graph[y][x] = 1
		}

		count := 0

		// DFS 실행
		for y := 0; y < N; y++ {
			for x := 0; x < M; x++ {
				if graph[y][x] == 1 && !visited[y][x] {
					dfs(x, y)
					count++
				}
			}
		}

		fmt.Fprintln(writer, count)
	}
}
