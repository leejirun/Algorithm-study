# 💡Github URL
https://github.com/leejirun/Algorithm-study/tree/main/third_week/worker5

# 💡**문제 분석**
문제 조건 : 
- 한 걸음 더 나아가기: X + 1 
- 한 걸음 뒤로 물러나기:X - 1
- 점프하다:2 * X

# 💡**알고리즘 설계**
BFS를 큐와 함께 사용하여 가능한 위치를 탐색하기
- queue: 선입선출

# 💡코드
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


# 💡틀린 부분
1.도달 체크 시 return 하지 않음
if current == k {
    return steps[current]
}
- 목표 위치에 도달하면 걸음 수를 반환해야하는데 하지 않았다.
- 문제에서 가장 짧은 경로를 찾으면 검색을 즉시 중단할 수 있기에 반드시 return 해줘야 한다!(실수)

2.걸음 수 세기 
steps[next] = steps[current] + 1
=> 이 부분에서 steps[current]만 했음


# 💡풀이 방법
bfs + quueu + 시간 복잡도

ex) 입력: 5, 17 가정한다.
[단계별 BFS 실행]
5-> 10 -> 9 -> 8 -> 16 -> 17: 총 4단계


# 💡느낀점
- bfs는 가장 짧은 경로를 찾는 문제에 적합하다.
- dfs는 최단 경로 보단 모든 위치를 다 거쳐야 하기 때문에 이 문제에는 적합하지 않다! 하지만, 풀라면 풀 수는 있음! 오래 걸릴 뿐
- 손으로 그림을 그리면서 했는데, 잘 못 생각해서 한참을 걸리면서 풀었다...