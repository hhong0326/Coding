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
	counts  []int
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

	// R 반복적인 패턴 파악
	// N-2, M-2 -> 회전 바퀴수도

	countingStar()

	for i := 0; i < len(counts); i++ {
		counts[i] = R % counts[i]
	}

	fmt.Println(counts)

	rotate2()

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {

			fmt.Fprintf(w, "%d ", graph[i][j])
		}
		fmt.Fprintf(w, "\n")
	}
}

func countingStar() {

	for r := 0; r < int(math.Min(float64(N), float64(M)))/2; r++ {
		x1, y1, x2, y2 := r, r, N-1-r, M-1-r
		count := 0
		//좌
		for i := y2 - 1; i >= y1; i-- {
			count++
		}
		//하
		for i := x1 + 1; i <= x2; i++ {
			count++
		}
		//우
		for i := y1 + 1; i <= y2; i++ {
			count++
		}
		//상
		for i := x2 - 1; i >= x1; i-- {
			count++
		}

		counts = append(counts, count)
	}
}

func rotate2() {
	temp := [301][301]int{}

	for r := 0; r < int(math.Min(float64(N), float64(M)))/2; r++ {
		for rr := 0; rr < counts[r]; rr++ {

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

			for i := 0; i < N; i++ {
				for j := 0; j < M; j++ {
					if temp[i][j] != 0 {

						graph[i][j] = temp[i][j]
					}
				}
			}
		}

	}

}
