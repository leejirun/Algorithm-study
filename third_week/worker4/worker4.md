# 💡Github URL


# 💡**문제 분석**
- N × N 크기의 2차원 지도에서 1(집)과 0(빈 공간)으로 이루어진 지도가 주어진다.
-연결된 1(집)들을 하나의 단지로 간주하며, 상하좌우로 연결된 집들만 같은 단지이다.
-단지의 개수를 찾고, 각 단지 내 집의 수를 오름차순으로 출력해야한다.

# 💡**알고리즘 설계**
1.그래프 표현 및 입력 처리
- 2차원 슬라이스([][]int)로 지도 정보를 저장한다.
- 방문 여부를 확인하기 위한 2차원 슬라이스([][]bool)를 준비한다.

2. DFS 또는 BFS 탐색을 통해 단지 찾기
-모든 좌표를 순회하면서 1(집)이면서 방문하지 않은 경우 새로운 단지를 탐색한다.
-DFS 또는 BFS를 이용해 해당 단지의 모든 집을 방문하며 개수를 세어 준다.

3. 결과 정렬 및 출력
-단지의 개수와 단지 내 집의 수를 저장한 배열을 정렬하여 출력한다.

# 💡코드
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


# 💡풀이 방법
-DFS
-시간 복잡도 O(N^2)

# 💡다른 풀이 방법: BFS 방식 + q 사용
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

// 큐를 위한 구조체 정의
type Point struct {
	x, y int
}

// BFS 함수 정의
func bfs(x, y int) int {
	queue := []Point{{x, y}} // 큐 초기화
	visited[x][y] = true
	count := 1 // 현재 위치도 집이므로 1부터 시작

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:] // 큐에서 첫 번째 원소 제거 (dequeue)

		// 네 방향으로 이동
		for i := 0; i < 4; i++ {
			nx, ny := curr.x+dx[i], curr.y+dy[i]

			// 범위 체크 & 방문 여부 체크 & 집(1)인지 확인
			if nx >= 0 && nx < N && ny >= 0 && ny < N {
				if !visited[nx][ny] && grid[nx][ny] == 1 {
					visited[nx][ny] = true
					queue = append(queue, Point{nx, ny}) // 큐에 추가 (enqueue)
					count++
				}
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
				complexes = append(complexes, bfs(i, j))
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



# 💡느낀점
