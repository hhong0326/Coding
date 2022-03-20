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
	n, m   int
	temp   [][]int
	answer int
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	defer w.Flush()

	fmt.Fscanf(r, "%d %d\n", &n, &m)

	rooms := make([][]int, n)

	for i := 0; i < n; i++ {
		rooms[i] = make([]int, m)

		for j := 0; j < m; j++ {
			fmt.Fscanf(r, "%d ", &rooms[i][j])
		}
	}

	// 벽을 막는 모든 경우의 수 -> BFS 탐색 -> 2의 확장 -> 갯수 결과

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if rooms[i][j] == 0 {
				temp = rooms

				temp[i][j] = 1
				makeWall(1)
				temp[i][j] = 0

			}
		}
	}

}

func makeWall(count int) {
	if count == 3 {
		bfs()
		return
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if temp[i][j] == 0 {

				temp[i][j] = 1
				makeWall(count + 1)
				temp[i][j] = 0
			}
		}
	}
}

func bfs() {

	vmap := temp

	q := [][]int{}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if vmap[i][j] == 2 {
				q = append(q, []int{i, j})
			}
		}
	}

	for len(q) > 0 {

		v := q[0]

		for i := 0; i < len(pos); i++ {
			dx := v[0] + pos[i][0]
			dy := v[1] + pos[i][1]

			if dx >= 0 && dx < n && dy >= 0 && dy < m {
				if vmap[dx][dy] == 0 {
					vmap[dx][dy] = 2
					q = append(q, []int{dx, dy})
				}
			}
		}
	}

	cnt := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if vmap[i][j] == 0 {
				cnt++
			}
		}
	}

}
