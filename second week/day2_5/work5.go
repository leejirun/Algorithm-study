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
