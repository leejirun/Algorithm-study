# 💡Github URL

https://github.com/leejirun/Algorithm-study/tree/main/third_week/worker3

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

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	graph  [][]int
	visited [][]bool
	dx     = []int{0, 0, -1, 1} // 상하좌우 이동 <-: -1, 0, ->:1 / ↓: -1, 0, ↑: 1
	dy     = []int{-1, 1, 0, 0}
	N, M   int
)

func dfs(x, y int) {
	visited[y][x] = true // 방문 표시 => 틀린 부분

	// 상하좌우 탐색
	for i := 0; i < 4; i++ {
		nx, ny := x+dx[i], y+dy[i] //범위 체크용 nx:가로 방향, ny:세로 방향

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
		graph = make([][]int, N) // N x M 크기의 2D 배열
		visited = make([][]bool, N) // N x M 크기의 2D 배열
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

```

# 💡틀린 이유

1. 방문 표시

```go
visited[x][y] => [y][x]로 헤야함
```

- y와 N: 세로 크기 기준
- x와 M: 가로 크기 기준

⇒ 문제에서 가로M, 세로N이라고 지정을 분명히 했기 때문에, x와 y는 뒤집어서 사용하면 잘못된 인덱싱이 된다. 

⇒ 예를 들어, 

```go
graph (N=3, M=5)
  0 1 2 3 4   (x 좌표)
0 . . 1 . . 
1 . 1 1 . . 
2 . . . 1 1 
(y 좌표)
```

- visited[y][x]를 사용하면, visited[1][2]는 올바른 위치, 배추가 있는 곳이고
- visited[x][y]를 사용하면, visited[2][1]은 없는 곳을 가리키게 된다.

2.상하좌우 이동 반복문 작성 오류

```go
for i := 0; i < 4; i++ {
    nx, ny := x+dx[i], y+dy[i]
```

- dx와 dy는 방향을 위한 용도이기 때문에 이를 사용해서 상하좌우 이동을 할 수 있게 해야 한다.
- 위에서 정의했던 걸 보면

```go
dx := []int{0, 0, -1, 1}  // 좌우 이동
dy := []int{-1, 1, 0, 0}  // 상하 이동
```

- i=0이면, dx[0]=0, dy[0] = -1
- i=1이면, dx[1]=0, dy[1]=1
- i=2이면, dx[2]=-1, dy[2]=0
- i=3이면, dx[3]=-1, dy[3]=0

1. 범위 체크 오류

```go
if nx >= 0 && nx < M && ny >= 0 && ny < N { // 범위 체크
```

- 배열 인덱스가 0부터 시작이고, 문제에서도 `(0 ≤ nx < M)` 및 `(0 ≤ ny < N)`를 정의 했기 때문에 체크가 필요하다. 이를 하지 않아서 array index out of bounds 오류가 발생했다.

1. 조건문 오류

```go
if !visited[ny][nx] && graph[ny][nx] == 1 {
    dfs(nx, ny) // 재귀 호출
}
```

- 나는 graph[nx][ny] == 1 이라고 했다. 방향은 틀렸지만, 우선 배추가 있으면 1일 거니까 이 설정은 완료헸다. 근데,  !visited[ny][nx] 이 부분을 하지 않았다. 이미 방문한 곳인지 확인하는 것을 생각하지 못했다.

# 💡다른 풀이

- 또, BFS 방법도 보았다.

```go
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
	N, M   int //N: 가로, M: 세로
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

```

# 💡DFS VS BFS 비교

비교항목:	DFS (깊이 우선 탐색)  ||	BFS (너비 우선 탐색)
탐색 방식:	깊이 먼저 탐색 ||	너비 먼저 탐색
구현 방식:	재귀 (스택) 사용 ||	큐 사용
속도 차이:	작은 그래프에서는 빠름	|| 큰 그래프에서도 일정한 속도
메모리 사용:	재귀 호출이 많으면 스택 오버플로우 가능	|| 큐를 사용해 비교적 안정적

둘 중 선택 방법?
:그래프가 작고, 재귀 호출이 부담되지 않으면 DFS
:그래프가 크거나, 재귀 호출이 많아질 경우 BFS
:연결된 모든 요소를 한 번만 방문하면 되는 문제는 BFS가 더 직관적

# 💡느낀점

이번주 1,2회차는 쉬운 문제였다면, 오늘은 한 두번 더 생각을 해야하는 부분이 많았어서 힘들었는데 이해가 돼서 재미있다. 특히 오늘은 방향값으로 탐색 순서를 정하는 조건이 추가로 생겨서 하나 더 배웠다.