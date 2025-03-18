# 💡Github URL


# 💡**문제 분석**
- 배추가 심어진 곳을 하나의 군집으로 보고, 몇 개의 군집이 있는지 개수를 세는 문제다.
- 입력: 가로 M, 세로 N, 배추 개수 K 와 배추 위치 리스트
- 출력: 필요한 최소한의 벌레(군집 수)

# 💡**알고리즘 설계**
1.DFS 또는 BFS을 이용한다.
2.방문 표시를 체크해야한다.
3.2차원 그래프 만들어서 배추가 있는 위치를 표시한다.
4.연결된 군집 개수 세기 위해, 총 탐색을 몇 번 했는 지 개수를 저장하면 되겠다.

# 💡코드
package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	graph  [][]int
	visited [][]bool
	dx     = []int{0, 0, -1, 1} // 상하좌우 이동
	dy     = []int{-1, 1, 0, 0}
	N, M   int
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


# 💡다른 풀이
- 또, BFS 방법도 보았다.
- 풀이 방법)
1.배추가 있는 곳을 찾는다.
2.큐에 해당 위치를 넣고 BFS 탐색을 시작한다.
3.큐에서 하나씩 꺼내며 인접한 배추를 큐에 넣고 방문 처리한다.
4.한 번의 BFS 실행이 끝날 때마다 군집 개수를 증가시킨다.
package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

var (
	graph  [][]int
	visited [][]bool
	dx     = []int{0, 0, -1, 1} // 상하좌우 이동
	dy     = []int{-1, 1, 0, 0}
	N, M   int
)

func bfs(x, y int) {
	queue := []Point{{x, y}}
	visited[y][x] = true // 방문 처리

	for len(queue) > 0 {
		curr := queue[0] // 큐에서 요소를 꺼냄
		queue = queue[1:]

		// 네 방향 탐색
		for i := 0; i < 4; i++ {
			nx, ny := curr.x+dx[i], curr.y+dy[i]

			if nx >= 0 && nx < M && ny >= 0 && ny < N { // 범위 체크
				if !visited[ny][nx] && graph[ny][nx] == 1 {
					visited[ny][nx] = true
					queue = append(queue, Point{nx, ny}) // 큐에 추가
				}
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

		// BFS 실행
		for y := 0; y < N; y++ {
			for x := 0; x < M; x++ {
				if graph[y][x] == 1 && !visited[y][x] {
					bfs(x, y)
					count++
				}
			}
		}

		fmt.Fprintln(writer, count)
	}
}


# 💡DFS VS BFS 비교
비교항목	DFS (깊이 우선 탐색)	BFS (너비 우선 탐색)
탐색 방식	깊이 먼저 탐색	너비 먼저 탐색
구현 방식	재귀 (스택) 사용	큐 사용
속도 차이	작은 그래프에서는 빠름	큰 그래프에서도 일정한 속도
메모리 사용	재귀 호출이 많으면 스택 오버플로우 가능	큐를 사용해 비교적 안정적

- 둘 중 선택 방법?
:그래프가 작고, 재귀 호출이 부담되지 않으면 DFS
:그래프가 크거나, 재귀 호출이 많아질 경우 BFS
:연결된 모든 요소를 한 번만 방문하면 되는 문제는 BFS가 더 직관적


# 💡느낀점
