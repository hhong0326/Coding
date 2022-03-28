package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	N, M, R int
	graph   [301][301]int
)

func main() {

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	defer w.Flush()

	fmt.Fscanf(r, "%d %d %d\n", &N, &M, &R)

	// graph = make([][]int, N)

	for i := 0; i < N; i++ {
		// graph[i] = make([]int, M)
		for j := 0; j < M; j++ {
			fmt.Fscanf(r, "%d ", &graph[i][j])
		}
	}

	for R > 0 {

		// for k := 0; k < int(math.Min(float64(N), float64(M)))/2; k++ {

		// 	rotate(k)
		// }

		rotate2()
		R--
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {

			fmt.Fprintf(w, "%d ", graph[i][j])
		}
		fmt.Fprintf(w, "\n")
	}
}

func rotate2() {
	temp := [301][301]int{}
	// temp := make([][]int, N)
	// for i := 0; i < N; i++ {
	// 	temp[i] = make([]int, M)
	// }

	for r := 0; r < int(math.Min(float64(N), float64(M)))/2; r++ {
		x1, y1, x2, y2 := r, r, N-1-r, M-1-r
		//좌
		for i := y2 - 1; i >= y1; i-- {
			temp[x1][i] = graph[x1][i+1]
		}
		//하
		for i := x1 + 1; i <= x2; i++ {
			temp[i][y1] = graph[i-1][y1]
		}
		//우
		for i := y1 + 1; i <= y2; i++ {
			temp[x2][i] = graph[x2][i-1]
		}
		//상
		for i := x2 - 1; i >= x1; i-- {
			temp[i][y2] = graph[i+1][y2]
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			graph[i][j] = temp[i][j]
		}
	}
}

func rotate(k int) {

	i, j, count := k, k, k
	temp := graph[i][j]

	for j < len(graph[0])-1-count {
		graph[i][j] = graph[i][j+1]
		j++
	}
	for i < len(graph)-1-count {
		graph[i][j] = graph[i+1][j]
		i++
	}
	for j > count {
		graph[i][j] = graph[i][j-1]
		j--
	}
	for i > count {
		if i == count+1 {
			graph[i][j] = temp
		} else {
			graph[i][j] = graph[i-1][j]
		}
		i--
	}
}
