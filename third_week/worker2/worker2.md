# 💡Github URL


# 💡**문제 분석**
- 무향 그래프에서 연결된 구성요소의 수를 찾기
- 첫번째 입력 줄: n(정점의 개수), m(모서리의 개수)
- m개 줄에는 모서리(연결된 노드의 쌍)가 나열됨
- 연결된 노드의 별도 그룹이 있는지 확인하기

# 💡**알고리즘 설계**
- 무형 그래프 => DFS나 BFS를 사용해야겠다.
- DFS를 적용하기 위해 인접 리스트를 이용해야겠다.

# 💡코드
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	adjList  [][]int
	visited  []bool
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
	n, _ := strconv.Atoi(parts[0])//정점
	m, _ := strconv.Atoi(parts[1])//간선

	adjList = make([][]int, n+1)//인접행렬
	visited = make([]bool, n+1)//방문확인

	for i := 0; i < m; i++ {
		line, _ := reader.ReadString('\n')
		parts := strings.Fields(line)
		u, _ := strconv.Atoi(parts[0])
		v, _ := strconv.Atoi(parts[1])
		adjList[u] = append(adjList[u], v)
		adjList[v] = append(adjList[v], u)
	}

    //연결된 개수 찾기
	components := 0
	for i := 1; i <= n; i++ {
        //재귀 사용
		if !visited[i] {
			dfs(i)
			components++
		}
	}

	fmt.Println(components)
}


# 💡기본 개념
- 시간 복잡도: O(N + M) (각 노드와 에지를 한 번씩 방문)
- O(N): 각 노드는 한 번씩 방문한다.
- O(M): 각 모서리는 한 번씩 처리한다.

# 💡다른 풀이
- 나는 DFS 방식을 적용했는데, 내 코드를 BFS 방식으로 변경해보라고 GPT에게 명령했다.

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

func bfs(start int) {
	queue := []int{start}
	visited[start] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, neighbor := range adjList[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
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
			bfs(i)
			components++
		}
	}

	fmt.Println(components)
}

정리)
- BFS 방식은 큐를 사용해서 연결 요소를 레벨로 탐색할 수 있다.
- DFS 방식과 동일하게 인접 리스트를 만들었다.
- 큐를 사용해서 방문하지 않았던 노드에서 BFS를 통해 방문하게 한다.

# 💡느낀점
어제 시간을 내서 그래프 개념을 공부하기도 했고, 어제 틀렸던 부분을 한번 더 풀어보고 시작했었다. 그래서 오늘은 문제를 푸는 데 한결 쉬웠다. DFS는 재귀 함수를 사용하고, BFS는 Queue를 사용한다는 것!을 이제 정확히 알겠다. 
그리고 문제를 풀면서 인접 리스트인 `adjList = make([][]int, n+1)`을 n+1로 설정해야하나 n으로 해야하나 엄청난 고민을 했다. 근데, test로 돌려보니
'Index Out of Range error'가 났고, 문제에서 정점은 1부터 시작이라고 되어있음을 확인하고 n+1로 돌려서 코드가 잘 돌아감을 확인했다.