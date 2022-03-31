package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	N, M, R int
	graph   [101][101]int
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

	temp := [101][101]int{}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			temp[i][j] = graph[N-1-i][j]
		}
	}
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			graph[i][j] = temp[i][j]
		}
	}
}

func calculate2() {

	temp := [101][101]int{}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			temp[i][j] = graph[i][M-1-j]
		}
	}
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			graph[i][j] = temp[i][j]
		}
	}
}

func calculate3() {

	temp := [101][101]int{}

	h := N

	tmp := N
	N = M
	M = tmp

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			temp[i][j] = graph[h-1-j][i]
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			graph[i][j] = temp[i][j]

		}
	}
}

func calculate4() {
	temp := [101][101]int{}

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

			graph[i][j] = temp[i][j]
		}
	}
}

func calculate5() {

	temp := [101][101]int{}

	n, m := N/2, M/2

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			temp[i][j+m] = graph[i][j]
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			temp[i+n][j+m] = graph[i][j+m]
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			temp[i+n][j] = graph[i+n][j+m]
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			temp[i][j] = graph[i+n][j]
		}
	}

	// for i := 0; i < n; i++ {
	// 	for j := 0; j < m; j++ {
	// 		temp[i][m+j] = graph[i][j]
	// 	}
	// }
	// for i := 0; i < n; i++ {
	// 	for j := m; j < M; j++ {
	// 		temp[n+i][j] = graph[i][j]
	// 	}
	// }
	// for i := n; i < N; i++ {
	// 	for j := m; j < M; j++ {
	// 		temp[i][j-m] = graph[i][j]
	// 	}
	// }
	// for i := n; i < N; i++ {
	// 	for j := 0; j < m; j++ {
	// 		temp[i-n][j] = graph[i][j]
	// 	}
	// }

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {

			graph[i][j] = temp[i][j]

		}
	}
}

func calculate6() {

	temp := [101][101]int{}

	n, m := N/2, M/2

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			temp[i+n][j+m] = graph[i+n][j]
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			temp[i][j+m] = graph[i+n][j+m]
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			temp[i+n][j] = graph[i][j]
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			temp[i][j] = graph[i][j+m]
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {

			graph[i][j] = temp[i][j]

		}
	}

	// for i := 0; i < N; i++ {
	// 	fmt.Println(graph[i])
	// }
}
