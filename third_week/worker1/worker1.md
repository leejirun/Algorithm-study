# 💡기본 개념

DFS, BFS의 개념을 학습하고 진행함

https://minhamina.tistory.com/36?category=837168

[BFS 수행 방법]

![bfs.gif](attachment:bc59c34d-037f-4c47-ba85-d66770dbf2f6:bfs.gif)

1. **a 노드(시작 노드)를 방문한다. (방문한 노드 체크)**
    - 큐에 방문된 노드를 삽입(enqueue)한다.
    - 초기 상태의 큐에는 시작 노드만이 저장
2. **큐에서 꺼낸 노드과 인접한 노드들을 큐에 추가한다. (모두 차례로 방문)**
    - 큐에서 꺼낸 노드를 방문한다.
    - 큐에서 꺼낸 노드과 인접한 노드들을 모두 방문한다.
    - 인접한 노드가 없다면 큐의 앞에서 노드를 꺼낸다(dequeue).
    - 큐에 방문된 노드를 삽입(enqueue)한다.
3. **큐가 공백 상태가 될 때까지 계속한다.**
- 미로 찾기 등 최단거리를 구해야 할 경우에 적합

[DFS 수행 방법]

![img.gif](attachment:9df0991c-a6fe-448e-8465-0091931c46ae:img.gif)

1.  노드(시작 노드)를 방문한다.
    - 방문한 노드는 방문했다고 표시한다.
2. a와 인접한 노드들을 차례로 순회한다.
    - a와 인접한 노드가 없다면 종료한다.

3, a와 이웃한 노드 b를 방문했다면, a와 인접한 또 다른 노드를 방문하기 전에 b의 이웃 노드들을 전부 방문해야 한다.

- b를 시작 정점으로 DFS를 다시 시작하여 b의 이웃 노드들을 방문한다. (단계 1 다시)
1. b의 분기를 전부 완벽하게 탐색했다면 다시 a에 인접한 정점들 중에서 아직 방문이 안 된 정점을 찾는다.
    - 즉, b의 분기를 전부 완벽하게 탐색한 뒤에야 a의 다른 이웃 노드를 방문할 수 있다는 뜻이다.
    - 아직 방문이 안 된 정점이 없으면 종료한다.
- 검색 대상 그래프가 정말 크다면 적합

# 💡Github URL

# 💡**문제 분석**

1. Depth-First Search (DFS)

- DFS는 백트래킹 전에 가능한 한 깊이 브랜치를 따라 탐색합니다. 재귀나 명시적 스택을 사용하여 구현
- **사용 사례** : 경로 찾기, 순환 감지, 그래프의 연결 구성 요소

2, **Breadth-First Search (BFS)**

- BFS는 큐를 사용하여 노드를 레벨별로 탐색하므로 가중치가 없는 그래프에서 최단 경로를 찾는 데 이상적
- **사용 사례**: 최단 경로 문제, 연결된 구성 요소 찾기

# 💡**알고리즘 설계**

1,BFS

필요한 변수

- 방문한 노드 체크 변수 ⇒ 무한 루프 방지용
- 간선 개수, 정점 개수, 탐색 시작할 정점의 개수

알고리즘 풀이 방법: 큐(Queue)를 이용 반복 형태로 구현 → 선입선출, 인접 리스트 사용

2,DFS

필요한 변수

- 방문한 노드 체크 변수 ⇒ 무한루프 방지용
- 간선개수, 정점 개수, 탐색 시작할 정점의 개수

알고리즘 풀이 방법

- 2차원 배열, 인접 리스트 사용
- 순환 호출 이용(재귀)하여 구현

# 💡코드

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//인접 리스트 저장 (각 노드의 연결된 이웃 정보)
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
```

# 💡 틀린 이유 & 틀린 부분 수정

1. 방문 배열 초기화 안함
- BFS 실행 전에 visited를 초기화하지 않으면, DFS에서 방문했던 노드를 BFS에서 다시 탐색할 수 없다.
- 따라서 DFS가 끝난 후 BFS를 실행하기 전에 **방문 기록을 초기화**해야 한다.

```go

1) 인접 리스트 초기화
	graph = make(map[int][]int)

2) 방문 배열 초기화
	visited = make([]bool, n+1)

3) 방문 배열 초기화 (DFS 후 BFS 실행)
visited = make([]bool, n+1)
```

1. 재귀 함수 호출 하지 않음 ⇒ 이 기능이 없어서 DFS는 첫 번째 노드에서 멈춰졌음

```go
for _, neighbor := range graph[node] {
		if !visited[neighbor] {
			DFS(neighbor) //방문하지 않은 노드면 DFS(neighbor)를 재귀 호출 
		}
	}
```

- 재귀 개념: 자기 자신을 호출하는 함수
- 이웃이 더 이상 없으면 완료 처리하기 위해 재귀를 사용함
- 이웃이 있으면 계속해서 함수를 호출해서 연결된 노드를 방문할 수 있게 함

1. 경로 탐색 방법 기억
- DFS는 한 경로를 끝까지 탐색 후 백트래킹
- BFS는 가까운 노드부터 탐색

# 💡느낀점

기본적인 개념부터 안다면 풀기 쉬운 문제여서, 정답률이 70-80으로 꽤나 높은 알고리즘 문제였다.