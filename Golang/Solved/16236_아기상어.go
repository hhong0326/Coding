package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type FishPos struct {
	dist int
	x    int
	y    int
}

var (
	n        int
	fishTank [][]int
	visited  [][]bool
	pos      = [][]int{
		{0, -1},
		{0, 1},
		{-1, 0},
		{1, 0},
	}
	w *bufio.Writer
)

func bfs(x, y, size int) *FishPos {

	visited = make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, n)
	}

	fishPos := []*FishPos{}

	q := []*FishPos{
		{0, x, y},
	}

	visited[x][y] = true

	for len(q) > 0 {

		dis, i, j := q[0].dist, q[0].x, q[0].y

		q = q[1:]

		for k := 0; k < len(pos); k++ {
			dx := i + pos[k][0]
			dy := j + pos[k][1]

			if dx < 0 || dx >= n || dy < 0 || dy >= n || visited[dx][dy] {
				continue
			}
			if fishTank[dx][dy] != 0 && size > fishTank[dx][dy] {
				// be able to eat
				fishPos = append(fishPos, &FishPos{dis + 1, dx, dy})
				visited[dx][dy] = true
			} else if size == fishTank[dx][dy] || fishTank[dx][dy] == 0 {
				// be able to pass
				q = append(q, &FishPos{dis + 1, dx, dy})
				visited[dx][dy] = true
			}

		}

	}

	if len(fishPos) == 0 {
		return nil
	}

	// close - top - left
	sort.Slice(fishPos, func(i, j int) bool {
		if fishPos[i].dist == fishPos[j].dist {
			if fishPos[i].x == fishPos[j].x {
				return fishPos[i].y < fishPos[j].y
			} else {
				return fishPos[i].x < fishPos[j].x
			}
		}
		return fishPos[i].dist < fishPos[j].dist
	})

	return fishPos[0]
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)

	defer w.Flush()

	fmt.Fscanln(r, &n)
	var X, Y int

	fishTank = make([][]int, n)
	for i := 0; i < n; i++ {
		fishTank[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscanf(r, "%d ", &fishTank[i][j])
			if fishTank[i][j] == 9 {
				X = i
				Y = j
			}
		}
	}

	timer := 0
	size, ate := 2, 0

	for {

		fishTank[X][Y] = 0
		shark := bfs(X, Y, size)
		if shark == nil {

			fmt.Fprintln(w, timer)
			break
		}

		ate++
		if ate == size {
			size++
			ate = 0
		}

		dis, x, y := shark.dist, shark.x, shark.y
		timer += dis
		X, Y = x, y
	}

}
