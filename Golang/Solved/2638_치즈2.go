package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	N, M    int
	dishes  [][]int
	visited [][]bool
	airs    [][]int
	pos     = [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
)

func main() {

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	defer w.Flush()

	fmt.Fscanf(r, "%d %d\n", &N, &M)

	dishes = make([][]int, N)
	airs = make([][]int, N)

	for i := 0; i < N; i++ {
		dishes[i] = make([]int, M)
		airs[i] = make([]int, M)
		for j := 0; j < M; j++ {
			fmt.Fscanf(r, "%d ", &dishes[i][j])
		}
	}

	timer := 0

	bfs := func() bool {

		q := [][]int{
			{0, 0},
		}

		for len(q) > 0 {

			x, y := q[0][0], q[0][1]
			q = q[1:]

			if visited[x][y] {
				continue
			}
			visited[x][y] = true

			for _, v := range pos {

				dx := x + v[0]
				dy := y + v[1]

				if dx < 0 || dx >= N || dy < 0 || dy >= M {
					continue
				}

				if dishes[dx][dy] == 1 {
					airs[dx][dy]++
				} else if !visited[dx][dy] {
					q = append(q, []int{dx, dy})
				}
			}

		}

		for i := 0; i < N; i++ {
			for j := 0; j < M; j++ {
				if airs[i][j] >= 2 {
					dishes[i][j] = 0
				}
			}
		}

		for i := 1; i < N-1; i++ {
			for j := 1; j < M-1; j++ {
				if dishes[i][j] == 1 {
					return false
				}
			}
		}

		return true
	}

	for {
		// 방문 초기화
		visited = make([][]bool, N)
		airs = make([][]int, N)
		for i := 0; i < N; i++ {
			visited[i] = make([]bool, M)
			airs[i] = make([]int, M)
		}
		bfs()
		if melt() {
			timer++
		} else {
			break
		}

	}

	fmt.Fprintln(w, timer)
}

func melt() (isMelted bool) {
	isMelted = false
	for i := 1; i < N-1; i++ {
		for j := 1; j < M-1; j++ {
			if airs[i][j] >= 2 {
				dishes[i][j] = 0
				isMelted = true
			}
		}
	}
	return
}
