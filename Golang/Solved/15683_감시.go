package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	w    *bufio.Writer
	room [][]int
	N, M int
	pos  = [][]int{
		{0, -1},
		{0, 1},
		{-1, 0},
		{1, 0},
	}
	walls   [][]int
	cameras [][]int
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)

	defer w.Flush()

	fmt.Fscanf(r, "%d %d\n", &N, &M)

	room = make([][]int, N)
	for i := 0; i < N; i++ {
		room[i] = make([]int, M)
		for j := 0; j < M; j++ {
			fmt.Fscanf(r, "%d ", &room[i][j])
			if room[i][j] == 6 {
				walls = append(walls, []int{i, j})
			} else if room[i][j] >= 1 && room[i][j] <= 5 {
				cameras = append(cameras, []int{room[i][j], i, j})
			}
		}
	}

	for i := 0; i < len(cameras); i++ {

	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if room[i][j] == 5 {
				five(i, j)
			} else {
				switch room[i][j] {
				case 1:
					// 4방향 일직선
					one(i, j)
				case 2:
					// 2방향 직선
					two(i, j)
				case 3:
					// 4방향 직각선
					three(i, j)
				case 4:
					// 4방향 삼각선
					four(i, j)
				}
			}
		}
	}

	count := 0
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if room[i][j] == 0 {
				count += 1
			}
		}
	}

	fmt.Fprintln(w, count)
}

func move_up(up, y int) int {
	count := 0
	for up >= 0 {
		if room[up][y] == 6 {
			break
		}

		if !checkVisted(up, y) {
			count++
		}
		up--

	}
	return count
}

func move_right(x, right int) int {
	count := 0
	for right < N {
		if room[x][right] == 6 {
			break
		}
		if !checkVisted(x, right) {
			count++
		}
		right++

	}
	return count
}

func move_down(down, y int) int {
	count := 0
	for down < N {
		if room[down][y] == 6 {
			break
		}
		if !checkVisted(down, y) {
			count++
		}
		down++

	}
	return count
}
func move_left(x, left int) int {
	count := 0
	for left >= 0 {
		if room[x][left] == 6 {
			break
		}

		if !checkVisted(x, left) {
			count++
		}
		left--

	}
	return count
}

func one(x, y int) {

	up, down, left, right := x-1, x+1, y-1, y+1

	counts := []int{0, 0, 0, 0}
	counts[0] = move_up(up, y)

	counts[1] = move_down(down, y)

	counts[2] = move_left(x, left)

	counts[3] = move_right(x, right)

	// up down left right
	max := 0
	maxIndex := -1
	for i := 0; i < 4; i++ {
		if max < counts[i] {
			max = counts[i]
			maxIndex = i
		}
	}

	switch maxIndex {
	case 0:
		up = x - 1
		for up >= 0 {
			if room[up][y] == 6 {
				break
			}

			if room[up][y] == 0 {

				room[up][y] = 7
			}
			up--
		}
	case 1:
		down = x + 1
		for down < N {
			if room[down][y] == 6 {
				break
			}
			if room[down][y] == 0 {
				room[down][y] = 7
			}
			down++
		}
	case 2:
		left = y - 1
		for left >= 0 {
			if room[x][left] == 6 {
				break
			}
			if room[x][left] == 0 {
				room[x][left] = 7
			}
			left--
		}
	case 3:
		right = y + 1
		for right < M {
			if room[x][right] == 6 {
				break
			}
			if room[x][right] == 0 {
				room[x][right] = 7
			}
			right++
		}
	}
}

func two(x, y int) {

	up, down, left, right := x-1, x+1, y-1, y+1

	c1 := move_up(up, y)
	c1 += move_down(down, y)
	for down < N {
		if room[down][y] == 6 {
			break
		}

		if !checkVisted(down, y) {
			c1++
		}
		down++
	}

	c2 := move_left(x, left)
	c2 += move_right(x, right)

	if c1 > c2 {
		up, down = x-1, x+1
		for up >= 0 {
			if room[up][y] == 6 {
				break
			}
			if room[up][y] == 0 {
				room[up][y] = 7
			}
			up--
		}
		for down < N {
			if room[down][y] == 6 {
				break
			}
			if room[down][y] == 0 {
				room[down][y] = 7
			}
			down++
		}
	} else {
		left, right = y-1, y+1
		for left >= 0 {
			if room[x][left] == 6 {
				break
			}
			if room[x][left] == 0 {
				room[x][left] = 7
			}
			left--
		}
		for right < M {
			if room[x][right] == 6 {
				break
			}
			if room[x][right] == 0 {
				room[x][right] = 7
			}
			right++
		}
	}
}

