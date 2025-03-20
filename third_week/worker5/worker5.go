package main

import (
	"bufio"
	"fmt"
	"os"
)

const MAX = 100000

func bfs(n, k int) int {
	if n >= k {
		return n - k // If N is already greater than or equal to K, we can only move back
	}

	//방문체크
	visited := make([]bool, MAX+1) // To track visited positions
	//bfs 사용 방법
	queue := make([]int, 0)
	steps := make([]int, MAX+1) // Store step count for each position

	queue = append(queue, n)
	visited[n] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		//목표 위치에 도달하면, 걸음 수 반환하기
		if current == k {
			return steps[current]
		}

		// Possible next positions
		nextPositions := []int{current - 1, current + 1, current * 2}

		for _, next := range nextPositions {
			//유효성 검사하기
			if next >= 0 && next <= MAX && !visited[next] {
				visited[next] = true
				queue = append(queue, next)
				//목표 도달하기 위해 몇번 이동했는지 체크하기
				steps[next] = steps[current] + 1
			}
		}
	}
	return -1
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	var n, k int
	fmt.Sscanf(input, "%d %d", &n, &k)

	result := bfs(n, k)
	fmt.Println(result)
}
