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
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	defer w.Flush()

	var n, m int

	fmt.Fscanf(r, "%d %d\n", &m, &n)

	boxes := make([][]int, n)

	for i := 0; i < n; i++ {
		boxes[i] = make([]int, m)

		for j := 0; j < m; j++ {
			fmt.Fscanf(r, "%d ", &boxes[i][j])
		}
	}

	q := [][]int{}

	total := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if boxes[i][j] == 1 || boxes[i][j] == 0 {
				total++
				if boxes[i][j] == 1 {

					q = append(q, []int{i, j})
				}
			}
		}
	}

	count := 0
	timer := 0

	// 최소 날짜 - bfs

	for len(q) > 0 {
		l := len(q)

		// 4방향 동시적으로 진행
		for k := 0; k < l; k++ {
			count++

			t := q[0]
			q = q[1:]

			for _, v := range pos {

				di := t[0] + v[0]
				dj := t[1] + v[1]

				if di < 0 || di >= n || dj < 0 || dj >= m || boxes[di][dj] != 0 {
					continue
				}

				boxes[di][dj] = 1
				q = append(q, []int{di, dj})

			}
		}

		timer++
	}

	if total == count {
		fmt.Fprintln(w, timer-1)
		return
	}

	fmt.Fprintln(w, -1)
}