func three(x, y int) {

	up, down, left, right := x-1, x+1, y-1, y+1

	counts := []int{0, 0, 0, 0}
	counts[0] = move_up(up, y)
	counts[0] += move_right(x, right)

	counts[1] = move_right(x, right)
	counts[1] += move_down(down, y)

	counts[2] = move_down(down, y)
	counts[2] += move_left(x, left)

	counts[3] = move_left(x, left)
	counts[3] += move_up(up, y)

	max := 0
	maxIndex := -1
	for i := 0; i < 4; i++ {
		if max < counts[i] {
			max = counts[i]
			maxIndex = i
		}
	}

	switch maxIndex {
	case 0:
		up, right = x-1, y+1
		for up >= 0 {
			if room[up][y] == 6 {
				break
			}
			if room[up][y] == 0 {
				room[up][y] = 7
			}
			up--
		}
		for right < M {
			if room[x][right] == 6 {
				break
			}
			if room[x][right] == 0 {
				room[x][right] = 7
			}
			right++
		}
	case 1:
		right, down = y+1, x+1
		for right < M {
			if room[x][right] == 6 {
				break
			}
			if room[x][right] == 0 {
				room[x][right] = 7
			}
			right++
		}
		for down < N {
			if room[down][y] == 6 {
				break
			}
			if room[down][y] == 0 {
				room[down][y] = 7
			}
			down++
		}
	case 2:

		down, left = x+1, y-1
		for down < N {
			if room[down][y] == 6 {
				break
			}
			if room[down][y] == 0 {
				room[down][y] = 7
			}
			down++
		}
		for left >= 0 {
			if room[x][left] == 6 {
				break
			}
			if room[x][left] == 0 {
				room[x][left] = 7
			}
			left--
		}
	case 3:
		left, up = y-1, x-1
		for left >= 0 {
			if room[x][left] == 6 {
				break
			}
			if room[x][left] == 0 {
				room[x][left] = 7
			}
			left--
		}
		for up >= 0 {
			if room[up][y] == 6 {
				break
			}
			if room[up][y] == 0 {
				room[up][y] = 7
			}

			up--
		}
	}
}

func four(x, y int) {

	up, down, left, right := x-1, x+1, y-1, y+1

	counts := []int{0, 0, 0, 0}

	counts[0] = move_up(x-1, y)
	counts[0] += move_right(x, y+1)
	counts[0] += move_down(x+1, y)

	counts[1] = move_right(x, y+1)
	counts[1] += move_down(x+1, y)
	counts[1] += move_left(x, y-1)

	counts[2] = move_down(x+1, y)
	counts[2] += move_left(x, y-1)
	counts[2] += move_up(x-1, y)

	counts[3] = move_left(x, y-1)
	counts[3] += move_up(x-1, y)
	counts[3] += move_right(x, y+1)

	max := 0
	maxIndex := -1
	for i := 0; i < 4; i++ {
		if max < counts[i] {
			max = counts[i]
			maxIndex = i
		}
	}

	switch maxIndex {
	case 0:
		up, right, down = x-1, y+1, x+1
		for up >= 0 {
			if room[up][y] == 6 {
				break
			}
			if room[up][y] == 0 {
				room[up][y] = 7
			}
			up--
		}
		for right < M {
			if room[x][right] == 6 {
				break
			}
			if room[x][right] == 0 {
				room[x][right] = 7
			}
			right++
		}
		for down < N {
			if room[down][y] == 6 {
				break
			}
			if room[down][y] == 0 {
				room[down][y] = 7
			}
			down++
		}
	case 1:
		right, down, left = y+1, x+1, y-1
		for right < M {
			if room[x][right] == 6 {
				break
			}
			if room[x][right] == 0 {
				room[x][right] = 7
			}
			right++
		}
		for down < N {
			if room[down][y] == 6 {
				break
			}
			if room[down][y] == 0 {
				room[down][y] = 7
			}
			down++
		}
		for left >= 0 {
			if room[x][left] == 6 {
				break
			}
			if room[x][left] == 0 {
				room[x][left] = 7
			}
			left--
		}
	case 2:

		down, left, up = x+1, y-1, x-1
		for down < N {
			if room[down][y] == 6 {
				break
			}
			if room[down][y] == 0 {
				room[down][y] = 7
			}
			down++
		}
		for left >= 0 {
			if room[x][left] == 6 {
				break
			}
			if room[x][left] == 0 {
				room[x][left] = 7
			}
			left--
		}
		for up >= 0 {
			if room[up][y] == 6 {
				break
			}
			if room[up][y] == 0 {
				room[up][y] = 7

			}
			up--
		}
	case 3:
		left, up, right = y-1, x-1, y+1
		for left >= 0 {
			if room[x][left] == 6 {
				break
			}
			if room[x][left] == 0 {
				room[x][left] = 7
			}
			left--
		}
		for up >= 0 {
			if room[up][y] == 6 {
				break
			}
			if room[up][y] == 0 {
				room[up][y] = 7
			}
			up--
		}
		for right < M {
			if room[x][right] == 6 {
				break
			}
			if room[x][right] == 0 {
				room[x][right] = 7
			}
			right++
		}

	}
}
func five(x, y int) {

	up, down, left, right := x-1, x+1, y-1, y+1

	for up >= 0 {
		if room[up][y] == 6 {
			break
		}
		if room[up][y] == 0 {
			room[up][y] = 7
		}
		up--
	}
	for down < N {
		if room[down][y] == 6 {
			break
		}
		if room[down][y] == 0 {
			room[down][y] = 7
		}
		down++
	}
	for left >= 0 {
		if room[x][left] == 6 {
			break
		}
		if room[x][left] == 0 {
			room[x][left] = 7
		}
		left--
	}
	for right < M {
		if room[x][right] == 6 {
			break
		}
		if room[x][right] == 0 {
			room[x][right] = 7
		}
		right++
	}

}

func checkVisted(i, j int) bool {
	if room[i][j] != 7 {
		return true
	}

	return false
}
