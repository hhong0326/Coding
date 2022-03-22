package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	w             *bufio.Writer
	N, M, R, C, K int
	dice          Dice
)

type Dice struct {
	top    int
	front  int
	bottom int
	back   int
	left   int
	right  int
}

func main() {

	r := bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)

	defer w.Flush()

	fmt.Fscanf(r, "%d %d %d %d %d\n", &N, &M, &R, &C, &K)

	graph := make([][]int, N)

	for i := 0; i < N; i++ {
		graph[i] = make([]int, M)
		for j := 0; j < M; j++ {
			fmt.Fscanf(r, "%d ", &graph[i][j])
		}
	}

	ways := make([]int, K)
	for i := 0; i < K; i++ {
		fmt.Fscanf(r, "%d ", &ways[i])
	}

	for _, v := range ways {

		switch v {
		case 1: // i, j+1

			if checkNoway(R, C+1) {
				continue
			}

			C += 1

			t := dice.top
			dice.top = dice.left
			dice.left = dice.bottom
			dice.bottom = dice.right
			dice.right = t

			action(&graph)

		case 2: // i, j-1

			if checkNoway(R, C-1) {
				continue
			}

			C -= 1

			t := dice.top
			dice.top = dice.right
			dice.right = dice.bottom
			dice.bottom = dice.left
			dice.left = t

			action(&graph)
		case 3:
			// i-1, j
			if checkNoway(R-1, C) {
				continue
			}

			R -= 1

			t := dice.top
			dice.top = dice.back
			dice.back = dice.bottom
			dice.bottom = dice.front
			dice.front = t

			action(&graph)
		case 4:
			// i+1, j
			if checkNoway(R+1, C) {
				continue
			}

			R += 1

			t := dice.top
			dice.top = dice.front
			dice.front = dice.bottom
			dice.bottom = dice.back
			dice.back = t

			action(&graph)
		}

		fmt.Fprintln(w, dice.top)
	}
}

func action(graph *[][]int) {
	if (*graph)[R][C] == 0 {
		(*graph)[R][C] = dice.bottom
	} else {
		dice.bottom = (*graph)[R][C]
		(*graph)[R][C] = 0
	}
}

func checkNoway(i, j int) bool {

	if i < 0 || i >= N || j < 0 || j >= M {
		return true
	}

	return false
}
