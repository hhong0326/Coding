package main

import (
	"bufio"
	"fmt"
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

	cal := make([]int, R)

	for i := 0; i < R; i++ {
		fmt.Fscanf(r, "%d ", &cal[i])
	}

	for _, v := range cal {

		switch v {
		case 1:
			calculate1()
		case 2:
			calculate2()
		case 3:
			calculate3()
		case 4:
			calculate4()
		case 5:
			calculate5()
		case 6:
			calculate6()
		}

	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if graph[i][j] != 0 {
				fmt.Fprintf(w, "%d ", graph[i][j])
			}
		}
		fmt.Fprintf(w, "\n")
	}

}

// caluculate1 is changing top bottom
func calculate1() {
	for i, j := 0, N-1; i < N/2; i, j = i+1, j-1 {
		for k := 0; k < M; k++ {
			temp := graph[i][k]
			graph[i][k] = graph[j][k]
			graph[j][k] = temp
		}
	}
}

func calculate2() {

	for k := 0; k < N; k++ {
		for i, j := 0, M-1; i < M/2; i, j = i+1, j-1 {
			temp := graph[k][i]
			graph[k][i] = graph[k][j]
			graph[k][j] = temp

		}
	}
}

func calculate3() {

	temp := [301][301]int{}

	h := N

	tmp := M
	M = N
	N = tmp

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			temp[i][j] = graph[h-1-j][i]
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if temp[i][j] != 0 {
				graph[i][j] = temp[i][j]
			}
		}
	}
}

func calculate4() {
	temp := [301][301]int{}

	w := M
	tmp := M
	M = N
	N = tmp

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			temp[i][j] = graph[j][w-1-i]
			// graph[i][j] = 0
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if temp[i][j] != 0 {

				graph[i][j] = temp[i][j]
			}
		}
	}
}

func calculate5() {

	temp := [301][301]int{}

	n, m := N/2, M/2

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			temp[i][m+j] = graph[i][j]
		}
	}
	for i := 0; i < n; i++ {
		for j := m; j < M; j++ {
			temp[n+i][j] = graph[i][j]
		}
	}
	for i := n; i < N; i++ {
		for j := m; j < M; j++ {
			temp[i][j-m] = graph[i][j]
		}
	}
	for i := n; i < N; i++ {
		for j := 0; j < m; j++ {
			temp[i-n][j] = graph[i][j]
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if temp[i][j] != 0 {

				graph[i][j] = temp[i][j]
			}
		}
	}
}

func calculate6() {

	temp := [301][301]int{}

	n, m := N/2, M/2

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			temp[i+n][j] = graph[i][j]
		}
	}
	for i := n; i < N; i++ {
		for j := 0; j < m; j++ {
			temp[i][j+m] = graph[i][j]
		}
	}
	for i := n; i < N; i++ {
		for j := m; j < M; j++ {
			temp[i-n][j] = graph[i][j]
		}
	}
	for i := 0; i < N; i++ {
		for j := m; j < M; j++ {
			temp[i][j-m] = graph[i][j]
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if temp[i][j] != 0 {

				graph[i][j] = temp[i][j]
			}
		}
	}

	// for i := 0; i < N; i++ {
	// 	fmt.Println(graph[i])
	// }
}
