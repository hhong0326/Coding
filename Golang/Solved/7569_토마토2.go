package main

import (
	"bufio"
	"fmt"
	"os"
)

var (

	// 6방향
	pos = [][]int{
		{0, 0, 1},
		{0, 0, -1},
		{0, 1, 0},
		{0, -1, 0},
		{1, 0, 0},  // up
		{-1, 0, 0}, // down
	}
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	defer w.Flush()

	var n, m, h int

	fmt.Fscanf(r, "%d %d %d\n", &m, &n, &h)

	boxes := make([][][]int, h)

	for k := 0; k < h; k++ {
		boxes[k] = make([][]int, n)
		for i := 0; i < n; i++ {
			boxes[k][i] = make([]int, m)
			for j := 0; j < m; j++ {
				fmt.Fscanf(r, "%d ", &boxes[k][i][j])

			}
		}
	}

	q := [][]int{}

	total := 0
	for k := 0; k < h; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {

				if boxes[k][i][j] == 1 || boxes[k][i][j] == 0 {
					total++
					if boxes[k][i][j] == 1 {

						q = append(q, []int{k, i, j})
					}

				}
			}
		}
	}

	count := 0
	timer := 0

	for len(q) > 0 {
		l := len(q)

		for k := 0; k < l; k++ {
			count++

			t := q[0]
			q = q[1:]

			for _, v := range pos {

				dk := t[0] + v[0]
				di := t[1] + v[1]
				dj := t[2] + v[2]

				if di < 0 || di >= n || dj < 0 || dj >= m || dk < 0 || dk >= h || boxes[dk][di][dj] != 0 {
					continue
				}

				boxes[dk][di][dj] = 1
				q = append(q, []int{dk, di, dj})

			}
		}
		fmt.Println(boxes)
		timer++
	}

	if total == count {
		fmt.Fprintln(w, timer-1)
		return
	}

	fmt.Fprintln(w, -1)
}
