package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 인접 리스트 저장 (각 노드의 연결된 이웃 정보)
var graph map[int][]int
var visited []bool

func DFS(node int) {
	fmt.Print(node, " ")

	//방문한 노드 표시
	visited[node] = true

	//현재 노드의 모든 이웃 노드(neighbor)를 반복문으로 확인
	for _, neighbor := range graph[node] {
		if !visited[neighbor] {
			DFS(neighbor) //방문하지 않은 노드면 DFS(neighbor)를 재귀 호출
		}
	}
}

func BFS(start int) {
	//큐를 선언하고 시작 노드를 추가
	queue := []int{start}
	//방문 표시
	visited[start] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Print(node, " ")

		//현재 노드의 이웃 탐색
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number of nodes and edges: ")
	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)
	//노드 개수
	n, _ := strconv.Atoi(parts[0])
	//간선 개수
	m, _ := strconv.Atoi(parts[1])

	//인접 리스트 => key: 노드, value: 노드의 이웃을 나타내는 정수

	//인접 리스트 초기화
	graph = make(map[int][]int)

	//방문 배열 초기화
	visited = make([]bool, n+1)

	fmt.Println("Enter edges:")
	for i := 0; i < m; i++ {
		line, _ := reader.ReadString('\n')
		parts := strings.Fields(line)

		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])

		//a → b 연결
		graph[a] = append(graph[a], b)
		//b → a 연결 (양방향 그래프)
		graph[b] = append(graph[b], a)
	}

	fmt.Println("DFS:")
	// DFS 탐색 수행 (1번 노드부터)
	DFS(1)

	fmt.Println("\nBFS:")
	visited = make([]bool, n+1) // 방문 배열 초기화 (DFS 후 BFS 실행)
	// BFS 탐색 수행 (1번 노드부터)
	BFS(1)
}
