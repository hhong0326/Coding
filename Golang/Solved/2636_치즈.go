package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	pos = [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	dishes  [][]int
	visited [][]bool
	result  int
	timer   int
)

func main() {

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	defer w.Flush()

	var N, M int
	fmt.Fscanf(r, "%d %d\n", &N, &M)

	dishes = make([][]int, N)
	visited = make([][]bool, N)

	for i := 0; i < N; i++ {
		dishes[i] = make([]int, M)
		visited[i] = make([]bool, M)
		for j := 0; j < M; j++ {
			fmt.Fscanf(r, "%d ", &dishes[i][j])

		}
	}

	bfs := func() bool {

		visited[0][0] = true
		q := [][]int{
			{0, 0},
		}

		timer++
		count := 0
		for len(q) > 0 {

			x, y := q[0][0], q[0][1]

			q = q[1:]

			for k := 0; k < len(pos); k++ {

				dx := x + pos[k][0]
				dy := y + pos[k][1]

				// 겉이 공기이므로 바깥에서 탐색 0 -> 1
				// 1에 닿으면 녹이고 count

				if dx < 0 || dx >= N || dy < 0 || dy >= M || visited[dx][dy] {
					continue
				}

				visited[dx][dy] = true

				if dishes[dx][dy] == 1 {
					dishes[dx][dy] = 0
					count++
				} else {
					q = append(q, []int{dx, dy})

				}

			}
		}

		if count == 0 {
			fmt.Fprintf(w, "%d\n%d\n", timer-1, result)
			return true
		} else {
			result = count
			return false
		}

	}

	for {
		if bfs() {
			break
		}

		// 방문 초기화
		visited = make([][]bool, N)
		for i := 0; i < N; i++ {
			visited[i] = make([]bool, M)
		}
	}
}
