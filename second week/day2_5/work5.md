## 💡문제 분석 요약
1.총 N명의 사람을 N/2명씩 두 팀으로 나눠야 한다.
2.팀의 능력치는 팀원 간의 시너지 값(S[i][j] + S[j][i])의 합이다.
3.두 팀의 능력치 차이를 최소화해야 한다.

## 💡알고리즘 설계
1.백트래킹(Brute Force)
- N명의 사람 중 N/2명을 선택해 한 팀을 구성하면, 나머지는 자동으로 상대 팀이 된다.
- 모든 가능한 팀 조합을 생성하고 각 팀의 능력치를 계산하여 차이의 최솟값을 구한다.
- N이 최대 20이므로, 최악의 경우 20C10 (약 18만)개의 경우를 탐색해야 하는데, 이는 브루트포스로 충분히 가능하다.

2.능력치 계산
-선택된 팀원의 능력치는 모든 조합 (i, j)에 대해 S[i][j] + S[j][i]를 더한 값이다.
-두 팀의 능력치 차이를 구하고, 최솟값을 갱신한다.


## 💡코드
`
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	N        int
	S        [][]int
	minDiff  int = math.MaxInt32
	selected []bool
)

func calculateAbility(team []int) int {
	ability := 0
	for i := 0; i < len(team); i++ {
		for j := i + 1; j < len(team); j++ {
			ability += S[team[i]][team[j]] + S[team[j]][team[i]]
		}
	}
	return ability
}

func backtrack(index, count int) {
	if count == N/2 {
		startTeam, linkTeam := []int{}, []int{}
		for i := 0; i < N; i++ {
			if selected[i] {
				startTeam = append(startTeam, i)
			} else {
				linkTeam = append(linkTeam, i)
			}
		}

		startAbility := calculateAbility(startTeam)
		linkAbility := calculateAbility(linkTeam)
		diff := int(math.Abs(float64(startAbility - linkAbility)))
		if diff < minDiff {
			minDiff = diff
		}
		return
	}

	for i := index; i < N; i++ {
		if !selected[i] {
			selected[i] = true
			backtrack(i+1, count+1)
			selected[i] = false
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &N)

	S = make([][]int, N)
	selected = make([]bool, N)
	for i := 0; i < N; i++ {
		S[i] = make([]int, N)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		nums := strings.Split(line, " ")
		for j := 0; j < N; j++ {
			S[i][j], _ = strconv.Atoi(nums[j])
		}
	}

	backtrack(0, 0)
	fmt.Println(minDiff)
}
`

## 💡 정리
1.입력 처리
-bufio.Reader를 사용해 빠르게 입력을 받는다.
-S[i][j]를 2D 슬라이스([][]int)에 저장

2.백트래킹을 통한 팀 조합 생성 (backtrack)
-selected[i]가 true이면 스타트 팀, false이면 링크 팀.
-N/2명을 선택하면, 나머지는 자동으로 다른 팀이 된다.
-모든 조합을 확인하고 능력치 차이의 최솟값을 갱신.

3.능력치 계산 (calculateAbility)
-모든 (i, j) 쌍을 탐색하여 능력치의 합을 구한다.

4.결과 출력
-최솟값(minDiff)을 출력

## 💡틀린 부분
1.능력치 계산 시 S[i][j] + S[j][i]를 고려하지 않았다.
-S[i][j]만 더하고 S[j][i]를 더하지 않았기 때문에 능력치를 잘못계산해버렸다.
-S[1][2] = 1, S[2][1] = 4라면, 올바른 능력치는 1 + 4 = 5여야 하지만, S[1][2]만 더해서 1이 되어버릴 수 있다.

2.자기 자신을 포함하여 잘못된 값을 더했다.
-i == j인 경우도 포함되므로 S[i][i](항상 0)을 더하게 되어 의미 없는 연산이 포함됨.


-내가 푼 코드
`
func calculateAbility(team []int) int {
	ability := 0
	for i := 0; i < len(team); i++ {
		for j := 0; j < len(team); j++ { // 잘못된 부분: i == j인 경우도 포함됨
			ability += S[team[i]][team[j]]
		}
	}
	return ability
}
`

-올바른 코드
`
func calculateAbility(team []int) int {
	ability := 0
	for i := 0; i < len(team); i++ {
		for j := i + 1; j < len(team); j++ { // i == j가 아니라 (i, j) 쌍을 고려
			ability += S[team[i]][team[j]] + S[team[j]][team[i]] // S[i][j] + S[j][i]
		}
	}
	return ability
}
`
정리)
1.S[i][j] + S[j][i]를 모두 더한다.
예를 들어, S[1][2] = 1, S[2][1] = 4라면 1 + 4 = 5로 정확한 값을 구할 수 있다/
2.i == j일 때 연산하지 않는다.
i == j인 경우를 제외하여 자기 자신과의 능력치를 더하지 않음
3.(i, j) 쌍을 한 번만 계산한다.
for j := i + 1로 i < j인 경우만 고려하여 중복 계산을 방지


## 💡개념
1.시간 복잡도 분석
-N이 최대 20일 때, 백트래킹 탐색은 O(2^N)이지만, N/2개를 선택하는 문제이므로 O(NC(N/2)) = O(184,756) ≈ O(2^20)
-능력치 계산 O(N^2), 하지만 팀이 절반씩 나뉘므로 O(N^2/2).
-최악의 경우 O(2^N * N^2/2) = O(2^20 * 10^2) ≈ 10^7 → 가능.


## 💡느낀점