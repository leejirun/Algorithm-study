# 💡Github URL
https://github.com/leejirun/Algorithm-study/tree/main/third_week/worker5

# 💡**문제 분석**
1.수빈: 현재 위치 N에서 출발하여 동생이 있는 위치 K로 이동해야한다.
2.이동 방법 정리
- X - 1 (뒤로 한 칸 이동)
- X + 1 (앞으로 한 칸 이동)
- X * 2 (순간이동)
3.가장 빠른 시간(최소 이동 횟수)으로 K에 도착하는 방법의 수를 구해야 한다.
4.K에 도달하는 최소 시간과 최소 시간으로 K에 도달하는 방법의 수를 출력해야한다.

# 💡**알고리즘 설계**
최단 거리를 구하는 문제이기에 BFS 적용하기
:최소 이동 횟수를 찾고, 해당 경로의 개수를 동시에 계산
- queue 활용
- 방문 여부 배열 
- 체크 변수

# 💡코드
package main

import (
	"bufio"
	"fmt"
	"os"
)

func findWays(N, K int) (int, int) {
	if N == K {
		return 0, 1
	}

	const MAX = 100000
	queue := []int{N}
	visited := make([]int, MAX+1) // 방문 시간을 저장 (0이면 미방문)
	ways := make([]int, MAX+1)    // 도착 방법의 수 저장
	visited[N] = 1                // 시작 위치 방문 처리
	ways[N] = 1                   // 시작 위치에서 도달하는 방법은 1개

	minTime := 0
	found := false

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			pos := queue[0]
			queue = queue[1:]

			// 현재 위치에서 이동 가능한 위치들
			nextPositions := []int{pos - 1, pos + 1, pos * 2}

			for _, next := range nextPositions {
				if next < 0 || next > MAX {
					continue
				}

				// 처음 방문하는 경우
				if visited[next] == 0 {
					visited[next] = visited[pos] + 1
					ways[next] = ways[pos] // 방법의 수 그대로 계승
					queue = append(queue, next)
				} else if visited[next] == visited[pos]+1 {
					// 같은 최소 시간으로 도달한 경우 방법의 수 증가
					ways[next] += ways[pos]
				}

				// 목표 위치에 도달한 경우
				if next == K {
					minTime = visited[next] - 1
					found = true
				}
			}
		}

		if found {
			break
		}
	}

	return minTime, ways[K]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N, K int
	fmt.Fscan(reader, &N, &K)

	minTime, count := findWays(N, K)
	fmt.Fprintln(writer, minTime)
	fmt.Fprintln(writer, count)
}

# 💡풀이 방법
BFS+O(N)의 시간 복잡도


# 💡틀린 부분





# 💡느낀점