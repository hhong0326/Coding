package main

import (
	"sort"
	"strings"
)

var (
	pos = [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
)

func main() {
	solution([]string{"SL", "LR"})
}

func solution(grid []string) []int {

	newGrid := make([][]string, len(grid))

	for i, v := range grid {

		newGrid[i] = strings.Split(v, "")
	}

	n, m := len(grid), len(grid[0])

	visited := make([][][4]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([][4]bool, m)
	}
	cycles := func(x, y, d int) int {

		count := 0

		for {

			visited[x][y][d] = true
			count++

			// go through
			x = (x + pos[d][0] + n) % n
			y = (y + pos[d][1] + m) % m

			if newGrid[x][y] == "L" {
				d = (d + 4 - 1) % 4

			} else if newGrid[x][y] == "R" {
				d = (d + 1) % 4
			}

			if visited[x][y][d] {
				break
			}
		}

		return count
	}

	answer := []int{}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < 4; k++ {
				if !visited[i][j][k] {

					cycle := cycles(i, j, k)
					answer = append(answer, cycle)
				}
			}
		}
	}

	// ascending
	sort.Ints(answer)
	return answer
}

// 참고: https://hoi-hoon.github.io/algorithm/Level2-35/
