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
