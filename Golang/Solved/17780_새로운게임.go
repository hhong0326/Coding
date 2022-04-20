package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	N, K       int
	board      [][]int
	qboard     [][][]*horseInfo
	horsesInfo [][]int
	// → ← ↑ ↓
	pos = [][]int{
		{0, 1},
		{0, -1},
		{-1, 0},
		{1, 0},
	}
	horseInfos []*horseInfo
)

type horseInfo struct {
	x    int
	y    int
	move int
}

func main() {

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	defer w.Flush()

	fmt.Fscanf(r, "%d %d\n", &N, &K)

	board = make([][]int, N)
	qboard = make([][][]*horseInfo, 0)
	for i := 0; i < N; i++ {
		board[i] = make([]int, N)
		qboard = append(qboard, [][]*horseInfo{})
		for j := 0; j < N; j++ {
			fmt.Fscanf(r, "%d ", &board[i][j])
		}
	}

	horseInfos = make([]*horseInfo, K)

	for i := 0; i < K; i++ {
		horseInfos[i] = &horseInfo{}
		fmt.Fscanf(r, "%d %d %d\n", &horseInfos[i].x, &horseInfos[i].y, &horseInfos[i].move)
		qboard[horseInfos[i].x][horseInfos[i].y] = append(qboard[horseInfos[i].x][horseInfos[i].y], horseInfos[i])
	}

	fmt.Println(qboard)
	// horsesInfo
	// 행, 열의 번호, 이동 방향

	// directions
	// 1 2 3 4
	// → ← ↑ ↓

	// white 0 -> stack
	// red 1 -> to queue
	// blue 2 or out of board -> opposite direction

	// 모든 말이 같은 포지션을 가질 때 까지 반복

	turns := 0

	for {

		if turns > 1000 {
			fmt.Fprintln(w, -1)
			return
		}

		// 종료
		for i := 0; i < len(horseInfos)-1; i++ {
			if horseInfos[i].x != horseInfos[i+1].x ||
				horseInfos[i+1].y != horseInfos[i+1].y {
				break
			}

			if i == len(horseInfos)-1 {
				fmt.Fprintln(w, turns)
				return
			}
		}

		// 턴마다 말 움직이기
		for _, h := range horseInfos {

			switch h.move {
			case 1, 2:
				h.y = h.y + pos[h.move-1][1]
			case 3, 4:
				h.x = h.x + pos[h.move-1][0]
			}

			// stack

			// blue 2 or out of board
			if board[h.x][h.y] == 2 || (h.x < 0 || h.y < 0 || h.x >= N || h.y >= N) {
				// direction change
				newd := oppositeDirection(h.move)
				h.move = newd

				var x, y int
				switch h.move {
				case 1, 2:
					y = h.y + pos[h.move-1][1]
				case 3, 4:
					x = h.x + pos[h.move-1][0]
				}

				if board[x][y] != 2 {
					h.x = x
					h.y = y
				}
			}
			// red 1

		}

		turns++

		fmt.Println(horsesInfo)

	}
}

func oppositeDirection(d int) int {

	if d == 1 {
		return 2
	} else if d == 2 {
		return 1
	} else if d == 3 {
		return 4
	} else {
		return 3
	}
}
