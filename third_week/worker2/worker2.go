package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	adjList [][]int
	visited []bool
)

func dfs(node int) {
	visited[node] = true
	for _, neighbor := range adjList[node] {
		if !visited[neighbor] {
			dfs(neighbor)
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)
	n, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])

	adjList = make([][]int, n+1)
	visited = make([]bool, n+1)

	for i := 0; i < m; i++ {
		line, _ := reader.ReadString('\n')
		parts := strings.Fields(line)
		u, _ := strconv.Atoi(parts[0])
		v, _ := strconv.Atoi(parts[1])
		adjList[u] = append(adjList[u], v)
		adjList[v] = append(adjList[v], u)
	}

	components := 0
	for i := 1; i <= n; i++ {
		if !visited[i] {
			dfs(i)
			components++
		}
	}

	fmt.Println(components)
}
