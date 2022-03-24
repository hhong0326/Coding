package main

import (
	"bufio"
	"fmt"
	"os"
)

//
type Cleaner struct {
	x  int
	xx int
}

var (
	R, C, T int
	cleaner Cleaner
	pos     = [][]int{
		{0, -1},
		{0, 1},
		{-1, 0},
		{1, 0},
	}
	dusts [][]int
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	defer w.Flush()

	var room [][]int
	fmt.Fscanf(r, "%d %d %d\n", &R, &C, &T)

	room = make([][]int, R)

	sum := 0
	for i := 0; i < R; i++ {
		room[i] = make([]int, C)
		for j := 0; j < C; j++ {
			fmt.Fscanf(r, "%d ", &room[i][j])
			sum += room[i][j]
			if room[i][j] == -1 && cleaner.x == 0 {
				cleaner.x = i
				cleaner.xx = i + 1
			} else if room[i][j] > 0 {
				dusts = append(dusts, []int{i, j})
			}
		}
	}

	for i := 0; i < T; i++ {
		spreadMicro(&room)

	}

	result := 0
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			result += room[i][j]
		}
	}

	fmt.Fprintln(w, result+2)
}

// T번
func moveAir(room *[][]int) {

	// reverse clock

	i := cleaner.x
	for i > 0 {
		if (*room)[i][0] != -1 {

			(*room)[i][0] = (*room)[i-1][0]

		}
		i--
	}

	j := 0
	for j < len((*room)[0])-1 {
		(*room)[0][j] = (*room)[0][j+1]
		j++
	}

	i = 0
	for i < cleaner.x {
		(*room)[i][len((*room)[0])-1] = (*room)[i+1][len((*room)[0])-1]
		i++

	}

	j = len((*room)[0]) - 1
	for j > 0 {
		if (*room)[cleaner.x][j-1] == -1 {
			(*room)[cleaner.x][j] = 0
		} else {
			(*room)[cleaner.x][j] = (*room)[cleaner.x][j-1]
		}
		j--
	}

	// clock

	i = cleaner.xx
	for i < len((*room))-1 {
		if (*room)[i][0] != -1 {
			(*room)[i][0] = (*room)[i+1][0]
		}
		i++
	}

	j = 0
	for j < len((*room)[0])-1 {
		(*room)[len((*room))-1][j] = (*room)[len((*room))-1][j+1]
		j++
	}

	i = len((*room)) - 1
	for i > cleaner.xx {
		(*room)[i][len((*room)[0])-1] = (*room)[i-1][len((*room)[0])-1]
		i--
	}

	j = len((*room)[0]) - 1
	for j > 0 {
		if (*room)[cleaner.xx][j-1] == -1 {
			(*room)[cleaner.xx][j] = 0
		} else {
			(*room)[cleaner.xx][j] = (*room)[cleaner.xx][j-1]
		}
		j--
	}
}

func spreadMicro(r *[][]int) {

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {

			if (*r)[i][j] > 0 {
				x, y := i, j
				v := (*r)[x][y] / 5
				count := 0
				for i := 0; i < len(pos); i++ {

					dx := x + pos[i][0]
					dy := y + pos[i][1]

					if dx < 0 || dx >= R || dy < 0 || dy >= C || (*r)[dx][dy] == -1 {
						continue
					}

					// spread
					(*r)[dx][dy] += v
					count++
				}
				(*r)[x][y] = (*r)[x][y] - v*count

			}
		}
	}

	moveAir(r)
}
